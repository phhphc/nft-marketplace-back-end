package controllers

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/services"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

type NftController interface {
	GetNftsOfCollection(c echo.Context) error
}

type nftController struct {
	lg           *log.Logger
	tokenService services.NftService
}

func NewNftController(db *sql.DB) NftController {
	return &nftController{
		lg:           log.GetLogger(),
		tokenService: services.NewNftService(db),
	}
}

func (ctl *nftController) GetNftsOfCollection(c echo.Context) error {
	nfts, err := ctl.tokenService.GetNftsByCollection(c.Request().Context())
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("err")
		return err
	}

	return c.JSON(200, nfts)
}
