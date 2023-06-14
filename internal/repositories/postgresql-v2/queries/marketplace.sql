-- name: GetMarketplaceLastSyncBlock :one
SELECT "last_sync_block"
FROM "marketplace"
LIMIT 1;

-- name: UpdateMarketplaceLastSyncBlock :exec
UPDATE "marketplace"
SET "last_sync_block" = $1
WHERE true;