package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2/gen"
	"github.com/tabbed/pqtype"
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

func (r *PostgresqlRepository) InsertMarketplaceSettings(
	ctx context.Context,
	marketplace common.Address,
	admin common.Address,
	signer common.Address,
	royalty float64,
	sighash []byte,
	signature []byte,
	jsonTypedData []byte,
	createdAt *big.Int,
) (*entities.MarketplaceSettings, error) {
	arg := gen.InsertMarketplaceSettingsParams{
		Marketplace: marketplace.Hex(),
		Admin:       admin.Hex(),
		Signer:      signer.Hex(),
		Royalty:     fmt.Sprintf("%f", royalty),
		Sighash:     sql.NullString{String: common.BytesToHash(sighash).String(), Valid: true},
		Signature:   sql.NullString{String: common.BytesToHash(signature).String(), Valid: true},
		TypedData:   pqtype.NullRawMessage{RawMessage: jsonTypedData, Valid: true},
		CreatedAt:   sql.NullString{String: createdAt.String(), Valid: true},
	}

	fmt.Printf("arg: %+v\n", arg)

	res, err := r.queries.InsertMarketplaceSettings(ctx, arg)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error insert marketplace settings")
		return nil, err
	}

	settings := &entities.MarketplaceSettings{
		Id:          int64(res.ID),
		Marketplace: marketplace,
		Admin:       admin,
		Signer:      signer,
		Royalty:     royalty,
		Sighash:     common.BytesToHash(sighash),
		CreatedAt:   *createdAt,
	}
	return settings, err
}
