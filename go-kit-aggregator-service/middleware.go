package main

import (
	"context"

	"github.com/pdrm26/toll-calculator/types"
)

type Middleware func(Service) Service

type loggingMiddleware struct {
	next Service
}

func newLoggingMiddleware() Middleware {
	return func(next Service) Service {
		return loggingMiddleware{next}
	}
}

func (l loggingMiddleware) Aggregate(ctx context.Context, distance types.Distance) (err error) {
	return nil
}

func (l loggingMiddleware) Calculate(ctx context.Context, id int) (invoice *types.Invoice, err error) {
	return nil, nil
}

// Instrumentation middleware
type instrumentingMiddleware struct {
	next Service
}

func newInstrumentingMiddleware() Middleware {
	return func(next Service) Service {
		return instrumentingMiddleware{next}
	}
}

func (imw instrumentingMiddleware) Aggregate(ctx context.Context, distance types.Distance) (err error) {
	return nil
}

func (imw instrumentingMiddleware) Calculate(ctx context.Context, id int) (invoice *types.Invoice, err error) {
	return nil, nil
}
