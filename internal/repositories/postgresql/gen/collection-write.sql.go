// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: collection-write.sql

package gen

import (
	"context"

	"github.com/tabbed/pqtype"
)

const insertCollection = `-- name: InsertCollection :one
INSERT INTO "collections" ("token", "owner", "name", "description","category", "metadata")
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING token, owner, name, description, metadata, category, created_at, last_sync_block
`

type InsertCollectionParams struct {
	Token       string
	Owner       string
	Name        string
	Description string
	Category    int32
	Metadata    pqtype.NullRawMessage
}

func (q *Queries) InsertCollection(ctx context.Context, arg InsertCollectionParams) (Collection, error) {
	row := q.queryRow(ctx, q.insertCollectionStmt, insertCollection,
		arg.Token,
		arg.Owner,
		arg.Name,
		arg.Description,
		arg.Category,
		arg.Metadata,
	)
	var i Collection
	err := row.Scan(
		&i.Token,
		&i.Owner,
		&i.Name,
		&i.Description,
		&i.Metadata,
		&i.Category,
		&i.CreatedAt,
		&i.LastSyncBlock,
	)
	return i, err
}

const updateCollectionLastSyncBlock = `-- name: UpdateCollectionLastSyncBlock :exec
UPDATE collections
SET "last_sync_block" = $2
WHERE token = $1
`

type UpdateCollectionLastSyncBlockParams struct {
	Token         string
	LastSyncBlock int64
}

func (q *Queries) UpdateCollectionLastSyncBlock(ctx context.Context, arg UpdateCollectionLastSyncBlockParams) error {
	_, err := q.exec(ctx, q.updateCollectionLastSyncBlockStmt, updateCollectionLastSyncBlock, arg.Token, arg.LastSyncBlock)
	return err
}
