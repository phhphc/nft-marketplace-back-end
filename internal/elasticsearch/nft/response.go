package nft

type CreateResponse struct {
	Token      string `json:"token"`
	Identifier string `json:"identifier"`
}

type UpdateResponse struct {
	Token      string `json:"token"`
	Identifier string `json:"identifier"`
}

type FindOneResponse struct {
	Token      string `json:"token"`
	Identifier string `json:"identifier"`
	Owner      string `json:"owner"`
}
