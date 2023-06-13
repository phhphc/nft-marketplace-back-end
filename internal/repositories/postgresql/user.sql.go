// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package postgresql

import (
	"context"
)

const getUserByAddress = `-- name: GetUserByAddress :one
SELECT public_address, nonce
FROM "users"
WHERE public_address ILIKE $1
`

func (q *Queries) GetUserByAddress(ctx context.Context, publicAddress string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAddress, publicAddress)
	var i User
	err := row.Scan(&i.PublicAddress, &i.Nonce)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO "users" (public_address, nonce)
VALUES ($1, $2)
RETURNING public_address, nonce
`

type InsertUserParams struct {
	PublicAddress string
	Nonce         string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser, arg.PublicAddress, arg.Nonce)
	var i User
	err := row.Scan(&i.PublicAddress, &i.Nonce)
	return i, err
}

const updateNonce = `-- name: UpdateNonce :one
UPDATE "users"
SET nonce = $1
WHERE public_address ILIKE $2
RETURNING public_address, nonce
`

type UpdateNonceParams struct {
	Nonce         string
	PublicAddress string
}

func (q *Queries) UpdateNonce(ctx context.Context, arg UpdateNonceParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateNonce, arg.Nonce, arg.PublicAddress)
	var i User
	err := row.Scan(&i.PublicAddress, &i.Nonce)
	return i, err
}
