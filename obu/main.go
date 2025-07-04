package main

import (
	"fmt"
	"math/rand"
	"time"
)

const sendInterval = time.Second

type OBUData struct {
	OBUID int     `json:"obuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

type Coord float64

func generateCoord() Coord {
	return Coord(rand.Float64() * 100)
}

func generateLocation() (Coord, Coord) {
	return generateCoord(), generateCoord()
}

func main() {

	for {
		fmt.Println(generateLocation())
		time.Sleep(sendInterval)
	}

}
