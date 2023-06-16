package services

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type NotificationService interface {
	CreateNotification(ctx context.Context, notification entities.NotificationPost) error
	GetListNotification(ctx context.Context, address common.Address, isViewed *bool) ([]entities.NotificationGet, error)
	UpdateNotification(ctx context.Context, notification entities.NotificationUpdate) error
}

func (s *Services) CreateNotification(ctx context.Context, notification entities.NotificationPost) error {
	return s.notificationWriter.InsertNotification(
		ctx,
		notification,
	)
}

func (s *Services) GetListNotification(ctx context.Context, address common.Address, isViewed *bool) (ns []entities.NotificationGet, err error) {
	return s.notificationReader.FindNotification(
		ctx,
		address,
		isViewed,
	)
}

func (s *Services) UpdateNotification(ctx context.Context, notification entities.NotificationUpdate) error {
	return s.notificationWriter.UpdateNotification(
		ctx,
		notification,
	)
}
