package infrastructure

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type NftListing struct {
	OrderHash  common.Hash
	ItemType   entities.EnumItemType
	StartPrice *big.Int
	EndPrice   *big.Int
	StartTime  *big.Int
	EndTime    *big.Int
}

type NftWithListing struct {
	entities.Nft
	Listings []NftListing
}

type NftReader interface {
	FindOneNft(
		ctx context.Context,
		token common.Address,
		identifier *big.Int,
	) (entities.Nft, error)
	FindNftsWithListings(
		ctx context.Context,
		token common.Address,
		identifier *big.Int,
		owner common.Address,
		isHidden *bool,
		offset int32,
		limit int32,
		listingLimit int32,
	) ([]NftWithListing, error)
}
