package services

import (
	"context"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order entities.Order) error
	FulFillOrder(ctx context.Context, order entities.Order) error
	GetOrder(
		ctx context.Context,
		offer entities.OfferItem,
		consideration entities.ConsiderationItem,
		orderHash common.Hash,
		offerer common.Address,
		IsFulfilled *bool,
		IsCancelled *bool,
		IsInvalid *bool,
	) ([]map[string]any, error)
	RemoveInvalidOrder(ctx context.Context, offerer common.Address, token common.Address, identifier *big.Int) error
	HandleOrderCancelled(ctx context.Context, orderHash common.Hash) error
	HandleCounterIncremented(ctx context.Context, offerer common.Address) error
	GetExpiredOrder(ctx context.Context) ([]entities.ExpiredOrder, error)
}

func (s *Services) CreateOrder(ctx context.Context, order entities.Order) (err error) {
	// TODO - use transaction

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

	// TODO - insert event listing or offer
	s.CreateEventsByOrder(ctx, order)

	return
}

func (s *Services) FulFillOrder(ctx context.Context, order entities.Order) error {
	// TODO - use tx

	err := s.repo.UpdateOrderStatus(ctx, postgresql.UpdateOrderStatusParams{
		OrderHash:   sql.NullString{String: order.OrderHash.Hex(), Valid: true},
		IsValidated: sql.NullBool{Bool: true, Valid: true},
		IsCancelled: sql.NullBool{Bool: false, Valid: true},
		IsFulfilled: sql.NullBool{Bool: true, Valid: true},
	})
	if err == nil {
		return nil
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
		return err
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
			return err
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
			return err
		}
	}

	return nil
}

func (s *Services) GetOrder(
	ctx context.Context,
	offer entities.OfferItem,
	consideration entities.ConsiderationItem,
	orderHash common.Hash,
	offerer common.Address,
	IsFulfilled *bool,
	IsCancelled *bool,
	IsInvalid *bool,
) ([]map[string]any, error) {
	params := postgresql.GetOrderParams{}
	if IsFulfilled != nil {
		params.IsFulfilled = sql.NullBool{
			Bool:  *IsFulfilled,
			Valid: true,
		}
	}
	if IsCancelled != nil {
		params.IsCancelled = sql.NullBool{
			Bool:  *IsCancelled,
			Valid: true,
		}
	}
	if IsInvalid != nil {
		params.IsInvalid = sql.NullBool{
			Bool:  *IsInvalid,
			Valid: true,
		}
	}

	if (consideration.Token != common.Address{}) {
		params.ConsiderationToken = sql.NullString{
			String: consideration.Token.Hex(),
			Valid:  true,
		}
	}
	if consideration.Identifier != nil {
		params.ConsiderationIdentifier = sql.NullString{
			String: consideration.Identifier.String(),
			Valid:  true,
		}
	}

	if (offer.Token != common.Address{}) {
		params.OfferToken = sql.NullString{
			String: offer.Token.Hex(),
			Valid:  true,
		}
	}
	if offer.Identifier != nil {
		params.OfferIdentifier = sql.NullString{
			String: offer.Identifier.String(),
			Valid:  true,
		}
	}

	if orderHash != (common.Hash{}) {
		params.OrderHash = sql.NullString{
			String: orderHash.Hex(),
			Valid:  true,
		}
	}
	if offerer != (common.Address{}) {
		params.Offerer = sql.NullString{
			String: offerer.Hex(),
			Valid:  true,
		}
	}

	js, err := s.repo.GetOrder(ctx, params)
	if err != nil {
		s.lg.Error().Caller().Err(err).Interface("param", params).Msg("Err")
		return nil, err
	}

	rs := make([]map[string]any, len(js))
	for i, j := range js {
		json.Unmarshal(j, &rs[i])
	}
	return rs, nil
}

func (s *Services) RemoveInvalidOrder(ctx context.Context, offerer common.Address, token common.Address, identifier *big.Int) error {
	err := s.repo.MarkOrderInvalid(ctx, postgresql.MarkOrderInvalidParams{
		Offerer:    offerer.Hex(),
		Token:      token.Hex(),
		Identifier: identifier.String(),
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("update error")
	}

	return err
}

func (s *Services) HandleOrderCancelled(ctx context.Context, orderHash common.Hash) error {
	err := s.repo.UpdateOrderStatus(ctx, postgresql.UpdateOrderStatusParams{
		OrderHash: sql.NullString{
			String: orderHash.Hex(),
			Valid:  true,
		},
		IsCancelled: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
		IsInvalid: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("update error")
	}

	return err
}

func (s *Services) HandleCounterIncremented(ctx context.Context, offerer common.Address) error {
	err := s.repo.UpdateOrderStatus(ctx, postgresql.UpdateOrderStatusParams{
		Offerer: sql.NullString{
			String: offerer.Hex(),
			Valid:  true,
		},
		IsCancelled: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
		IsInvalid: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("update error")
	}

	return err
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

func (s *Services) GetExpiredOrder(ctx context.Context) (expiredOrderList []entities.ExpiredOrder, err error) {
	eos, err := s.repo.GetExpiredOrder(ctx)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("get expired orders error")
		return
	}

	for _, eo := range eos {
		expiredOrder := entities.ExpiredOrder{
			EventName: 		eo.Name,
			OrderHash: 		common.HexToHash(eo.OrderHash),
			EndTime: 		ToBigInt(eo.EndTime.String),
			IsCancelled:	eo.IsCancelled,
			IsInvalid:		eo.IsInvalid,
			Offerer:		common.HexToAddress(eo.Offerer),
		}

		// Update expired order's isInvalid to true
		s.repo.UpdateOrderStatus(ctx, postgresql.UpdateOrderStatusParams{
			OrderHash: sql.NullString{
				Valid: true,
				String: eo.OrderHash,
			},
			IsInvalid: sql.NullBool{
				Valid: true,
				Bool: true,
			},
		})
		
		expiredOrderList = append(expiredOrderList, expiredOrder)
	}
	return
}