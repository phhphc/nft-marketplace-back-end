-- name: InsertOrder :exec
INSERT INTO "orders" ("order_hash", "offerer", "recipient", "salt", "start_time",
                      "end_time",
                      "signature", "is_validated", "is_cancelled", "is_fulfilled")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);


-- name: InsertOrderOfferItem :exec
INSERT INTO "offer_items" ("order_hash", "item_type", "token", "identifier", "amount", "start_amount", "end_amount")
VALUES ($1, $2, $3, $4, $5, $6, $7);


-- name: InsertOrderConsiderationItem :exec
INSERT INTO "consideration_items" ("order_hash", "item_type", "token", "identifier", "amount", "start_amount",
                                   "end_amount",
                                   "recipient")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: UpdateOrderStatus :exec
UPDATE "orders"
SET "is_validated" = COALESCE(sqlc.narg('is_validated'), "is_validated"),
    "is_cancelled" = COALESCE(sqlc.narg('is_cancelled'), "is_cancelled"),
    "is_fulfilled" = COALESCE(sqlc.narg('is_fulfilled'), "is_fulfilled"),
    "is_invalid"   = COALESCE(sqlc.narg('is_invalid'), "is_invalid")
WHERE "order_hash" = COALESCE(sqlc.narg('order_hash'), "order_hash")
  AND "offerer" = COALESCE(sqlc.narg('offerer'), "offerer")
RETURNING "order_hash";

-- name: GetOrder :many
SELECT json_build_object(
               'orderHash', o.order_hash,
               'offerer', o.offerer,
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
                         AND o.offerer ILIKE COALESCE(sqlc.narg('offerer'), o.offerer)
GROUP BY o.order_hash, o.offerer, o.signature, o.start_time, o.end_time, o.salt, o.is_fulfilled, o.is_cancelled,
         o.is_invalid;

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

-- name: GetExpiredOrder :many
SELECT DISTINCT e.name, o.order_hash, o.end_time, o.is_cancelled, o.is_invalid, o.offerer
FROM events e
JOIN orders o ON e.order_hash = o.order_hash
WHERE (e.name = 'listing' OR e.name = 'offer')
AND o.is_cancelled = false
AND o.is_invalid = false
AND o.end_time < round(EXTRACT(EPOCH FROM now()));