package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pdrm26/toll-calculator/types"
)

const sendInterval = time.Second

func generateOBUID() uint64 {
	return rand.Uint64()
}

func generateCoord() types.Coord {
	return types.Coord(rand.Float64() * 100)
}

func generateLocation() (types.Coord, types.Coord) {
	return generateCoord(), generateCoord()
}

func main() {

	for {
		lat, long := generateLocation()
		obu := types.OBU{
			ID:   generateOBUID(),
			Lat:  lat,
			Long: long,
		}
		fmt.Printf("%+v\n", obu)
		time.Sleep(sendInterval)
	}

}
