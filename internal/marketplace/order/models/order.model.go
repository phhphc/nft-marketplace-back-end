package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
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

	CreatedAt  time.Time
	ModifiedAt time.Time
}

type OfferItem struct {
	ItemType     *big.Int
	TokenId      *big.Int
	TokenAddress common.Address
	StartAmount  *big.Int
	EndAmount    *big.Int
}

type ConsiderationItem struct {
	ItemType     *big.Int
	TokenId      *big.Int
	TokenAddress common.Address
	StartAmount  *big.Int
	EndAmount    *big.Int
	Recipient    common.Address
}

const (
	NAIVE   = 0
	ERC20   = 1
	ERC721  = 2
	ERC1155 = 3
)

const (
	FULL_OPEN          = 0
	FULL_RESTRICTED    = 1
	PARTIAL_RESTRICTED = 2
	CONTRACT           = 3
)

func (o *Order) GetOrderHash() string {
	return o.OrderHash
}

func (o *Order) UpdateOrderIsCancelled(isCancelled bool) {
	o.IsCancelled = isCancelled
}

func (o *Order) UpdateOrderIsValidated(isValidated bool) {
	o.IsValidated = isValidated
}
