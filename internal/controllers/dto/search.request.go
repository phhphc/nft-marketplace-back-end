package dto

type SearchNftReq struct {
	Q          string `query:"q" validate:"omitempty"`
	Token      string `query:"token" validate:"omitempty,eth_addr"`
	Owner      string `query:"owner" validate:"omitempty,eth_addr"`
	Identifier string `query:"identifier" validate:"omitempty"`
	IsHidden   *bool  `query:"isHidden"`
	Offset     int32  `query:"offset" validate:"gte=0"`
	Limit      int32  `query:"limit" validate:"gte=0,lte=100"`
}
