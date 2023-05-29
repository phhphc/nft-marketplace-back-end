package controllers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

func (ctl *Controls) GetNotification(c echo.Context) error {
	var req dto.GetNotificationReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	ns, err := ctl.service.GetListNotification(c.Request().Context(), common.HexToAddress(req.Address))
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("controller cannot get list notification")
		return err
	}

	res := dto.GetNotificationRes{}
	for _, n := range ns {
		notification := dto.NotificationRes{
			IsViewed:  n.IsViewed,
			Info:      n.Info,
			EventName: n.EventName,
			OrderHash: n.OrderHash.Hex(),
			Address:   n.Address.Hex(),
			Token:     n.Token.Hex(),
			TokenId:   n.TokenId.String(),
			Quantity:  n.Quantity,
			Type:      n.Type,
			Price:     n.Price.String(),
			From:      n.From.Hex(),
			To:        n.To.Hex(),
			Date:      n.Date,
			Owner:     n.OrderHash.Hex(),
			NftImage:  n.NftImage,
			NftName:   n.NftName,
		}
		res.Notifications = append(res.Notifications, notification)
	}

	return c.JSON(200, dto.Response{
		Data:      res,
		IsSuccess: true,
	})
}

func (ctl *Controls) UpdateNotification(c echo.Context) error {
	var req dto.UpdateNotificationReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	err = ctl.service.UpdateNotification(c.Request().Context(), entities.NotificationUpdate{
		EventName: req.EventName,
		OrderHash: common.HexToHash(req.OrderHash),
	})

	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("controller cannot update notification")
		return err
	}

	return nil
}
