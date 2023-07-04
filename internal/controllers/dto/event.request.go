package dto

type GetEventReq struct {
	Token     string `query:"token" validate:"omitempty,eth_addr"`
	TokenId   string `query:"token_id" validate:"omitempty,hexadecimal"`
	Name      string `query:"name" validate:"omitempty"`
	Address   string `query:"address" validate:"omitempty,eth_addr"`
	Type      string `query:"type" validate:"omitempty"`
	StartDate string `query:"start_date" validate:"omitempty"`
	EndDate   string `query:"end_date" validate:"omitempty"`
}
