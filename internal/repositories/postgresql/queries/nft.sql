-- name: UpdateNft :exec
INSERT INTO "nfts" ("token", "identifier", "owner", "is_burned", "metadata", "block_number", "tx_index")
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT ("token", "identifier") DO UPDATE
    SET "owner"=$3,
        "is_burned"=$4,
        "metadata"=$5,
        "block_number"=$6,
        "tx_index"=$7
WHERE $6 > nfts."block_number"
   OR ($6 = nfts."block_number" AND $7 > nfts."tx_index");