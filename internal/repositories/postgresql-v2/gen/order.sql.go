// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: order.sql

package gen

import (
	"context"
	"database/sql"
	"encoding/json"
)

const getExpiredOrder = `-- name: GetExpiredOrder :many
SELECT DISTINCT e.name, o.order_hash, o.end_time, o.is_cancelled, o.is_invalid, o.offerer
FROM events e
JOIN orders o ON e.order_hash = o.order_hash
WHERE (e.name = 'listing' OR e.name = 'offer')
AND o.is_cancelled = false
AND o.is_invalid = false
AND o.end_time < round(EXTRACT(EPOCH FROM now()))
`

type GetExpiredOrderRow struct {
	Name        string
	OrderHash   string
	EndTime     sql.NullString
	IsCancelled bool
	IsInvalid   bool
	Offerer     string
}

func (q *Queries) GetExpiredOrder(ctx context.Context) ([]GetExpiredOrderRow, error) {
	rows, err := q.query(ctx, q.getExpiredOrderStmt, getExpiredOrder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetExpiredOrderRow{}
	for rows.Next() {
		var i GetExpiredOrderRow
		if err := rows.Scan(
			&i.Name,
			&i.OrderHash,
			&i.EndTime,
			&i.IsCancelled,
			&i.IsInvalid,
			&i.Offerer,
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

const getOrder = `-- name: GetOrder :many
SELECT json_build_object(
               'orderHash', o.order_hash,
               'offerer', o.offerer,
               'offer',
               (SELECT json_agg(
                               json_build_object(
                                       'itemType', offer.item_type,
                                       'token', offer.token,
                                       'identifier', offer.identifier::VARCHAR,
                                       'startAmount', offer.start_amount::VARCHAR,
                                       'endAmount', offer.end_amount::VARCHAR
                                   )
                           )
                FROM offer_items offer
                WHERE o.order_hash = offer.order_hash),
               'consideration',
               (SELECT json_agg(
                               json_build_object(
                                       'itemType', cons.item_type,
                                       'token', cons.token,
                                       'identifier', cons.identifier::VARCHAR,
                                       'startAmount', cons.start_amount::VARCHAR,
                                       'endAmount', cons.end_amount::VARCHAR,
                                       'recipient', cons.recipient
                                   )
                           )
                FROM consideration_items cons
                WHERE o.order_hash = cons.order_hash),
               'signature', o.signature,
               'startTime', o.start_time::VARCHAR,
               'endTime', o.end_time::VARCHAR,
               'salt', o.salt,
               'status', json_build_object(
                       'isFulfilled', o.is_fulfilled,
                       'isCancelled', o.is_cancelled,
                       'isInvalid', o.is_invalid
                   )
           )
FROM orders o
WHERE o.order_hash in (SELECT DISTINCT o.order_hash
                       FROM orders o
                                JOIN consideration_items ci on ci.order_hash = o.order_hash
                                JOIN offer_items oi on oi.order_hash = o.order_hash
                       WHERE (o.order_hash ILIKE $1 OR $1 IS NULL)
                         AND (o.is_cancelled = $2 OR $2 IS NULL)
                         AND (o.is_fulfilled = $3 OR $3 IS NULL)
                         AND (o.is_invalid = $4 OR $4 IS NULL)
                         AND (ci.token ILIKE $5 OR
                              $5 IS NULL)
                         AND (ci.identifier = $6 OR
                              $6 IS NULL)
                         AND (oi.token ILIKE $7 OR $7 IS NULL)
                         AND (oi.identifier = $8 OR $8 IS NULL))
                         AND o.offerer ILIKE COALESCE($9, o.offerer)
GROUP BY o.order_hash
`

type GetOrderParams struct {
	OrderHash               sql.NullString
	IsCancelled             sql.NullBool
	IsFulfilled             sql.NullBool
	IsInvalid               sql.NullBool
	ConsiderationToken      sql.NullString
	ConsiderationIdentifier sql.NullString
	OfferToken              sql.NullString
	OfferIdentifier         sql.NullString
	Offerer                 sql.NullString
}

func (q *Queries) GetOrder(ctx context.Context, arg GetOrderParams) ([]json.RawMessage, error) {
	rows, err := q.query(ctx, q.getOrderStmt, getOrder,
		arg.OrderHash,
		arg.IsCancelled,
		arg.IsFulfilled,
		arg.IsInvalid,
		arg.ConsiderationToken,
		arg.ConsiderationIdentifier,
		arg.OfferToken,
		arg.OfferIdentifier,
		arg.Offerer,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []json.RawMessage{}
	for rows.Next() {
		var json_build_object json.RawMessage
		if err := rows.Scan(&json_build_object); err != nil {
			return nil, err
		}
		items = append(items, json_build_object)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertOrder = `-- name: InsertOrder :exec
INSERT INTO "orders" ("order_hash", "offerer", "recipient", "salt", "start_time",
                      "end_time",
                      "signature", "is_validated", "is_cancelled", "is_fulfilled")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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

func (q *Queries) InsertOrder(ctx context.Context, arg InsertOrderParams) error {
	_, err := q.exec(ctx, q.insertOrderStmt, insertOrder,
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
	return err
}

const insertOrderConsiderationItem = `-- name: InsertOrderConsiderationItem :exec
INSERT INTO "consideration_items" ("order_hash", "item_type", "token", "identifier", "amount", "start_amount",
                                   "end_amount",
                                   "recipient")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
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

func (q *Queries) InsertOrderConsiderationItem(ctx context.Context, arg InsertOrderConsiderationItemParams) error {
	_, err := q.exec(ctx, q.insertOrderConsiderationItemStmt, insertOrderConsiderationItem,
		arg.OrderHash,
		arg.ItemType,
		arg.Token,
		arg.Identifier,
		arg.Amount,
		arg.StartAmount,
		arg.EndAmount,
		arg.Recipient,
	)
	return err
}

const insertOrderOfferItem = `-- name: InsertOrderOfferItem :exec
INSERT INTO "offer_items" ("order_hash", "item_type", "token", "identifier", "amount", "start_amount", "end_amount")
VALUES ($1, $2, $3, $4, $5, $6, $7)
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

func (q *Queries) InsertOrderOfferItem(ctx context.Context, arg InsertOrderOfferItemParams) error {
	_, err := q.exec(ctx, q.insertOrderOfferItemStmt, insertOrderOfferItem,
		arg.OrderHash,
		arg.ItemType,
		arg.Token,
		arg.Identifier,
		arg.Amount,
		arg.StartAmount,
		arg.EndAmount,
	)
	return err
}

const markOrderInvalid = `-- name: MarkOrderInvalid :exec
UPDATE orders o
SET is_invalid = true
WHERE o.order_hash in (SELECT DISTINCT o.order_hash
                       FROM orders o
                                JOIN offer_items oi on o.order_hash = oi.order_hash
                       WHERE o.is_invalid = false
                         AND o.offerer = $1
                         AND oi.token = $2
                         AND oi.identifier = $3)
`

type MarkOrderInvalidParams struct {
	Offerer    string
	Token      string
	Identifier string
}

func (q *Queries) MarkOrderInvalid(ctx context.Context, arg MarkOrderInvalidParams) error {
	_, err := q.exec(ctx, q.markOrderInvalidStmt, markOrderInvalid, arg.Offerer, arg.Token, arg.Identifier)
	return err
}

const updateOrderStatus = `-- name: UpdateOrderStatus :exec
UPDATE "orders"
SET "is_validated" = COALESCE($1, "is_validated"),
    "is_cancelled" = COALESCE($2, "is_cancelled"),
    "is_fulfilled" = COALESCE($3, "is_fulfilled"),
    "is_invalid"   = COALESCE($4, "is_invalid")
WHERE "order_hash" = COALESCE($5, "order_hash")
  AND "offerer" = COALESCE($6, "offerer")
RETURNING "order_hash"
`

type UpdateOrderStatusParams struct {
	IsValidated sql.NullBool
	IsCancelled sql.NullBool
	IsFulfilled sql.NullBool
	IsInvalid   sql.NullBool
	OrderHash   sql.NullString
	Offerer     sql.NullString
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) error {
	_, err := q.exec(ctx, q.updateOrderStatusStmt, updateOrderStatus,
		arg.IsValidated,
		arg.IsCancelled,
		arg.IsFulfilled,
		arg.IsInvalid,
		arg.OrderHash,
		arg.Offerer,
	)
	return err
}
