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
	Quantity int32
	Type     string
	Price    *big.Int
	From     common.Address
	To       common.Address
	Date     time.Time
	Link     string
}

type EventRead struct {
	Name    string
	Token   common.Address
	TokenId *big.Int
	Type    string
	Address common.Address
}
