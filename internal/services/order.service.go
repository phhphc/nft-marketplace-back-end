package services

import (
	"context"
	"database/sql"
	"encoding/hex"
	"reflect"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order entities.Order) error
	FulFillOrder(ctx context.Context, order entities.Order) error
}

func (s *Services) CreateOrder(ctx context.Context, order entities.Order) (err error) {
	// TODO - use transaction

	var orderType sql.NullInt32
	if order.OrderType != nil {
		orderType = sql.NullInt32{
			Valid: true,
			Int32: int32(*order.OrderType),
		}
	}

	var salt sql.NullString
	if order.Salt != nil {
		salt = sql.NullString{
			Valid:  true,
			String: order.Salt.Hex(),
		}
	}

	var signature sql.NullString
	if order.Signature != nil {
		signature = sql.NullString{
			Valid:  true,
			String: "0x" + hex.EncodeToString(order.Signature),
		}
	}

	var recipient sql.NullString
	if order.Recipient != nil {
		recipient = sql.NullString{
			Valid:  true,
			String: order.Recipient.Hex(),
		}
	}

	err = s.repo.InsertOrder(ctx, postgresql.InsertOrderParams{
		OrderHash:   order.OrderHash.Hex(),
		Offerer:     order.Offerer.Hex(),
		Recipient:   recipient,
		Zone:        order.Zone.Hex(),
		OrderType:   orderType,
		ZoneHash:    order.ZoneHash.Hex(),
		Salt:        salt,
		StartTime:   ToNullString(order.StartTime),
		EndTime:     ToNullString(order.EndTime),
		Signature:   signature,
		IsValidated: false,
		IsCancelled: false,
		IsFulfilled: false,
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error create order")
		return
	}

	for _, offerItem := range order.Offer {
		params := postgresql.InsertOrderOfferItemParams{
			OrderHash:   order.OrderHash.Hex(),
			ItemType:    int32(offerItem.ItemType),
			Token:       offerItem.Token.Hex(),
			Identifier:  offerItem.Identifier.String(),
			Amount:      ToNullString(offerItem.Amount),
			StartAmount: ToNullString(offerItem.StartAmount),
			EndAmount:   ToNullString(offerItem.EndAmount),
		}
		err = s.repo.InsertOrderOfferItem(ctx, params)
		if err != nil {
			s.lg.Error().Caller().Interface("param", params).Err(err).Msg("error create order")
			return
		}
	}

	for _, considerationItem := range order.Consideration {
		err = s.repo.InsertOrderConsiderationItem(ctx, postgresql.InsertOrderConsiderationItemParams{
			OrderHash:   order.OrderHash.Hex(),
			ItemType:    int32(considerationItem.ItemType),
			Token:       considerationItem.Token.Hex(),
			Identifier:  considerationItem.Identifier.String(),
			Amount:      ToNullString(considerationItem.Amount),
			StartAmount: ToNullString(considerationItem.StartAmount),
			EndAmount:   ToNullString(considerationItem.EndAmount),
			Recipient:   considerationItem.Recipient.Hex(),
		})
		if err != nil {
			s.lg.Error().Caller().Err(err).Msg("error create order")
			return
		}
	}

	return
}

func (s *Services) FulFillOrder(ctx context.Context, order entities.Order) (err error) {
	// TODO - use tx

	orderHash, err := s.repo.UpdateOrderStatus(ctx, postgresql.UpdateOrderStatusParams{
		OrderHash:   order.OrderHash.Hex(),
		IsValidated: sql.NullBool{Bool: true, Valid: true},
		IsCancelled: sql.NullBool{Bool: false, Valid: true},
		IsFulfilled: sql.NullBool{Bool: true, Valid: true},
	})
	s.lg.Info().Caller().Msg(orderHash)
	if err == nil {
		return
	}
	s.lg.Error().Caller().Err(err).Msg(orderHash)

	var orderType sql.NullInt32
	if order.OrderType != nil {
		orderType = sql.NullInt32{
			Valid: true,
			Int32: int32(*order.OrderType),
		}
	}

	var salt sql.NullString
	if order.Salt != nil {
		salt = sql.NullString{
			Valid:  true,
			String: order.Salt.Hex(),
		}
	}

	var signature sql.NullString
	if order.Signature != nil {
		signature = sql.NullString{
			Valid:  true,
			String: "0x" + hex.EncodeToString(order.Signature),
		}
	}

	var recipient sql.NullString
	if order.Recipient != nil {
		recipient = sql.NullString{
			Valid:  true,
			String: order.Recipient.Hex(),
		}
	}

	err = s.repo.InsertOrder(ctx, postgresql.InsertOrderParams{
		OrderHash:   order.OrderHash.Hex(),
		Offerer:     order.Offerer.Hex(),
		Recipient:   recipient,
		Zone:        order.Zone.Hex(),
		OrderType:   orderType,
		ZoneHash:    order.ZoneHash.Hex(),
		Salt:        salt,
		StartTime:   ToNullString(order.StartTime),
		EndTime:     ToNullString(order.EndTime),
		Signature:   signature,
		IsValidated: true,
		IsCancelled: false,
		IsFulfilled: true,
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error create order")
		return
	}

	for _, offerItem := range order.Offer {
		err = s.repo.InsertOrderOfferItem(ctx, postgresql.InsertOrderOfferItemParams{
			OrderHash:   order.OrderHash.Hex(),
			ItemType:    int32(offerItem.ItemType),
			Token:       offerItem.Token.Hex(),
			Identifier:  offerItem.Identifier.String(),
			StartAmount: ToNullString(offerItem.StartAmount),
			EndAmount:   ToNullString(offerItem.EndAmount),
		})
		if err != nil {
			s.lg.Error().Caller().Err(err).Msg("error create order")
			return
		}
	}

	for _, considerationItem := range order.Consideration {
		err = s.repo.InsertOrderConsiderationItem(ctx, postgresql.InsertOrderConsiderationItemParams{
			OrderHash:   order.OrderHash.Hex(),
			ItemType:    int32(considerationItem.ItemType),
			Token:       considerationItem.Token.Hex(),
			Identifier:  considerationItem.Identifier.String(),
			StartAmount: ToNullString(considerationItem.StartAmount),
			EndAmount:   ToNullString(considerationItem.EndAmount),
			Recipient:   considerationItem.Recipient.Hex(),
		})
		if err != nil {
			s.lg.Error().Caller().Err(err).Msg("error create order")
			return
		}
	}

	return
}

type Stringer interface {
	String() string
}

func ToNullString(s Stringer) (ns sql.NullString) {
	if !reflect.ValueOf(s).IsNil() {
		ns.Valid = true
		ns.String = s.String()
	}
	return
}
