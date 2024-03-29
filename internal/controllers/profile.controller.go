package controllers

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type ProfileController interface {
	GetProfile(c echo.Context) error
	PostProfile(c echo.Context) error
	GetOffer(c echo.Context) error
}

type GetProfileReq struct {
	Address string `param:"address" validation:"eth_addr"`
}

type GetProfileResp struct {
	Address   string         `json:"address,omitempty"`
	Username  string         `json:"username"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	Signature string         `json:"signature,omitempty"`
}

type PostProfileReq struct {
	Address   string         `json:"address" validation:"eth_addr"`
	Username  string         `json:"username"`
	Metadata  map[string]any `json:"metadata"`
	Signature string         `json:"signature" validation:"hexadecimal,startswith=0x"`
}

type PostProfileResp struct {
	Address   string         `json:"address"`
	Username  string         `json:"username,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	Signature string         `json:"signature,omitempty"`
}

type GetOffer struct {
	Owner string `query:"owner" validate:"omitempty,eth_addr"`
	From  string `query:"from" validate:"omitempty,eth_addr"`
}

type OfferRes struct {
	Name        string    `json:"name"`
	Token       string    `json:"token"`
	TokenId     string    `json:"token_id"`
	Quantity    int       `json:"quantity,omitempty"`
	Type        string    `json:"type"`
	Price       string    `json:"price,omitempty"`
	From        string    `json:"from"`
	To          string    `json:"to,omitempty"`
	Link        string    `json:"link,omitempty"`
	OrderHash   string    `json:"order_hash,omitempty"`
	NftImage    string    `json:"nft_image"`
	NftName     string    `json:"nft_name"`
	EndTime     string    `json:"end_time,omitempty"`
	IsCancelled bool      `json:"is_cancelled"`
	IsFulfilled bool      `json:"is_fulfilled"`
	IsExpired 	bool      `json:"is_expired"`
	Owner		string	  `json:"owner,omitempty"`
}

type GetOfferRes struct {
	OfferList []OfferRes `json:"offer_list"`
}

func (ctl *Controls) GetProfile(c echo.Context) error {
	var req GetProfileReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	resp, err := ctl.service.GetProfile(c.Request().Context(), req.Address)
	if err != nil {
		if err.Error() == "not found" {
			return dto.NewHTTPError(404, err)
		}
	}

	profile := GetProfileResp{
		Address:   resp.Address.Hex(),
		Username:  resp.Username,
		Metadata:  resp.Metadata,
		Signature: string(resp.Signature),
	}

	return c.JSON(200, dto.Response{
		Data:      profile,
		IsSuccess: true,
	})
}

func (ctl *Controls) PostProfile(c echo.Context) error {
	var req PostProfileReq
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	resp, err := ctl.service.UpsertProfile(context.TODO(), entities.Profile{
		Address:   common.HexToAddress(req.Address),
		Username:  req.Username,
		Metadata:  req.Metadata,
		Signature: []byte(req.Signature),
	})
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("cannot create collection")
		return err
	}

	profile := PostProfileResp{
		Address:   resp.Address.Hex(),
		Username:  resp.Username,
		Metadata:  resp.Metadata,
		Signature: string(resp.Signature),
	}

	return c.JSON(200, dto.Response{
		Data:      profile,
		IsSuccess: true,
	})
}

func (ctl *Controls) GetOffer(c echo.Context) error {

	var req GetOffer
	var err error
	if err = c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err = c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	offerList, err := ctl.service.GetOffer(c.Request().Context(),
		common.HexToAddress(req.Owner),
		common.HexToAddress(req.From))
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("err")
		return err
	}

	res := GetOfferRes{}
	for _, offer := range offerList {
		newOffer := OfferRes{
			Name:      offer.Name,
			Token:     offer.Token.Hex(),
			TokenId:   offer.TokenId.String(),
			Quantity:  int(offer.Quantity),
			NftImage:  offer.NftImage,
			NftName:   offer.NftName,
			Type:      offer.Type,
			OrderHash: offer.OrderHash.String(),
			Price:     offer.Price.String(),
			Owner:     offer.Owner.Hex(),
			From:      offer.From.Hex(),
			EndTime:   offer.EndTime.String(),
			IsFulfilled: offer.IsFulfilled,
			IsCancelled: offer.IsCancelled,
			IsExpired: 	 offer.IsExpired,
		}
		res.OfferList = append(res.OfferList, newOffer)
	}

	return c.JSON(200, dto.Response{
		Data:      res,
		IsSuccess: true,
	})
}