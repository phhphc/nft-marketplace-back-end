-- name: UpdateMarketplaceSettings :one
UPDATE "marketplace_settings"
SET "marketplace" = coalesce(sqlc.narg('n_marketplace'), "marketplace"),
    "beneficiary" = coalesce(sqlc.narg('n_beneficiary'), "beneficiary"),
    "royalty" = coalesce(sqlc.narg('n_royalty'), "royalty")
WHERE "marketplace" = sqlc.arg('marketplace')
RETURNING *;

-- -- name: InsertMarketplaceSettings :one
-- INSERT INTO "marketplace_settings" ("marketplace", "admin", "signer", "royalty", "typed_data", "sighash", "signature", "created_at")
-- VALUES (sqlc.arg('marketplace'), sqlc.arg('admin'), sqlc.arg('signer'), sqlc.arg('royalty'), sqlc.arg('typed_data'), sqlc.arg('sighash'), sqlc.arg('signature'), sqlc.arg('created_at'))
-- RETURNING *;
