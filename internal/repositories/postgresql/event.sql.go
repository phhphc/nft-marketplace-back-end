// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: event.sql

package postgresql

import (
	"context"
	"database/sql"
)

const getEvent = `-- name: GetEvent :many
SELECT e.name, e.token, e.token_id, e.quantity, e.type, e.price, e.from, e.to, e.date, e.link,
    CAST(n.metadata ->> 'image' AS VARCHAR) AS nft_image,
	CAST(n.metadata ->> 'name' AS VARCHAR) AS nft_name,
    o.end_time, o.is_cancelled, o.is_fulfilled
FROM "events" e 
JOIN "nfts" n ON e.token = n.token AND e.token_id = CAST(n.identifier AS varchar(78))
LEFT JOIN "orders" o ON e.order_hash = o.order_hash
WHERE (e.name ILIKE $1 OR $1 IS NULL)
AND (e.token ILIKE $2 OR $2 IS NULL)
AND (e.token_id ILIKE $3 OR $3 IS NULL)
AND (e.type ILIKE $4 OR $4 IS NULL)
AND ((e.from ILIKE $5 OR $5 IS NULL) OR (e.to ILIKE $5 OR $5 IS NULL))
AND (extract(month from e.date) = $6::int OR $6::int IS NULL)
AND (extract(year from e.date) = $7::int OR $7::int IS NULL)
`

type GetEventParams struct {
	Name    sql.NullString
	Token   sql.NullString
	TokenID sql.NullString
	Type    sql.NullString
	Address sql.NullString
	Month   sql.NullInt32
	Year    sql.NullInt32
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
	Link        sql.NullString
	NftImage    string
	NftName     string
	EndTime     sql.NullString
	IsCancelled sql.NullBool
	IsFulfilled sql.NullBool
}

func (q *Queries) GetEvent(ctx context.Context, arg GetEventParams) ([]GetEventRow, error) {
	rows, err := q.db.QueryContext(ctx, getEvent,
		arg.Name,
		arg.Token,
		arg.TokenID,
		arg.Type,
		arg.Address,
		arg.Month,
		arg.Year,
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
			&i.Link,
			&i.NftImage,
			&i.NftName,
			&i.EndTime,
			&i.IsCancelled,
			&i.IsFulfilled,
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

const insertEvent = `-- name: InsertEvent :one
INSERT INTO "events" ("name", "token", "token_id", "quantity", "type", "price", "from", "to", "link", "order_hash")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
RETURNING id, name, token, token_id, quantity, type, price, "from", "to", date, link, order_hash
`

type InsertEventParams struct {
	Name      string
	Token     string
	TokenID   string
	Quantity  sql.NullInt32
	Type      sql.NullString
	Price     sql.NullString
	From      string
	To        sql.NullString
	Link      sql.NullString
	OrderHash sql.NullString
}

func (q *Queries) InsertEvent(ctx context.Context, arg InsertEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, insertEvent,
		arg.Name,
		arg.Token,
		arg.TokenID,
		arg.Quantity,
		arg.Type,
		arg.Price,
		arg.From,
		arg.To,
		arg.Link,
		arg.OrderHash,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.TokenID,
		&i.Quantity,
		&i.Type,
		&i.Price,
		&i.From,
		&i.To,
		&i.Date,
		&i.Link,
		&i.OrderHash,
	)
	return i, err
}
