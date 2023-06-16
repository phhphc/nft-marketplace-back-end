package postgresql

import "context"

func (r *PostgresqlRepository) GetMarketplaceLastSyncBlock(
	ctx context.Context,
) (uint64, error) {
	lastSyncBlock, err := r.queries.GetMarketplaceLastSyncBlock(ctx)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error get last block")
	}
	return uint64(lastSyncBlock), err
}
