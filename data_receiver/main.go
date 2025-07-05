package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/pdrm26/toll-calculator/types"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial failed: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to OBU sender server")

	for {
		var obuData types.OBU
		if err := conn.ReadJSON(&obuData); err != nil {
			log.Println("Read error:", err)
			break
		}
		fmt.Printf("Received: %+v\n", obuData)
	}

}
