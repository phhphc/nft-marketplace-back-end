-- name: GetMarketplaceSettings :one
SELECT ms.id, ms.marketplace, ms.admin, ms.signer, ms.royalty, ms.sighash, ms.signature, ms.created_at
FROM "marketplace_settings" ms
WHERE (ms.marketplace = sqlc.narg('marketplace') OR sqlc.narg('marketplace') IS NULL)
AND (ms.id = sqlc.narg('id') OR sqlc.narg('id') IS NULL)
ORDER BY ms.id DESC
LIMIT 1;

-- name: GetValidMarketplaceSettings :one
SELECT ms.id, ms.marketplace, ms.admin, ms.signer, ms.royalty, ms.sighash, ms.signature, ms.created_at
FROM "marketplace_settings" ms
WHERE ms.marketplace = sqlc.arg('marketplace')
AND ms.signature IS NOT NULL OR ms.id = 1
ORDER BY ms.id DESC
LIMIT 1;