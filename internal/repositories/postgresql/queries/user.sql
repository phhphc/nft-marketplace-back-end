-- name: GetUserByAddress :one
SELECT public_address, nonce, is_block
FROM "users"
WHERE public_address ILIKE sqlc.arg('public_address');

-- name: GetUsers :many
SELECT fu.public_address, fu.nonce, r.id as role_id, r.name as role, fu.is_block
FROM (
    SELECT * FROM "users" u
    WHERE (u.public_address ILIKE sqlc.narg('public_address') OR sqlc.narg('public_address') IS NULL)
    AND (u.is_block = sqlc.narg('is_block') OR sqlc.narg('is_block') IS NULL)
    ORDER BY public_address ASC
    LIMIT sqlc.arg('limit')
    OFFSET sqlc.arg('offset')
     ) fu
LEFT JOIN "user_roles" ur on fu.public_address = ur.address
LEFT JOIN "roles" r on r.id = ur.role_id
WHERE (r.name = sqlc.narg('role') OR sqlc.narg('role') IS NULL);

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

-- name: GetAllRoles :many
SELECT * FROM "roles";