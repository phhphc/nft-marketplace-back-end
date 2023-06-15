package infrastructure

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type FindOrderOffer struct {
	Token      common.Address
	Identifier *big.Int
}

type FindOrderConsideration struct {
	Token      common.Address
	Identifier *big.Int
}

type OrderReader interface {
	FindOrder(
		ctx context.Context,
		offer FindOrderOffer,
		consideration FindOrderConsideration,
		orderHash common.Hash,
		offerer common.Address,
		IsFulfilled *bool,
		IsCancelled *bool,
		IsInvalid *bool,
	) ([]entities.Order, error)
	FindExpiredOrder(
		ctx context.Context,
	) ([]entities.ExpiredOrder, error)
}
