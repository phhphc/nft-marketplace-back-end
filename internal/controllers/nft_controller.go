package controllers

import (
	"encoding/json"
	"fmt"
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
	owner := common.HexToAddress(req.Owner)

	nfts, err := ctl.service.GetNFTsWithListings(c.Request().Context(), token, owner, req.Offset, req.Limit)

	nftsResponse := make([]*dto.GetNftRes, len(nfts))

	for i, nft := range nfts {
		nftsResponse[i] = &dto.GetNftRes{
			Token:       nft.Token.String(),
			Identifier:  nft.Identifier.String(),
			Owner:       nft.Owner.String(),
			Image:       nft.Image,
			Name:        nft.Name,
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

func (ctl *Controls) GetNFTWithListings(c echo.Context) error {
	var req dto.GetNftReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	token := common.HexToAddress(req.Token)
	identifier, ok := big.NewInt(0).SetString(req.Identifier, 10)
	if !ok {
		return dto.NewHTTPError(400, fmt.Errorf("invalid type in identifier"))
	}

	nft, err := ctl.service.GetNFTWithListings(c.Request().Context(), token, identifier)

	if err != nil {
		return dto.NewHTTPError(400, err)
	}
	if nft == nil {
		return c.JSON(http.StatusOK, dto.Response{
			Data:      []byte(""),
			IsSuccess: true,
		})
	}

	var metadata map[string]any
	err = json.Unmarshal(nft.Metadata, &metadata)
	if err != nil {
		ctl.lg.Panic().Caller().Err(err).Msg("panic")
	}

	nftResponse := &dto.GetNftRes{
		Token:       nft.Token.String(),
		Identifier:  nft.Identifier.String(),
		Owner:       nft.Owner.String(),
		Image:       nft.Image,
		Name:        nft.Name,
		Description: nft.Description,
		Metadata:    metadata,
	}

	nftResponse.Listings = make([]*dto.GetNftListingRes, len(nft.Listings))

	if len(nft.Listings) == 0 {
		return c.JSON(http.StatusOK, dto.Response{
			Data:      nftResponse,
			IsSuccess: true,
		})
	}

	for j, listing := range nft.Listings {
		nftResponse.Listings[j] = &dto.GetNftListingRes{
			OrderHash:  listing.OrderHash.String(),
			ItemType:   listing.ItemType.Int(),
			StartPrice: listing.StartPrice.String(),
			EndPrice:   listing.EndPrice.String(),
			StartTime:  listing.StartTime.String(),
			EndTime:    listing.EndTime.String(),
		}
	}

	return c.JSON(http.StatusOK, dto.Response{
		Data:      nftResponse,
		IsSuccess: true,
	})
}
