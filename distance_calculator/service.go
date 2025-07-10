package main

import "github.com/pdrm26/toll-calculator/types"

type calculatorServicer interface {
	CalculateDistance(types.OBU) (float64, error)
}

type CalculateService struct{}

func NewCalculateService() *CalculateService {
	return &CalculateService{}
}

func (s *CalculateService) CalculateDistance(data *types.OBU) (float64, error) {
	return 0.0, nil
}
