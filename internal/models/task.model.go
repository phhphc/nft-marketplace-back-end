package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type NewErc721Task struct {
	Token      common.Address `json:"token"`
	Identifier *big.Int       `json:"idientifier"`
}

type NewCollectionTask struct {
	Address      common.Address 	`json:"address"`
}
