// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: event-read.sql

package gen

import (
	"context"
	"database/sql"
)

const getEvent = `-- name: GetEvent :many
SELECT e.name, e.token, e.token_id, e.quantity, e.type, e.price, e.from, e.to, e.date, e.tx_hash,
  COALESCE(n.metadata ->> 'image', '')::VARCHAR AS nft_image,
	COALESCE(n.metadata ->> 'name', '')::VARCHAR AS nft_name,
    o.end_time, o.is_cancelled, o.is_fulfilled, o.order_hash
FROM "events" e 
JOIN "nfts" n ON e.token = n.token AND e.token_id = CAST(n.identifier AS varchar(78))
LEFT JOIN "orders" o ON e.order_hash = o.order_hash
WHERE (e.name ILIKE $1 OR $1 IS NULL)
AND (e.token ILIKE $2 OR $2 IS NULL)
AND (e.token_id ILIKE $3 OR $3 IS NULL)
AND (e.type ILIKE $4 OR $4 IS NULL)
AND ((e.from ILIKE $5 OR $5 IS NULL) OR (e.to ILIKE $5 OR $5 IS NULL))
AND (e.date >= TO_TIMESTAMP($6, 'YYYY-MM-DD') OR $6 IS NULL)
AND (e.date <= TO_TIMESTAMP($7, 'YYYY-MM-DD') OR $7 IS NULL)
`

type GetEventParams struct {
	Name      sql.NullString
	Token     sql.NullString
	TokenID   sql.NullString
	Type      sql.NullString
	Address   sql.NullString
	StartDate sql.NullString
	EndDate   sql.NullString
}

type GetEventRow struct {
	Name        string
	Token       string
	TokenID     string
	Quantity    sql.NullInt32
	Type        sql.NullString
	Price       sql.NullString
	From        string
	To          sql.NullString
	Date        sql.NullTime
	TxHash      sql.NullString
	NftImage    string
	NftName     string
	EndTime     sql.NullString
	IsCancelled sql.NullBool
	IsFulfilled sql.NullBool
	OrderHash   sql.NullString
}

func (q *Queries) GetEvent(ctx context.Context, arg GetEventParams) ([]GetEventRow, error) {
	rows, err := q.query(ctx, q.getEventStmt, getEvent,
		arg.Name,
		arg.Token,
		arg.TokenID,
		arg.Type,
		arg.Address,
		arg.StartDate,
		arg.EndDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetEventRow{}
	for rows.Next() {
		var i GetEventRow
		if err := rows.Scan(
			&i.Name,
			&i.Token,
			&i.TokenID,
			&i.Quantity,
			&i.Type,
			&i.Price,
			&i.From,
			&i.To,
			&i.Date,
			&i.TxHash,
			&i.NftImage,
			&i.NftName,
			&i.EndTime,
			&i.IsCancelled,
			&i.IsFulfilled,
			&i.OrderHash,
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

const getOffer = `-- name: GetOffer :many
SELECT e.name,
 e.token, 
 e.token_id,
  e.quantity,
  COALESCE(n.metadata ->> 'image', '')::VARCHAR AS nft_image, 
  COALESCE(n.metadata ->> 'name', '')::VARCHAR AS nft_name,
	e.type,
  o.order_hash,
  e.price,
  n.owner,
  e.from,
  o.start_time,
  o.end_time,
  o.is_fulfilled, 
  o.is_cancelled, 
  (o.end_time < round(EXTRACT(EPOCH FROM now()))) as is_expired
FROM "events" e 
JOIN "nfts" n ON e.token = n.token AND e.token_id = CAST(n.identifier AS varchar(78))
LEFT JOIN "orders" o ON e.order_hash = o.order_hash
WHERE e.name ILIKE 'offer'
AND o.start_time <= round(EXTRACT(EPOCH FROM now()))
AND (n.owner ILIKE $1 OR $1 IS NULL)
AND (e.from ILIKE $2 OR $2 IS NULL)
`

type GetOfferParams struct {
	Owner sql.NullString
	From  sql.NullString
}

type GetOfferRow struct {
	Name        string
	Token       string
	TokenID     string
	Quantity    sql.NullInt32
	NftImage    string
	NftName     string
	Type        sql.NullString
	OrderHash   sql.NullString
	Price       sql.NullString
	Owner       string
	From        string
	StartTime   sql.NullString
	EndTime     sql.NullString
	IsFulfilled sql.NullBool
	IsCancelled sql.NullBool
	IsExpired   bool
}

func (q *Queries) GetOffer(ctx context.Context, arg GetOfferParams) ([]GetOfferRow, error) {
	rows, err := q.query(ctx, q.getOfferStmt, getOffer, arg.Owner, arg.From)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOfferRow{}
	for rows.Next() {
		var i GetOfferRow
		if err := rows.Scan(
			&i.Name,
			&i.Token,
			&i.TokenID,
			&i.Quantity,
			&i.NftImage,
			&i.NftName,
			&i.Type,
			&i.OrderHash,
			&i.Price,
			&i.Owner,
			&i.From,
			&i.StartTime,
			&i.EndTime,
			&i.IsFulfilled,
			&i.IsCancelled,
			&i.IsExpired,
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
