-- name: InsertEvent :one
INSERT INTO "events" ("name", "token", "token_id", "quantity", "price", "from", "to", "link")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: GetEvent :many
SELECT e.name, e.token, e.token_id, e.quantity, e.price, e.from, e.to, e.date, e.link
FROM "events" e
WHERE (e.name ILIKE sqlc.narg('name') OR sqlc.narg('name') IS NULL)
AND (e.token ILIKE sqlc.narg('token') OR sqlc.narg('token') IS NULL)
AND (e.token_id ILIKE sqlc.narg('token_id') OR sqlc.narg('token_id') IS NULL)
AND ((e.from ILIKE sqlc.narg('address') OR sqlc.narg('address') IS NULL) OR (e.to ILIKE sqlc.narg('address') OR sqlc.narg('address') IS NULL));