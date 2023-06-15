package dto

type GetOrderStruct struct {
	OrderHash     string               `json:"orderHash"`
	Offerer       string               `json:"offerer"`
	StartTime     string               `json:"startTime"`
	EndTime       string               `json:"endTime"`
	Salt          string               `json:"salt"`
	Signature     string               `json:"signature"`
	Status        OrderStatus          `json:"status"`
	Offer         []OrderOffer         `json:"offer"`
	Consideration []OrderConsideration `json:"consideration"`
}

type PostOrderRes struct {
	OrderHash string `json:"order_hash"`
}

type OrderStatus struct {
	IsCancelled bool `json:"isCancelled"`
	IsFulfilled bool `json:"isFulfilled"`
	IsInvalid   bool `json:"isInvalid"`
}

type OrderOffer struct {
	ItemType    int    `json:"itemType"`
	Token       string `json:"token"`
	Identifier  string `json:"identifier"`
	StartAmount string `json:"startAmount"`
	EndAmount   string `json:"endAmount"`
}

type OrderConsideration struct {
	ItemType    int    `json:"itemType"`
	Token       string `json:"token"`
	Identifier  string `json:"identifier"`
	StartAmount string `json:"startAmount"`
	EndAmount   string `json:"endAmount"`
	Recipient   string `json:"recipient"`
}
