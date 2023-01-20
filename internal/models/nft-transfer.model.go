package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type NftTransfer struct {
	ContractAddr common.Address
	TokenId      *big.Int
	From         common.Address
	To           common.Address
}
