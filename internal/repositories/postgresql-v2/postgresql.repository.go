package postgresql

import (
	"context"

	"database/sql"

	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2/gen"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"

	_ "github.com/lib/pq"
)

type PostgresqlRepository struct {
	db      *sql.DB
	queries *gen.Queries
	lg      *log.Logger
}

func (r *PostgresqlRepository) Close() error {
	return r.queries.Close()
}

func NewPostgresqlRepository(
	ctx context.Context,
	dataSourceName string,
) (*PostgresqlRepository, error) {
	lg := log.GetLogger()

	db, err := getDb(ctx, dataSourceName)
	if err != nil {
		lg.Error().Err(err).Caller().Msg("error get db")
		return nil, err
	}

	queries, err := gen.Prepare(ctx, db)
	if err != nil {
		lg.Error().Err(err).Caller().Msg("error prepared")
		return nil, err
	}

	r := PostgresqlRepository{
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
