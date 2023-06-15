package infrastructure

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type UpdateOrderCondition struct {
	OrderHash common.Hash
	Offerer   common.Address
}

type UpdateOrderValue struct {
	IsValidated *bool
	IsCancelled *bool
	IsFulfilled *bool
	IsInvalid   *bool
}

type OrderWritter interface {
	InsertOneOrder(
		ctx context.Context,
		order entities.Order,
	) (entities.Order, error)
	UpdateOrderStatus(
		ctx context.Context,
		condition UpdateOrderCondition,
		value UpdateOrderValue,
	) error
	UpdateOrderStatusByOffer(
		ctx context.Context,
		condition UpdateOrderStatusByOfferCondition,
		value UpdateOrderValue,
	) error
}

type UpdateOrderStatusByOfferCondition struct {
	Offerer    common.Address
	Token      common.Address
	Identifier *big.Int
}
