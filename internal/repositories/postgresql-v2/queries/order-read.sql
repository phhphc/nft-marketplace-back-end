-- name: GetExpiredOrder :many
SELECT DISTINCT e.name, o.order_hash, o.end_time, o.is_cancelled, o.is_invalid, o.offerer
FROM events e
         JOIN orders o ON e.order_hash = o.order_hash
WHERE (e.name = 'listing' OR e.name = 'offer')
  AND o.is_cancelled = false
  AND o.is_invalid = false
  AND o.end_time < @now;

-- name: GetOrder :many
SELECT json_build_object(
               'order_hash', o.order_hash,
               'offerer', o.offerer,
               'signature', o.signature,
               'start_time', o.start_time::VARCHAR,
               'end_time', o.end_time::VARCHAR,
               'salt', o.salt,
               'status', json_build_object(
                       'is_fulfilled', o.is_fulfilled,
                       'is_cancelled', o.is_cancelled,
                       'is_invalid', o.is_invalid
                   ),
               'offer', (SELECT json_agg(
                                        json_build_object(
                                                'item_type', offer.item_type,
                                                'token', offer.token,
                                                'identifier', offer.identifier::VARCHAR,
                                                'start_amount', offer.start_amount::VARCHAR,
                                                'end_amount', offer.end_amount::VARCHAR
                                            )
                                    )
                         FROM offer_items offer
                         WHERE o.order_hash = offer.order_hash),
               'consideration', (SELECT json_agg(
                                                json_build_object(
                                                        'item_type', cons.item_type,
                                                        'token', cons.token,
                                                        'identifier', cons.identifier::VARCHAR,
                                                        'start_amount', cons.start_amount::VARCHAR,
                                                        'end_amount', cons.end_amount::VARCHAR,
                                                        'recipient', cons.recipient
                                                    )
                                            )
                                 FROM consideration_items cons
                                 WHERE o.order_hash = cons.order_hash)
           )
FROM orders o
WHERE o."order_hash" ILIKE COALESCE(sqlc.narg('order_hash'), o."order_hash")
  AND o.offerer ILIKE COALESCE(sqlc.narg('offerer'), o.offerer)
  AND o."is_cancelled" = COALESCE(sqlc.narg('is_cancelled'), o."is_cancelled")
  AND o."is_fulfilled" = COALESCE(sqlc.narg('is_fulfilled'), o."is_fulfilled")
  AND o."is_invalid" = COALESCE(sqlc.narg('is_invalid'), o."is_invalid")
  AND o.order_hash IN (SELECT DISTINCT od.order_hash
                       FROM orders od
                                JOIN consideration_items ci on ci.order_hash = od.order_hash
                                JOIN offer_items oi on oi.order_hash = od.order_hash
                       WHERE ci."token" ILIKE COALESCE(sqlc.narg('consideration_token'), ci."token")
                         AND ci.identifier = COALESCE(sqlc.narg('consideration_identifier'), ci.identifier)
                         AND oi.token ILIKE COALESCE(sqlc.narg('offer_token'), oi.token)
                         AND oi.identifier = COALESCE(sqlc.narg('offer_identifier'), oi.identifier))
;
