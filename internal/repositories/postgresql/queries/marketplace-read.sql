-- name: GetMarketplaceLastSyncBlock :one
SELECT "last_sync_block"
FROM "marketplace"
LIMIT 1;
