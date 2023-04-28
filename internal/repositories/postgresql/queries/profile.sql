-- name: UpsertProfile :one
INSERT INTO "profiles" ("address", "username", "metadata", "signature")
VALUES (sqlc.arg('address'), sqlc.arg('username'), sqlc.arg('metadata'), sqlc.arg('signature'))
ON CONFLICT ("address") DO UPDATE SET
  "username" = sqlc.arg('username'),
  "metadata" = sqlc.arg('metadata'),
  "signature" = sqlc.arg('signature')
RETURNING *;

-- name: DeleteProfile :exec
DELETE FROM "profiles"
WHERE "address" = sqlc.arg('address');

-- name: GetProfile :one
SELECT "address", "username", "metadata", "signature"
FROM "profiles"
WHERE "address" = sqlc.arg('address')
LIMIT 1;