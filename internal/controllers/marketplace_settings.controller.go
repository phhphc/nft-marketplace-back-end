package controllers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
)

type MarketplaceSettingsController interface {
	GetMarketplaceSettings(c echo.Context) error
	UpdateMarketplaceSettings(c echo.Context) error
}

type GetMarketplaceSettingsReq struct {
	Marketplace string `query:"marketplace" validate:"eth_addr"`
}

type GetMarketplaceSettingsResp struct {
	Id          int64   `json:"id"`
	Marketplace string  `json:"marketplace"`
	Beneficiary string  `json:"beneficiary"`
	Royalty     float64 `json:"royalty"`
}

type UpdateMarketplaceSettingsReq struct {
	Marketplace string  `json:"marketplace" validate:"eth_addr"`
	Beneficiary string  `json:"beneficiary" validate:"eth_addr"`
	Royalty     float64 `json:"royalty" validate:"required,gte=0,lte=1"`
}

type UpdateMarketplaceSettingsRes struct {
	Id          int64   `json:"id"`
	Marketplace string  `json:"marketplace"`
	Beneficiary string  `json:"beneficiary"`
	Royalty     float64 `json:"royalty"`
}

func (ctl *Controls) GetMarketplaceSettings(c echo.Context) error {
	var req GetMarketplaceSettingsReq

	if err := c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	if err := c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	resp, err := ctl.service.GetMarketplaceSettings(c.Request().Context(), common.HexToAddress(req.Marketplace))
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("controller cannot get marketplace settings")
		return dto.NewHTTPError(400, err)
	}

	res := GetMarketplaceSettingsResp{
		Id:          resp.Id,
		Marketplace: resp.Marketplace.Hex(),
		Beneficiary: resp.Beneficiary.Hex(),
		Royalty:     resp.Royalty,
	}

	return c.JSON(200, dto.Response{
		Data:      res,
		IsSuccess: true,
	})
}

func (ctl *Controls) UpdateMarketplaceSettings(c echo.Context) error {
	var req UpdateMarketplaceSettingsReq

	if err := c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err := c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	settings, err := ctl.service.UpdateMarketplaceSettings(c.Request().Context(), common.HexToAddress(req.Marketplace), common.HexToAddress(req.Beneficiary), req.Royalty)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("controller cannot update marketplace settings")
		return dto.NewHTTPError(400, err)
	}
	return c.JSON(200, dto.Response{
		Data: UpdateMarketplaceSettingsRes{
			Id:          settings.Id,
			Marketplace: settings.Marketplace.Hex(),
			Beneficiary: settings.Beneficiary.Hex(),
			Royalty:     settings.Royalty,
		},
		IsSuccess: true,
	})
}

//func (ctl *Controls) UpdateMarketplaceSettings(c echo.Context) error {
//	var (
//		req CreateMarketplaceSettingsReq
//		err error
//	)
//	if err = c.Bind(&req); err != nil {
//		return dto.NewHTTPError(400, err)
//	}
//	//fmt.Printf("req: %+v\n", req)
//	if err = c.Validate(&req); err != nil {
//		return dto.NewHTTPError(400, err)
//	}
//
//	signature, err := hexutil.Decode(req.Signature)
//	if err != nil {
//		ctl.lg.Error().Caller().Err(err).Msg("controller cannot decode signature")
//		return dto.NewHTTPError(400, err)
//	}
//
//	//fmt.Println("SIG:", hexutil.Encode(signature))
//
//	typedDataBytes, err := base64.StdEncoding.DecodeString(req.TypedData)
//	if err != nil {
//		ctl.lg.Error().Caller().Err(err).Msg("controller cannot decode typed data")
//		return dto.NewHTTPError(400, err)
//	}
//
//	typedData := apitypes.TypedData{}
//	if err = json.Unmarshal(typedDataBytes, &typedData); err != nil {
//		ctl.lg.Error().Caller().Err(err).Msg("controller cannot unmarshal typed data: ")
//		return dto.NewHTTPError(400, err)
//	}
//
//	settings, err := ctl.service.CreateMarketplaceSettings(c.Request().Context(), typedData, signature)
//	if err != nil {
//		ctl.lg.Error().Caller().Err(err).Msg("controller cannot create marketplace settings: ")
//		return dto.NewHTTPError(400, err)
//	}
//
//	return c.JSON(200, dto.Response{
//		Data: CreateMarketplaceSettingsRes{
//			Id:          settings.Id,
//			Marketplace: settings.Marketplace.Hex(),
//			Admin:       settings.Admin.Hex(),
//			Signer:      settings.Signer.Hex(),
//			Royalty:     strconv.FormatFloat(settings.Royalty, 'f', -1, 64),
//			Signature:   string(settings.Signature),
//			CreatedAt:   strconv.FormatInt(settings.CreatedAt.Int64(), 10),
//		},
//		IsSuccess: true,
//	})
//}
