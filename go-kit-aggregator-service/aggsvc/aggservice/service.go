package aggservice

import (
	"context"

	"github.com/pdrm26/toll-calculator/types"
)

const basePrice = 1.12

type Service interface {
	Aggregate(context.Context, types.Distance) error
	Invoice(context.Context, int) (*types.Invoice, error)
}

type basicService struct {
	store Storer
}

func newBasicService(store Storer) Service {
	return &basicService{store: store}
}

func (s basicService) Aggregate(ctx context.Context, distance types.Distance) error {
	return s.store.Insert(distance)
}

func (s basicService) Invoice(ctx context.Context, obuID int) (*types.Invoice, error) {
	dist, err := s.store.Get(obuID)
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

func NewAggregator() Service {
	var svc Service
	{
		svc = newBasicService(NewMemoryStore())
		svc = newLoggingMiddleware()(svc)
		svc = newInstrumentingMiddleware()(svc)
	}
	return svc
}
