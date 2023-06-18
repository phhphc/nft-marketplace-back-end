-- name: GetCollection :many
SELECT token, owner, co.name, ca.name as category, description, metadata, created_at
FROM collections co
         JOIN categories ca on co.category = ca.id
WHERE token ILIKE COALESCE(sqlc.narg('token'), token)
  AND owner ILIKE COALESCE(sqlc.narg('owner'), owner)
  AND co.name ILIKE COALESCE(sqlc.narg('name'), co.name)
  AND ca.name ILIKE COALESCE(sqlc.narg('category'), ca.name)
OFFSET sqlc.arg('offset') LIMIT sqlc.arg('limit');

-- name: GetCollectionLastSyncBlock :one
SELECT "last_sync_block"
FROM collections
WHERE token = $1;
