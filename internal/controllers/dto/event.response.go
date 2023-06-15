package dto

import "time"

type EventRes struct {
	Name        string    `json:"name"`
	Token       string    `json:"token"`
	TokenId     string    `json:"token_id"`
	Quantity    int       `json:"quantity,omitempty"`
	Type        string    `json:"type"`
	Price       string    `json:"price,omitempty"`
	From        string    `json:"from"`
	To          string    `json:"to,omitempty"`
	Date        time.Time `json:"date"`
	TxHash      string    `json:"tx_hash,omitempty"`
	OrderHash   string    `json:"order_hash"`
	NftImage    string    `json:"nft_image"`
	NftName     string    `json:"nft_name"`
	EndTime     string    `json:"end_time,omitempty"`
	IsCancelled bool      `json:"is_cancelled"`
	IsFulfilled bool      `json:"is_fulfilled"`
}

type GetEventRes struct {
	Events []EventRes `json:"events"`
}
