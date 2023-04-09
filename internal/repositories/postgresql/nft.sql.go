// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: nft.sql

package postgresql

import (
	"context"

	"github.com/tabbed/pqtype"
)

const updateNft = `-- name: UpdateNft :exec
INSERT INTO "nfts" ("token", "identifier", "owner", "is_burned", "block_number", "tx_index")
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT ("token", "identifier") DO UPDATE
    SET "owner"=$3,
        "is_burned"=$4,
        "block_number"=$5,
        "tx_index"=$6
WHERE $5 > nfts."block_number"
   OR ($5 = nfts."block_number" AND $6 > nfts."tx_index")
`

type UpdateNftParams struct {
	Token       string
	Identifier  string
	Owner       string
	IsBurned    bool
	BlockNumber string
	TxIndex     int64
}

func (q *Queries) UpdateNft(ctx context.Context, arg UpdateNftParams) error {
	_, err := q.db.ExecContext(ctx, updateNft,
		arg.Token,
		arg.Identifier,
		arg.Owner,
		arg.IsBurned,
		arg.BlockNumber,
		arg.TxIndex,
	)
	return err
}

const updateNftMetadata = `-- name: UpdateNftMetadata :exec
UPDATE "nfts"
SET "metadata" = $1
WHERE "token" = $2
  AND "identifier" = $3
`

type UpdateNftMetadataParams struct {
	Metadata   pqtype.NullRawMessage
	Token      string
	Identifier string
}

func (q *Queries) UpdateNftMetadata(ctx context.Context, arg UpdateNftMetadataParams) error {
	_, err := q.db.ExecContext(ctx, updateNftMetadata, arg.Metadata, arg.Token, arg.Identifier)
	return err
}
