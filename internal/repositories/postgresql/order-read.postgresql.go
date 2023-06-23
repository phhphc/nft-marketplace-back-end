package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	helpsql "github.com/phhphc/nft-marketplace-back-end/internal/repositories/help-sql"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql/gen"
	"github.com/phhphc/nft-marketplace-back-end/internal/services/infrastructure"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

type dbOffer struct {
	ItemType    int    `json:"item_type"`
	Token       string `json:"token"`
	Identifier  string `json:"identifier"`
	StartAmount string `json:"start_amount"`
	EndAmount   string `json:"end_amount"`
}

type dbConsideration struct {
	ItemType    int    `json:"item_type"`
	Token       string `json:"token"`
	Identifier  string `json:"identifier"`
	StartAmount string `json:"start_amount"`
	EndAmount   string `json:"end_amount"`
	Recipient   string `json:"recipient"`
}

type dbOrderStatus struct {
	IsFulfilled bool `json:"is_fulfilled"`
	IsCancelled bool `json:"is_cancelled"`
	IsInvalid   bool `json:"is_invalid"`
}

type dbOrder struct {
	OrderHash     string            `json:"order_hash"`
	Offerer       string            `json:"offerer"`
	Signature     string            `json:"signature"`
	StartTime     string            `json:"start_time"`
	EndTime       string            `json:"end_time"`
	Salt          string            `json:"salt"`
	Status        dbOrderStatus     `json:"status"`
	Offer         []dbOffer         `json:"offer"`
	Consideration []dbConsideration `json:"consideration"`
}

func (r *PostgresqlRepository) FindOrder(
	ctx context.Context,
	offer infrastructure.FindOrderOffer,
	consideration infrastructure.FindOrderConsideration,
	orderHash common.Hash,
	offerer common.Address,
	IsFulfilled *bool,
	IsCancelled *bool,
	IsInvalid *bool,
) ([]entities.Order, error) {

	res, err := r.queries.GetOrder(
		ctx,
		gen.GetOrderParams{
			OrderHash: helpsql.HashToNullString(orderHash),
			Offerer:   helpsql.AddressToNullString(offerer),

			OfferToken:      helpsql.AddressToNullString(offer.Token),
			OfferIdentifier: helpsql.PointerBigIntToNullString(offer.Identifier),

			ConsiderationToken:      helpsql.AddressToNullString(consideration.Token),
			ConsiderationIdentifier: helpsql.PointerBigIntToNullString(consideration.Identifier),

			IsCancelled: helpsql.PointerBoolToNullBool(IsCancelled),
			IsFulfilled: helpsql.PointerBoolToNullBool(IsFulfilled),
			IsInvalid:   helpsql.PointerBoolToNullBool(IsInvalid),
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error queries")
		return nil, err
	}

	os := make([]entities.Order, len(res))
	for i, row := range res {
		var dbo dbOrder
		err = json.Unmarshal(row, &dbo)
		if err != nil {
			r.lg.Error().Caller().Err(err).Msg("error unmashal")
			return nil, err
		}

		os[i] = dbOrderToEntityOrder(dbo)
	}

	return os, nil
}

func (r *PostgresqlRepository) FindExpiredOrder(
	ctx context.Context,
) ([]entities.ExpiredOrder, error) {

	res, err := r.queries.GetExpiredOrder(
		ctx,
		sql.NullString{
			String: strconv.FormatInt(time.Now().Unix(), 10),
			Valid:  true,
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("get expired orders error")
		return nil, err
	}

	eos := make([]entities.ExpiredOrder, len(res))
	for i, row := range res {
		eos[i] = entities.ExpiredOrder{
			EventName:   row.Name,
			OrderHash:   common.HexToHash(row.OrderHash),
			EndTime:     util.MustStringToBigInt(row.EndTime.String),
			IsCancelled: row.IsCancelled,
			IsInvalid:   row.IsInvalid,
			Offerer:     common.HexToAddress(row.Offerer),
		}
	}

	return eos, nil
}

func dbOrderToEntityOrder(
	dbo dbOrder,
) (o entities.Order) {
	dboOffer := dbo.Offer
	ois := make([]entities.OfferItem, len(dboOffer))
	for i, oi := range dboOffer {
		ois[i] = entities.OfferItem{
			ItemType:    entities.EnumItemType(oi.ItemType),
			Token:       common.HexToAddress(oi.Token),
			Identifier:  util.MustStringToBigInt(oi.Identifier),
			StartAmount: util.MustStringToBigInt(oi.StartAmount),
			EndAmount:   util.MustStringToBigInt(oi.EndAmount),
		}
	}

	dboConsideration := dbo.Consideration
	cis := make([]entities.ConsiderationItem, len(dboConsideration))
	for i, ci := range dboConsideration {
		cis[i] = entities.ConsiderationItem{
			ItemType:    entities.EnumItemType(ci.ItemType),
			Token:       common.HexToAddress(ci.Token),
			Identifier:  util.MustStringToBigInt(ci.Identifier),
			StartAmount: util.MustStringToBigInt(ci.StartAmount),
			EndAmount:   util.MustStringToBigInt(ci.EndAmount),
			Recipient:   common.HexToAddress(ci.Recipient),
		}
	}

	dboStatus := dbo.Status
	o = entities.Order{
		OrderHash: common.HexToHash(dbo.OrderHash),
		Offerer:   common.HexToAddress(dbo.Offerer),
		StartTime: util.MustStringToBigInt(dbo.StartTime),
		EndTime:   util.MustStringToBigInt(dbo.EndTime),
		Salt:      &[]common.Hash{common.HexToHash(dbo.Salt)}[0],
		Signature: util.MustHexToBytes(dbo.Signature),
		Status: entities.OrderStatus{
			IsFulfilled: dboStatus.IsFulfilled,
			IsCancelled: dboStatus.IsCancelled,
			IsValidated: dboStatus.IsInvalid,
		},
		Offer:         ois,
		Consideration: cis,
	}
	return
}
