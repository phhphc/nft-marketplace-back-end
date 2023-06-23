package entities

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Nft struct {
	Token      common.Address
	Identifier *big.Int
	Owner      common.Address
	TokenUri   string

	Metadata map[string]any

	IsBurned bool
	IsHidden bool
}
