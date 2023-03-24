-- name: InsertOrder :exec
INSERT INTO "orders" ("order_hash", "offerer","recipient", "zone", "order_type", "zone_hash", "salt", "start_time", "end_time",
                      "signature", "is_validated", "is_cancelled", "is_fulfilled")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);


-- name: InsertOrderOfferItem :exec
INSERT INTO "offer_items" ("order_hash", "item_type", "token", "identifier","amount", "start_amount", "end_amount")
VALUES ($1, $2, $3, $4, $5, $6,$7);


-- name: InsertOrderConsiderationItem :exec
INSERT INTO "consideration_items" ("order_hash", "item_type", "token", "identifier","amount", "start_amount", "end_amount",
                                   "recipient")
VALUES ($1, $2, $3, $4, $5, $6, $7,$8);

-- name: UpdateOrderStatus :one
UPDATE "orders"
SET
    "is_validated" = COALESCE(sqlc.narg('is_validated'), "is_validated"),
    "is_cancelled" = COALESCE(sqlc.narg('is_cancelled'), "is_cancelled"),
    "is_fulfilled" = COALESCE(sqlc.narg('is_fulfilled'), "is_fulfilled")
WHERE "order_hash" = @order_hash
RETURNING "order_hash";