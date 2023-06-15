package services

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/phhphc/nft-marketplace-back-end/internal/services/infrastructure"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

func (s *Services) TransferNft(
	ctx context.Context,
	transfer models.NftTransfer,
	blockNumber uint64,
	txIndex uint,
) error {
	_, err := s.nftWriter.UpsertNftLatest(
		ctx,
		transfer.Token,
		transfer.Identifier,
		transfer.To,
		transfer.To == util.ZeroAddress,
		blockNumber,
		txIndex,
	)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("upsert nft fail")
		return err
	}

	value, err := json.Marshal(models.NewErc721Task{
		Token:      transfer.Token,
		Identifier: transfer.Identifier,
	})
	if err != nil {
		s.lg.Panic().Caller().Err(err).Msg("cannot marshal")
		return err
	}
	if (transfer.From == common.Address{}) {
		s.EmitTask(context.TODO(), models.TaskNewErc721, value)
	}

	err = s.RemoveInvalidOrder(ctx, transfer.From, transfer.Token, transfer.Identifier)
	if err != nil {
		s.lg.Fatal().Caller().Err(err).Msg("remove error")
	}
	return err
}

func (s *Services) UpdateNftMetadata(
	ctx context.Context,
	token common.Address,
	identifier *big.Int,
	metadata map[string]any,
) (err error) {
	_, err = s.nftWriter.UpdateNft(
		ctx,
		token,
		identifier,
		infrastructure.UpdateNftNewValue{
			Metadata: metadata,
		},
	)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("update nft metadata fail")
		return
	}
	return
}
