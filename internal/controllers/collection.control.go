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

	// var metadata map[string]any
	// if err = json.Unmarshal([]byte(req.Metadata), &metadata); err != nil {
	// 	return dto.NewHTTPError(400, err)
	// }
	collection := entities.Collection{
		Token:       common.HexToAddress(req.Token),
		Owner:       common.HexToAddress(req.Owner),
		Name:        req.Name,
		Description: req.Description,
		Metadata:    req.Metadata,
		Category:    req.Category,
	}
	collection, err = ctl.service.CreateCollection(context.TODO(), collection)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("cannot create collection")
		return err
	}

	return c.JSON(200, dto.Response{
		Data: dto.PostCollectionRes{
			Token:       collection.Token.Hex(),
			Owner:       collection.Owner.Hex(),
			Name:        collection.Name,
			Description: collection.Description,
			Metadata:    collection.Metadata,
			Category:    collection.Category,
			CreatedAt:   collection.CreatedAt,
		},
		IsSuccess: true,
	})
}

func (ctl *Controls) GetCollection(c echo.Context) error {
	var req dto.GetCollectionReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	query := entities.Collection{
		Token:    common.HexToAddress(req.Token),
		Owner:    common.HexToAddress(req.Owner),
		Name:     req.Name,
		Category: req.Category,
	}
	cs, err := ctl.service.GetListCollection(c.Request().Context(), query, 0, 9999)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("cannot")
		return err
	}

	collection := dto.GetCollectionRes{
		PageSize: 9999,
		Page:     0,
	}
	for _, c := range cs {
		collection.Collections = append(collection.Collections, dto.Collection{
			Token:       c.Token.Hex(),
			Owner:       c.Owner.Hex(),
			Name:        c.Name,
			Description: c.Description,
			Metadata:    c.Metadata,
			Category:    c.Category,
			CreatedAt:   c.CreatedAt,
		})
	}

	return c.JSON(200, dto.Response{
		Data:      collection,
		IsSuccess: true,
	})
}

func (ctl *Controls) GetCollectionWithCategory(c echo.Context) error {
	var req dto.GetCollectionWithCategoryReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	cs, err := ctl.service.GetListCollectionWithCategory(c.Request().Context(), req.Category, 0, 9999)
	if err != nil {
		return dto.NewHTTPError(400, err)
	}

	collection := dto.GetCollectionWithCategoryRes{
		PageSize: 9999,
		Page:     0,
	}
	for _, c := range cs {
		collection.Collections = append(collection.Collections, dto.Collection{
			Token:       c.Token.Hex(),
			Owner:       c.Owner.Hex(),
			Name:        c.Name,
			Description: c.Description,
			Metadata:    c.Metadata,
			Category:    c.Category,
			CreatedAt:   c.CreatedAt,
		})
	}

	return c.JSON(200, dto.Response{
		Data:      collection,
		IsSuccess: true,
	})
}

