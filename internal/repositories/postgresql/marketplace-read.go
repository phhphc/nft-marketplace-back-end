package postgresql

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

func (r *PostgresqlRepository) GetMarketplaceLastSyncBlock(
	ctx context.Context,
) (uint64, error) {
	lastSyncBlock, err := r.queries.GetMarketplaceLastSyncBlock(ctx)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error get last block")
	}
	return uint64(lastSyncBlock), err
}

func (r *PostgresqlRepository) GetMarketplaceSettings(
	ctx context.Context,
	marketplaceAddress common.Address,
) (*entities.MarketplaceSettings, error) {
	res, err := r.queries.GetMarketplaceSettings(ctx, marketplaceAddress.Hex())
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error get admin address")
		return nil, err
	}

	royalty, err := strconv.ParseFloat(res.Royalty, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	settings := &entities.MarketplaceSettings{
		Id:          int64(res.ID),
		Marketplace: marketplaceAddress,
		Beneficiary: common.HexToAddress(res.Beneficiary),
		Royalty:     royalty,
	}

	return settings, err
}
