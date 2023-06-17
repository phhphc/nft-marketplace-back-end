package postgresql

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2/gen"
)

func (r *PostgresqlRepository) UpdateMarketplaceLastSyncBlock(
	ctx context.Context,
	block uint64,
) error {
	err := r.queries.UpdateMarketplaceLastSyncBlock(ctx, int64(block))
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error update block")
	}
	return err
}

func (r *PostgresqlRepository) UpdateMarketplaceSettings(
	ctx context.Context,
	marketplace common.Address,
	beneficiary common.Address,
	royalty float64,
) (*entities.MarketplaceSettings, error) {

	arg := gen.UpdateMarketplaceSettingsParams{
		Marketplace:  marketplace.Hex(),
		NBeneficiary: sql.NullString{String: beneficiary.Hex(), Valid: true},
		NRoyalty:     sql.NullString{String: strconv.FormatFloat(royalty, 'f', 6, 64), Valid: true},
	}

	res, err := r.queries.UpdateMarketplaceSettings(ctx, arg)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error update marketplace settings")
		return nil, err
	}

	settings := &entities.MarketplaceSettings{
		Id:          int64(res.ID),
		Marketplace: common.HexToAddress(res.Marketplace),
		Beneficiary: common.HexToAddress(res.Beneficiary),
		Royalty:     royalty,
	}
	return settings, nil
}
