-- name: InsertOrder :exec
INSERT INTO
    marketplace_order(order_hash,
                      offerer,
                      is_cancelled,
                      is_validated,
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
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, now(), now());

-- name: InsertOrderOffer :exec
INSERT INTO marketplace_order_offer(order_hash, type_number, token_id, token_address, start_amount, end_amount)
VALUES($1, $2, $3, $4, $5, $6);

-- name: InsertOrderConsideration :exec
INSERT INTO marketplace_order_consideration(order_hash, type_number, token_id, token_address, start_amount, end_amount, recipient)
VALUES($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateOrderIsCancelled :exec
UPDATE marketplace_order
SET is_cancelled = $2 , modified_at = now()
WHERE order_hash = $1;

-- name: UpdateOrderIsValidated :exec
UPDATE marketplace_order
SET is_validated = $2 , modified_at = now()
WHERE order_hash = $1;

-- name: DestroyOrders :exec
UPDATE marketplace_order
SET is_cancelled = true, modified_at = now()
WHERE counter != $2 and offerer = $1;