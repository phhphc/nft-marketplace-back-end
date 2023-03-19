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
	StartTime string `json:"start_time" validate:"hexadecimal,startswith=0x"`
	EndTime   string `json:"end_time" validate:"hexadecimal,startswith=0x"`

	Signature string `json:"signature" validate:"hexadecimal,startswith=0x"`
}

type OfferItemReq struct {
	ItemType entities.EnumItemType `json:"item_type" validate:"gte=0"`

	Token       string `json:"token" validate:"eth_addr"`
	Identifier  string `json:"identifier" validate:"hexadecimal,startswith=0x"`
	StartAmount string `json:"start_amount" validate:"hexadecimal,startswith=0x"`
	EndAmount   string `json:"end_amount" validate:"hexadecimal,startswith=0x"`
}

type ConsiderationItemReq struct {
	ItemType entities.EnumItemType `json:"item_type" validate:"gte=0"`

	Token       string `json:"token" validate:"eth_addr"`
	Identifier  string `json:"identifier" validate:"hexadecimal,startswith=0x"`
	StartAmount string `json:"start_amount" validate:"hexadecimal,startswith=0x"`
	EndAmount   string `json:"end_amount" validate:"hexadecimal,startswith=0x"`
	Recipient   string `json:"recipient" validate:"eth_addr"`
}
