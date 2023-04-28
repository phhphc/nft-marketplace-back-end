package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/tabbed/pqtype"
)

func (s *Services) TransferNft(ctx context.Context, transfer models.NftTransfer, blockNumber uint64, txIndex uint) (err error) {
	params := postgresql.UpdateNftParams{
		Token:       transfer.Token.Hex(),
		Identifier:  transfer.Identifier.String(),
		Owner:       transfer.To.Hex(),
		IsBurned:    transfer.To == common.Address{},
		BlockNumber: fmt.Sprintf("%v", blockNumber),
		TxIndex:     int64(txIndex),
	}
	err = s.repo.UpdateNft(ctx, params)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("update nft fail")
		return
	}

	value, err := json.Marshal(models.NewErc721Task{
		Token:      transfer.Token,
		Identifier: transfer.Identifier,
	})
	if err != nil {
		s.lg.Panic().Caller().Err(err).Msg("cannot marshal")
		return
	}
	if (transfer.From == common.Address{}) {
		s.EmitTask(context.TODO(), models.TaskNewErc721, value)
	}

	err = s.RemoveInvalidOrder(ctx, transfer.From, transfer.Token, transfer.Identifier)
	if err != nil {
		s.lg.Fatal().Caller().Err(err).Msg("remove error")
	}
	return
}

func (s *Services) UpdateNftMetadata(ctx context.Context, token common.Address, identifier *big.Int, metadata json.RawMessage) (err error) {
	err = s.repo.UpdateNftMetadata(ctx, postgresql.UpdateNftMetadataParams{
		Metadata: pqtype.NullRawMessage{
			RawMessage: metadata,
			Valid:      len(metadata) > 0,
		},
		Token:      token.Hex(),
		Identifier: identifier.String(),
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("update nft metadata fail")
		return
	}
	return
}
