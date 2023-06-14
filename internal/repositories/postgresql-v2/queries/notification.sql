-- name: InsertNotification :one
INSERT INTO "notifications" ("info", "event_name", "order_hash", "address")
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetNotification :many
SELECT n.is_viewed, n.info, n.event_name, n.order_hash, n.address,
    e.token, e.token_id, e.quantity, e.type, e.price, e.from, e.to, e.date,
    nft.owner,
    CAST(nft.metadata ->> 'image' AS VARCHAR) AS nft_image,
	CAST(nft.metadata ->> 'name' AS VARCHAR) AS nft_name
FROM "notifications" n
JOIN "events" e ON n.event_name = e.name AND n.order_hash = e.order_hash
JOIN "nfts" nft ON e.token = nft.token AND e.token_id = CAST(nft.identifier AS varchar(78))
WHERE (n.address = sqlc.narg('address') OR sqlc.narg('address') IS NULL)
AND (n.is_viewed = sqlc.narg('is_viewed') OR sqlc.narg('is_viewed') IS NULL);

-- name: UpdateNotification :one
UPDATE "notifications"
SET "is_viewed" = true
WHERE "event_name" = sqlc.narg('event_name')
AND "order_hash" = sqlc.narg('order_hash')
RETURNING *;