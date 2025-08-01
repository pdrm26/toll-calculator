package aggservice

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
	return l.next.Aggregate(ctx, distance)
}

func (l loggingMiddleware) Invoice(ctx context.Context, id int) (invoice *types.Invoice, err error) {
	return l.next.Invoice(ctx, id)
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
	return imw.next.Aggregate(ctx, distance)
}

func (imw instrumentingMiddleware) Invoice(ctx context.Context, id int) (invoice *types.Invoice, err error) {
	return imw.next.Invoice(ctx, id)
}
