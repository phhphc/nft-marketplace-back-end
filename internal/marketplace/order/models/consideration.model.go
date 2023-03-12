package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ConsiderationItem struct {
	TypeNumber   *big.Int       `json:"type_number"`
	TokenId      *big.Int       `json:"token_id"`
	TokenAddress common.Address `json:"token_address"`
	StartAmount  *big.Int       `json:"start_amount"`
	EndAmount    *big.Int       `json:"end_amount"`
	Recipient    common.Address `json:"recipient"`
}
