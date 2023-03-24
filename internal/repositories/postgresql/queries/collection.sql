-- name: InsertCollection :one
INSERT INTO "collections" ("token", "owner", "name", "description","category", "metadata")
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetListCollection :many
SELECT "token", c."name", "description", "owner",k."name" as "category"
FROM "collections" c
JOIN categories k on k.id = c.category
ORDER BY "created_at" DESC
OFFSET $1
LIMIT $2;