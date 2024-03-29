package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type NftTransfer struct {
	Token      common.Address
	Identifier *big.Int
	From       common.Address
	To         common.Address
}
