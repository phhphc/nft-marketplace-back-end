package dto

import "time"

type EventRes struct {
	Name     string    `json:"name"`
	Token    string    `json:"token"`
	TokenId  string    `json:"token_id"`
	Quantity int       `json:"quantity,omitempty"`
	Price    string    `json:"price,omitempty"`
	From     string    `json:"from"`
	To       string    `json:"to,omitempty"`
	Date     time.Time `json:"date"`
	Link     string    `json:"link,omitempty"`
}

type GetEventRes struct {
	Events []EventRes `json:"events"`
}
