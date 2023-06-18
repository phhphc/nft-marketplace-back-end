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

-- name: GetUserRoles :many
SELECT *
FROM "roles" r
JOIN "user_roles" ur ON ur.role_id = r.id
WHERE ur.address = $1;

-- name: GetAllRoles :many
SELECT * FROM "roles";