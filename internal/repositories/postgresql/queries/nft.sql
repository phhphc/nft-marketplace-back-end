-- name: UpsertNft :exec
INSERT INTO "nfts" (token_id, contract_addr, owner, is_burned, metadata, block_number, tx_index)
VALUES ($1,$2,$3,$4,$5,$6,$7)
ON CONFLICT (token_id, contract_addr) DO UPDATE
SET owner=$3, is_burned=$4, metadata=$5, block_number=$6, tx_index=$7
WHERE $6 > nfts.block_number OR ($6 = nfts.block_number AND $7 > nfts.tx_index);

-- name: SearchValidListingNft :many
SELECT n.token_id, n.contract_addr, n.owner, n.is_burned, n.metadata, l.listing_id, l.seller, l.price
FROM "nfts" n
LEFT JOIN "listings" l ON n.token_id = l.token_id AND n.contract_addr = l.contract_addr
WHERE (l.status = 'Open' OR l.listing_id IS NULL)
AND (n.contract_addr ILIKE sqlc.narg('contract_addr') OR sqlc.narg('contract_addr') IS NULL)
AND (n.owner ILIKE sqlc.narg('owner') OR l.seller ILIKE sqlc.narg('owner') OR sqlc.narg('owner') IS NULL)
OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');

-- name: GetNft :one
SELECT n.token_id, n.contract_addr, n.owner, n.is_burned, n.metadata, l.listing_id, l.seller, l.price
FROM "nfts" n
LEFT JOIN "listings" l ON n.token_id = l.token_id AND n.contract_addr = l.contract_addr
WHERE (n.contract_addr ILIKE sqlc.narg('contract_addr')) AND (n.token_id = CAST(sqlc.narg('token_id') AS NUMERIC));

