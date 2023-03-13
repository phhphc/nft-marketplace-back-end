package repository

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/models"
)

type Reader interface {
	GetOrder(ctx context.Context, orderHash string) (models.Order, error)
	GetOrderList(ctx context.Context, offset, limit int) ([]models.Order, error)
	GetOrderByItemConsideration(ctx context.Context, tokenAddress common.Address, tokenId *big.Int) ([]models.Order, error)
	GetOrderByItemOffer(ctx context.Context, tokenAddress common.Address, tokenId *big.Int) ([]models.Order, error)
}

type Writer interface {
	InsertOrder(ctx context.Context, order models.Order) error
	SetOrderCancelled(ctx context.Context, orderHash string) error
	SetOrderValidated(ctx context.Context, orderHash string) error
	SetOrderFulfilled(ctx context.Context, orderHash string) error
	SetAllOrderCancelled(ctx context.Context, offerer string, counter *big.Int) error
}

type OrderRepository interface {
	Reader
	Writer
}

var _ OrderRepository = (*orderRepository)(nil)
