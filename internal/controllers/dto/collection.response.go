package dto

import "time"

type PostCollectionRes struct {
	Token string `json:"token"`
	Owner string `json:"owner"`

	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
}
