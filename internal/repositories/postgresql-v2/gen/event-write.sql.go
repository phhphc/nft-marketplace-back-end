// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: event-write.sql

package gen

import (
	"context"
	"database/sql"
)

const insertEvent = `-- name: InsertEvent :one
INSERT INTO "events" ("name", "token", "token_id", "quantity", "type", "price", "from", "to", "tx_hash", "order_hash")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
RETURNING id, name, token, token_id, quantity, type, price, "from", "to", date, tx_hash, order_hash
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
	TxHash    sql.NullString
	OrderHash sql.NullString
}

func (q *Queries) InsertEvent(ctx context.Context, arg InsertEventParams) (Event, error) {
	row := q.queryRow(ctx, q.insertEventStmt, insertEvent,
		arg.Name,
		arg.Token,
		arg.TokenID,
		arg.Quantity,
		arg.Type,
		arg.Price,
		arg.From,
		arg.To,
		arg.TxHash,
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
		&i.TxHash,
		&i.OrderHash,
	)
	return i, err
}
