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
	Listings    []*GetNftListingRes `json:"listings"`
}

type GetNftListingRes struct {
	OrderHash          string `json:"order_hash"`
	ItemType           string `json:"item_type"`
	ConsiderationPrice string `json:"price"`
}
