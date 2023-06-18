-- name: GetEvent :many
SELECT e.name, e.token, e.token_id, e.quantity, e.type, e.price, e.from, e.to, e.date, e.tx_hash,
  COALESCE(n.metadata ->> 'image', '')::VARCHAR AS nft_image,
	COALESCE(n.metadata ->> 'name', '')::VARCHAR AS nft_name,
    o.end_time, o.is_cancelled, o.is_fulfilled, o.order_hash
FROM "events" e 
JOIN "nfts" n ON e.token = n.token AND e.token_id = CAST(n.identifier AS varchar(78))
LEFT JOIN "orders" o ON e.order_hash = o.order_hash
WHERE (e.name ILIKE sqlc.narg('name') OR sqlc.narg('name') IS NULL)
AND (e.token ILIKE sqlc.narg('token') OR sqlc.narg('token') IS NULL)
AND (e.token_id ILIKE sqlc.narg('token_id') OR sqlc.narg('token_id') IS NULL)
AND (e.type ILIKE sqlc.narg('type') OR sqlc.narg('type') IS NULL)
AND ((e.from ILIKE sqlc.narg('address') OR sqlc.narg('address') IS NULL) OR (e.to ILIKE sqlc.narg('address') OR sqlc.narg('address') IS NULL))
AND (extract(month from e.date) = sqlc.narg('month')::int OR sqlc.narg('month')::int IS NULL)
AND (extract(year from e.date) = sqlc.narg('year')::int OR sqlc.narg('year')::int IS NULL);


-- name: GetOffer :many
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
AND (n.owner ILIKE sqlc.narg('owner') OR sqlc.narg('owner') IS NULL)
AND (e.from ILIKE sqlc.narg('from') OR sqlc.narg('from') IS NULL);