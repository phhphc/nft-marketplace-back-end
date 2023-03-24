package dto

type PostOrderRes struct {
	OrderHash string `json:"order_hash"`
}

type GetOrderHashes struct {
	OrderHashes []string `json:"order_hashes"`
	PageSize    int      `json:"page_size"`
	Page        int      `json:"page"`
}
