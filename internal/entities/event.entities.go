package entities

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

type Event struct {
	Name     string
	Token    common.Address
	TokenId  *big.Int
	Quantity *big.Int
	Price    *big.Int
	From     common.Address
	To       common.Address
	Date     time.Time
	Link     string
}
