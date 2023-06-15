package dto

import "time"

type NotificationRes struct {
	IsViewed  bool		`json:"is_viewed"`
	Info      string    `json:"info"`
	EventName string    `json:"event_name"`
	OrderHash string    `json:"order_hash"`
	Address   string	`json:"address"`
	Token     string	`json:"token"`
	TokenId   string    `json:"token_id"`
	Quantity  int32     `json:"quantity"`
	Type      string    `json:"type"`
	Price     string    `json:"price"`
	From      string	`json:"from"`
	To        string	`json:"to"`
	Date      time.Time `json:"date"`
	Owner     string	`json:"owner"`
	NftImage  string    `json:"nft_image"`
	NftName   string    `json:"nft_name"`
}

type GetNotificationRes struct {
	Notifications []NotificationRes	`json:"notifications"`
}
