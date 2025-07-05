package main

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pdrm26/toll-calculator/types"
)

const sendInterval = time.Second

type OBUSender struct {
	conn *websocket.Conn
}

func generateOBUID() int {
	return rand.Intn(math.MaxInt)
}

func generateCoord() types.Coord {
	return types.Coord(rand.Float64() * 100)
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func generateLocation() (types.Coord, types.Coord) {
	return generateCoord(), generateCoord()
}

func (o *OBUSender) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Upgrade failed: ", err)
		return
	}
	defer conn.Close()
	o.conn = conn

	o.handleOBUSender()
}

func (o *OBUSender) handleOBUSender() {
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
	obuSender := OBUSender{}
	http.HandleFunc("/ws", obuSender.handleWS)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
