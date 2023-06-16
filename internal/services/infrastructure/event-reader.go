package infrastructure

import (
	"context"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type EventReader interface {
	FindEvent(
		ctx context.Context,
		query entities.EventRead,
	) (events []entities.Event, err error)
}
