package types

type OBU struct {
	ID   uint64 `json:"id"`
	Lat  Coord  `json:"lat"`
	Long Coord  `json:"long"`
}

type Coord float64
