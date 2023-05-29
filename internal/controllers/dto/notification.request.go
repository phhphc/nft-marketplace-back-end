package dto

type GetNotificationReq struct {
	Address string `query:"address" validation:"eth_addr,omitempty"`
}

type UpdateNotificationReq struct {
	EventName string `json:"event_name"`
	OrderHash string `json:"order_hash" validate:"eth_hash"`
}
