package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pdrm26/toll-calculator/types"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

var kafkaTopic string = "obudata"

type OBUReceiver struct {
	msgch    chan types.OBU
	conn     *websocket.Conn
	producer DataProducer
}

func NewOBUReceiver() (*OBUReceiver, error) {
	p, err := NewKafkaProducer()
	if err != nil {
		return nil, err
	}
	return &OBUReceiver{
		msgch:    make(chan types.OBU, 10),
		producer: p,
	}, nil
}

func (or *OBUReceiver) produceData(data types.OBU) error {
	return or.producer.ProduceData(data)
}

func (or *OBUReceiver) wsReceiveLoop() {
	fmt.Println("New OBU connected client connected!")

	for {
		var obu types.OBU
		if err := or.conn.ReadJSON(&obu); err != nil {
			log.Println("Read error:", err)
			break
		}
		if err := or.produceData(obu); err != nil {
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
