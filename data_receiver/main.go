package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pdrm26/toll-calculator/types"
)

type OBUReceiver struct {
	conn *websocket.Conn
}

func NewOBUReceiver() (*OBUReceiver, error) {
	return &OBUReceiver{}, nil
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func (or *OBUReceiver) wsReceiveLoop() {
	fmt.Println("New OBU connected client connected!")

	for {
		var obu types.OBU
		if err := or.conn.ReadJSON(&obu); err != nil {
			log.Println("Read error:", err)
			break
		}
		fmt.Printf("Received: %+v\n", obu)
	}
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Upgrade failed: ", err)
		return
	}
	defer conn.Close()

	receiver := &OBUReceiver{conn: conn}
	receiver.wsReceiveLoop()
}

func main() {
	http.HandleFunc("/ws", handleWS)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
