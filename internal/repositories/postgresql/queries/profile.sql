-- name: UpsertProfile :one
INSERT INTO "profiles" ("address", "username", "metadata", "signature")
VALUES (sqlc.arg('address'), sqlc.arg('username'), sqlc.arg('metadata'), sqlc.arg('signature'))
ON CONFLICT ("address") DO UPDATE SET
  "username" = sqlc.arg('username'),
  "metadata" = sqlc.arg('metadata'),
  "signature" = sqlc.arg('signature')
RETURNING *;

-- name: DeleteProfile :exec
DELETE FROM "profiles"
WHERE "address" = sqlc.arg('address');

-- name: GetProfile :one
SELECT "address", "username", "metadata", "signature"
FROM "profiles"
WHERE "address" = sqlc.arg('address')
LIMIT 1;

-- name: GetOffer :many
SELECT e.name, e.token, e.token_id, e.quantity,
  CAST(n.metadata ->> 'image' AS VARCHAR) AS nft_image, CAST(n.metadata ->> 'name' AS VARCHAR) AS nft_name,
	e.type, o.order_hash, e.price, n.owner, e.from, o.start_time, o.end_time
FROM "events" e 
JOIN "nfts" n ON e.token = n.token AND e.token_id = CAST(n.identifier AS varchar(78))
LEFT JOIN "orders" o ON e.order_hash = o.order_hash
WHERE e.name ILIKE 'offer'
AND o.start_time <= round(EXTRACT(EPOCH FROM now()))
AND o.end_time >= round(EXTRACT(EPOCH FROM now()))
AND o.is_fulfilled = FALSE
AND o.is_cancelled = FALSE
AND o.is_invalid = FALSE
AND (n.owner ILIKE sqlc.narg('owner') OR sqlc.narg('owner') IS NULL)
AND (e.from ILIKE sqlc.narg('from') OR sqlc.narg('from') IS NULL);