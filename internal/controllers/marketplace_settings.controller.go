package controllers

import (
	"encoding/base64"
	"encoding/json"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
)

type MarketplaceSettingsController interface {
	GetMarketplaceSettings(c echo.Context) error
	CreateMarketplaceSettings(c echo.Context) error
}

type GetMarketplaceSettingsReq struct {
	Marketplace string `param:"marketplace" validate:"eth_addr"`
}

type GetMarketplaceSettingsResp struct {
	Id          int64  `json:"id"`
	Marketplace string `json:"marketplace"`
	Admin       string `json:"admin_address"`
	Signer      string `json:"signer_address"`
	Royalty     string `json:"royalty"`
	Signature   string `json:"signature"`
	CreatedAt   string `json:"created_at"`
}

type CreateMarketplaceSettingsReq struct {
	TypedData string `json:"typed_data" validate:"required"`
	Signature string `json:"signature" validate:"hexadecimal,startswith=0x"`
}

type CreateMarketplaceSettingsRes struct {
	Id          int64  `json:"id"`
	Marketplace string `json:"marketplace"`
	Admin       string `json:"admin_address"`
	Signer      string `json:"signer_address"`
	Royalty     string `json:"royalty"`
	Signature   string `json:"signature"`
	CreatedAt   string `json:"created_at"`
}

func (ctl *Controls) GetMarketplaceSettings(c echo.Context) error {
	var (
		req GetMarketplaceSettingsReq
		err error
	)
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	resp, err := ctl.service.GetValidMarketplaceSettings(c.Request().Context(), common.HexToAddress(req.Marketplace))
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("controller cannot get marketplace settings")
		return dto.NewHTTPError(400, err)
	}

	res := GetMarketplaceSettingsResp{
		Id:          resp.Id,
		Marketplace: resp.Marketplace.Hex(),
		Admin:       resp.Admin.Hex(),
		Signer:      resp.Signer.Hex(),
		Royalty:     strconv.FormatFloat(resp.Royalty, 'f', -1, 64),
		Signature:   string(resp.Signature),
		CreatedAt:   strconv.FormatInt(resp.CreatedAt.Int64(), 10),
	}

	return c.JSON(200, dto.Response{
		Data:      res,
		IsSuccess: true,
	})
}

func (ctl *Controls) CreateMarketplaceSettings(c echo.Context) error {
	var (
		req CreateMarketplaceSettingsReq
		err error
	)
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	//fmt.Printf("req: %+v\n", req)
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	signature, err := hexutil.Decode(req.Signature)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("controller cannot decode signature")
		return dto.NewHTTPError(400, err)
	}

	//fmt.Println("SIG:", hexutil.Encode(signature))

	typedDataBytes, err := base64.StdEncoding.DecodeString(req.TypedData)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("controller cannot decode typed data")
		return dto.NewHTTPError(400, err)
	}

	typedData := apitypes.TypedData{}
	if err = json.Unmarshal(typedDataBytes, &typedData); err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("controller cannot unmarshal typed data: ")
		return dto.NewHTTPError(400, err)
	}

	settings, err := ctl.service.CreateMarketplaceSettings(c.Request().Context(), typedData, signature)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("controller cannot create marketplace settings: ")
		return dto.NewHTTPError(400, err)
	}

	return c.JSON(200, dto.Response{
		Data: CreateMarketplaceSettingsRes{
			Id:          settings.Id,
			Marketplace: settings.Marketplace.Hex(),
			Admin:       settings.Admin.Hex(),
			Signer:      settings.Signer.Hex(),
			Royalty:     strconv.FormatFloat(settings.Royalty, 'f', -1, 64),
			Signature:   string(settings.Signature),
			CreatedAt:   strconv.FormatInt(settings.CreatedAt.Int64(), 10),
		},
		IsSuccess: true,
	})
}
