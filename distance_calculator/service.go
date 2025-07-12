package main

import (
	"math"

	"github.com/pdrm26/toll-calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.OBU) (float64, error)
}

type CalculatorService struct{}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.OBU) (float64, error) {
	distance, err := calculateDistance(float64(data.Lat), float64(data.Long))
	if err != nil {
		return 0.0, err
	}

	return distance, nil
}

func calculateDistance(lat, long float64) (float64, error) {
	// Placeholder logic — replace with accurate distance formula if needed
	delta := lat - long
	return math.Sqrt(2 * math.Pow(delta, 2)), nil
}
