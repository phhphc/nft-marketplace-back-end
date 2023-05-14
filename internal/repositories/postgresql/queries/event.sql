-- name: InsertEvent :one
INSERT INTO "events" ("name", "token", "token_id", "quantity", "type", "price", "from", "to", "link", "order_hash")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
RETURNING *;

-- name: GetEvent :many
SELECT e.name, e.token, e.token_id, e.quantity, e.type, e.price, e.from, e.to, e.date, e.link,
    CAST(n.metadata ->> 'image' AS VARCHAR) AS nft_image,
	CAST(n.metadata ->> 'name' AS VARCHAR) AS nft_name,
    o.end_time, o.is_cancelled, o.is_fulfilled
FROM "events" e 
JOIN "nfts" n ON e.token = n.token AND e.token_id = CAST(n.identifier AS varchar(78))
LEFT JOIN "orders" o ON e.order_hash = o.order_hash
WHERE (e.name ILIKE sqlc.narg('name') OR sqlc.narg('name') IS NULL)
AND (e.token ILIKE sqlc.narg('token') OR sqlc.narg('token') IS NULL)
AND (e.token_id ILIKE sqlc.narg('token_id') OR sqlc.narg('token_id') IS NULL)
AND (e.type ILIKE sqlc.narg('type') OR sqlc.narg('type') IS NULL)
AND ((e.from ILIKE sqlc.narg('address') OR sqlc.narg('address') IS NULL) OR (e.to ILIKE sqlc.narg('address') OR sqlc.narg('address') IS NULL));