-- name: GetProfile :one
SELECT "address", "username", "metadata", "signature"
FROM "profiles"
WHERE "address" = sqlc.arg('address')
LIMIT 1;