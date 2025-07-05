package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pdrm26/toll-calculator/types"
)

const sendInterval = time.Second

func generateOBUID() uint64 {
	return rand.Uint64()
}

func generateCoord() types.Coord {
	return types.Coord(rand.Float64() * 100)
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func generateLocation() (types.Coord, types.Coord) {
	return generateCoord(), generateCoord()
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Upgrade failed: ", err)
		return
	}
	defer conn.Close()

	for {
		lat, long := generateLocation()
		obu := types.OBU{
			ID:   generateOBUID(),
			Lat:  lat,
			Long: long,
		}
		if err := o.conn.WriteJSON(obu); err != nil {
			log.Fatal("Send error:", err)
		}
		time.Sleep(sendInterval)
	}

}

func main() {
	http.HandleFunc("/ws", serveWS)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
