package entities

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type NotificationPost struct {
	Info      string         `json:"info"`
	EventName string         `json:"event_name"`
	OrderHash common.Hash    `json:"order_hash"`
	Address   common.Address `json:"address"`
}

type NotificationGet struct {
	IsViewed  bool           `json:"is_viewed"`
	Info      string         `json:"info"`
	EventName string         `json:"event_name"`
	OrderHash common.Hash    `json:"order_hash"`
	Address   common.Address `json:"address"`
	Token     common.Address `json:"token"`
	TokenId   *big.Int       `json:"token_id"`
	Quantity  int32          `json:"quantity"`
	Type      string         `json:"type"`
	Price     *big.Int       `json:"price"`
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	Date      time.Time      `json:"date"`
	Owner     common.Address `json:"owner"`
	NftImage  string         `json:"nft_image"`
	NftName   string         `json:"nft_name"`
}

type NotificationUpdate struct {
	EventName string         `json:"event_name"`
	OrderHash common.Hash    `json:"order_hash"`
}