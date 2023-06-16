// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package postgresql

import (
	"context"
	"database/sql"
)

const deleteUserRole = `-- name: DeleteUserRole :exec
DELETE FROM "user_roles"
WHERE address = $1 AND role_id = $2
`

type DeleteUserRoleParams struct {
	Address string
	RoleID  int32
}

func (q *Queries) DeleteUserRole(ctx context.Context, arg DeleteUserRoleParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserRole, arg.Address, arg.RoleID)
	return err
}

const getAllRoles = `-- name: GetAllRoles :many
SELECT id, name FROM "roles"
`

func (q *Queries) GetAllRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.QueryContext(ctx, getAllRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Role{}
	for rows.Next() {
		var i Role
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const getUserByAddress = `-- name: GetUserByAddress :one
SELECT public_address, nonce, is_block
FROM "users"
WHERE public_address ILIKE $1
`

func (q *Queries) GetUserByAddress(ctx context.Context, publicAddress string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAddress, publicAddress)
	var i User
	err := row.Scan(&i.PublicAddress, &i.Nonce, &i.IsBlock)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT fu.public_address, fu.nonce, r.id as role_id, r.name as role, fu.is_block
FROM (
    SELECT public_address, nonce, is_block FROM "users" u
    WHERE (u.public_address ILIKE $1 OR $1 IS NULL)
    AND (u.is_block = $2 OR $2 IS NULL)
    ORDER BY public_address ASC
    LIMIT $4
    OFFSET $3
     ) fu
LEFT JOIN "user_roles" ur on fu.public_address = ur.address
LEFT JOIN "roles" r on r.id = ur.role_id
WHERE (r.name = $5 OR $5 IS NULL)
`

type GetUsersParams struct {
	PublicAddress sql.NullString
	IsBlock       sql.NullBool
	Offset        int32
	Limit         int32
	Role          sql.NullString
}

type GetUsersRow struct {
	PublicAddress string
	Nonce         string
	RoleID        sql.NullInt32
	Role          sql.NullString
	IsBlock       bool
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsers,
		arg.PublicAddress,
		arg.IsBlock,
		arg.Offset,
		arg.Limit,
		arg.Role,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUsersRow{}
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.PublicAddress,
			&i.Nonce,
			&i.RoleID,
			&i.Role,
			&i.IsBlock,
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

const insertUser = `-- name: InsertUser :one
INSERT INTO "users" (public_address, nonce)
VALUES ($1, $2)
RETURNING public_address, nonce, is_block
`

type InsertUserParams struct {
	PublicAddress string
	Nonce         string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser, arg.PublicAddress, arg.Nonce)
	var i User
	err := row.Scan(&i.PublicAddress, &i.Nonce, &i.IsBlock)
	return i, err
}

const insertUserRole = `-- name: InsertUserRole :one
INSERT INTO "user_roles" (address, role_id)
VALUES ($1, $2)
RETURNING address, role_id
`

type InsertUserRoleParams struct {
	Address string
	RoleID  int32
}

func (q *Queries) InsertUserRole(ctx context.Context, arg InsertUserRoleParams) (UserRole, error) {
	row := q.db.QueryRowContext(ctx, insertUserRole, arg.Address, arg.RoleID)
	var i UserRole
	err := row.Scan(&i.Address, &i.RoleID)
	return i, err
}

const updateNonce = `-- name: UpdateNonce :one
UPDATE "users"
SET nonce = $1
WHERE public_address ILIKE $2
RETURNING public_address, nonce, is_block
`

type UpdateNonceParams struct {
	Nonce         string
	PublicAddress string
}

func (q *Queries) UpdateNonce(ctx context.Context, arg UpdateNonceParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateNonce, arg.Nonce, arg.PublicAddress)
	var i User
	err := row.Scan(&i.PublicAddress, &i.Nonce, &i.IsBlock)
	return i, err
}

const updateUserBlockState = `-- name: UpdateUserBlockState :one
UPDATE "users"
SET is_block = $1
WHERE public_address ILIKE $2
RETURNING public_address, nonce, is_block
`

type UpdateUserBlockStateParams struct {
	IsBlock       bool
	PublicAddress string
}

func (q *Queries) UpdateUserBlockState(ctx context.Context, arg UpdateUserBlockStateParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserBlockState, arg.IsBlock, arg.PublicAddress)
	var i User
	err := row.Scan(&i.PublicAddress, &i.Nonce, &i.IsBlock)
	return i, err
}
