package models

type GetNftsReq struct {
	ContractAddr string `query:"contract_addr" validate:"omitempty,eth_addr"`
	Owner        string `query:"owner" validate:"omitempty,eth_addr"`
	Offset       int32  `query:"offset" validate:"gte=0"`
}

type GetNftReq struct {
	ContractAddr string `query:"contract_addr"`
	TokenId      string `param:"tokenId"`
}
