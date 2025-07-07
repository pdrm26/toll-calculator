package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gorilla/websocket"
	"github.com/pdrm26/toll-calculator/types"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

var kafkaTopic string = "obudata"

type OBUReceiver struct {
	msgch    chan types.OBU
	conn     *websocket.Conn
	producer *kafka.Producer
}

func NewOBUReceiver() (*OBUReceiver, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return nil, err
	}
	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	return &OBUReceiver{
		msgch:    make(chan types.OBU, 10),
		producer: p,
	}, nil
}

func (or *OBUReceiver) producetData(data types.OBU) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = or.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)

	return err
}

func (or *OBUReceiver) wsReceiveLoop() {
	fmt.Println("New OBU connected client connected!")

	for {
		var obu types.OBU
		if err := or.conn.ReadJSON(&obu); err != nil {
			log.Println("Read error:", err)
			break
		}
		if err := or.producetData(obu); err != nil {
			fmt.Println("kafka produce error:", err)
		}
	}
}

func (or *OBUReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Upgrade failed: ", err)
		return
	}
	defer conn.Close()

	or.conn = conn
	or.wsReceiveLoop()
}

func main() {
	rec, err := NewOBUReceiver()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", rec.handleWS)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
