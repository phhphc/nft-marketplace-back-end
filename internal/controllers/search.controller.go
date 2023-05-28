package controllers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
	"net/http"
)

type SearchController interface {
	SearchNFTs(c echo.Context) error
}

func (ctl *Controls) SearchNFTs(c echo.Context) error {
	var req dto.SearchNftReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	token := common.HexToAddress(req.Token)
	owner := common.HexToAddress(req.Owner)

	nfts, err := ctl.service.SearchNFTsWithListings(c.Request().Context(), token, owner, req.Q, req.IsHidden, req.Offset, req.Limit)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("error")
	}

	nftsResponse := make([]*dto.GetNftRes, len(nfts))

	for i, nft := range nfts {
		nftsResponse[i] = &dto.GetNftRes{
			Token:       nft.Token.String(),
			Identifier:  nft.Identifier.String(),
			Owner:       nft.Owner.String(),
			Image:       nft.Image,
			Name:        nft.Name,
			IsHidden:    nft.IsHidden,
			Description: nft.Description,
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
