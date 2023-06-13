-- name: UpdateNft :exec
INSERT INTO "nfts" ("token", "identifier", "owner", "is_burned", "block_number", "tx_index")
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT ("token", "identifier") DO UPDATE
    SET "owner"=$3,
        "is_burned"=$4,
        "block_number"=$5,
        "tx_index"=$6
WHERE $5 > nfts."block_number"
   OR ($5 = nfts."block_number" AND $6 > nfts."tx_index");


-- name: UpdateNftMetadata :exec
UPDATE "nfts"
SET "metadata" = sqlc.narg('metadata')
WHERE "token" = sqlc.arg('token')
  AND "identifier" = sqlc.arg('identifier');

-- name: UpdateNftStatus :one
UPDATE "nfts"
SET "is_burned" = COALESCE(sqlc.narg('is_burned'), "is_burned"),
    "is_hidden"= COALESCE(sqlc.narg('is_hidden'), "is_hidden")
WHERE token = @token
  AND identifier = @identifier
RETURNING *;

-- name: GetNft :one
SELECT *
FROM "nfts"
WHERE "token" = sqlc.arg('token')
  AND "identifier" = sqlc.arg('identifier');