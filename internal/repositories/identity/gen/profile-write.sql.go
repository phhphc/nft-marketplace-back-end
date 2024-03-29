// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: profile-write.sql

package gen

import (
	"context"
	"database/sql"

	"github.com/tabbed/pqtype"
)

const deleteProfile = `-- name: DeleteProfile :exec
DELETE FROM "profiles"
WHERE "address" = $1
`

func (q *Queries) DeleteProfile(ctx context.Context, address string) error {
	_, err := q.exec(ctx, q.deleteProfileStmt, deleteProfile, address)
	return err
}

const upsertProfile = `-- name: UpsertProfile :one
INSERT INTO "profiles" ("address", "username", "metadata", "signature")
VALUES ($1, $2, $3, $4)
ON CONFLICT ("address") DO UPDATE SET
  "username" = $2,
  "metadata" = $3,
  "signature" = $4
RETURNING address, username, metadata, signature
`

type UpsertProfileParams struct {
	Address   string
	Username  sql.NullString
	Metadata  pqtype.NullRawMessage
	Signature string
}

func (q *Queries) UpsertProfile(ctx context.Context, arg UpsertProfileParams) (Profile, error) {
	row := q.queryRow(ctx, q.upsertProfileStmt, upsertProfile,
		arg.Address,
		arg.Username,
		arg.Metadata,
		arg.Signature,
	)
	var i Profile
	err := row.Scan(
		&i.Address,
		&i.Username,
		&i.Metadata,
		&i.Signature,
	)
	return i, err
}
