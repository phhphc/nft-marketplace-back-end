-- name: GetJsonOrder :one
SELECT
    json_build_object(
        'order_hash', o.order_hash,
        'offerer', o.offerer,
        'zone' ,o.zone,
        'offer', json_agg(
            json_build_object(
                'token', offer.token_address,
                'identifier', offer.token_id,
                'start_amount', offer.start_amount,
                'end_amount', offer.end_amount
                )
            ),
        'consideration', json_agg(
            json_build_object(
                'token', cons.token_address,
                'identifier', cons.token_id,
                'start_amount', cons.start_amount,
                'end_amount', cons.end_amount,
                'recipient', cons.recipient
                )
            ),
        'signature', o.signature,
        'order_type', o.order_type,
        'start_time', o.start_time,
        'end_time', o.end_time,
        'salt', o.salt,
        'counter', o.counter
    )
FROM marketplace_order o
JOIN marketplace_order_offer offer ON o.order_hash = offer.order_hash
JOIN marketplace_order_consideration cons ON o.order_hash = cons.order_hash
WHERE o.order_hash = $1
GROUP BY o.order_hash, o.offerer, o.zone, o.order_type, o.start_time, o.end_time, o.salt, o.counter;

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
    o.zone_hash
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

-- name: GetOrderHashByItemConsideration :many
SELECT DISTINCT
    order_hash
FROM marketplace_order_consideration
WHERE token_address = $1 AND token_id = $2;

-- name: GetOrderHashByItemOffer :many
SELECT DISTINCT
    order_hash
FROM marketplace_order_offer
WHERE token_address = $1 AND token_id = $2;