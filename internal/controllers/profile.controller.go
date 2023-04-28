package controllers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type ProfileController interface {
	GetProfile(c echo.Context) error
	PostProfile(c echo.Context) error
}

type GetProfileReq struct {
	Address string `param:"address" validation:"eth_addr"`
}

type GetProfileResp struct {
	Address   string         `json:"address,omitempty"`
	Username  string         `json:"username"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	Signature string         `json:"signature,omitempty"`
}

type PostProfileReq struct {
	Address   string         `json:"address" validation:"eth_addr"`
	Username  string         `json:"username"`
	Metadata  map[string]any `json:"metadata"`
	Signature string         `json:"signature" validation:"hexadecimal,startswith=0x"`
}

type PostProfileResp struct {
	Address   string         `json:"address"`
	Username  string         `json:"username,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	Signature string         `json:"signature,omitempty"`
}

func (ctl *Controls) GetProfile(c echo.Context) error {
	var req GetProfileReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	resp, err := ctl.service.GetProfile(c.Request().Context(), req.Address)
	if err != nil {
		return dto.NewHTTPError(400, err)
	}

	profile := GetProfileResp{
		Address:   resp.Address.Hex(),
		Username:  resp.Username,
		Metadata:  resp.Metadata,
		Signature: string(resp.Signature),
	}

	return c.JSON(200, dto.Response{
		Data:      profile,
		IsSuccess: true,
	})
}

func (ctl *Controls) PostProfile(c echo.Context) error {
	var req PostProfileReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	resp, err := ctl.service.UpsertProfile(c.Request().Context(), entities.Profile{
		Address:   common.HexToAddress(req.Address),
		Username:  req.Username,
		Metadata:  req.Metadata,
		Signature: []byte(req.Signature),
	})
	if err != nil {
		return dto.NewHTTPError(400, err)
	}

	profile := PostProfileResp{
		Address:   resp.Address.Hex(),
		Username:  resp.Username,
		Metadata:  resp.Metadata,
		Signature: string(resp.Signature),
	}

	return c.JSON(200, dto.Response{
		Data:      profile,
		IsSuccess: true,
	})
}
