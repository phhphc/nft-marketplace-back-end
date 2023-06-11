package dto

type GetNotificationReq struct {
	Address string `query:"address" validate:"omitempty,eth_addr"`
	IsViewed *bool  `query:"is_viewed" validate:"omitempty"`
}

type UpdateNotificationReq struct {
	EventName string `json:"event_name"`
	OrderHash string `json:"order_hash" validate:"eth_hash"`
}
