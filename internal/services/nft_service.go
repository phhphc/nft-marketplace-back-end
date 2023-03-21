package services

import (
	"bytes"
	"context"
	"database/sql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"math/big"
)

type NftNewService interface {
	GetListNft(ctx context.Context, token common.Address, owner common.Address, offset int32, limit int32) ([]entities.Nft, error)
	GetNft(ctx context.Context, token common.Address, identifier *big.Int) (entities.Nft, error)
	TransferNft(ctx context.Context, nft entities.Nft, from common.Address, to common.Address, blockNumber *big.Int, txIndex *big.Int) error
}

func (s *Services) GetListNft(ctx context.Context, token common.Address, owner common.Address, offset int32, limit int32) ([]entities.Nft, error) {
	listNft := make([]entities.Nft, 0)

	res, err := s.repo.GetListValidNft(ctx, postgresql.GetListValidNftParams{
		Token:  ToNullString(token),
		Owner:  ToNullString(owner),
		Offset: offset,
		Limit:  limit,
	})

	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error in query nft")
		return listNft, err
	}

	for _, row := range res {
		identifier, _ := big.NewInt(0).SetString(row.Identifier, 10)
		nft := entities.Nft{
			Token:      common.HexToAddress(row.Token),
			Identifier: identifier,
			Owner:      common.HexToAddress(row.Owner),
			TokenUri:   row.TokenUri.String,
			IsBurned:   row.IsBurned,
		}
		listNft = append(listNft, nft)
	}

	return listNft, nil
}

func (s *Services) GetNft(ctx context.Context, token common.Address, identifier *big.Int) (entities.Nft, error) {
	nft := entities.Nft{}

	res, err := s.repo.GetValidNft(ctx, postgresql.GetValidNftParams{
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
	err := s.repo.UpsertNftV2(ctx, postgresql.UpsertNftV2Params{
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
	err := s.repo.UpsertNftV2(ctx, postgresql.UpsertNftV2Params{
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
	err := s.repo.UpsertNftV2(ctx, postgresql.UpsertNftV2Params{
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
