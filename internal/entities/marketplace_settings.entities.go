package entities

import (
	"github.com/ethereum/go-ethereum/common"
)

type MarketplaceSettings struct {
	Id          int64
	Marketplace common.Address
	Beneficiary common.Address
	Royalty     float64
}
