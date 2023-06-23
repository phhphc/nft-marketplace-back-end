package infrastructure

import (
	"context"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type ProfileReader interface {
	FindOneProfile(
		ctx context.Context,
		address string,
	) (entities.Profile, error)
}
