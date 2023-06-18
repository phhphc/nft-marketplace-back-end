package postgresql

import (
	"context"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	helpsql "github.com/phhphc/nft-marketplace-back-end/internal/repositories/help-sql"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql/gen"
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
			OrderHash:   helpsql.HashToNullString(condition.OrderHash),
			Offerer:     helpsql.AddressToNullString(condition.Offerer),
			IsValidated: helpsql.PointerBoolToNullBool(value.IsValidated),
			IsCancelled: helpsql.PointerBoolToNullBool(value.IsCancelled),
			IsFulfilled: helpsql.PointerBoolToNullBool(value.IsFulfilled),
			IsInvalid:   helpsql.PointerBoolToNullBool(value.IsInvalid),
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
			IsValidated: helpsql.PointerBoolToNullBool(value.IsValidated),
			IsCancelled: helpsql.PointerBoolToNullBool(value.IsCancelled),
			IsFulfilled: helpsql.PointerBoolToNullBool(value.IsFulfilled),
			IsInvalid:   helpsql.PointerBoolToNullBool(value.IsInvalid),
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
			StartTime:   helpsql.PointerBigIntToNullString(order.StartTime),
			EndTime:     helpsql.PointerBigIntToNullString(order.EndTime),
			Salt:        helpsql.HashToNullString(*order.Salt),
			Signature:   helpsql.BytesToNullString(order.Signature),
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
				StartAmount: helpsql.PointerBigIntToNullString(oi.StartAmount),
				EndAmount:   helpsql.PointerBigIntToNullString(oi.EndAmount),
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
				StartAmount: helpsql.PointerBigIntToNullString(ci.StartAmount),
				EndAmount:   helpsql.PointerBigIntToNullString(ci.EndAmount),
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
