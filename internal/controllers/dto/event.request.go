package dto

type GetEventReq struct {
	Token   string `query:"token" validate:"omitempty,eth_addr"`
	TokenId string `query:"token_id" validate:"omitempty,hexadecimal"`
	Name    string `query:"name" validate:"omitempty"`
	From    string `query:"from" validate:"omitempty,eth_addr"`
	To      string `query:"to" validate:"omitempty,eth_addr"`
}
