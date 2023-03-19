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
		Zone:          common.HexToAddress(req.Zone),
		Offer:         offerItems,
		Consideration: considerationItems,
		OrderType:     &req.OrderType,
		ZoneHash:      common.HexToHash(req.ZoneHash),
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
