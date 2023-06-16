package infrastructure

import "context"

type MarketplaceReader interface {
	GetMarketplaceLastSyncBlock(
		ctx context.Context,
	) (uint64, error)
}