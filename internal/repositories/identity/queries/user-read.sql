-- name: GetUserByAddress :one
SELECT public_address, nonce, is_block
FROM "users"
WHERE public_address ILIKE sqlc.arg('public_address');

-- name: GetUsers :many
SELECT json_build_object(
               'address', u."public_address",
               'nonce', u."nonce",
               'is_block', u."is_block",
               'roles', (SELECT json_agg(
                                        json_build_object(
                                                'role_id', uwr.id,
                                                'role', uwr.name
                                            )
                                    )
                         FROM (SELECT r.*
                               FROM "user_roles" ur
                                        JOIN "roles" r ON r.id = ur.role_id
                               WHERE ur.address = u.public_address
                               ORDER BY r.id ASC) uwr)
           )
FROM "users" u
WHERE u.public_address ILIKE COALESCE(sqlc.narg('public_address'), u.public_address)
  AND u.is_block = COALESCE(sqlc.narg('is_block'), u.is_block)
  AND u.public_address IN (SELECT DISTINCT us.public_address
                           FROM "users" us
                                    LEFT JOIN "user_roles" ur on us.public_address = ur.address
                                    LEFT JOIN "roles" r on r.id = ur.role_id
                           WHERE r.name ILIKE COALESCE(sqlc.narg('role'), r.name))
ORDER BY public_address ASC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetUserRoles :many
SELECT *
FROM "roles" r
         JOIN "user_roles" ur ON ur.role_id = r.id
WHERE ur.address = $1;

-- name: GetAllRoles :many
SELECT *
FROM "roles";