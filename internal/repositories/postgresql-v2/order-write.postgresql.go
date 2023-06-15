package postgresql

import (
	"context"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2/gen"
	"github.com/phhphc/nft-marketplace-back-end/internal/services/infrastructure"
)

func (r *PostgresqlRepository) UpdateOrderStatus(
	ctx context.Context,
	condition infrastructure.UpdateOrderCondition,
	value infrastructure.UpdateOrderValue,
) error {
	_, err := r.queries.UpdateOrderStatus(
		ctx,
		gen.UpdateOrderStatusParams{
			OrderHash:   HashToNullString(condition.OrderHash),
			Offerer:     AddressToNullString(condition.Offerer),
			IsValidated: PointerBoolToNullBool(value.IsValidated),
			IsCancelled: PointerBoolToNullBool(value.IsCancelled),
			IsFulfilled: PointerBoolToNullBool(value.IsFulfilled),
			IsInvalid:   PointerBoolToNullBool(value.IsInvalid),
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error update")
		return err
	}
	return nil
}

func (r *PostgresqlRepository) UpdateOrderStatusByOffer(
	ctx context.Context,
	condition infrastructure.UpdateOrderStatusByOfferCondition,
	value infrastructure.UpdateOrderValue,
) error {
	_, err := r.queries.UpdateOrderStatusByOffer(
		ctx,
		gen.UpdateOrderStatusByOfferParams{
			Offerer:     condition.Offerer.Hex(),
			Token:       condition.Token.Hex(),
			Identifier:  condition.Identifier.String(),
			IsValidated: PointerBoolToNullBool(value.IsValidated),
			IsCancelled: PointerBoolToNullBool(value.IsCancelled),
			IsFulfilled: PointerBoolToNullBool(value.IsFulfilled),
			IsInvalid:   PointerBoolToNullBool(value.IsInvalid),
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error update")
		return err
	}
	return nil
}

func (r *PostgresqlRepository) InsertOneOrder(
	ctx context.Context,
	order entities.Order,
) (o entities.Order, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()
	qtx := r.queries.WithTx(tx)

	_, err = qtx.InsertOrder(
		ctx,
		gen.InsertOrderParams{
			OrderHash:   order.OrderHash.Hex(),
			Offerer:     order.Offerer.Hex(),
			StartTime:   PointerBigIntToNullString(order.StartTime),
			EndTime:     PointerBigIntToNullString(order.EndTime),
			Salt:        HashToNullString(*order.Salt),
			Signature:   BytesToNullString(order.Signature),
			IsValidated: order.Status.IsValidated,
			IsCancelled: order.Status.IsCancelled,
			IsFulfilled: order.Status.IsFulfilled,
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error insert")
		return
	}

	orderOffer := order.Offer
	for _, oi := range orderOffer {
		_, err = qtx.InsertOrderOfferItem(
			ctx,
			gen.InsertOrderOfferItemParams{
				OrderHash:   order.OrderHash.Hex(),
				ItemType:    int32(oi.ItemType),
				Token:       oi.Token.Hex(),
				Identifier:  oi.Identifier.String(),
				StartAmount: PointerBigIntToNullString(oi.StartAmount),
				EndAmount:   PointerBigIntToNullString(oi.EndAmount),
			},
		)
		if err != nil {
			r.lg.Error().Caller().Err(err).Msg("error insert")
			return
		}
	}

	orderConsideration := order.Consideration
	for _, ci := range orderConsideration {
		_, err = qtx.InsertOrderConsiderationItem(
			ctx,
			gen.InsertOrderConsiderationItemParams{
				OrderHash:   order.OrderHash.Hex(),
				ItemType:    int32(ci.ItemType),
				Token:       ci.Token.Hex(),
				Identifier:  ci.Identifier.String(),
				StartAmount: PointerBigIntToNullString(ci.StartAmount),
				EndAmount:   PointerBigIntToNullString(ci.EndAmount),
				Recipient:   ci.Recipient.Hex(),
			},
		)
		if err != nil {
			r.lg.Error().Caller().Err(err).Msg("error insert")
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error commit")
		return
	}

	o = order
	return
}
