package dto

type PostCollectionReq struct {
	Token string `json:"token" validate:"eth_addr"`
	Owner string `json:"owner" validate:"eth_addr"`

	Name        string `json:"name" validate:"required"`
	Description string `json:"desctiption" validate:"required"`
	Category    string `json:"category" validate:"alphanum"`
}
