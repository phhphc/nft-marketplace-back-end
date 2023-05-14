package dto

type GetListNftReq struct {
	Token      string `query:"token" validate:"omitempty,eth_addr"`
	Owner      string `query:"owner" validate:"omitempty,eth_addr"`
	Identifier string `query:"identifier" validate:"omitempty,eth_addr"`
	IsHidden   *bool  `query:"isHidden"`
	Offset     int32  `query:"offset" validate:"gte=0"`
	Limit      int32  `query:"limit" validate:"gte=0,lte=100"`
}

type UpdateNftStatusReq struct {
	Token      string `param:"token" validate:"required,eth_addr"`
	Identifier string `param:"identifier" validate:"required,hexadecimal"`
	IsHidden   bool   `json:"isHidden"`
}

type UpdateNftStatusRes struct {
	Token      string `json:"token"`
	Identifier string `json:"identifier"`
	IsHidden   bool   `json:"isHidden"`
}

type GetNftReq struct {
	Token      string `param:"token"`
	Identifier string `param:"identifier"`
}
