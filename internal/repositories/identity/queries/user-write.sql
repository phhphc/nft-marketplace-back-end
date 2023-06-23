-- name: UpdateNonce :one
UPDATE "users"
SET nonce = sqlc.arg('nonce')
WHERE public_address ILIKE sqlc.arg('public_address')
RETURNING *;

-- name: UpdateUserBlockState :one
UPDATE "users"
SET is_block = sqlc.arg('is_block')
WHERE public_address ILIKE sqlc.arg('public_address')
RETURNING *;

-- name: InsertUser :one
INSERT INTO "users" (public_address, nonce)
VALUES (sqlc.arg('public_address'), sqlc.arg('nonce'))
RETURNING *;

-- name: InsertUserRole :one
INSERT INTO "user_roles" (address, role_id)
VALUES (sqlc.arg('address'), sqlc.arg('role_id'))
RETURNING *;

-- name: DeleteUserRole :exec
DELETE FROM "user_roles"
WHERE address = sqlc.arg('address') AND role_id = sqlc.arg('role_id');