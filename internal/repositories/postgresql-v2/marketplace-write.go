package postgresql

import "context"

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
