package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/models"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order models.Order) error
	GetOrder(ctx context.Context, orderHash string) (models.Order, error)
	// Dummy name, need rework :))
	GetAllOrderByOfferItem(ctx context.Context, tokenAddress common.Address, tokenId *big.Int) ([]models.Order, error)
	GetAllOrderByConsiderationItem(ctx context.Context, tokenAddress common.Address, tokenId *big.Int) ([]models.Order, error)

	UpdateOrderIsCancelled(ctx context.Context, orderHash string) error
	UpdateOrderIsValidated(ctx context.Context, orderHash string) error
	UpdateOrderIsFulfilled(ctx context.Context, orderHash string) error
}
