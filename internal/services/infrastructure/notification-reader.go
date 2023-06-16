package infrastructure

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type NotificationReader interface {
	FindNotification(
		ctx context.Context,
		address common.Address,
		isViewed *bool,
	) (ns []entities.NotificationGet, err error)
}
