-- name: InsertCollection :one
INSERT INTO "collections" ("token", "owner", "name", "description","category", "metadata")
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetCollection :many
SELECT token, owner, co.name, ca.name as category, description, metadata, created_at
FROM collections co
         JOIN categories ca on co.category = ca.id
WHERE (token ILIKE sqlc.narg('token') or sqlc.narg('token') IS NULL)
  AND (owner ILIKE sqlc.narg('owner') or sqlc.narg('owner') IS NULL)
  AND (co.name ILIKE sqlc.narg('name') or sqlc.narg('name') IS NULL)
  AND (ca.name ILIKE sqlc.narg('category') or sqlc.narg('category') IS NULL)
OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');

-- name: GetCollectionWithCategory :many
SELECT token, owner, co.name, ca.name as category, description, metadata, created_at
FROM collections co
         JOIN categories ca on co.category = ca.id
WHERE ca.name ILIKE sqlc.narg('category')
OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');

-- name: UpdateCollectionLastSyncBlock :exec
UPDATE collections
SET "last_sync_block" = $2
WHERE token = $1;

-- name: GetCollectionLastSyncBlock :one
SELECT "last_sync_block"
FROM collections
WHERE token = $1;
