-- name: GetNft :one
SELECT *
FROM "nfts"
WHERE "token" = sqlc.arg('token')
  AND "identifier" = sqlc.arg('identifier');

-- name: ListNftWithListing :many
SELECT json_build_object(
               'token', n.token,
               'identifier', n.identifier::VARCHAR,
               'owner', n.owner,
               'token_uri', n.token_uri,
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
LIMIT sqlc.arg('limit_nft') OFFSET sqlc.arg('offset_nft');
