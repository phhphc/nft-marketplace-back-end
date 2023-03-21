-- name: InsertCollection :one
INSERT INTO "collections" ("token", "owner", "name", "description","category", "metadata")
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;
