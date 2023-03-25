package services

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
)

type NftNewService interface {
	GetNFTWithListings(ctx context.Context, token common.Address, identifier *big.Int) (*entities.NftRead, error)
	GetNFTsWithListings(ctx context.Context, token common.Address, owner common.Address, offset int32, limit int32) ([]*entities.NftRead, error)
}

func ToBigInt(str string) *big.Int {
	bigInt := big.NewInt(0)
	bigInt.SetString(str, 10)
	return bigInt
}

func (s *Services) GetNFTsWithListings(ctx context.Context, token common.Address, owner common.Address, offset int32, limit int32) ([]*entities.NftRead, error) {
	tokenValid := true
	ownerValid := true
	if bytes.Equal(token.Bytes(), common.Address{}.Bytes()) {
		tokenValid = false
	}

	if bytes.Equal(owner.Bytes(), common.Address{}.Bytes()) {
		ownerValid = false
	}

	res, err := s.repo.GetNFTsWithPricesPaginated(ctx, postgresql.GetNFTsWithPricesPaginatedParams{
		Offset: offset,
		Limit:  limit,
		Token:  sql.NullString{String: token.Hex(), Valid: tokenValid},
		Owner:  sql.NullString{String: owner.Hex(), Valid: ownerValid},
	})

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
				Image:       fmt.Sprintf("%v", nft.Image),
				Name:        fmt.Sprintf("%v", nft.Name),
				Description: fmt.Sprintf("%v", nft.Description),
				Listings:    make([]*entities.ListingRead, 0),
			}
		}

		if nft.StartPrice.Valid || nft.EndPrice.Valid {
			nftRes := nftsMap[nft.Identifier]
			nftRes.Listings = append(nftRes.Listings, &entities.ListingRead{
				OrderHash:  common.HexToHash(nft.OrderHash.String),
				ItemType:   entities.EnumItemType(nft.ItemType.Int32),
				StartPrice: ToBigInt(nft.StartPrice.String),
				EndPrice:   ToBigInt(nft.EndPrice.String),
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
			nft = &entities.NftRead{
				Token:       common.HexToAddress(order.Token),
				Identifier:  ToBigInt(order.Identifier),
				Owner:       common.HexToAddress(order.Owner),
				Image:       fmt.Sprintf("%v", order.Image),
				Name:        fmt.Sprintf("%v", order.Name),
				Description: fmt.Sprintf("%v", order.Description),
				Metadata:    order.Metadata.RawMessage,
				Listings:    make([]*entities.ListingRead, 0),
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
