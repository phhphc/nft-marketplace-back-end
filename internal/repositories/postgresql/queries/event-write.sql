-- name: InsertEvent :one
INSERT INTO "events" ("name", "token", "token_id", "quantity", "type", "price", "from", "to", "tx_hash", "order_hash")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
RETURNING *;