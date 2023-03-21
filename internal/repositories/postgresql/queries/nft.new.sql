-- name: GetListValidNft :many
SELECT
    n.token, n.identifier, n.owner, n.token_uri, n.metadata, n.is_burned
FROM "nfts" n
WHERE
    n.is_burned = FALSE
  AND
    (n.token ILIKE sqlc.narg('token') OR sqlc.narg('token') IS NULL)
  AND
    (n.owner ILIKE sqlc.narg('owner') OR sqlc.narg('owner') IS NULL)
ORDER BY n.token ASC, n.identifier ASC
OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');

-- name: GetValidNft :one
SELECT
    n.token, n.identifier, n.owner, n.token_uri, n.metadata, n.is_burned
FROM "nfts" n
WHERE
    n.is_burned = FALSE
  AND
    n.token = sqlc.arg('token')
  AND
    n.identifier = sqlc.arg('identifier');

-- name: UpsertNftV2 :exec
INSERT INTO "nfts" (token, identifier, owner, token_uri, metadata, is_burned)
VALUES (sqlc.arg('token'), sqlc.arg('identifier'), sqlc.arg('owner'), sqlc.arg('token_uri'), sqlc.arg('metadata'), sqlc.arg('is_burned'))
ON CONFLICT (token, identifier) DO UPDATE SET
    owner = sqlc.arg('owner'),
    token_uri = sqlc.arg('token_uri'),
    metadata = sqlc.arg('metadata'),
    is_burned = sqlc.arg('is_burned')
WHERE nfts.block_number < sqlc.arg('block_number') OR (nfts.block_number = sqlc.arg('block_number') AND nfts.tx_index < sqlc.arg('tx_index'));
