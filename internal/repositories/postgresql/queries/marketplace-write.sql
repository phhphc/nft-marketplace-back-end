

-- name: UpdateMarketplaceLastSyncBlock :exec
UPDATE "marketplace"
SET "last_sync_block" = $1
WHERE true;