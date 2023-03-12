package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Order struct {
	Offer         []OfferItem         `json:"offer"`
	Consideration []ConsiderationItem `json:"consideration"`
	OrderHash     string              `json:"order_hash"`

	Offerer     common.Address `json:"offerer"`
	Signature   string         `json:"signature,omitempty"`
	OrderType   *big.Int       `json:"order_type"`
	StartTime   *big.Int       `json:"start_time"`
	EndTime     *big.Int       `json:"end_time"`
	Counter     *big.Int       `json:"counter"`
	Salt        string         `json:"salt"`
	IsCancelled bool           `json:"is_cancelled"`
	IsValidated bool           `json:"is_validated"`

	Zone     common.Address `json:"zone,omitempty"`
	ZoneHash string         `json:"zone_hash,omitempty"`
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
