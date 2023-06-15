-- name: GetCategoryByName :one
SELECT * FROM "categories"
WHERE "name" = $1;