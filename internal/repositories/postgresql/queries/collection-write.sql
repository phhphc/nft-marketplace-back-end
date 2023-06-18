-- name: InsertCollection :one
INSERT INTO "collections" ("token", "owner", "name", "description","category", "metadata")
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: UpdateCollectionLastSyncBlock :exec
UPDATE collections
SET "last_sync_block" = $2
WHERE token = $1;