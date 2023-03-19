package dto

type Response struct {
	Data      any  `json:"data,omitempty"`
	IsSuccess bool `json:"is_success"`
	Error     any  `json:"error,omitempty"`
}
