package entities

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type MarketplaceSettings struct {
	Id          int64
	Marketplace common.Address
	Admin       common.Address
	Signer      common.Address
	Royalty     float64
	Sighash     common.Hash
	Signature   []byte
	CreatedAt   big.Int
}
