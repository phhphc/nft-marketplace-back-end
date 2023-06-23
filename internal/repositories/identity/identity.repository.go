package identity

import (
	"context"

	"database/sql"

	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/identity/gen"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"

	_ "github.com/lib/pq"
)

type IdentityRepository struct {
	db      *sql.DB
	queries *gen.Queries
	lg      *log.Logger
}

func (r *IdentityRepository) Close() error {
	// return r.queries.Close()
	return nil
}

func NewIdentityRepository(
	ctx context.Context,
	dataSourceName string,
) (*IdentityRepository, error) {
	lg := log.GetLogger()

	db, err := getDb(ctx, dataSourceName)
	if err != nil {
		lg.Error().Err(err).Caller().Msg("error get db")
		return nil, err
	}

	lg.Debug().Caller().Interface("db", db).Str("source", dataSourceName).Msg("debug")
	queries, err := gen.Prepare(ctx, db)
	if err != nil {
		lg.Error().Err(err).Caller().Msg("error prepared")
		return nil, err
	}
	// queries := gen.New(db)

	r := IdentityRepository{
		db:      db,
		queries: queries,
		lg:      lg,
	}
	return &r, nil
}

func getDb(
	ctx context.Context,
	dataSourceName string,
) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
