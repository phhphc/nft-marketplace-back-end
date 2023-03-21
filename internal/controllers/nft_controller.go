package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
)

func (ctl *Controls) GetListNftOfCollection(c echo.Context) error {
	var req dto.GetListNftReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	//nfts, err := ctl.service.GetListNft(c.Request().Context(), req.Token, req.Owner, req.Offset, req.Limit)

	return nil
}

func (ctl *Controls) GetNft(c echo.Context) error {

	return nil
}
