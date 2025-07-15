package main

import (
	"fmt"

	"github.com/pdrm26/toll-calculator/types"
)

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (s *MemoryStore) Insert(distance types.Distance) error {
	s.data[distance.OBUID] += distance.Value
	return nil
}

func (s *MemoryStore) Get(obuID int) (float64, error) {
	dist, ok := s.data[obuID]
	if !ok {
		return 0.0, fmt.Errorf("could not find distance for obuID %d", obuID)
	}
	return dist, nil
}
