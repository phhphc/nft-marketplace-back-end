package postgresql

import (
	"context"
	"database/sql"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql/gen"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

func (r *PostgresqlRepository) FindNotification(
	ctx context.Context,
	address common.Address,
	isViewed *bool,
) (ns []entities.NotificationGet, err error) {

	params := gen.GetNotificationParams{}
	if address != (common.Address{}) {
		params.Address = sql.NullString{
			Valid:  true,
			String: address.Hex(),
		}
	}
	if isViewed != nil {
		params.IsViewed = sql.NullBool{
			Valid: true,
			Bool:  *isViewed,
		}
	}

	notificationList, err := r.queries.GetNotification(ctx, params)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("cannot get list notification")
		return
	}

	for _, notification := range notificationList {
		newNotification := entities.NotificationGet{
			IsViewed:  notification.IsViewed.Bool,
			Info:      notification.Info,
			EventName: notification.EventName,
			OrderHash: common.HexToHash(notification.OrderHash),
			Address:   common.HexToAddress(notification.Address),
			Token:     common.HexToAddress(notification.Token),
			TokenId:   util.MustStringToBigInt(notification.TokenID),
			Quantity:  notification.Quantity.Int32,
			Type:      notification.Type.String,
			Price:     util.MustStringToBigInt(notification.Price.String),
			From:      common.HexToAddress(notification.From),
			To:        common.HexToAddress(notification.To.String),
			Date:      notification.Date.Time,
			Owner:     common.HexToAddress(notification.Owner),
			NftImage:  notification.NftImage,
			NftName:   notification.NftName,
		}

		ns = append(ns, newNotification)
	}
	return
}
