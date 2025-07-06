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

func (or *OBUReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Upgrade failed: ", err)
		return
	}
	defer conn.Close()
	or.conn = conn
	for {
		var obuData types.OBU
		if err := conn.ReadJSON(&obuData); err != nil {
			log.Println("Read error:", err)
			break
		}
		fmt.Printf("Received: %+v\n", obuData)
	}
}

func main() {
	recv, err := NewOBUReceiver()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.handleWS)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
