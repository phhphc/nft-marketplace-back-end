-- name: GetUserByAddress :one
SELECT public_address, nonce
FROM "users"
WHERE public_address ILIKE sqlc.arg('public_address');

-- name: GetUserRolesByAddress :many
SELECT ur.user_id r.name as role
FROM "user_roles" ur
JOIN "roles" r on r.id = ur.role_id
WHERE public_address ILIKE sqlc.arg('public_address');

-- name: UpdateNonce :one
UPDATE "users"
SET nonce = sqlc.arg('nonce')
WHERE public_address ILIKE sqlc.arg('public_address')
RETURNING *;

-- name: InsertUser :one
INSERT INTO "users" (public_address, nonce)
VALUES (sqlc.arg('public_address'), sqlc.arg('nonce'))
RETURNING *;
