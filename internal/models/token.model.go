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
	Metadata     *NftMetadata   `json:"metadata,omitempty"`
}

type NftListing struct {
	ListingId *big.Int       `json:"listing_id"`
	Seller    common.Address `json:"seller"`
	Price     *big.Int       `json:"price"`
}

type NftMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
