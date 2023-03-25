package dto

type PostCollectionReq struct {
	Token string `json:"token" validate:"eth_addr"`
	Owner string `json:"owner" validate:"eth_addr"`

	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Metadata    map[string]any `json:"metadata" validate:"omitempty"`
	Category    string         `json:"category" validate:"alphanum"`
}

type GetCollectionReq struct {
	Token    string `query:"token" validate:"omitempty,eth_addr"`
	Owner    string `query:"owner" validate:"omitempty,eth_addr"`
	Name     string `query:"name" validate:"omitempty,required"`
	Category string `query:"category" validate:"omitempty,alphanum"`
}
