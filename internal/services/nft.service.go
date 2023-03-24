package services

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
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
	return
}
