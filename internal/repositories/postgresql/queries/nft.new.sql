-- name: GetListValidNFT :many
SELECT
    n.token, n.identifier, n.owner, n.metadata, n.is_burned
FROM "nfts" n
WHERE
    n.is_burned = FALSE
  AND
    (n.token ILIKE sqlc.narg('token') OR sqlc.narg('token') IS NULL)
  AND
    (n.owner ILIKE sqlc.narg('owner') OR sqlc.narg('owner') IS NULL)
ORDER BY n.token ASC, n.identifier ASC
OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');

-- name: GetValidNFT :one
SELECT
    n.token, n.identifier, n.owner, n.metadata, n.is_burned
FROM "nfts" n
WHERE
    n.is_burned = FALSE
  AND
    n.token = sqlc.arg('token')
  AND
    n.identifier = sqlc.arg('identifier');

-- name: UpsertNFTV2 :exec
INSERT INTO "nfts" (token, identifier, owner, metadata, is_burned)
VALUES (sqlc.arg('token'), sqlc.arg('identifier'), sqlc.arg('owner'), sqlc.arg('metadata'), sqlc.arg('is_burned'))
ON CONFLICT (token, identifier) DO UPDATE SET
    owner = sqlc.arg('owner'),
    metadata = sqlc.arg('metadata'),
    is_burned = sqlc.arg('is_burned')
WHERE nfts.block_number < sqlc.arg('block_number') OR (nfts.block_number = sqlc.arg('block_number') AND nfts.tx_index < sqlc.arg('tx_index'));

-- name: GetNFTsWithPricesPaginated :many
SELECT paged_nfts.token,
       paged_nfts.identifier,
       paged_nfts.owner,
       paged_nfts.metadata -> 'image' AS image,
       paged_nfts.metadata -> 'name' AS name,
       paged_nfts.metadata -> 'description' AS description,
       ci.order_hash,
       ci.item_type,
       ci.start_amount AS start_price,
       ci.end_amount AS end_price,
       o.start_time AS start_time,
       o.end_time AS end_time
FROM (
        SELECT * FROM nfts
        WHERE nfts.is_burned = FALSE
        AND (nfts.token ILIKE sqlc.narg('token') OR sqlc.narg('token') IS NULL)
        AND (nfts.owner ILIKE sqlc.narg('owner') OR sqlc.narg('owner') IS NULL)
        OFFSET sqlc.arg('offset') LIMIT sqlc.arg('limit')
     ) AS paged_nfts
        LEFT JOIN offer_items oi ON paged_nfts.token = oi.token AND paged_nfts.identifier = oi.identifier
        LEFT JOIN consideration_items ci ON oi.order_hash = ci.order_hash
        LEFT JOIN orders o ON oi.order_hash = o.order_hash
ORDER BY paged_nfts.token, paged_nfts.identifier, ci.order_hash;
