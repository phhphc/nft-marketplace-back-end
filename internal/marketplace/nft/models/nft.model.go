package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Nft struct {
	TokenId      *big.Int       `json:"token_id"`
	TokenAddress common.Address `json:"token_address"`
	TokenUri     string         `json:"token_uri"`
	Owner        common.Address `json:"owner"`
}
