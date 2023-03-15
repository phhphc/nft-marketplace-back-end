package controllers

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/models"
	"time"
)

type (
	GetOrderResponse struct {
		OrderHash     string                 `json:"order_hash"`
		Offer         []GetOfferItem         `json:"offer"`
		Consideration []GetConsiderationItem `json:"consideration"`
		Offerer       string                 `json:"offerer"`
		Signature     string                 `json:"signature"`
		OrderType     string                 `json:"order_type"`
		StartTime     string                 `json:"start_time"`
		EndTime       string                 `json:"end_time"`
		Counter       string                 `json:"counter"`
		Salt          string                 `json:"salt"`
		Zone          string                 `json:"zone,omitempty"`
		ZoneHash      string                 `json:"zone_hash,omitempty"`
		CreatedAt     string                 `json:"created_at"`
	}

	GetOfferItem struct {
		ItemType     string `json:"item_type"`
		TokenId      string `json:"identifier"`
		TokenAddress string `json:"token"`
		StartAmount  string `json:"start_amount"`
		EndAmount    string `json:"end_amount"`
	}

	GetConsiderationItem struct {
		ItemType     string `json:"item_type"`
		TokenId      string `json:"identifier"`
		TokenAddress string `json:"token"`
		StartAmount  string `json:"start_amount"`
		EndAmount    string `json:"end_amount"`
		Recipient    string `json:"recipient"`
	}
)

func NewOrderResponse(order models.Order) *GetOrderResponse {
	response := &GetOrderResponse{}
	response.MapFromDomainOrder(order)
	return response
}

func (o *GetOfferItem) MapFromDomainOfferItem(offerItem models.OfferItem) {
	o.ItemType = hexutil.EncodeBig(offerItem.ItemType)
	o.TokenId = hexutil.EncodeBig(offerItem.TokenId)
	o.TokenAddress = offerItem.TokenAddress.Hex()
	o.StartAmount = hexutil.EncodeBig(offerItem.StartAmount)
	o.EndAmount = hexutil.EncodeBig(offerItem.EndAmount)
}

func (c *GetConsiderationItem) MapFromDomainConsiderationItem(considerationItem models.ConsiderationItem) {
	c.ItemType = hexutil.EncodeBig(considerationItem.ItemType)
	c.TokenId = hexutil.EncodeBig(considerationItem.TokenId)
	c.TokenAddress = considerationItem.TokenAddress.Hex()
	c.StartAmount = hexutil.EncodeBig(considerationItem.StartAmount)
	c.EndAmount = hexutil.EncodeBig(considerationItem.EndAmount)
	c.Recipient = considerationItem.Recipient.Hex()
}

func (c *GetOrderResponse) MapFromDomainOrder(order models.Order) {
	c.OrderHash = order.OrderHash
	c.Offerer = order.Offerer.Hex()
	c.Signature = order.Signature
	c.OrderType = hexutil.EncodeBig(order.OrderType)
	c.StartTime = hexutil.EncodeBig(order.StartTime)
	c.EndTime = hexutil.EncodeBig(order.EndTime)
	c.Counter = hexutil.EncodeBig(order.Counter)
	c.Salt = order.Salt
	c.Zone = order.Zone.Hex()
	c.ZoneHash = order.ZoneHash
	c.CreatedAt = order.CreatedAt.Format(time.RFC3339)

	c.Offer = make([]GetOfferItem, len(order.Offer))
	for i, offerItem := range order.Offer {
		c.Offer[i].MapFromDomainOfferItem(offerItem)
	}

	c.Consideration = make([]GetConsiderationItem, len(order.Consideration))
	for i, considerationItem := range order.Consideration {
		c.Consideration[i].MapFromDomainConsiderationItem(considerationItem)
	}
}
