-- name: FullTextSearch :many
SELECT paged_nfts.block_number,
       paged_nfts.token,
       paged_nfts.identifier,
       paged_nfts.owner,
       paged_nfts.is_hidden,
       paged_nfts.metadata -> 'image'       AS image,
       paged_nfts.metadata -> 'name'        AS name,
       paged_nfts.metadata -> 'description' AS description,
       ci.order_hash,
       ci.item_type,
       ci.start_amount                      AS start_price,
       ci.end_amount                        AS end_price,
       o.start_time                         AS start_time,
       o.end_time                           AS end_time
FROM (SELECT *
      FROM nfts, plainto_tsquery('simple', sqlc.narg('q')) AS q
      WHERE nfts.is_burned = FALSE
        AND nfts."is_hidden" = COALESCE(sqlc.narg('is_hidden'), "nfts"."is_hidden")
        AND (nfts.owner ILIKE sqlc.narg('owner') OR sqlc.narg('owner') IS NULL)
        AND (nfts.token ILIKE sqlc.narg('token') OR sqlc.narg('token') IS NULL)
        AND (nfts.tsv @@ q OR sqlc.narg('q') IS NULL)
      LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset')) AS paged_nfts
         LEFT JOIN offer_items oi
                   ON paged_nfts.token ILIKE oi.token AND paged_nfts.identifier = oi.identifier
         LEFT JOIN consideration_items ci ON oi.order_hash ILIKE ci.order_hash
         LEFT JOIN (SELECT *
                    FROM orders
                    WHERE orders.is_fulfilled = FALSE
                      AND orders.is_cancelled = FALSE
                      AND orders.is_invalid = FALSE
                      AND orders.start_time <= round(EXTRACT(EPOCH FROM now()))
                      AND orders.end_time >= round(EXTRACT(EPOCH FROM now()))) o
                   ON oi.order_hash ILIKE o.order_hash
ORDER BY paged_nfts.block_number, paged_nfts.tx_index, ci.id, paged_nfts.token, paged_nfts.identifier;