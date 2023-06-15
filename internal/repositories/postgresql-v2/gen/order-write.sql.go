// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: order-write.sql

package gen

import (
	"context"
	"database/sql"
)

const insertOrder = `-- name: InsertOrder :one
INSERT INTO "orders" ("order_hash",
                      "offerer",
                      "recipient",
                      "salt",
                      "start_time",
                      "end_time",
                      "signature",
                      "is_validated",
                      "is_cancelled",
                      "is_fulfilled")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING order_hash, offerer, recipient, salt, start_time, end_time, signature, is_cancelled, is_validated, is_fulfilled, is_invalid
`

type InsertOrderParams struct {
	OrderHash   string
	Offerer     string
	Recipient   sql.NullString
	Salt        sql.NullString
	StartTime   sql.NullString
	EndTime     sql.NullString
	Signature   sql.NullString
	IsValidated bool
	IsCancelled bool
	IsFulfilled bool
}

func (q *Queries) InsertOrder(ctx context.Context, arg InsertOrderParams) (Order, error) {
	row := q.queryRow(ctx, q.insertOrderStmt, insertOrder,
		arg.OrderHash,
		arg.Offerer,
		arg.Recipient,
		arg.Salt,
		arg.StartTime,
		arg.EndTime,
		arg.Signature,
		arg.IsValidated,
		arg.IsCancelled,
		arg.IsFulfilled,
	)
	var i Order
	err := row.Scan(
		&i.OrderHash,
		&i.Offerer,
		&i.Recipient,
		&i.Salt,
		&i.StartTime,
		&i.EndTime,
		&i.Signature,
		&i.IsCancelled,
		&i.IsValidated,
		&i.IsFulfilled,
		&i.IsInvalid,
	)
	return i, err
}

const insertOrderConsiderationItem = `-- name: InsertOrderConsiderationItem :one
INSERT INTO "consideration_items" ("order_hash",
                                   "item_type",
                                   "token",
                                   "identifier",
                                   "amount",
                                   "start_amount",
                                   "end_amount",
                                   "recipient")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, order_hash, item_type, token, identifier, amount, start_amount, end_amount, recipient
`

type InsertOrderConsiderationItemParams struct {
	OrderHash   string
	ItemType    int32
	Token       string
	Identifier  string
	Amount      sql.NullString
	StartAmount sql.NullString
	EndAmount   sql.NullString
	Recipient   string
}

func (q *Queries) InsertOrderConsiderationItem(ctx context.Context, arg InsertOrderConsiderationItemParams) (ConsiderationItem, error) {
	row := q.queryRow(ctx, q.insertOrderConsiderationItemStmt, insertOrderConsiderationItem,
		arg.OrderHash,
		arg.ItemType,
		arg.Token,
		arg.Identifier,
		arg.Amount,
		arg.StartAmount,
		arg.EndAmount,
		arg.Recipient,
	)
	var i ConsiderationItem
	err := row.Scan(
		&i.ID,
		&i.OrderHash,
		&i.ItemType,
		&i.Token,
		&i.Identifier,
		&i.Amount,
		&i.StartAmount,
		&i.EndAmount,
		&i.Recipient,
	)
	return i, err
}

const insertOrderOfferItem = `-- name: InsertOrderOfferItem :one
INSERT INTO "offer_items" ("order_hash",
                           "item_type",
                           "token",
                           "identifier",
                           "amount",
                           "start_amount",
                           "end_amount")
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, order_hash, item_type, token, identifier, amount, start_amount, end_amount
`

type InsertOrderOfferItemParams struct {
	OrderHash   string
	ItemType    int32
	Token       string
	Identifier  string
	Amount      sql.NullString
	StartAmount sql.NullString
	EndAmount   sql.NullString
}

