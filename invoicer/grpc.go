package main

import (
	"context"

	"github.com/pdrm26/toll-calculator/types"
)

type GRPCServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewGRPCServer(svc Aggregator) *GRPCServer {
	return &GRPCServer{
		svc: svc,
	}
}

func (s *GRPCServer) Aggregate(ctx context.Context, dist *types.AggregatorDistance) (*types.None, error) {
	distance := types.Distance{
		OBUID:     int(dist.Obuid),
		Timestamp: dist.UnixTimestamp,
		Value:     dist.Value,
	}
	if err := s.svc.AggregateDistance(distance); err != nil {
		return nil, err
	}

	return &types.None{}, nil
}
