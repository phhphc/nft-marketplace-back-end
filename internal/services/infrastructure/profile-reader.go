package infrastructure

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type ProfileReader interface {
	FindOneProfile(
		ctx context.Context,
		address string,
	) (entities.Profile, error)

	GetOffer(
		ctx context.Context,
		owner common.Address,
		from common.Address,
	) (offers []entities.Event, err error)
}
