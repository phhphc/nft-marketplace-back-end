package controllers

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

func (ctl *Controls) GetEvent(c echo.Context) error {
	var req dto.GetEventReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	query := entities.EventRead{
		Name:  req.Name,
		Token: common.HexToAddress(req.Token),
		// TokenId
		Address: common.HexToAddress(req.Address),
		Type:    req.Type,
		StartDate: 	req.StartDate,
		EndDate:	req.EndDate,	
	}
	tokenId, ok := big.NewInt(0).SetString(req.TokenId, 10)
	if ok {
		query.TokenId = tokenId
	}

	es, err := ctl.service.GetListEvent(c.Request().Context(), query)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("cannot")
		return err
	}

	event := dto.GetEventRes{}
	for _, e := range es {
		newEvent := dto.EventRes{
			Name:     e.Name,
			Token:    e.Token.Hex(),
			TokenId:  e.TokenId.String(),
			Quantity: int(e.Quantity),
			Type:     e.Type,
			// price
			From: e.From.Hex(),
			// to
			Date: e.Date,
			// link
			NftImage: e.NftImage,
			NftName:  e.NftName,
			OrderHash: e.OrderHash.Hex(),
		}

		if e.Name == "listing" || e.Name == "offer" || e.Name == "sale" {
			newEvent.Price = e.Price.String()
		}
		if e.Name == "sale" || e.Name == "transfer" {
			newEvent.To = e.To.Hex()
			newEvent.TxHash = e.TxHash
		}
		if e.Name == "listing" || e.Name == "offer" {
			newEvent.EndTime = e.EndTime.String()
			newEvent.IsCancelled = e.IsCancelled
			newEvent.IsFulfilled = e.IsFulfilled
		}

		event.Events = append(event.Events, newEvent)
	}

	return c.JSON(200, dto.Response{
		Data:      event,
		IsSuccess: true,
	})
}
