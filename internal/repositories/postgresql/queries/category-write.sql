-- name: InsertCategory :one
INSERT INTO "categories" ("name")
VALUES ($1)
RETURNING *;
