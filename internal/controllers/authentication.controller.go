package controllers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
)

type AuthenticationController interface {
	GetUserNonce(c echo.Context) error
	Login(c echo.Context) error
	Test(c echo.Context) error
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
	Nonce           string `json:"nonce"`
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
			Nonce:           req.Message,
			AuthToken:       token,
		},
		IsSuccess: true,
	})
}

func (ctl *Controls) Test(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	address := claims["address"].(string)
	return c.JSON(200, dto.Response{
		Data: AuthenticateRes{
			IsAuthenticated: true,
			Address:         address,
			Nonce:           claims["nonce"].(string),
			AuthToken:       user.Raw,
		},
		IsSuccess: true,
	})
}
