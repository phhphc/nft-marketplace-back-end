-- name: UpsertNFTV2 :exec
INSERT INTO "nfts" (token, identifier, owner, metadata, is_burned)
VALUES (sqlc.arg('token'), sqlc.arg('identifier'), sqlc.arg('owner'), sqlc.arg('metadata'), sqlc.arg('is_burned'))
ON CONFLICT (token, identifier) DO UPDATE SET
                                              owner = sqlc.arg('owner'),
                                              metadata = sqlc.arg('metadata'),
                                              is_burned = sqlc.arg('is_burned')
WHERE nfts.block_number < sqlc.arg('block_number') OR (nfts.block_number = sqlc.arg('block_number') AND nfts.tx_index < sqlc.arg('tx_index'));

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

-- name: GetNFTValidConsiderations :many
SELECT
    selected_nft.block_number,
    selected_nft.token,
    selected_nft.identifier,
    selected_nft.owner,
    selected_nft.metadata ->> 'image' AS image,
    selected_nft.metadata ->> 'name' AS name,
    selected_nft.metadata ->> 'description' AS description,
    selected_nft.metadata AS metadata,
    ci.order_hash,
    ci.item_type,
    ci.start_amount AS start_price,
    ci.end_amount AS end_price,
    o.start_time AS start_time,
    o.end_time AS end_time
FROM (
         SELECT * FROM nfts WHERE nfts.token ILIKE sqlc.arg('token') AND nfts.identifier = sqlc.arg('identifier')
     ) selected_nft
         LEFT JOIN "offer_items" oi ON oi.token ILIKE selected_nft.token AND oi.identifier = selected_nft.identifier
         LEFT JOIN "consideration_items" ci ON ci.order_hash ILIKE oi.order_hash
         LEFT JOIN (
    SELECT * FROM orders WHERE orders.is_fulfilled = FALSE AND orders.is_cancelled = FALSE
) o ON oi.order_hash ILIKE o.order_hash;

-- name: GetNFTsWithPricesPaginated :many
SELECT
    paged_nfts.block_number,
    paged_nfts.token,
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
         AND (nfts.owner ILIKE sqlc.narg('owner') OR sqlc.narg('owner') IS NULL)
         AND (nfts.token ILIKE sqlc.narg('token') OR sqlc.narg('token') IS NULL)
         LIMIT sqlc.arg('limit')
         OFFSET sqlc.arg('offset')
     ) AS paged_nfts
         LEFT JOIN offer_items oi
                   ON paged_nfts.token ILIKE oi.token AND paged_nfts.identifier = oi.identifier
         LEFT JOIN consideration_items ci ON oi.order_hash ILIKE ci.order_hash
         LEFT JOIN (
    SELECT * FROM orders WHERE orders.is_fulfilled = FALSE AND orders.is_cancelled = FALSE
) o ON oi.order_hash ILIKE o.order_hash
ORDER BY paged_nfts.block_number, paged_nfts.tx_index, ci.id, paged_nfts.token, paged_nfts.identifier;
