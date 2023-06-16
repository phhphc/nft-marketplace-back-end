-- name: GetMarketplaceSettings :one
SELECT ms.id, ms.marketplace, ms.beneficiary, ms.royalty
FROM "marketplace_settings" ms
WHERE ms.marketplace = sqlc.arg('marketplace');
--
-- -- name: GetValidMarketplaceSettings :one
-- SELECT ms.id, ms.marketplace, ms.admin, ms.signer, ms.royalty, ms.sighash, ms.signature, ms.created_at
-- FROM "marketplace_settings" ms
-- WHERE ms.marketplace = sqlc.arg('marketplace')
-- AND ms.signature IS NOT NULL OR ms.id = 1
-- ORDER BY ms.id DESC
-- LIMIT 1;

-- -- name: InsertMarketplaceSettings :one
-- INSERT INTO "marketplace_settings" ("marketplace", "admin", "signer", "royalty", "typed_data", "sighash", "signature", "created_at")
-- VALUES (sqlc.arg('marketplace'), sqlc.arg('admin'), sqlc.arg('signer'), sqlc.arg('royalty'), sqlc.arg('typed_data'), sqlc.arg('sighash'), sqlc.arg('signature'), sqlc.arg('created_at'))
-- RETURNING *;

-- name: UpdateMarketplaceSettings :one
UPDATE "marketplace_settings"
SET "marketplace" = coalesce(sqlc.narg('n_marketplace'), "marketplace"),
    "beneficiary" = coalesce(sqlc.narg('n_beneficiary'), "beneficiary"),
    "royalty" = coalesce(sqlc.narg('n_royalty'), "royalty")
WHERE "marketplace" = sqlc.arg('marketplace')
RETURNING *;