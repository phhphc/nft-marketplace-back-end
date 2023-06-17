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


