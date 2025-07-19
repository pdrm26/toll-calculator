package client

import (
	"context"

	"github.com/pdrm26/toll-calculator/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Endpoint string
	Client   types.AggregatorClient
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := types.NewAggregatorClient(conn)

	return &GRPCClient{
		Endpoint: endpoint,
		Client:   client,
	}, nil
}

func (c *GRPCClient) Aggregate(ctx context.Context, req *types.AggregatorDistance) error {
	_, err := c.Client.Aggregate(ctx, req)
	return err
}

func (c *GRPCClient) GetInvoice(ctx context.Context, obuID int) (*types.Invoice, error) {
	return &types.Invoice{
		OBUID:         obuID,
		TotalDistance: 1242098.124,
		TotalPrice:    10.10001294,
	}, nil
}
