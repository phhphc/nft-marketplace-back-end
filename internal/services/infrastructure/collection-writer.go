package infrastructure

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type CollectionWriter interface {
	CreateCollection(
		ctx context.Context,
		collection entities.Collection,
	) (ec entities.Collection, err error)
	UpdateCollectionLastSyncBlock(
		ctx context.Context,
		token common.Address,
		block uint64,
	) error
}
