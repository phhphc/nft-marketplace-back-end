package postgresql

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2/gen"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

func (r *PostgresqlRepository) FindEvent(
	ctx context.Context,
	query entities.EventRead,
) (events []entities.Event, err error) {

	eventList, err := r.queries.GetEvent(
		ctx,
		gen.GetEventParams{
			Name:    StringToNullString(query.Name),
			Token:   AddressToNullString(query.Address),
			TokenID: PointerBigIntToNullString(query.TokenId),
			Address: AddressToNullString(query.Address),
			Type:    StringToNullString(query.Type),
			Month:   PointerIntToNullInt32(query.Month),
			Year:    PointerIntToNullInt32(query.Year),
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
