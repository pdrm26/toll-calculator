package types

type Coord float64
type OBU struct {
	ID        int    `json:"id"`
	Lat       Coord  `json:"lat"`
	Long      Coord  `json:"long"`
	RequestID string `json:"requestID"`
}

type Distance struct {
	OBUID     int     `json:"obuID"`
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type Invoice struct {
	OBUID         int     `json:"obuID"`
	TotalDistance float64 `json:"totalDistance"`
	TotalPrice    float64 `json:"totalPrice"`
}
