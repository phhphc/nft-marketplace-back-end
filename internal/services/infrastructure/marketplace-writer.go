package infrastructure

import "context"

type MarketplaceWriter interface {
	UpdateMarketplaceLastSyncBlock(
		ctx context.Context,
		block uint64,
	) error
}
