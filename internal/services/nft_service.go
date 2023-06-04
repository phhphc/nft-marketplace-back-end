package services

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
)

type NftNewService interface {
	GetNFTWithListings(ctx context.Context, token common.Address, identifier *big.Int) (*entities.NftRead, error)
	GetNFTsWithListings(ctx context.Context, token common.Address, owner common.Address, isHidden *bool, offset int32, limit int32) ([]*entities.NftRead, error)
	UpdateNftStatus(ctx context.Context, token common.Address, identifier *big.Int, isHidden bool) error
}

func ToBigInt(str string) *big.Int {
	bigInt := big.NewInt(0)
	bigInt.SetString(str, 10)
	return bigInt
}

func (s *Services) GetNFTsWithListings(ctx context.Context, token common.Address, owner common.Address, isHidden *bool, offset int32, limit int32) ([]*entities.NftRead, error) {
	tokenValid := true
	ownerValid := true
	if bytes.Equal(token.Bytes(), common.Address{}.Bytes()) {
		tokenValid = false
	}

	if bytes.Equal(owner.Bytes(), common.Address{}.Bytes()) {
		ownerValid = false
	}

	params := postgresql.GetNFTsWithPricesPaginatedParams{
		Offset: offset,
		Limit:  limit,
		Token:  sql.NullString{String: token.Hex(), Valid: tokenValid},
		Owner:  sql.NullString{String: owner.Hex(), Valid: ownerValid},
	}
	if isHidden != nil {
		params.IsHidden = sql.NullBool{
			Bool:  *isHidden,
			Valid: true,
		}
	}
	res, err := s.repo.GetNFTsWithPricesPaginated(ctx, params)

	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error in query nfts with prices")
		return nil, err
	}

	nftsMap := make(map[string]*entities.NftRead)
	for _, nft := range res {
		if _, ok := nftsMap[nft.Identifier]; !ok {
			nftsMap[nft.Identifier] = &entities.NftRead{
				Token:       common.HexToAddress(nft.Token),
				Identifier:  ToBigInt(nft.Identifier),
				Owner:       common.HexToAddress(nft.Owner),
				Image:       FromInterfaceString2String(nft.Image),
				Name:        FromInterfaceString2String(nft.Name),
				Description: FromInterfaceString2String(nft.Description),
				IsHidden:    nft.IsHidden,
				Listings:    make([]*entities.ListingRead, 0),
			}
		}

		if nft.StartTime.Valid || nft.EndTime.Valid {
			nftRes := nftsMap[nft.Identifier]
			nftRes.Listings = append(nftRes.Listings, &entities.ListingRead{
				OrderHash:  common.HexToHash(nft.OrderHash.String),
				ItemType:   entities.EnumItemType(nft.ItemType),
				StartPrice: new(big.Int).SetInt64(nft.StartPrice),
				EndPrice:   new(big.Int).SetInt64(nft.EndPrice),
				StartTime:  ToBigInt(nft.StartTime.String),
				EndTime:    ToBigInt(nft.EndTime.String),
			})
		}
	}

	nfts := make([]*entities.NftRead, 0)
	for _, nft := range nftsMap {
		nfts = append(nfts, nft)
	}

	return nfts, nil
}

func (s *Services) GetNFTWithListings(ctx context.Context, token common.Address, identifier *big.Int) (*entities.NftRead, error) {
	res, err := s.repo.GetNFTValidConsiderations(ctx, postgresql.GetNFTValidConsiderationsParams{
		Token:      token.Hex(),
		Identifier: identifier.String(),
	})

	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error in query nft")
		return nil, err
	}

	// Lay len thong tin cua nft
	// Lay len danh sach cac order ma nft nay la offer item (order valid)
	var nft *entities.NftRead
	for i, order := range res {
		if i == 0 {
			s.lg.Debug().Caller().Interface("order", order).Msg("order")

			nft = &entities.NftRead{
				Token:      common.HexToAddress(order.Token),
				Identifier: ToBigInt(order.Identifier),
				Owner:      common.HexToAddress(order.Owner),
				Metadata:   order.Metadata.RawMessage,
				Listings:   make([]*entities.ListingRead, 0),
			}
		}

		if order.OrderHash.Valid {
			listing := &entities.ListingRead{
				OrderHash:  common.HexToHash(order.OrderHash.String),
				ItemType:   entities.EnumItemType(order.ItemType.Int32),
				StartPrice: ToBigInt(order.StartPrice.String),
				EndPrice:   ToBigInt(order.EndPrice.String),
				StartTime:  ToBigInt(order.StartTime.String),
				EndTime:    ToBigInt(order.EndTime.String),
			}
			nft.Listings = append(nft.Listings, listing)
		}
	}
	return nft, nil
}

func (s *Services) UpdateNftStatus(ctx context.Context, token common.Address, identifier *big.Int, isHidden bool) error {
	_, err := s.repo.UpdateNftStatus(ctx, postgresql.UpdateNftStatusParams{
		IsHidden: sql.NullBool{
			Bool:  isHidden,
			Valid: true,
		},
		Token:      token.Hex(),
		Identifier: identifier.String(),
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("update error")
		return err
	}

	return nil
}

func FromInterfaceString2String(bstr interface{}) string {
	if bstr == nil {
		return ""
	}
	byteArrayStr := fmt.Sprint(bstr)

	byteArrayStrArr := strings.Split(byteArrayStr, " ")
	subByteArr := byteArrayStrArr[1 : len(byteArrayStrArr)-1]
	var byteArray []byte
	for _, val := range subByteArr {
		b, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		byteArray = append(byteArray, byte(b))
	}

	str := string(byteArray)
	return str
}
