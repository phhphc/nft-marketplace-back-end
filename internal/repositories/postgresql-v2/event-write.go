package postgresql

import (
	"context"
	"database/sql"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2/gen"
)

func (r *PostgresqlRepository) InsertEvent(
	ctx context.Context,
	event entities.Event,
) (ee entities.Event, err error) {
	r.lg.Info().Caller().
		Str("name", event.Name).
		Str("token", event.Token.Hex()).
		Str("token_id", event.TokenId.String()).
		Str("event_type", event.Type).
		Str("from", event.From.Hex()).
		Str("to", event.To.Hex()).
		Str("order_hash", event.OrderHash.Hex()).
		Msg("create event")

	dbEvent, err := r.queries.InsertEvent(
		ctx,
		gen.InsertEventParams{
			Quantity: sql.NullInt32{
				Valid: true,
				Int32: event.Quantity,
			},
			Type: sql.NullString{
				Valid:  true,
				String: event.Type,
			},
			Price: sql.NullString{
				Valid:  true,
				String: event.Price.String(),
			},
			From: event.From.Hex(),
			To: sql.NullString{
				Valid:  true,
				String: event.To.Hex(),
			},
			TxHash: sql.NullString{
				Valid:  true,
				String: event.TxHash,
			},
			OrderHash: sql.NullString{
				Valid:  true,
				String: event.OrderHash.Hex(),
			},
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error insert")
		return
	}

	ee.Name = dbEvent.Name
	ee.Token = common.HexToAddress(dbEvent.Token)
	// token_id
	tokenId, ok := big.NewInt(0).SetString(dbEvent.TokenID, 10)
	if ok {
		ee.TokenId = tokenId
	}
	// quantity
	if dbEvent.Quantity.Valid {
		ee.Quantity = dbEvent.Quantity.Int32
	}
	// is_bundle
	if dbEvent.Type.Valid {
		ee.Type = dbEvent.Type.String
	}
	// price
	price, ok := big.NewInt(0).SetString(dbEvent.Price.String, 10)
	if dbEvent.Price.Valid && ok {
		ee.Price = price
	}
	// from
	ee.From = common.HexToAddress(dbEvent.From)
	// to
	if dbEvent.To.Valid {
		ee.To = common.HexToAddress(dbEvent.To.String)
	}
	// date
	ee.Date = dbEvent.Date.Time
	//link
	ee.TxHash = dbEvent.TxHash.String
	//order hash
	if dbEvent.OrderHash.Valid {
		ee.OrderHash = common.HexToHash(dbEvent.OrderHash.String)
	}

	return
}
