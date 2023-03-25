package entities

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type NftRead struct {
	Token       common.Address `json:"token"`
	Identifier  *big.Int       `json:"identifier"`
	Owner       common.Address `json:"owner"`
	Image       string         `json:"image"`
	Name        string         `json:"name"`
	Description string         `json:"description"`

	BlockNumber *big.Int       `json:"block_number"`
	TxIndex     *big.Int       `json:"tx_index"`
	Listings    []*ListingRead `json:"prices"`
}

type ListingRead struct {
	OrderHash common.Hash  `json:"order_hash"`
	ItemType  EnumItemType `json:"item_type"`
	Price     *big.Int     `json:"price"`
}