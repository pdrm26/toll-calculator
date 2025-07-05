package types

type Coord float64
type OBU struct {
	ID   int   `json:"id"`
	Lat  Coord `json:"lat"`
	Long Coord `json:"long"`
}
