package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Listing struct {
	ListingId    *big.Int       `json:"listing_id"`
	ContractAddr common.Address `json:"contract_addr"`
	TokenId      *big.Int       `json:"token_id"`
	Seller       common.Address `json:"seller"`
	Price        *big.Int       `json:"price"`
}
