package services

import (
	"context"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type Servicer interface {
	CreateOrder(ctx context.Context, order entities.Order) error
	FulFillOrder(ctx context.Context, order entities.Order) error

	CreateCollection(ctx context.Context, collection entities.Collection) (entities.Collection, error)
}

var _ Servicer = (*Services)(nil)
