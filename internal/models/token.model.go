package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Nft struct {
	TokenId      *big.Int       `json:"token_id"`
	ContractAddr common.Address `json:"contract_addr"`
	Owner        common.Address `json:"owner"`
	Listing      *NftListing    `json:"listing,omitempty"`
}

type NftListing struct {
	ListingId *big.Int       `json:"listing_id"`
	Seller    common.Address `json:"seller"`
	Price     *big.Int       `json:"price"`
}
