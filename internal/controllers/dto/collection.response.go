package dto

import "time"

type PostCollectionRes struct {
	Token string `json:"token"`
	Owner string `json:"owner"`

	Name        string         `json:"name"`
	Description string         `json:"description"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	Category    string         `json:"category"`
	CreatedAt   time.Time      `json:"created_at"`
}

type GetCollectionRes struct {
	Collections []Collection `json:"collections"`
	PageSize    int          `json:"page_size"`
	Page        int          `json:"page"`
}

type Collection PostCollectionRes
