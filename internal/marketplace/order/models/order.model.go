package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Order struct {
	Offer         []OfferItem
	Consideration []ConsiderationItem
	OrderHash     string

	Offerer     common.Address
	Signature   string
	OrderType   *big.Int
	StartTime   *big.Int
	EndTime     *big.Int
	Counter     *big.Int
	Salt        string
	IsCancelled bool
	IsValidated bool

	Zone     common.Address
	ZoneHash string
}

func (o *Order) GetOrderHash() string {
	return o.OrderHash
}

func (o *Order) UpdateOrderIsCancelled(isCancelled bool) {
	o.IsCancelled = isCancelled
}

func (o *Order) UpdateOrderIsValidated(isValidated bool) {
	o.IsValidated = isValidated
}
