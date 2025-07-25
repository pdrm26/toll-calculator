package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/pdrm26/toll-calculator/types"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

type OBUReceiver struct {
	msgch    chan types.OBU
	conn     *websocket.Conn
	producer DataProducer
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	rec, err := NewOBUReceiver()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", rec.handleWS)
	log.Fatal(http.ListenAndServe(os.Getenv("DATA_RECEIVER_ENDPOINT"), nil))
}

func NewOBUReceiver() (*OBUReceiver, error) {
	p, err := NewKafkaProducer(os.Getenv("KAFKA_TOPIC_NAME"))
	if err != nil {
		return nil, err
	}
	p = NewLogMiddleware(p)
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
		obu.RequestID = uuid.New().String()
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
