package infrastructure

import (
	"context"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type EventWriter interface {
	InsertEvent(
		ctx context.Context,
		event entities.Event,
	) (ee entities.Event, err error)
}
