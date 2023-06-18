package infrastructure

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type EventReader interface {
	FindEvent(
		ctx context.Context,
		query entities.EventRead,
	) (events []entities.Event, err error)

	GetOffer(
		ctx context.Context,
		owner common.Address,
		from common.Address,
	) (offers []entities.Event, err error)
}
