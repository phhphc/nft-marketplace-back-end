// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: marketplace_settings.sql

package postgresql

import (
	"context"
	"database/sql"

	"github.com/tabbed/pqtype"
)

const getMarketplaceSettings = `-- name: GetMarketplaceSettings :one
SELECT ms.id, ms.marketplace, ms.admin, ms.signer, ms.royalty, ms.sighash, ms.signature, ms.created_at
FROM "marketplace_settings" ms
WHERE (ms.marketplace = $1 OR $1 IS NULL)
AND (ms.id = $2 OR $2 IS NULL)
ORDER BY ms.id DESC
LIMIT 1
`

type GetMarketplaceSettingsParams struct {
	Marketplace sql.NullString
	ID          sql.NullInt32
}

type GetMarketplaceSettingsRow struct {
	ID          int32
	Marketplace string
	Admin       string
	Signer      string
	Royalty     string
	Sighash     sql.NullString
	Signature   sql.NullString
	CreatedAt   sql.NullString
}

func (q *Queries) GetMarketplaceSettings(ctx context.Context, arg GetMarketplaceSettingsParams) (GetMarketplaceSettingsRow, error) {
	row := q.db.QueryRowContext(ctx, getMarketplaceSettings, arg.Marketplace, arg.ID)
	var i GetMarketplaceSettingsRow
	err := row.Scan(
		&i.ID,
		&i.Marketplace,
		&i.Admin,
		&i.Signer,
		&i.Royalty,
		&i.Sighash,
		&i.Signature,
		&i.CreatedAt,
	)
	return i, err
}

const getValidMarketplaceSettings = `-- name: GetValidMarketplaceSettings :one
SELECT ms.id, ms.marketplace, ms.admin, ms.signer, ms.royalty, ms.sighash, ms.signature, ms.created_at
FROM "marketplace_settings" ms
WHERE ms.marketplace = $1
AND ms.signature IS NOT NULL OR ms.id = 1
ORDER BY ms.id DESC
LIMIT 1
`

type GetValidMarketplaceSettingsRow struct {
	ID          int32
	Marketplace string
	Admin       string
	Signer      string
	Royalty     string
	Sighash     sql.NullString
	Signature   sql.NullString
	CreatedAt   sql.NullString
}

func (q *Queries) GetValidMarketplaceSettings(ctx context.Context, marketplace string) (GetValidMarketplaceSettingsRow, error) {
	row := q.db.QueryRowContext(ctx, getValidMarketplaceSettings, marketplace)
	var i GetValidMarketplaceSettingsRow
	err := row.Scan(
		&i.ID,
		&i.Marketplace,
		&i.Admin,
		&i.Signer,
		&i.Royalty,
		&i.Sighash,
		&i.Signature,
		&i.CreatedAt,
	)
	return i, err
}

const insertMarketplaceSettings = `-- name: InsertMarketplaceSettings :one
INSERT INTO "marketplace_settings" ("marketplace", "admin", "signer", "royalty", "typed_data", "sighash", "signature", "created_at")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, marketplace, admin, signer, royalty, typed_data, sighash, signature, created_at
`

type InsertMarketplaceSettingsParams struct {
	Marketplace string
	Admin       string
	Signer      string
	Royalty     string
	TypedData   pqtype.NullRawMessage
	Sighash     sql.NullString
	Signature   sql.NullString
	CreatedAt   sql.NullString
}

func (q *Queries) InsertMarketplaceSettings(ctx context.Context, arg InsertMarketplaceSettingsParams) (MarketplaceSetting, error) {
	row := q.db.QueryRowContext(ctx, insertMarketplaceSettings,
		arg.Marketplace,
		arg.Admin,
		arg.Signer,
		arg.Royalty,
		arg.TypedData,
		arg.Sighash,
		arg.Signature,
		arg.CreatedAt,
	)
	var i MarketplaceSetting
	err := row.Scan(
		&i.ID,
		&i.Marketplace,
		&i.Admin,
		&i.Signer,
		&i.Royalty,
		&i.TypedData,
		&i.Sighash,
		&i.Signature,
		&i.CreatedAt,
	)
	return i, err
}