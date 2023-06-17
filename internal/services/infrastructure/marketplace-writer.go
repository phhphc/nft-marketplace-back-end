package infrastructure

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type MarketplaceWriter interface {
	UpdateMarketplaceLastSyncBlock(
		ctx context.Context,
		block uint64,
	) error

	InsertMarketplaceSettings(
		ctx context.Context,
		marketplace common.Address,
		admin common.Address,
		signer common.Address,
		royalty float64,
		sighash []byte,
		signature []byte,
		jsonTypedData []byte,
		createdAt *big.Int,
	) (*entities.MarketplaceSettings, error)
}
