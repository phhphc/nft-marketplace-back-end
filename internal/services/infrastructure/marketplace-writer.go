package infrastructure

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type MarketplaceWriter interface {
	UpdateMarketplaceLastSyncBlock(
		ctx context.Context,
		block uint64,
	) error

	UpdateMarketplaceSettings(
		ctx context.Context,
		marketplace common.Address,
		beneficiary common.Address,
		royalty float64,
	) (*entities.MarketplaceSettings, error)
}
