package postgresql

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	helpsql "github.com/phhphc/nft-marketplace-back-end/internal/repositories/help-sql"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql/gen"
)

func (r *PostgresqlRepository) CreateCollection(
	ctx context.Context,
	collection entities.Collection,
) (ec entities.Collection, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()
	qtx := r.queries.WithTx(tx)

	c, err := qtx.GetCategoryByName(
		ctx,
		collection.Category,
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("err get category")
		return
	}

	_, err = qtx.InsertCollection(
		ctx,
		gen.InsertCollectionParams{
			Token:       collection.Token.Hex(),
			Owner:       collection.Owner.Hex(),
			Name:        collection.Name,
			Description: collection.Description,
			Category:    c.ID,
			Metadata:    helpsql.MustMapJsonToNullRawMessage(collection.Metadata),
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("err insert")
		return
	}

	err = tx.Commit()
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error commit")
		return
	}

	ec = collection
	return
}

func (r *PostgresqlRepository) UpdateCollectionLastSyncBlock(
	ctx context.Context,
	token common.Address,
	block uint64,
) error {
	err := r.queries.UpdateCollectionLastSyncBlock(
		ctx,
		gen.UpdateCollectionLastSyncBlockParams{
			Token:         token.Hex(),
			LastSyncBlock: int64(block),
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error update")
		return err
	}

	return nil
}
