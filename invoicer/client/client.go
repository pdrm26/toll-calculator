package client

import (
	"context"

	"github.com/pdrm26/toll-calculator/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregatorDistance) error
}
