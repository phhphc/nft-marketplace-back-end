-- name: GetOrder :one
SELECT
    o.order_hash,
    o.offerer,
    o.zone,
    o.is_cancelled,
    o.is_validated,
    o.signature,
    o.order_type,
    o.start_time,
    o.end_time,
    o.salt,
    o.counter,
    o.zone,
    o.zone_hash,
    o.created_at,
    o.modified_at
FROM marketplace_order o
WHERE o.order_hash = $1;

-- name: GetOrderOffer :many
SELECT
    offer.type_number,
    offer.token_address,
    offer.token_id,
    offer.start_amount,
    offer.end_amount
FROM marketplace_order_offer offer
WHERE offer.order_hash = $1;

-- name: GetOrderConsideration :many
SELECT
    cons.type_number,
    cons.token_address,
    cons.token_id,
    cons.start_amount,
    cons.end_amount,
    cons.recipient
FROM marketplace_order_consideration cons
WHERE cons.order_hash = $1;

-- name: GetOrderHashByConsiderationItem :many
SELECT DISTINCT
    c.order_hash,
    o.created_at
FROM marketplace_order_consideration c
LEFT JOIN (
    SELECT created_at, order_hash, is_cancelled FROM marketplace_order
    ) o
ON c.order_hash = o.order_hash
WHERE c.token_address = sqlc.arg('token_address')
AND c.token_id = sqlc.arg('token_id')
AND (is_cancelled = sqlc.narg('is_cancelled') OR is_cancelled IS NULL)
ORDER BY o.created_at DESC;

-- name: GetOrderHashByOfferItem :many
SELECT DISTINCT
    offer.order_hash,
    o.created_at
FROM marketplace_order_offer offer
LEFT JOIN (
    SELECT created_at, order_hash, is_cancelled FROM marketplace_order
    ) o
ON offer.order_hash = o.order_hash
WHERE offer.token_address = sqlc.arg('token_address')
AND offer.token_id = sqlc.arg('token_id')
AND (is_cancelled = sqlc.narg('is_cancelled') OR is_cancelled IS NULL)
ORDER BY o.created_at DESC;