-- name: InsertCategory :one
INSERT INTO "categories" ("name")
VALUES ($1)
RETURNING *;

-- name: GetCategoryByName :one
SELECT * FROM "categories"
WHERE "name" = $1;