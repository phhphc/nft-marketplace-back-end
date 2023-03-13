-- name: InsertOrder :exec
INSERT INTO
    marketplace_order(order_hash,
                      offerer,
                      is_cancelled,
                      is_validated,
                      is_fulFilled,
                      signature,
                      order_type,
                      start_time,
                      end_time,
                      counter,
                      salt,
                      zone,
                      zone_hash,
                      created_at,
                      modified_at)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12,$13, now(), now());

-- name: InsertOrderOffer :exec
INSERT INTO marketplace_order_offer(order_hash, type_number, token_id, token_address, start_amount, end_amount)
VALUES($1, $2, $3, $4, $5, $6);

-- name: InsertOrderConsideration :exec
INSERT INTO marketplace_order_consideration(order_hash, type_number, token_id, token_address, start_amount, end_amount, recipient)
VALUES($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateOrderStatus :exec
UPDATE marketplace_order
SET 
    is_cancelled = coalesce(sqlc.narg(is_cancelled), is_cancelled),
    is_validated = coalesce(sqlc.narg(is_validated), is_validated),
    is_fulfilled = coalesce(sqlc.narg(is_fulfilled), is_fulfilled),
    modified_at = now()
WHERE order_hash = @order_hash;

-- name: DestroyOrders :exec
UPDATE marketplace_order
SET is_cancelled = true, modified_at = now()
WHERE counter != $2 and offerer = $1;