package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ConsiderationItem struct {
	TypeNumber   *big.Int
	TokenId      *big.Int
	TokenAddress common.Address
	StartAmount  *big.Int
	EndAmount    *big.Int
	Recipient    common.Address
}
