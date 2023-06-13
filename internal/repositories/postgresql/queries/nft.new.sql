-- name: UpsertNFTV2 :exec
INSERT INTO "nfts" (token, identifier, owner, metadata, is_burned)
VALUES (sqlc.arg('token'), sqlc.arg('identifier'), sqlc.arg('owner'), sqlc.arg('metadata'), sqlc.arg('is_burned'))
ON CONFLICT (token, identifier) DO UPDATE SET owner     = sqlc.arg('owner'),
                                              metadata  = sqlc.arg('metadata'),
                                              is_burned = sqlc.arg('is_burned')
WHERE nfts.block_number < sqlc.arg('block_number')
   OR (nfts.block_number = sqlc.arg('block_number') AND nfts.tx_index < sqlc.arg('tx_index'));

-- name: GetListValidNFT :many
SELECT n.token,
       n.identifier,
       n.owner,
       n.metadata,
       n.is_burned,
       n.is_hidden
FROM "nfts" n
WHERE n.is_burned = FALSE
  AND (n.token ILIKE sqlc.narg('token') OR sqlc.narg('token') IS NULL)
  AND (n.owner ILIKE sqlc.narg('owner') OR sqlc.narg('owner') IS NULL)
ORDER BY n.token ASC, n.identifier ASC
OFFSET sqlc.arg('offset') LIMIT sqlc.arg('limit');

-- name: GetNFTValidConsiderations :many
SELECT selected_nft.block_number,
       selected_nft.token,
       selected_nft.identifier,
       selected_nft.owner,
       selected_nft.is_hidden,
       selected_nft.metadata ->> 'image'       AS image,
       selected_nft.metadata ->> 'name'        AS name,
       selected_nft.metadata ->> 'description' AS description,
       selected_nft.metadata                   AS metadata,
       ci.order_hash,
       ci.item_type,
       ci.start_amount                         AS start_price,
       ci.end_amount                           AS end_price,
       o.start_time                            AS start_time,
       o.end_time                              AS end_time
FROM (SELECT *
      FROM nfts
      WHERE nfts.token ILIKE sqlc.arg('token')
        AND nfts.identifier = sqlc.arg('identifier')) selected_nft
         LEFT JOIN "offer_items" oi ON oi.token ILIKE selected_nft.token AND oi.identifier = selected_nft.identifier
         LEFT JOIN "consideration_items" ci ON ci.order_hash ILIKE oi.order_hash
         LEFT JOIN (SELECT *
                    FROM orders
                    WHERE orders.is_fulfilled = FALSE
                      AND orders.is_cancelled = FALSE
                      AND orders.start_time <= round(EXTRACT(EPOCH FROM now()))
                      AND orders.end_time >= round(EXTRACT(EPOCH FROM now()))) o
                   ON oi.order_hash ILIKE o.order_hash;

-- name: ListNftWithListing :many
SELECT json_build_object(
               'token', n.token,
               'identifier', n.identifier::VARCHAR,
               'owner', n.owner,
               'metadata', n.metadata,
               'is_hidden', n.is_hidden,
               'listing',
               (SELECT json_agg(
                               json_build_object(
                                       'order_hash', l.order_hash,
                                       'item_type', l.item_type,
                                       'start_time', l.start_time::VARCHAR,
                                       'end_time', l.end_time::VARCHAR,
                                       'start_price', l.start_price::VARCHAR,
                                       'end_price', l.end_price::VARCHAR
                                   )
                           )
                FROM (SELECT o.order_hash,
                             ci.item_type,
                             o.start_time         AS start_time,
                             o.end_time           AS end_time,
                             SUM(ci.start_amount) AS start_price,
                             SUM(ci.end_amount)   AS end_price
                      FROM orders o
                               JOIN offer_items oi on o.order_hash = oi.order_hash
                               JOIN consideration_items ci on o.order_hash = ci.order_hash
                      WHERE o.order_hash NOT IN (SELECT DISTINCT c.order_hash
                                                 FROM consideration_items c
                                                 WHERE c.item_type != @item_type)
                        AND o.is_fulfilled = FALSE
                        AND o.is_cancelled = FALSE
                        AND o.is_invalid = FALSE
                        AND o.start_time <= sqlc.arg('now')
                        AND o.end_time > sqlc.arg('now')
                        AND oi.token ILIKE n.token
                        AND oi.identifier = n.identifier
                      GROUP BY o.order_hash,
                               ci.item_type,
                               o.start_time,
                               o.end_time
                      LIMIT sqlc.arg('limit_listing')) as l)
           )
FROM nfts n
WHERE n."is_burned" = FALSE
  AND n."is_hidden" = COALESCE(sqlc.narg('is_hidden'), n."is_hidden")
  AND n."owner" ILIKE COALESCE(sqlc.narg('owner'), n."owner")
  AND n."token" ILIKE COALESCE(sqlc.narg('token'), n."token")
  AND n."identifier" = COALESCE(sqlc.narg('identifier'), n."identifier")
GROUP BY n.token, n.identifier
LIMIT sqlc.arg('limit_nft') OFFSET sqlc.arg('offset_nft');