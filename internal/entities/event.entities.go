package entities

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

type Event struct {
	Name     string         `json:"name"`
	Token    common.Address `json:"token"`
	TokenId  *big.Int       `json:"token_id"`
	Quantity int32          `json:"quantity"`
	Type     string         `json:"type"`
	Price    *big.Int       `json:"price"`
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Date     time.Time      `json:"date"`
	Link     string         `json:"link"`
	Owner    common.Address	`json:"owner,omitempty"`
	OrderHash   common.Hash `json:"order_hash"`
	NftImage    string      `json:"nft_image"`
	NftName     string      `json:"nft_name"`
	EndTime     *big.Int    `json:"end_time"`
	IsCancelled bool        `json:"is_cancelled"`
	IsFulfilled bool        `json:"is_fulfilled"`
	IsExpired 	bool		`json:"is_expired"`
}

type EventRead struct {
	Name    string
	Token   common.Address
	TokenId *big.Int
	Type    string
	Address common.Address
}
