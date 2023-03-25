package controllers

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
	"net/http"
)

func (ctl *Controls) GetNFTsWithPrices(c echo.Context) error {
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

	fmt.Println("token: ", token.String())
	fmt.Println("owner: ", owner.String())

	nfts, err := ctl.service.GetNFTsWithPrices(c.Request().Context(), token, owner, req.Offset, req.Limit)

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

func (ctl *Controls) GetNft(c echo.Context) error {

	return nil
}
