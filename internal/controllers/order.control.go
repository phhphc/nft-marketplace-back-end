package controllers

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

var ErrParseBigInt error = errors.New("parse error")

func HexToBigInt(s string) (*big.Int, bool) {
	return new(big.Int).SetString(s, 0)
}

func (ctl *Controls) PostOrder(c echo.Context) error {
	var req dto.PostOrderReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	offerItems := make([]entities.OfferItem, len(req.Offer))
	for i, offer := range req.Offer {
		identifier, ok := HexToBigInt(offer.Identifier)
		if !ok {
			ctl.lg.Error().Caller().Msg(offer.Identifier)
			return ErrParseBigInt
		}
		startAmount, ok := HexToBigInt(offer.StartAmount)
		if !ok {
			ctl.lg.Error().Caller().Msg("err")
			return ErrParseBigInt
		}
		endAmount, ok := HexToBigInt(offer.EndAmount)
		if !ok {
			ctl.lg.Error().Caller().Msg("err")
			return ErrParseBigInt
		}

		offerItems[i] = entities.OfferItem{
			ItemType:    offer.ItemType,
			Token:       common.HexToAddress(offer.Token),
			Identifier:  identifier,
			StartAmount: startAmount,
			EndAmount:   endAmount,
		}
	}

	considerationItems := make([]entities.ConsiderationItem, len(req.Consideration))
	for i, consideration := range req.Consideration {
		identifier, ok := HexToBigInt(consideration.Identifier)
		if !ok {
			ctl.lg.Error().Caller().Msg("err")
			return ErrParseBigInt
		}
		startAmount, ok := HexToBigInt(consideration.StartAmount)
		if !ok {
			ctl.lg.Error().Caller().Msg("err")
			return ErrParseBigInt
		}
		endAmount, ok := HexToBigInt(consideration.EndAmount)
		if !ok {
			ctl.lg.Error().Caller().Msg("err")
			return ErrParseBigInt
		}

		considerationItems[i] = entities.ConsiderationItem{
			ItemType:    consideration.ItemType,
			Token:       common.HexToAddress(consideration.Token),
			Identifier:  identifier,
			StartAmount: startAmount,
			EndAmount:   endAmount,
			Recipient:   common.HexToAddress(consideration.Recipient),
		}
	}

	startTime, ok := HexToBigInt(req.StartTime)
	if !ok {
		ctl.lg.Error().Caller().Msg("err")
		return ErrParseBigInt
	}
	endTime, ok := HexToBigInt(req.EndTime)
	if !ok {
		ctl.lg.Error().Caller().Msg("err")
		return ErrParseBigInt
	}
	signature, err := hex.DecodeString(strings.TrimPrefix(req.Signature, "0x"))
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("err")
		return err
	}

	salt := common.HexToHash(req.Salt)

	order := entities.Order{
		OrderHash:     common.HexToHash(req.OrderHash),
		Offerer:       common.HexToAddress(req.Offerer),
		Offer:         offerItems,
		Consideration: considerationItems,
		Salt:          &salt,
		StartTime:     startTime,
		EndTime:       endTime,
		Signature:     signature,
	}

	err = ctl.service.CreateOrder(context.TODO(), order)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("err")
		return err
	}

	return c.JSON(200, dto.Response{
		Data: dto.PostOrderRes{
			OrderHash: req.OrderHash,
		},
		IsSuccess: true,
	})
}

func (ctl *Controls) GetOrder(c echo.Context) error {

	var req dto.GetOrder
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	offer := entities.OfferItem{}
	if len(req.OfferToken) > 0 {
		offer.Token = common.HexToAddress(req.OfferToken)
	}
	if len(req.OfferIdentifier) > 0 {
		identifier, ok := new(big.Int).SetString(req.OfferIdentifier, 0)
		if ok {
			offer.Identifier = identifier
		}
	}

	consideration := entities.ConsiderationItem{}
	if len(req.ConsiderationToken) > 0 {
		consideration.Token = common.HexToAddress(req.ConsiderationToken)
	}
	if len(req.ConsiderationIdentifier) > 0 {
		identifier, ok := new(big.Int).SetString(req.ConsiderationIdentifier, 0)
		if ok {
			consideration.Identifier = identifier
		}
	}

	orderHash := common.HexToHash(req.OrderHash)

	os, err := ctl.service.GetOrder(
		c.Request().Context(),
		offer,
		consideration,
		orderHash,
		common.HexToAddress(req.Offerer),
		req.IsFulfilled,
		req.IsCancelled,
		req.IsInvalid,
	)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("err")
		return err
	}

	res := make([]dto.GetOrderStruct, len(os))
	for i, o := range os {
		oOffer := o.Offer
		offer := make([]dto.OrderOffer, len(oOffer))
		for i, oi := range oOffer {
			offer[i] = dto.OrderOffer{
				ItemType:    oi.ItemType.Int(),
				Token:       oi.Token.Hex(),
				Identifier:  oi.Identifier.String(),
				StartAmount: oi.StartAmount.String(),
				EndAmount:   oi.EndAmount.String(),
			}
		}

		oConsideration := o.Consideration
		consideration := make([]dto.OrderConsideration, len(oConsideration))
		for i, ci := range oConsideration {
			consideration[i] = dto.OrderConsideration{
				ItemType:    ci.ItemType.Int(),
				Token:       ci.Token.Hex(),
				Identifier:  ci.Identifier.String(),
				StartAmount: ci.StartAmount.String(),
				EndAmount:   ci.EndAmount.String(),
				Recipient:   ci.Recipient.Hex(),
			}
		}

		oStatus := o.Status
		res[i] = dto.GetOrderStruct{
			OrderHash: o.OrderHash.Hex(),
			Offerer:   o.Offerer.Hex(),
			StartTime: o.StartTime.String(),
			EndTime:   o.EndTime.String(),
			Salt:      o.Salt.Hex(),
			Signature: util.BytesToHex(o.Signature),
			Status: dto.OrderStatus{
				IsCancelled: oStatus.IsCancelled,
				IsFulfilled: oStatus.IsFulfilled,
				IsInvalid:   oStatus.IsValidated,
			},
			Offer:         offer,
			Consideration: consideration,
		}
	}

	return c.JSON(200, dto.Response{
		Data: dto.PagedRespond[dto.GetOrderStruct]{
			PageSize:    99999,
			CurrentPage: 0,
			Content:     res,
		},
		IsSuccess: true,
	})
}
