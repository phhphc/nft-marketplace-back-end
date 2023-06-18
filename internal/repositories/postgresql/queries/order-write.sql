-- name: InsertOrder :one
INSERT INTO "orders" ("order_hash",
                      "offerer",
                      "recipient",
                      "salt",
                      "start_time",
                      "end_time",
                      "signature",
                      "is_validated",
                      "is_cancelled",
                      "is_fulfilled")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;


-- name: InsertOrderOfferItem :one
INSERT INTO "offer_items" ("order_hash",
                           "item_type",
                           "token",
                           "identifier",
                           "amount",
                           "start_amount",
                           "end_amount")
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;


-- name: InsertOrderConsiderationItem :one
INSERT INTO "consideration_items" ("order_hash",
                                   "item_type",
                                   "token",
                                   "identifier",
                                   "amount",
                                   "start_amount",
                                   "end_amount",
                                   "recipient")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateOrderStatus :many
UPDATE "orders"
SET "is_validated" = COALESCE(sqlc.narg('is_validated'), "is_validated"),
    "is_cancelled" = COALESCE(sqlc.narg('is_cancelled'), "is_cancelled"),
    "is_fulfilled" = COALESCE(sqlc.narg('is_fulfilled'), "is_fulfilled"),
    "is_invalid"   = COALESCE(sqlc.narg('is_invalid'), "is_invalid")
WHERE "order_hash" = COALESCE(sqlc.narg('order_hash'), "order_hash")
  AND "offerer" = COALESCE(sqlc.narg('offerer'), "offerer")
RETURNING *;

-- name: UpdateOrderStatusByOffer :many
UPDATE orders o
SET "is_validated" = COALESCE(sqlc.narg('is_validated'), "is_validated"),
    "is_cancelled" = COALESCE(sqlc.narg('is_cancelled'), "is_cancelled"),
    "is_fulfilled" = COALESCE(sqlc.narg('is_fulfilled'), "is_fulfilled"),
    "is_invalid"   = COALESCE(sqlc.narg('is_invalid'), "is_invalid")
WHERE o.order_hash IN (SELECT DISTINCT o.order_hash
                       FROM orders o
                                JOIN offer_items oi on o.order_hash = oi.order_hash
                       WHERE o.is_invalid = false
                         AND o.offerer = @offerer
                         AND oi.token = @token
                         AND oi.identifier = @identifier)
RETURNING *;