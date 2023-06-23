package services

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type SearchService interface {
	SearchNFTsWithListings(ctx context.Context, token common.Address, owner common.Address, q string, isHidden *bool, offset int32, limit int32) ([]*entities.NftRead, error)
}

func (s *Services) SearchNFTsWithListings(
	ctx context.Context,
	token common.Address,
	owner common.Address,
	q string,
	isHidden *bool,
	offset int32,
	limit int32,
) ([]*entities.NftRead, error) {
	return s.searcher.FullTextSearch(
		ctx,
		token,
		owner,
		q,
		isHidden,
		offset,
		limit,
	)
}
