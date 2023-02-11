package models

type GetNftsRes struct {
	Nfts   []Nft `json:"nfts"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

// type GetNftReq struct {
// 	ContractAddr string `query:"contract_addr"`
// 	TokenId      string `param:"tokenId"`
// }
