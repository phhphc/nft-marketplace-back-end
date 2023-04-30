-- name: InsertOrder :exec
INSERT INTO "orders" ("order_hash", "offerer", "recipient", "zone", "order_type", "zone_hash", "salt", "start_time",
                      "end_time",
                      "signature", "is_validated", "is_cancelled", "is_fulfilled")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);


-- name: InsertOrderOfferItem :exec
INSERT INTO "offer_items" ("order_hash", "item_type", "token", "identifier", "amount", "start_amount", "end_amount")
VALUES ($1, $2, $3, $4, $5, $6, $7);


-- name: InsertOrderConsiderationItem :exec
INSERT INTO "consideration_items" ("order_hash", "item_type", "token", "identifier", "amount", "start_amount",
                                   "end_amount",
                                   "recipient")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: UpdateOrderStatus :one
UPDATE "orders"
SET "is_validated" = COALESCE(sqlc.narg('is_validated'), "is_validated"),
    "is_cancelled" = COALESCE(sqlc.narg('is_cancelled'), "is_cancelled"),
    "is_fulfilled" = COALESCE(sqlc.narg('is_fulfilled'), "is_fulfilled")
WHERE "order_hash" = @order_hash
RETURNING "order_hash";

-- name: GetJsonOrderByHash :one
SELECT json_build_object(
               'order_hash', o.order_hash,
               'offerer', o.offerer,
               'zone', o.zone,
               'offer', json_agg(
                       json_build_object(
                               'item_type', offer.item_type,
                               'token', offer.token,
                               'identifier', offer.identifier::VARCHAR,
                               'start_amount', offer.start_amount::VARCHAR,
                               'end_amount', offer.end_amount::VARCHAR
                           )
                   ),
               'consideration', json_agg(
                       json_build_object(
                               'item_type', cons.item_type,
                               'token', cons.token,
                               'identifier', cons.identifier::VARCHAR,
                               'start_amount', cons.start_amount::VARCHAR,
                               'end_amount', cons.end_amount::VARCHAR,
                               'recipient', cons.recipient
                           )
                   ),
               'order_type', o.order_type,
               'zone_hash', o.zone_hash,
               'signature', o.signature,
               'start_time', o.start_time::VARCHAR,
               'end_time', o.end_time::VARCHAR,
               'salt', o.salt
           )
FROM orders o
         JOIN offer_items offer ON o.order_hash = offer.order_hash
         JOIN consideration_items cons ON o.order_hash = cons.order_hash
WHERE o.order_hash ILIKE $1
  AND o.is_fulfilled = false
  AND o.is_cancelled = false
  AND o.is_invalid = false
GROUP BY o.order_hash, o.offerer, o.zone, o.order_type, o.start_time, o.end_time;

-- name: GetOrderHash :many
SELECT DISTINCT o.order_hash
FROM orders o
         JOIN consideration_items ci on ci.order_hash = o.order_hash
         JOIN offer_items oi on oi.order_hash = o.order_hash
WHERE o.is_cancelled = false
  AND o.is_fulfilled = false
  AND o.is_invalid = false
  AND (ci.token ILIKE sqlc.narg('consideration_token') OR sqlc.narg('consideration_token') IS NULL)
  AND (ci.identifier = sqlc.narg('consideration_identifier') OR sqlc.narg('consideration_identifier') IS NULL)
  AND (oi.token ILIKE sqlc.narg('offer_token') OR sqlc.narg('offer_token') IS NULL)
  AND (oi.identifier = sqlc.narg('offer_identifier') OR sqlc.narg('offer_identifier') IS NULL);

-- name: GetOrder :many
SELECT json_build_object(
               'orderHash', o.order_hash,
               'offerer', o.offerer,
               'zone', o.zone,
               'offer',
               (SELECT json_agg(
                               json_build_object(
                                       'itemType', offer.item_type,
                                       'token', offer.token,
                                       'identifier', offer.identifier::VARCHAR,
                                       'startAmount', offer.start_amount::VARCHAR,
                                       'endAmount', offer.end_amount::VARCHAR
                                   )
                           )
                FROM offer_items offer
                WHERE o.order_hash = offer.order_hash),
               'consideration',
               (SELECT json_agg(
                               json_build_object(
                                       'itemType', cons.item_type,
                                       'token', cons.token,
                                       'identifier', cons.identifier::VARCHAR,
                                       'startAmount', cons.start_amount::VARCHAR,
                                       'endAmount', cons.end_amount::VARCHAR,
                                       'recipient', cons.recipient
                                   )
                           )
                FROM consideration_items cons
                WHERE o.order_hash = cons.order_hash),
               'orderType', o.order_type,
               'zoneHash', o.zone_hash,
               'signature', o.signature,
               'startTime', o.start_time::VARCHAR,
               'endTime', o.end_time::VARCHAR,
               'salt', o.salt,
               'status', json_build_object(
                       'isFulfilled', o.is_fulfilled,
                       'isCancelled', o.is_cancelled,
                       'isInvalid', o.is_invalid
                   )
           )
FROM orders o
WHERE o.order_hash in (SELECT DISTINCT o.order_hash
                       FROM orders o
                                JOIN consideration_items ci on ci.order_hash = o.order_hash
                                JOIN offer_items oi on oi.order_hash = o.order_hash
                       WHERE (o.order_hash ILIKE sqlc.narg('order_hash') OR sqlc.narg('order_hash') IS NULL)
                         AND (o.is_cancelled = sqlc.narg('is_cancelled') OR sqlc.narg('is_cancelled') IS NULL)
                         AND (o.is_fulfilled = sqlc.narg('is_fulfilled') OR sqlc.narg('is_fulfilled') IS NULL)
                         AND (o.is_invalid = sqlc.narg('is_invalid') OR sqlc.narg('is_invalid') IS NULL)
                         AND (ci.token ILIKE sqlc.narg('consideration_token') OR
                              sqlc.narg('consideration_token') IS NULL)
                         AND (ci.identifier = sqlc.narg('consideration_identifier') OR
                              sqlc.narg('consideration_identifier') IS NULL)
                         AND (oi.token ILIKE sqlc.narg('offer_token') OR sqlc.narg('offer_token') IS NULL)
                         AND (oi.identifier = sqlc.narg('offer_identifier') OR sqlc.narg('offer_identifier') IS NULL))
GROUP BY o.order_hash;

-- name: MarkOrderInvalid :exec
UPDATE orders o
SET is_invalid = true
WHERE o.order_hash in (SELECT DISTINCT o.order_hash
                       FROM orders o
                                JOIN offer_items oi on o.order_hash = oi.order_hash
                       WHERE o.is_invalid = false
                         AND o.offerer = $1
                         AND oi.token = $2
                         AND oi.identifier = $3);