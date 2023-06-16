package infrastructure

import (
	"context"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type NotificationWriter interface {
	InsertNotification(
		ctx context.Context,
		notification entities.NotificationPost,
	) error
	UpdateNotification(
		ctx context.Context,
		notification entities.NotificationUpdate,
	) error
}
