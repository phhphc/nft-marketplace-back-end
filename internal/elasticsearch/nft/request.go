package nft

type FindOneRequest struct {
	Token      string `json:"token"`
	Identifier string `json:"identifier"`
}

type CreateRequest struct {
	Token      string `json:"token"`
	Identifier string `json:"identifier"`
	Owner      string `json:"owner"`
}

type UpdateRequest struct {
	Token      string `json:"token"`
	Identifier string `json:"identifier"`
	Owner      string `json:"owner"`
}

type DeleteRequest struct {
	Token      string `json:"token"`
	Identifier string `json:"identifier"`
}
