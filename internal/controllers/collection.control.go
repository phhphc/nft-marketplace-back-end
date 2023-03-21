package controllers

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

func (ctl *Controls) PostCollection(c echo.Context) error {
	var req dto.PostCollectionReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	collection := entities.Collection{
		Token:       common.HexToAddress(req.Token),
		Owner:       common.HexToAddress(req.Owner),
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
	}
	collection, err = ctl.service.CreateCollection(context.TODO(), collection)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("cannot create collection")
		return err
	}

	return c.JSON(200, dto.PostCollectionRes{
		Token:       collection.Token.Hex(),
		Owner:       collection.Owner.Hex(),
		Name:        collection.Name,
		Description: collection.Description,
		Category:    collection.Category,
		CreatedAt:   collection.CreatedAt,
	})
}
