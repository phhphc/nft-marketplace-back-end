package entities

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Nft struct {
	Token      common.Address `json:"token"`
	Identifier *big.Int       `json:"identifier"`
	Owner      common.Address `json:"owner"`
	TokenUri   string         `json:"token_uri"`
	IsBurned   bool           `json:"is_burned"`

	BlockNumber *big.Int `json:"block_number"`
	TxIndex     *big.Int `json:"tx_index"`
}