func (q *Queries) InsertOrderOfferItem(ctx context.Context, arg InsertOrderOfferItemParams) (OfferItem, error) {
	row := q.queryRow(ctx, q.insertOrderOfferItemStmt, insertOrderOfferItem,
		arg.OrderHash,
		arg.ItemType,
		arg.Token,
		arg.Identifier,
		arg.Amount,
		arg.StartAmount,
		arg.EndAmount,
	)
	var i OfferItem
	err := row.Scan(
		&i.ID,
		&i.OrderHash,
		&i.ItemType,
		&i.Token,
		&i.Identifier,
		&i.Amount,
		&i.StartAmount,
		&i.EndAmount,
	)
	return i, err
}

const updateOrderStatus = `-- name: UpdateOrderStatus :many
UPDATE "orders"
SET "is_validated" = COALESCE($1, "is_validated"),
    "is_cancelled" = COALESCE($2, "is_cancelled"),
    "is_fulfilled" = COALESCE($3, "is_fulfilled"),
    "is_invalid"   = COALESCE($4, "is_invalid")
WHERE "order_hash" = COALESCE($5, "order_hash")
  AND "offerer" = COALESCE($6, "offerer")
RETURNING order_hash, offerer, recipient, salt, start_time, end_time, signature, is_cancelled, is_validated, is_fulfilled, is_invalid
`

type UpdateOrderStatusParams struct {
	IsValidated sql.NullBool
	IsCancelled sql.NullBool
	IsFulfilled sql.NullBool
	IsInvalid   sql.NullBool
	OrderHash   sql.NullString
	Offerer     sql.NullString
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) ([]Order, error) {
	rows, err := q.query(ctx, q.updateOrderStatusStmt, updateOrderStatus,
		arg.IsValidated,
		arg.IsCancelled,
		arg.IsFulfilled,
		arg.IsInvalid,
		arg.OrderHash,
		arg.Offerer,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderHash,
			&i.Offerer,
			&i.Recipient,
			&i.Salt,
			&i.StartTime,
			&i.EndTime,
			&i.Signature,
			&i.IsCancelled,
			&i.IsValidated,
			&i.IsFulfilled,
			&i.IsInvalid,
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

const updateOrderStatusByOffer = `-- name: UpdateOrderStatusByOffer :many
UPDATE orders o
SET "is_validated" = COALESCE($1, "is_validated"),
    "is_cancelled" = COALESCE($2, "is_cancelled"),
    "is_fulfilled" = COALESCE($3, "is_fulfilled"),
    "is_invalid"   = COALESCE($4, "is_invalid")
WHERE o.order_hash IN (SELECT DISTINCT o.order_hash
                       FROM orders o
                                JOIN offer_items oi on o.order_hash = oi.order_hash
                       WHERE o.is_invalid = false
                         AND o.offerer = $5
                         AND oi.token = $6
                         AND oi.identifier = $7)
RETURNING order_hash, offerer, recipient, salt, start_time, end_time, signature, is_cancelled, is_validated, is_fulfilled, is_invalid
`

type UpdateOrderStatusByOfferParams struct {
	IsValidated sql.NullBool
	IsCancelled sql.NullBool
	IsFulfilled sql.NullBool
	IsInvalid   sql.NullBool
	Offerer     string
	Token       string
	Identifier  string
}

func (q *Queries) UpdateOrderStatusByOffer(ctx context.Context, arg UpdateOrderStatusByOfferParams) ([]Order, error) {
	rows, err := q.query(ctx, q.updateOrderStatusByOfferStmt, updateOrderStatusByOffer,
		arg.IsValidated,
		arg.IsCancelled,
		arg.IsFulfilled,
		arg.IsInvalid,
		arg.Offerer,
		arg.Token,
		arg.Identifier,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderHash,
			&i.Offerer,
			&i.Recipient,
			&i.Salt,
			&i.StartTime,
			&i.EndTime,
			&i.Signature,
			&i.IsCancelled,
			&i.IsValidated,
			&i.IsFulfilled,
			&i.IsInvalid,
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