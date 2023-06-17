package infrastructure

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type ProfileWriter interface {
	UpsertProfile(
		ctx context.Context,
		profile entities.Profile,
	) (entities.Profile, error)

	DeleteProfile(
		ctx context.Context,
		address common.Address,
	) error

	GetOffer(
		ctx context.Context,
		owner common.Address,
		from common.Address,
	) (offers []entities.Event, err error)
}
