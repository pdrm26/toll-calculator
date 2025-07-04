package main

import (
	"fmt"
	"math/rand"
	"time"
)

const sendInterval = time.Second

type OBU struct {
	ID   uint64 `json:"id"`
	Lat  Coord  `json:"lat"`
	Long Coord  `json:"long"`
}

type Coord float64

func generateOBUID() uint64 {
	return rand.Uint64()
}

func generateCoord() Coord {
	return Coord(rand.Float64() * 100)
}

func generateLocation() (Coord, Coord) {
	return generateCoord(), generateCoord()
}

func main() {

	for {
		lat, long := generateLocation()
		obu := OBU{
			ID:   generateOBUID(),
			Lat:  lat,
			Long: long,
		}
		fmt.Printf("%+v\n", obu)
		time.Sleep(sendInterval)
	}

}
