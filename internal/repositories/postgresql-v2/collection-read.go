package postgresql

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2/gen"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

func (r *PostgresqlRepository) FindCollection(
	ctx context.Context,
	query entities.Collection,
	offset int,
	limit int,
) ([]entities.Collection, error) {

	dbc, err := r.queries.GetCollection(
		ctx,
		gen.GetCollectionParams{
			Token:    AddressToNullString(query.Token),
			Owner:    AddressToNullString(query.Owner),
			Name:     StringToNullString(query.Name),
			Category: StringToNullString(query.Category),
			Offset:   int32(offset),
			Limit:    int32(limit),
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error find")
		return nil, err
	}

	cs := make([]entities.Collection, len(dbc))
	for i, c := range dbc {
		cs[i] = entities.Collection{
			Token:       common.HexToAddress(c.Token),
			Owner:       common.HexToAddress(c.Owner),
			Name:        c.Name,
			Description: c.Description,
			Metadata:    util.MustBytesToMapJson(c.Metadata.RawMessage),
			Category:    c.Category,
		}
	}
	return cs, nil
}

func (r *PostgresqlRepository) GetCollectionLastSyncBlock(
	ctx context.Context,
	token common.Address,
) (uint64, error) {

	dbc, err := r.queries.GetCollectionLastSyncBlock(
		ctx,
		token.Hex(),
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error find one")
		return 0, err
	}

	return uint64(dbc), nil
}
