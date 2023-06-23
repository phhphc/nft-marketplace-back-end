// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: notification-write.sql

package gen

import (
	"context"
	"database/sql"
)

const insertNotification = `-- name: InsertNotification :one
INSERT INTO "notifications" ("info", "event_name", "order_hash", "address")
VALUES ($1,$2,$3,$4)
RETURNING id, info, event_name, order_hash, address, is_viewed
`

type InsertNotificationParams struct {
	Info      string
	EventName string
	OrderHash string
	Address   string
}

func (q *Queries) InsertNotification(ctx context.Context, arg InsertNotificationParams) (Notification, error) {
	row := q.queryRow(ctx, q.insertNotificationStmt, insertNotification,
		arg.Info,
		arg.EventName,
		arg.OrderHash,
		arg.Address,
	)
	var i Notification
	err := row.Scan(
		&i.ID,
		&i.Info,
		&i.EventName,
		&i.OrderHash,
		&i.Address,
		&i.IsViewed,
	)
	return i, err
}

const updateNotification = `-- name: UpdateNotification :one
UPDATE "notifications"
SET "is_viewed" = true
WHERE "event_name" = $1
AND "order_hash" = $2
RETURNING id, info, event_name, order_hash, address, is_viewed
`

type UpdateNotificationParams struct {
	EventName sql.NullString
	OrderHash sql.NullString
}

func (q *Queries) UpdateNotification(ctx context.Context, arg UpdateNotificationParams) (Notification, error) {
	row := q.queryRow(ctx, q.updateNotificationStmt, updateNotification, arg.EventName, arg.OrderHash)
	var i Notification
	err := row.Scan(
		&i.ID,
		&i.Info,
		&i.EventName,
		&i.OrderHash,
		&i.Address,
		&i.IsViewed,
	)
	return i, err
}
