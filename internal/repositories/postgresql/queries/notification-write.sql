-- name: InsertNotification :one
INSERT INTO "notifications" ("info", "event_name", "order_hash", "address")
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: UpdateNotification :one
UPDATE "notifications"
SET "is_viewed" = true
WHERE "event_name" = sqlc.narg('event_name')
AND "order_hash" = sqlc.narg('order_hash')
RETURNING *;