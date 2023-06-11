package dto

type GetEventReq struct {
	Token   string `query:"token" validate:"omitempty,eth_addr"`
	TokenId string `query:"token_id" validate:"omitempty,hexadecimal"`
	Name    string `query:"name" validate:"omitempty"`
	Address string `query:"address" validate:"omitempty,eth_addr"`
	Type    string `query:"type" validate:"omitempty"`
	Month   *int   `query:"month" validate:"omitempty"`
	Year    *int   `query:"year" validate:"omitempty"`
}
