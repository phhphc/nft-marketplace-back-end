package services

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type ProfileService interface {
	GetProfile(ctx context.Context, address string) (entities.Profile, error)
	UpsertProfile(ctx context.Context, profile entities.Profile) (entities.Profile, error)
	DeleteProfile(ctx context.Context, address common.Address) error
	GetOffer(ctx context.Context, owner common.Address, from common.Address) ([]entities.Event, error)
}

func (s *Services) GetProfile(ctx context.Context, address string) (entities.Profile, error) {
	return s.profileReader.FindOneProfile(
		ctx,
		address,
	)
}

func (s *Services) UpsertProfile(ctx context.Context, profile entities.Profile) (entities.Profile, error) {
	return s.profileWriter.UpsertProfile(
		ctx,
		profile,
	)
}

func (s *Services) DeleteProfile(ctx context.Context, address common.Address) error {
	return s.profileWriter.DeleteProfile(
		ctx,
		address,
	)
}

func (s *Services) GetOffer(ctx context.Context, owner common.Address, from common.Address) (offers []entities.Event, err error) {
	return s.eventReader.GetOffer(
		ctx,
		owner,
		from,
	)
}
