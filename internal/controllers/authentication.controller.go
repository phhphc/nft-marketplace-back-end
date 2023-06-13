package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
)

type AuthenticationController interface {
	GetUserNonce(c echo.Context) error
	Login(c echo.Context) error
}

type GetUserNonceReq struct {
	Address string `param:"address" validate:"required,eth_addr"`
}

type GetUserNonceRes struct {
	Nonce   string `json:"nonce"`
	Address string `json:"address"`
}

type AuthenticateReq struct {
	Address   string `json:"address" validate:"eth_addr"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type AuthenticateRes struct {
	IsAuthenticated bool   `json:"is_authenticated"`
	Address         string `json:"address"`
	AuthToken       string `json:"auth_token"`
}

func (ctl *Controls) GetUserNonce(c echo.Context) error {
	var (
		req GetUserNonceReq
		err error
	)

	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	nonce, err := ctl.service.GetUserNonce(c.Request().Context(), req.Address)
	if err != nil {
		return dto.NewHTTPError(400, err)
	}

	return c.JSON(200, dto.Response{
		Data: GetUserNonceRes{
			Nonce:   nonce,
			Address: req.Address,
		},
		IsSuccess: true,
	})
}

// Authenticate is a function that returns a JSON response with the following structure:
func (ctl *Controls) Login(c echo.Context) error {
	var (
		req AuthenticateReq
		err error
	)

	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	token, err := ctl.service.Login(c.Request().Context(), req.Address, req.Message, req.Signature)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("cannot authenticate")
		return dto.NewHTTPError(400, err)
	}

	return c.JSON(200, dto.Response{
		Data: AuthenticateRes{
			IsAuthenticated: true,
			Address:         req.Address,
			AuthToken:       token,
		},
		IsSuccess: true,
	})
}
