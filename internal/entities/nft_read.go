package entities

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type NftRead struct {
	Token      common.Address `json:"token"`
	Identifier *big.Int       `json:"identifier"`
	Owner      common.Address `json:"owner"`
	Metadata   map[string]any `json:"metadata"`
	IsHidden   bool           `json:"isHidden"`

	Image       string `json:"image,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	BlockNumber *big.Int       `json:"block_number"`
	TxIndex     *big.Int       `json:"tx_index"`
	Listings    []*ListingRead `json:"prices"`
}

type ListingRead struct {
	OrderHash  common.Hash  `json:"order_hash"`
	ItemType   EnumItemType `json:"item_type"`
	StartPrice *big.Int     `json:"start_price"`
	EndPrice   *big.Int     `json:"end_price"`
	StartTime  *big.Int     `json:"start_time"`
	EndTime    *big.Int     `json:"end_time"`
}
