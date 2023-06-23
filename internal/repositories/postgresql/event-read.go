package postgresql

import (
	"context"
	"database/sql"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	helpsql "github.com/phhphc/nft-marketplace-back-end/internal/repositories/help-sql"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql/gen"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

func (r *PostgresqlRepository) FindEvent(
	ctx context.Context,
	query entities.EventRead,
) (events []entities.Event, err error) {

	eventList, err := r.queries.GetEvent(
		ctx,
		gen.GetEventParams{
			Name:    helpsql.StringToNullString(query.Name),
			Token:   helpsql.AddressToNullString(query.Token),
			TokenID: helpsql.PointerBigIntToNullString(query.TokenId),
			Address: helpsql.AddressToNullString(query.Address),
			Type:    helpsql.StringToNullString(query.Type),
			Month:   helpsql.PointerIntToNullInt32(query.Month),
			Year:    helpsql.PointerIntToNullInt32(query.Year),
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error find event")
		return
	}

	for _, event := range eventList {
		newEvent := entities.Event{
			Name:      event.Name,
			Token:     common.HexToAddress(event.Token),
			TokenId:   util.MustStringToBigInt(event.TokenID),
			From:      common.HexToAddress(event.From),
			Date:      event.Date.Time,
			TxHash:    event.TxHash.String,
			NftImage:  event.NftImage,
			NftName:   event.NftName,
			OrderHash: common.HexToHash(event.OrderHash.String),
		}

		if event.Quantity.Valid {
			newEvent.Quantity = event.Quantity.Int32
		}

		price, ok := big.NewInt(0).SetString(event.Price.String, 10)
		if event.Price.Valid && ok {

			newEvent.Price = price
		}

		if event.To.Valid {
			newEvent.To = common.HexToAddress(event.To.String)
		}

		if event.Type.Valid {
			newEvent.Type = event.Type.String
		}

		if event.EndTime.Valid {
			newEvent.EndTime = util.MustStringToBigInt(event.EndTime.String)
		}

		if event.IsCancelled.Valid {
			newEvent.IsCancelled = event.IsCancelled.Bool
		}

		if event.IsFulfilled.Valid {
			newEvent.IsFulfilled = event.IsFulfilled.Bool
		}

		events = append(events, newEvent)
	}
	return
}

func (r *PostgresqlRepository) GetOffer(
	ctx context.Context,
	owner common.Address,
	from common.Address,
) (offers []entities.Event, err error) {
	params := gen.GetOfferParams{}
	if owner != (common.Address{}) {
		params.Owner = sql.NullString{
			String: owner.Hex(),
			Valid:  true,
		}
	}
	if from != (common.Address{}) {
		params.From = sql.NullString{
			String: from.Hex(),
			Valid:  true,
		}
	}

	offerList, err := r.queries.GetOffer(ctx, params)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("cannot get list offer")
		return
	}

	for _, offer := range offerList {
		newOffer := entities.Event{
			Name:        offer.Name,
			Token:       common.HexToAddress(offer.Token),
			TokenId:     util.MustStringToBigInt(offer.TokenID),
			Quantity:    offer.Quantity.Int32,
			NftImage:    offer.NftImage,
			NftName:     offer.NftName,
			Type:        offer.Type.String,
			OrderHash:   common.HexToHash(offer.OrderHash.String),
			Price:       util.MustStringToBigInt(offer.Price.String),
			Owner:       common.HexToAddress(offer.Owner),
			From:        common.HexToAddress(offer.From),
			EndTime:     util.MustStringToBigInt(offer.EndTime.String),
			IsFulfilled: offer.IsFulfilled.Bool,
			IsCancelled: offer.IsCancelled.Bool,
			IsExpired:   offer.IsExpired,
		}

		offers = append(offers, newOffer)
	}
	return
}
