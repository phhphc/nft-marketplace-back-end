-- name: UpsertListing :exec
INSERT INTO "listings" (listing_id, collection, token_id, seller, price, status, block_number, tx_index)
VALUES ($1,$2,$3,$4,$5,$6, $7, $8)
ON CONFLICT (collection, token_id) DO UPDATE
SET seller=$4,price=$5,status=$6, block_number=$7, tx_index=$8
WHERE $7 > listings.block_number or ($7 = listings.block_number and $8 > listings.tx_index);