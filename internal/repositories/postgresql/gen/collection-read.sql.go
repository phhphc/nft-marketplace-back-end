// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: collection-read.sql

package gen

import (
	"context"
	"database/sql"

	"github.com/tabbed/pqtype"
)

const getCollection = `-- name: GetCollection :many
SELECT token, owner, co.name, ca.name as category, description, metadata, created_at
FROM collections co
         JOIN categories ca on co.category = ca.id
WHERE token ILIKE COALESCE($1, token)
  AND owner ILIKE COALESCE($2, owner)
  AND co.name ILIKE COALESCE($3, co.name)
  AND ca.name ILIKE COALESCE($4, ca.name)
OFFSET $5 LIMIT $6
`

type GetCollectionParams struct {
	Token    sql.NullString
	Owner    sql.NullString
	Name     sql.NullString
	Category sql.NullString
	Offset   int32
	Limit    int32
}

type GetCollectionRow struct {
	Token       string
	Owner       string
	Name        string
	Category    string
	Description string
	Metadata    pqtype.NullRawMessage
	CreatedAt   sql.NullTime
}

func (q *Queries) GetCollection(ctx context.Context, arg GetCollectionParams) ([]GetCollectionRow, error) {
	rows, err := q.query(ctx, q.getCollectionStmt, getCollection,
		arg.Token,
		arg.Owner,
		arg.Name,
		arg.Category,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetCollectionRow{}
	for rows.Next() {
		var i GetCollectionRow
		if err := rows.Scan(
			&i.Token,
			&i.Owner,
			&i.Name,
			&i.Category,
			&i.Description,
			&i.Metadata,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCollectionLastSyncBlock = `-- name: GetCollectionLastSyncBlock :one
SELECT "last_sync_block"
FROM collections
WHERE token = $1
`

func (q *Queries) GetCollectionLastSyncBlock(ctx context.Context, token string) (int64, error) {
	row := q.queryRow(ctx, q.getCollectionLastSyncBlockStmt, getCollectionLastSyncBlock, token)
	var last_sync_block int64
	err := row.Scan(&last_sync_block)
	return last_sync_block, err
}