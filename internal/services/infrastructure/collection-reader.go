package infrastructure

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type CollectionReader interface {
	FindCollection(
		ctx context.Context,
		query entities.Collection,
		offset int,
		limit int,
	) (ec []entities.Collection, err error)
	GetCollectionLastSyncBlock(
		ctx context.Context,
		token common.Address,
	) (uint64, error)
}
