package services

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"math/big"
)

type NftNewService interface {
	GetNft(ctx context.Context, token common.Address, identifier *big.Int) (entities.Nft, error)
	TransferNft(ctx context.Context, nft entities.Nft, from common.Address, to common.Address, blockNumber *big.Int, txIndex *big.Int) error
	GetNFTsWithPrices(ctx context.Context, token common.Address, owner common.Address, offset int32, limit int32) ([]*entities.NftRead, error)
}

func ToBigInt(str string) *big.Int {
	bigInt := big.NewInt(0)
	bigInt.SetString(str, 10)
	return bigInt
}

func (s *Services) GetNFTsWithPrices(ctx context.Context, token common.Address, owner common.Address, offset int32, limit int32) ([]*entities.NftRead, error) {
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

		if nft.Price.Valid {
			nftRes := nftsMap[nft.Identifier]
			nftRes.Listings = append(nftRes.Listings, &entities.ListingRead{
				OrderHash: common.HexToHash(nft.OrderHash.String),
				ItemType:  entities.EnumItemType(nft.ItemType.Int32),
				Price:     ToBigInt(nft.Price.String),
			})
		}
	}

	nfts := make([]*entities.NftRead, 0)
	for _, nft := range nftsMap {
		nfts = append(nfts, nft)
	}

	return nfts, nil
}

func (s *Services) GetNft(ctx context.Context, token common.Address, identifier *big.Int) (entities.Nft, error) {
	nft := entities.Nft{}

	res, err := s.repo.GetValidNFT(ctx, postgresql.GetValidNFTParams{
		Token:      token.String(),
		Identifier: identifier.String(),
	})

	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error in query nft")
		return nft, err
	}

	nft = entities.Nft{
		Token:      token,
		Identifier: identifier,
		Owner:      common.HexToAddress(res.Owner),
		TokenUri:   res.TokenUri.String,
		IsBurned:   res.IsBurned,
	}

	return nft, nil
}

func (s *Services) TransferNft(ctx context.Context, nft entities.Nft, from common.Address, to common.Address, blockNumber *big.Int, txIndex *big.Int) error {
	// if from address = 0 -> mint nft
	// if to address = 0 -> burn nft
	// else -> transfer nft
	if bytes.Equal(from.Bytes(), common.Address{}.Bytes()) {
		err := s.mintNft(ctx, nft, blockNumber, txIndex)
		if err != nil {
			s.lg.Error().Caller().Err(err).Msg("error in save mint nft")
			return err
		}

	} else if bytes.Equal(to.Bytes(), common.Address{}.Bytes()) {
		err := s.burnNft(ctx, nft.Token, nft.Identifier, blockNumber, txIndex)
		if err != nil {
			s.lg.Error().Caller().Err(err).Msg("error in burn nft")
			return err
		}
	} else {
		err := s.transferNft(ctx, nft.Token, nft.Identifier, to, blockNumber, txIndex)
		if err != nil {
			s.lg.Error().Caller().Err(err).Msg("error in transfer nft")
			return err
		}
	}
	return nil
}

func (s *Services) burnNft(ctx context.Context, token common.Address, identifier *big.Int, blockNumber *big.Int, txIndex *big.Int) error {
	err := s.repo.UpsertNFTV2(ctx, postgresql.UpsertNFTV2Params{
		Token:       token.String(),
		Identifier:  identifier.String(),
		Owner:       common.Address{}.String(),
		TokenUri:    sql.NullString{Valid: false},
		IsBurned:    true,
		BlockNumber: blockNumber.String(),
		TxIndex:     txIndex.Int64(),
	})
	if err != nil {
		return err
	}

	// TODO - update order status to cancel
	return nil
}

func (s *Services) mintNft(ctx context.Context, nft entities.Nft, blockNumber *big.Int, txIndex *big.Int) error {
	err := s.repo.UpsertNFTV2(ctx, postgresql.UpsertNFTV2Params{
		Token:       nft.Token.String(),
		Identifier:  nft.Identifier.String(),
		Owner:       nft.Owner.String(),
		TokenUri:    sql.NullString{String: nft.TokenUri, Valid: true},
		IsBurned:    false,
		BlockNumber: blockNumber.String(),
		TxIndex:     txIndex.Int64(),
	})

	if err != nil {
		return err
	}

	// TODO - update the nft list of user
	return nil
}

func (s *Services) transferNft(ctx context.Context, token common.Address, identifier *big.Int, owner common.Address, blockNumber *big.Int, txIndex *big.Int) error {
	err := s.repo.UpsertNFTV2(ctx, postgresql.UpsertNFTV2Params{
		Token:       token.String(),
		Identifier:  identifier.String(),
		Owner:       owner.String(),
		BlockNumber: blockNumber.String(),
		TxIndex:     txIndex.Int64(),
	})

	if err != nil {
		return err
	}

	return nil
}
