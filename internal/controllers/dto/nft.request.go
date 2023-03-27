package dto

type GetListNftReq struct {
	Token      string `query:"token" validate:"omitempty,eth_addr"`
	Owner      string `query:"owner" validate:"omitempty,eth_addr"`
	Identifier string `query:"identifier" validate:"omitempty,eth_addr"`
	Offset     int32  `query:"offset" validate:"gte=0"`
	Limit      int32  `query:"limit" validate:"gte=0,lte=100"`
}

type GetNftReq struct {
	Token      string `param:"token"`
	Identifier string `param:"identifier"`
}
