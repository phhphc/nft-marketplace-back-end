package entities

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Order struct {
	OrderHash common.Hash

	Offerer       common.Address
	Offer         []OfferItem
	Consideration []ConsiderationItem
	Salt          *common.Hash
	StartTime     *big.Int
	EndTime       *big.Int

	Recipient *common.Address
	Signature []byte
}

type OfferItem struct {
	ItemType   EnumItemType
	Token      common.Address
	Identifier *big.Int

	Amount      *big.Int
	StartAmount *big.Int
	EndAmount   *big.Int
}

type ConsiderationItem struct {
	ItemType    EnumItemType
	Token       common.Address
	Identifier  *big.Int
	StartAmount *big.Int

	Amount    *big.Int
	EndAmount *big.Int
	Recipient common.Address
}

type EnumItemType int

const (
	NAIVE EnumItemType = iota
	ERC20
	ERC721
	ERC1155
)

func (e *EnumItemType) Int() int {
	return int(*e)
}

type ExpiredOrder struct {
	EventName		string
	OrderHash		common.Hash
	EndTime       	*big.Int
	IsCancelled		bool
	IsInvalid		bool
	Offerer			common.Address
}