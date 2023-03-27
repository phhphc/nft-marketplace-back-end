package dto

type GetNftsRes struct {
	Nfts   []*GetNftRes `json:"nfts"`
	Limit  int32        `json:"limit"`
	Offset int32        `json:"offset"`
}

type GetNftRes struct {
	Token       string              `json:"token"`
	Identifier  string              `json:"identifier"`
	Owner       string              `json:"owner"`
	Image       string              `json:"image"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Metadata    []byte              `json:"metadata"`
	Listings    []*GetNftListingRes `json:"listings"`
}

type GetNftListingRes struct {
	OrderHash  string `json:"order_hash"`
	ItemType   int    `json:"item_type"`
	StartPrice string `json:"start_price"`
	EndPrice   string `json:"end_price"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}
