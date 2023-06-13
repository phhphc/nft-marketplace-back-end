package controllers

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
)

func (ctl *Controls) GetNFTsWithListings(c echo.Context) error {
	var req dto.GetListNftReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	token := common.HexToAddress(req.Token)
	identifier, ok := new(big.Int).SetString(req.Identifier, 0)
	if !ok {
		identifier = nil
	}
	owner := common.HexToAddress(req.Owner)

	nfts, err := ctl.service.ListNftsWithListings(
		c.Request().Context(),
		token,
		identifier,
		owner,
		req.IsHidden,
		req.Offset,
		req.Limit,
	)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("error")
	}

	nftsResponse := make([]*dto.GetNftRes, len(nfts))

	for i, nft := range nfts {
		nftsResponse[i] = &dto.GetNftRes{
			Token:       nft.Token.String(),
			Identifier:  nft.Identifier.String(),
			Owner:       nft.Owner.String(),
			Metadata:    nft.Metadata,
			Image:       nft.Image,
			Name:        nft.Name,
			Description: nft.Description,
			IsHidden:    nft.IsHidden,
			Listings:    make([]*dto.GetNftListingRes, len(nft.Listings)),
		}
		for j, listing := range nft.Listings {
			nftsResponse[i].Listings[j] = &dto.GetNftListingRes{
				OrderHash:  listing.OrderHash.String(),
				ItemType:   listing.ItemType.Int(),
				StartPrice: listing.StartPrice.String(),
				EndPrice:   listing.EndPrice.String(),
				StartTime:  listing.StartTime.String(),
				EndTime:    listing.EndTime.String(),
			}
		}
	}

	return c.JSON(http.StatusOK, dto.Response{
		Data: dto.GetNftsRes{
			Nfts:   nftsResponse,
			Limit:  req.Limit,
			Offset: req.Offset,
		},
		IsSuccess: true,
	})
}

func (ctl *Controls) UpdateNftStatus(c echo.Context) error {
	var req dto.UpdateNftStatusReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	identifier, ok := new(big.Int).SetString(req.Identifier, 0)
	if !ok {
		ctl.lg.Error().Caller().Msg("err")
		return ErrParseBigInt
	}

	err = ctl.service.UpdateNftStatus(c.Request().Context(), common.HexToAddress(req.Token), identifier, req.IsHidden)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("err")
		return err
	}

	return c.JSON(http.StatusOK, dto.Response{
		Data:      dto.UpdateNftStatusRes(req),
		IsSuccess: true,
	})
}
