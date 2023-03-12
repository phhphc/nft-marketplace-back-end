package controllers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/models"
	"math/big"
)

type (
	OrderForm struct {
		OrderHash     string                  `json:"order_hash"`
		Offer         []OfferItemForm         `json:"offer"`
		Consideration []ConsiderationItemForm `json:"consideration"`
		Offerer       string                  `json:"offerer"`
		Signature     string                  `json:"signature,omitempty"`
		OrderType     string                  `json:"order_type"`
		StartTime     string                  `json:"start_time"`
		EndTime       string                  `json:"end_time"`
		Counter       string                  `json:"counter"`
		Salt          string                  `json:"salt"`
		Zone          string                  `json:"zone,omitempty"`
		ZoneHash      string                  `json:"zone_hash,omitempty"`
	}

	OfferItemForm struct {
		TypeNumber   string `json:"type_number"`
		TokenId      string `json:"token_id"`
		TokenAddress string `json:"token_address"`
		StartAmount  string `json:"start_amount"`
		EndAmount    string `json:"end_amount"`
	}

	ConsiderationItemForm struct {
		TypeNumber   string `json:"type_number"`
		TokenId      string `json:"token_id"`
		TokenAddress string `json:"token_address"`
		StartAmount  string `json:"start_amount"`
		EndAmount    string `json:"end_amount"`
		Recipient    string `json:"recipient"`
	}
)

func (o *OfferItemForm) MapToDomainOfferItem() models.OfferItem {
	typeNumber, _ := big.NewInt(0).SetString(o.TypeNumber, 10)
	tokenId, _ := hexutil.DecodeBig(o.TokenId)
	tokenAddress := common.HexToAddress(o.TokenAddress)
	startAmount, _ := hexutil.DecodeBig(o.StartAmount)
	endAmount, _ := hexutil.DecodeBig(o.EndAmount)

	return models.OfferItem{
		TypeNumber:   typeNumber,
		TokenId:      tokenId,
		TokenAddress: tokenAddress,
		StartAmount:  startAmount,
		EndAmount:    endAmount,
	}
}

func (c *ConsiderationItemForm) MapToDomainConsiderationItem() models.ConsiderationItem {
	typeNumber, _ := big.NewInt(0).SetString(c.TypeNumber, 10)
	tokenId, _ := hexutil.DecodeBig(c.TokenId)
	tokenAddress := common.HexToAddress(c.TokenAddress)
	startAmount, _ := hexutil.DecodeBig(c.StartAmount)
	endAmount, _ := hexutil.DecodeBig(c.EndAmount)
	recipient := common.HexToAddress(c.Recipient)

	return models.ConsiderationItem{
		TypeNumber:   typeNumber,
		TokenId:      tokenId,
		TokenAddress: tokenAddress,
		StartAmount:  startAmount,
		EndAmount:    endAmount,
		Recipient:    recipient,
	}
}

func (o *OrderForm) MapToDomainOrder() models.Order {
	offer := make([]models.OfferItem, len(o.Offer))
	consideration := make([]models.ConsiderationItem, len(o.Consideration))
	for i, v := range o.Offer {
		offer[i] = v.MapToDomainOfferItem()
	}
	for i, v := range o.Consideration {
		consideration[i] = v.MapToDomainConsiderationItem()
	}

	offerer := common.HexToAddress(o.Offerer)
	orderType, _ := big.NewInt(0).SetString(o.OrderType, 10)
	startTime, _ := hexutil.DecodeBig(o.StartTime)
	endTime, _ := hexutil.DecodeBig(o.EndTime)
	counter, _ := hexutil.DecodeBig(o.Counter)
	zone := common.HexToAddress(o.Zone)

	return models.Order{
		OrderHash:     o.OrderHash,
		Offer:         offer,
		Consideration: consideration,
		Offerer:       offerer,
		Signature:     o.Signature,
		OrderType:     orderType,
		StartTime:     startTime,
		EndTime:       endTime,
		Counter:       counter,
		Salt:          o.Salt,
		Zone:          zone,
		ZoneHash:      o.ZoneHash,
	}
}
