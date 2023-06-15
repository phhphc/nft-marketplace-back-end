package services

import (
	"context"
	"database/sql"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/services/infrastructure"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
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
	) ([]entities.Order, error)
	RemoveInvalidOrder(ctx context.Context, offerer common.Address, token common.Address, identifier *big.Int) error
	HandleOrderCancelled(ctx context.Context, orderHash common.Hash) error
	HandleCounterIncremented(ctx context.Context, offerer common.Address) error
	GetExpiredOrder(ctx context.Context) ([]entities.ExpiredOrder, error)
}

func (s *Services) CreateOrder(ctx context.Context, order entities.Order) (err error) {

	o, err := s.orderWriter.InsertOneOrder(
		ctx,
		order,
	)

	// TODO - insert event listing or offer
	s.CreateEventsByOrder(ctx, o)

	return
}

func (s *Services) FulFillOrder(ctx context.Context, order entities.Order) error {
	err := s.orderWriter.UpdateOrderStatus(
		ctx,
		infrastructure.UpdateOrderCondition{
			OrderHash: order.OrderHash,
		},
		infrastructure.UpdateOrderValue{
			IsValidated: util.TruePointer,
			IsCancelled: util.FalsePointer,
			IsFulfilled: util.TruePointer,
		},
	)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update")
		return err
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
) ([]entities.Order, error) {

	os, err := s.orderReader.FindOrder(
		ctx,
		infrastructure.FindOrderOffer{
			Token:      offer.Token,
			Identifier: offer.Identifier,
		},
		infrastructure.FindOrderConsideration{
			Token:      consideration.Token,
			Identifier: consideration.Identifier,
		},
		orderHash,
		offerer,
		IsFulfilled,
		IsCancelled,
		IsInvalid,
	)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error find")
		return nil, err
	}
	return os, err
}

func (s *Services) RemoveInvalidOrder(
	ctx context.Context,
	offerer common.Address,
	token common.Address,
	identifier *big.Int,
) error {
	err := s.orderWriter.UpdateOrderStatusByOffer(
		ctx,
		infrastructure.UpdateOrderStatusByOfferCondition{
			Offerer:    offerer,
			Token:      token,
			Identifier: identifier,
		},
		infrastructure.UpdateOrderValue{
			IsInvalid: util.TruePointer,
		},
	)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("update error")
	}

	return err
}

func (s *Services) HandleOrderCancelled(ctx context.Context, orderHash common.Hash) error {
	err := s.orderWriter.UpdateOrderStatus(
		ctx,
		infrastructure.UpdateOrderCondition{
			OrderHash: orderHash,
		},
		infrastructure.UpdateOrderValue{
			IsCancelled: util.TruePointer,
			IsInvalid:   util.TruePointer,
		},
	)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update")
		return err
	}

	return nil
}

func (s *Services) HandleCounterIncremented(ctx context.Context, offerer common.Address) error {
	err := s.orderWriter.UpdateOrderStatus(
		ctx,
		infrastructure.UpdateOrderCondition{
			Offerer: offerer,
		},
		infrastructure.UpdateOrderValue{
			IsCancelled: util.TruePointer,
			IsInvalid:   util.TruePointer,
		},
	)
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
	eos, err := s.orderReader.FindExpiredOrder(ctx)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("get expired orders error")
		return
	}

	for _, eo := range eos {
		err = s.orderWriter.UpdateOrderStatus(
			ctx,
			infrastructure.UpdateOrderCondition{
				OrderHash: eo.OrderHash,
			},
			infrastructure.UpdateOrderValue{
				IsInvalid: util.TruePointer,
			},
		)
		if err != nil {
			s.lg.Error().Caller().Err(err).Msg("error update")
			return
		}
	}

	return eos, err
}
