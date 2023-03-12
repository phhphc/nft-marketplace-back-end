// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: read.sql

package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
)

const getJsonOrder = `-- name: GetJsonOrder :one
SELECT
    json_build_object(
        'order_hash', o.order_hash,
        'offerer', o.offerer,
        'zone' ,o.zone,
        'offer', json_agg(
            json_build_object(
                'token', offer.token_address,
                'identifier', offer.token_id,
                'start_amount', offer.start_amount,
                'end_amount', offer.end_amount
                )
            ),
        'consideration', json_agg(
            json_build_object(
                'token', cons.token_address,
                'identifier', cons.token_id,
                'start_amount', cons.start_amount,
                'end_amount', cons.end_amount,
                'recipient', cons.recipient
                )
            ),
        'signature', o.signature,
        'order_type', o.order_type,
        'start_time', o.start_time,
        'end_time', o.end_time,
        'salt', o.salt,
        'counter', o.counter
    )
FROM marketplace_order o
JOIN marketplace_order_offer offer ON o.order_hash = offer.order_hash
JOIN marketplace_order_consideration cons ON o.order_hash = cons.order_hash
WHERE o.order_hash = $1
GROUP BY o.order_hash, o.offerer, o.zone, o.order_type, o.start_time, o.end_time, o.salt, o.counter
`

func (q *Queries) GetJsonOrder(ctx context.Context, orderHash string) (json.RawMessage, error) {
	row := q.db.QueryRowContext(ctx, getJsonOrder, orderHash)
	var json_build_object json.RawMessage
	err := row.Scan(&json_build_object)
	return json_build_object, err
}

const getOrder = `-- name: GetOrder :one
SELECT
    o.order_hash,
    o.offerer,
    o.zone,
    o.is_cancelled,
    o.is_validated,
    o.signature,
    o.order_type,
    o.start_time,
    o.end_time,
    o.salt,
    o.counter,
    o.zone,
    o.zone_hash
FROM marketplace_order o
WHERE o.order_hash = $1
`

type GetOrderRow struct {
	OrderHash   string
	Offerer     string
	Zone        sql.NullString
	IsCancelled bool
	IsValidated bool
	Signature   sql.NullString
	OrderType   string
	StartTime   string
	EndTime     string
	Salt        string
	Counter     string
	Zone_2      sql.NullString
	ZoneHash    sql.NullString
}

func (q *Queries) GetOrder(ctx context.Context, orderHash string) (GetOrderRow, error) {
	row := q.db.QueryRowContext(ctx, getOrder, orderHash)
	var i GetOrderRow
	err := row.Scan(
		&i.OrderHash,
		&i.Offerer,
		&i.Zone,
		&i.IsCancelled,
		&i.IsValidated,
		&i.Signature,
		&i.OrderType,
		&i.StartTime,
		&i.EndTime,
		&i.Salt,
		&i.Counter,
		&i.Zone_2,
		&i.ZoneHash,
	)
	return i, err
}

const getOrderConsideration = `-- name: GetOrderConsideration :many
SELECT
    cons.type_number,
    cons.token_address,
    cons.token_id,
    cons.start_amount,
    cons.end_amount,
    cons.recipient
FROM marketplace_order_consideration cons
WHERE cons.order_hash = $1
`

type GetOrderConsiderationRow struct {
	TypeNumber   string
	TokenAddress string
	TokenID      string
	StartAmount  string
	EndAmount    string
	Recipient    string
}

func (q *Queries) GetOrderConsideration(ctx context.Context, orderHash string) ([]GetOrderConsiderationRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrderConsideration, orderHash)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrderConsiderationRow{}
	for rows.Next() {
		var i GetOrderConsiderationRow
		if err := rows.Scan(
			&i.TypeNumber,
			&i.TokenAddress,
			&i.TokenID,
			&i.StartAmount,
			&i.EndAmount,
			&i.Recipient,
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

const getOrderHashByItemConsideration = `-- name: GetOrderHashByItemConsideration :many
SELECT DISTINCT
    order_hash
FROM marketplace_order_consideration
WHERE token_address = $1 AND token_id = $2
`

type GetOrderHashByItemConsiderationParams struct {
	TokenAddress string
	TokenID      string
}

func (q *Queries) GetOrderHashByItemConsideration(ctx context.Context, arg GetOrderHashByItemConsiderationParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getOrderHashByItemConsideration, arg.TokenAddress, arg.TokenID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var order_hash string
		if err := rows.Scan(&order_hash); err != nil {
			return nil, err
		}
		items = append(items, order_hash)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderHashByItemOffer = `-- name: GetOrderHashByItemOffer :many
SELECT DISTINCT
    order_hash
FROM marketplace_order_offer
WHERE token_address = $1 AND token_id = $2
`

type GetOrderHashByItemOfferParams struct {
	TokenAddress string
	TokenID      string
}

func (q *Queries) GetOrderHashByItemOffer(ctx context.Context, arg GetOrderHashByItemOfferParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getOrderHashByItemOffer, arg.TokenAddress, arg.TokenID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var order_hash string
		if err := rows.Scan(&order_hash); err != nil {
			return nil, err
		}
		items = append(items, order_hash)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderOffer = `-- name: GetOrderOffer :many
SELECT
    offer.type_number,
    offer.token_address,
    offer.token_id,
    offer.start_amount,
    offer.end_amount
FROM marketplace_order_offer offer
WHERE offer.order_hash = $1
`

type GetOrderOfferRow struct {
	TypeNumber   string
	TokenAddress string
	TokenID      string
	StartAmount  string
	EndAmount    string
}

func (q *Queries) GetOrderOffer(ctx context.Context, orderHash string) ([]GetOrderOfferRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrderOffer, orderHash)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrderOfferRow{}
	for rows.Next() {
		var i GetOrderOfferRow
		if err := rows.Scan(
			&i.TypeNumber,
			&i.TokenAddress,
			&i.TokenID,
			&i.StartAmount,
			&i.EndAmount,
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
