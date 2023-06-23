package infrastructure

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type MarketplaceReader interface {
	GetMarketplaceLastSyncBlock(
		ctx context.Context,
	) (uint64, error)

	GetMarketplaceSettings(
		ctx context.Context,
		marketplaceAddress common.Address,
	) (*entities.MarketplaceSettings, error)
}
