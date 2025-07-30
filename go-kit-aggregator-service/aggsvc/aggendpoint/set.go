package aggendpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/pdrm26/toll-calculator/go-kit-aggregator-service/aggsvc/aggservice"
	"github.com/pdrm26/toll-calculator/types"
)

type Set struct {
	AggregateEndpoint endpoint.Endpoint
	InvoiceEndpoint   endpoint.Endpoint
}

func makeAggregateEndpoint(s aggservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AggregateRequest)
		err = s.Aggregate(ctx, req.Distance)
		return AggregateResponse{Err: err}, nil
	}
}

func makeInvoiceEndpoint(s aggservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(InvoiceRequest)
		v, err := s.Calculate(ctx, req.OBUID)
		return InvoiceResponse{Invoice: *v, Err: err}, nil
	}
}

func (s Set) Aggregate(ctx context.Context, distance types.Distance) error {
	resp, err := s.AggregateEndpoint(ctx, AggregateRequest{Distance: distance})
	if err != nil {
		return err
	}
	response := resp.(AggregateResponse)
	return response.Err
}

func (s Set) Invoice(ctx context.Context, obuID int) (*types.Invoice, error) {
	resp, err := s.InvoiceEndpoint(ctx, InvoiceRequest{OBUID: obuID})
	if err != nil {
		return nil, err
	}
	response := resp.(InvoiceResponse)
	return &response.Invoice, response.Err

}

type AggregateRequest struct {
	Distance types.Distance
}

type AggregateResponse struct {
	Err error `json:"-"`
}

type InvoiceRequest struct {
	OBUID int `json:"obuID"`
}

type InvoiceResponse struct {
	Invoice types.Invoice `json:"invoice"`
	Err     error         `json:"-"`
}
