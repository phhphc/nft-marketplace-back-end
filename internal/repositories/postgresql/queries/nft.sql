-- name: UpsertNft :exec
INSERT INTO "nfts" (token_id, contract_addr, owner, is_burned, metadata, block_number, tx_index)
VALUES ($1,$2,$3,$4,$5,$6,$7)
ON CONFLICT (token_id, contract_addr) DO UPDATE
SET owner=$3, is_burned=$4, metadata=$5, block_number=$6, tx_index=$7
WHERE $6 > nfts.block_number OR ($6 = nfts.block_number AND $7 > nfts.tx_index);

-- name: GetListNft :many
SELECT n.token_id, n.contract_addr, n.owner, n.is_burned, n.metadata, l.listing_id, l.seller, l.price
FROM "nfts" n
LEFT JOIN (SELECT listing_id, seller, price FROM "listings" WHERE status = 'Open') AS l USING (token_id, contract_addr)
WHERE (n.contract_addr ILIKE sqlc.narg('contract_addr') OR sqlc.narg('contract_addr') IS NULL)
AND (n.owner ILIKE sqlc.narg('owner') OR l.seller ILIKE sqlc.narg('owner') OR sqlc.narg('owner') IS NULL)
ORDER BY n.contract_addr ASC, n.token_id ASC
OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');

-- name: GetNftDetail :one
SELECT n.token_id, n.contract_addr, n.owner, n.is_burned, n.metadata, l.listing_id, l.seller, l.price
FROM "nfts" n
LEFT JOIN "listings" l ON n.token_id = l.token_id AND n.contract_addr = l.contract_addr
WHERE n.contract_addr ILIKE sqlc.narg('contract_addr') AND n.token_id = sqlc.narg('token_id');

