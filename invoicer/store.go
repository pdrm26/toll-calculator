package main

import "github.com/pdrm26/toll-calculator/types"

type MemoryStore struct {
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) Insert(distance types.Distance) error {
	return nil
}
