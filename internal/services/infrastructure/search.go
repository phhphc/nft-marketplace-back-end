package infrastructure

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type Searcher interface {
	FullTextSearch(
		ctx context.Context,
		token common.Address,
		owner common.Address,
		q string,
		isHidden *bool,
		offset int32,
		limit int32,
	) ([]*entities.NftRead, error)
}
