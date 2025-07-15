package main

import (
	"github.com/pdrm26/toll-calculator/types"
)

const basePrice = 1.12

type Aggregator interface {
	AggregateDistance(distance types.Distance) error
	CalculateInvoice(obuID int) (*types.Invoice, error)
}

type Storer interface {
	Insert(types.Distance) error
	Get(int) (float64, error)
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{store: store}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	return i.store.Insert(distance)
}

func (i *InvoiceAggregator) CalculateInvoice(obuID int) (*types.Invoice, error) {
	dist, err := i.store.Get(obuID)
	if err != nil {
		return nil, err
	}
	invoce := &types.Invoice{
		OBUID:         obuID,
		TotalDistance: dist,
		TotalPrice:    basePrice * dist,
	}

	return invoce, nil
}
