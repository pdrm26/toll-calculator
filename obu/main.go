package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/pdrm26/toll-calculator/types"
)

const sendInterval = time.Second

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	u := url.URL{Scheme: "ws", Host: os.Getenv("DATA_RECEIVER_ENDPOINT"), Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial failed: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to OBU receiver server")

	for {
		lat, long := generateLocation()
		obu := types.OBU{
			ID:   generateOBUID(),
			Lat:  lat,
			Long: long,
		}
		if err := conn.WriteJSON(obu); err != nil {
			log.Fatal("Send error:", err)
		}
		time.Sleep(sendInterval)
	}
}

func generateOBUID() int {
	return rand.Intn(math.MaxInt)
}

func generateCoord() types.Coord {
	return types.Coord(rand.Float64() * 100)
}

func generateLocation() (types.Coord, types.Coord) {
	return generateCoord(), generateCoord()
}
