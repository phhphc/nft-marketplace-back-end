-- name: InsertMarketplaceSettings :one
INSERT INTO "marketplace_settings" ("marketplace", "admin", "signer", "royalty", "typed_data", "sighash", "signature", "created_at")
VALUES (sqlc.arg('marketplace'), sqlc.arg('admin'), sqlc.arg('signer'), sqlc.arg('royalty'), sqlc.arg('typed_data'), sqlc.arg('sighash'), sqlc.arg('signature'), sqlc.arg('created_at'))
RETURNING *;