package postgresql

import (
	"context"
	"database/sql"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2/gen"
)

func (r *PostgresqlRepository) InsertNotification(
	ctx context.Context,
	notification entities.NotificationPost,
) error {
	r.lg.Info().Caller().
		Str("info", notification.Info).
		Str("event_name", notification.EventName).
		Str("order_hash", notification.OrderHash.Hex()).
		Str("address", notification.Address.Hex()).
		Msg("create notification")

	_, err := r.queries.InsertNotification(ctx, gen.InsertNotificationParams{
		Info:      notification.Info,
		EventName: notification.EventName,
		OrderHash: notification.OrderHash.Hex(),
		Address:   notification.Address.Hex(),
	})

	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("cannot create notification")
		return err
	}

	return nil
}

func (r *PostgresqlRepository) UpdateNotification(
	ctx context.Context,
	notification entities.NotificationUpdate,
) error {
	r.lg.Info().Caller().
		Str("event_name", notification.EventName).
		Str("order_hash", notification.OrderHash.Hex()).
		Msg("update notification")

	_, err := r.queries.UpdateNotification(ctx, gen.UpdateNotificationParams{
		EventName: sql.NullString{
			Valid:  true,
			String: notification.EventName,
		},
		OrderHash: sql.NullString{
			Valid:  true,
			String: notification.OrderHash.Hex(),
		},
	})

	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("cannot update notification")
		return err
	}

	return nil
}
