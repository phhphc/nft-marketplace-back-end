package services

import (
	"context"
	"database/sql"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
)

type NotificationService interface {
	CreateNotification(ctx context.Context, notification entities.NotificationPost) error
	GetListNotification(ctx context.Context, address common.Address, isViewed *bool) ([]entities.NotificationGet, error)
	UpdateNotification(ctx context.Context, notification entities.NotificationUpdate) error
}

func (s *Services) CreateNotification(ctx context.Context, notification entities.NotificationPost) error {
	s.lg.Info().Caller().
		Str("info",notification.Info).
		Str("event_name", notification.EventName).
		Str("order_hash", notification.OrderHash.Hex()).
		Str("address", notification.Address.Hex()).
		Msg("create notification")
	
	_, err := s.repo.InsertNotification(ctx, postgresql.InsertNotificationParams{
		Info: notification.Info,
		EventName: notification.EventName,
		OrderHash: notification.OrderHash.Hex(),
		Address: notification.Address.Hex(),	
	})

	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot create notification")
		return err
	}
	
	return nil
}

func (s *Services) GetListNotification(ctx context.Context, address common.Address, isViewed *bool) (ns []entities.NotificationGet, err error) {
	
	params := postgresql.GetNotificationParams {}
	if address != (common.Address{}) {
		params.Address = sql.NullString{
			Valid: true,
			String: address.Hex(),
		}
	}
	if isViewed != nil {
		params.IsViewed = sql.NullBool{
			Valid: true,
			Bool: *isViewed,
		}
	}

	notificationList, err := s.repo.GetNotification(ctx, params)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot get list notification")
		return
	}

	for _, notification := range notificationList {
		newNotification := entities.NotificationGet {
			IsViewed: 	notification.IsViewed.Bool,
			Info: 		notification.Info,
			EventName: 	notification.EventName,
			OrderHash: 	common.HexToHash(notification.OrderHash),
			Address:   	common.HexToAddress(notification.Address),
			Token:     	common.HexToAddress(notification.Token),
			TokenId:  	ToBigInt(notification.TokenID),
			Quantity: 	notification.Quantity.Int32,
			Type: 		notification.Type.String,
			Price: 		ToBigInt(notification.Price.String),
			From: 		common.HexToAddress(notification.From),
			To: 		common.HexToAddress(notification.To.String),
			Date: 		notification.Date.Time,
			Owner: 		common.HexToAddress(notification.Owner),
			NftImage: 	notification.NftImage,
			NftName: 	notification.NftName,
		}

		ns = append(ns, newNotification)
	}
	return
}

func (s *Services) UpdateNotification(ctx context.Context, notification entities.NotificationUpdate) error {
	s.lg.Info().Caller().
		Str("event_name", notification.EventName).
		Str("order_hash", notification.OrderHash.Hex()).
		Msg("update notification")
	
	_, err := s.repo.UpdateNotification(ctx, postgresql.UpdateNotificationParams{
		EventName: sql.NullString{
			Valid: true,
			String: notification.EventName,
		},
		OrderHash: sql.NullString{
			Valid: true,
			String: notification.OrderHash.Hex(),
		},
	})

	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot update notification")
		return err
	}
	
	return nil
}