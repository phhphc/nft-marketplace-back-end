package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type AppEvent struct {
	Key   []byte
	Value []byte
}

type NewErc721Event struct {
	Token      common.Address `json:"token"`
	Identifier *big.Int       `json:"idientifier"`
}

type NewCollectionEvent struct {
	Address      common.Address 	`json:"address"`
}
