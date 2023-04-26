package dto

import "github.com/phhphc/nft-marketplace-back-end/internal/entities"

type PostOrderReq struct {
	OrderHash string `json:"order_hash" validate:"eth_hash"`

	Offerer string `json:"offerer" validate:"eth_addr"`
	Zone    string `json:"zone" validate:"eth_addr"`

	Offer         []OfferItemReq         `json:"offer" validate:"required,dive"`
	Consideration []ConsiderationItemReq `json:"consideration" validate:"required,dive"`

	OrderType entities.EnumOrderType `json:"order_type" validate:"gte=0"`

	ZoneHash  string `json:"zone_hash" validate:"eth_hash"`
	Salt      string `json:"salt" validate:"eth_hash"`
	StartTime string `json:"start_time" validate:"hexadecimal"`
	EndTime   string `json:"end_time" validate:"hexadecimal"`

	Signature string `json:"signature" validate:"hexadecimal,startswith=0x"`
}

type OfferItemReq struct {
	ItemType entities.EnumItemType `json:"item_type" validate:"gte=0"`

	Token       string `json:"token" validate:"eth_addr"`
	Identifier  string `json:"identifier" validate:"hexadecimal"`
	StartAmount string `json:"start_amount" validate:"hexadecimal"`
	EndAmount   string `json:"end_amount" validate:"hexadecimal"`
}

type ConsiderationItemReq struct {
	ItemType entities.EnumItemType `json:"item_type" validate:"gte=0"`

	Token       string `json:"token" validate:"eth_addr"`
	Identifier  string `json:"identifier" validate:"hexadecimal"`
	StartAmount string `json:"start_amount" validate:"hexadecimal"`
	EndAmount   string `json:"end_amount" validate:"hexadecimal"`
	Recipient   string `json:"recipient" validate:"eth_addr"`
}

type GetOrderHash struct {
	ConsiderationToken      string `query:"consideration_token" validate:"omitempty,eth_addr"`
	ConsiderationIdentifier string `query:"consideration_identifier" validate:"omitempty,hexadecimal"`
	OfferToken              string `query:"offer_token" validate:"omitempty,eth_addr"`
	OfferIdentifier         string `query:"offer_identifier" validate:"omitempty,hexadecimal"`
}

type GetOrder struct {
	OrderHash string `query:"orderHash" validate:"omitempty,eth_hash"`

	ConsiderationToken      string `query:"considerationToken" validate:"omitempty,eth_addr"`
	ConsiderationIdentifier string `query:"considerationIdentifier" validate:"omitempty,hexadecimal"`

	OfferToken      string `query:"offerToken" validate:"omitempty,eth_addr"`
	OfferIdentifier string `query:"offerIdentifier" validate:"omitempty,hexadecimal"`

	IsFulfilled *bool `query:"isFulfilled"`
	IsCancelled *bool `query:"isCancelled"`
	IsInvalid   *bool `query:"isInvalid"`
}

type GetOrderV1 struct {
	OrderHash string `query:"order_hash" validate:"omitempty,eth_hash"`
}
