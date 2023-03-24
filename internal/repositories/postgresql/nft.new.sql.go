// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: nft.new.sql

package postgresql

import (
	"context"
	"database/sql"

	"github.com/tabbed/pqtype"
)

const getListValidNFT = `-- name: GetListValidNFT :many
SELECT
    n.token, n.identifier, n.owner, n.token_uri, n.metadata, n.is_burned
FROM "nfts" n
WHERE
    n.is_burned = FALSE
  AND
    (n.token ILIKE $1 OR $1 IS NULL)
  AND
    (n.owner ILIKE $2 OR $2 IS NULL)
ORDER BY n.token ASC, n.identifier ASC
OFFSET $3
LIMIT $4
`

type GetListValidNFTParams struct {
	Token  sql.NullString
	Owner  sql.NullString
	Offset int32
	Limit  int32
}

type GetListValidNFTRow struct {
	Token      string
	Identifier string
	Owner      string
	TokenUri   sql.NullString
	Metadata   pqtype.NullRawMessage
	IsBurned   bool
}

func (q *Queries) GetListValidNFT(ctx context.Context, arg GetListValidNFTParams) ([]GetListValidNFTRow, error) {
	rows, err := q.db.QueryContext(ctx, getListValidNFT,
		arg.Token,
		arg.Owner,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListValidNFTRow{}
	for rows.Next() {
		var i GetListValidNFTRow
		if err := rows.Scan(
			&i.Token,
			&i.Identifier,
			&i.Owner,
			&i.TokenUri,
			&i.Metadata,
			&i.IsBurned,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNFTsWithPricesPaginated = `-- name: GetNFTsWithPricesPaginated :many
SELECT paged_nfts.token,
       paged_nfts.identifier,
       paged_nfts.owner,
       paged_nfts.token_uri,
       paged_nfts.metadata -> 'image' AS image,
       paged_nfts.metadata -> 'name' AS name,
       paged_nfts.metadata -> 'description' AS description,
       ci.order_hash,
       ci.item_type,
       ci.amount AS price
FROM (
        SELECT token, identifier, owner, is_burned, token_uri, metadata, block_number, tx_index FROM nfts
        WHERE nfts.is_burned = FALSE
        AND (nfts.token ILIKE $1 OR $1 IS NULL)
        AND (nfts.owner ILIKE $2 OR $2 IS NULL)
        OFFSET $3 LIMIT $4
     ) AS paged_nfts
        LEFT JOIN offer_items oi ON paged_nfts.token = oi.token AND paged_nfts.identifier = oi.identifier
        LEFT JOIN consideration_items ci ON oi.order_hash = ci.order_hash
ORDER BY paged_nfts.token, paged_nfts.identifier, ci.order_hash
`

type GetNFTsWithPricesPaginatedParams struct {
	Token  sql.NullString
	Owner  sql.NullString
	Offset int32
	Limit  int32
}

type GetNFTsWithPricesPaginatedRow struct {
	Token       string
	Identifier  string
	Owner       string
	TokenUri    sql.NullString
	Image       interface{}
	Name        interface{}
	Description interface{}
	OrderHash   sql.NullString
	ItemType    sql.NullInt32
	Price       sql.NullString
}

func (q *Queries) GetNFTsWithPricesPaginated(ctx context.Context, arg GetNFTsWithPricesPaginatedParams) ([]GetNFTsWithPricesPaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, getNFTsWithPricesPaginated,
		arg.Token,
		arg.Owner,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetNFTsWithPricesPaginatedRow{}
	for rows.Next() {
		var i GetNFTsWithPricesPaginatedRow
		if err := rows.Scan(
			&i.Token,
			&i.Identifier,
			&i.Owner,
			&i.TokenUri,
			&i.Image,
			&i.Name,
			&i.Description,
			&i.OrderHash,
			&i.ItemType,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getValidNFT = `-- name: GetValidNFT :one
SELECT
    n.token, n.identifier, n.owner, n.token_uri, n.metadata, n.is_burned
FROM "nfts" n
WHERE
    n.is_burned = FALSE
  AND
    n.token = $1
  AND
    n.identifier = $2
`

type GetValidNFTParams struct {
	Token      string
	Identifier string
}

type GetValidNFTRow struct {
	Token      string
	Identifier string
	Owner      string
	TokenUri   sql.NullString
	Metadata   pqtype.NullRawMessage
	IsBurned   bool
}

func (q *Queries) GetValidNFT(ctx context.Context, arg GetValidNFTParams) (GetValidNFTRow, error) {
	row := q.db.QueryRowContext(ctx, getValidNFT, arg.Token, arg.Identifier)
	var i GetValidNFTRow
	err := row.Scan(
		&i.Token,
		&i.Identifier,
		&i.Owner,
		&i.TokenUri,
		&i.Metadata,
		&i.IsBurned,
	)
	return i, err
}

const upsertNFTV2 = `-- name: UpsertNFTV2 :exec
INSERT INTO "nfts" (token, identifier, owner, token_uri, metadata, is_burned)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (token, identifier) DO UPDATE SET
    owner = $3,
    token_uri = $4,
    metadata = $5,
    is_burned = $6
WHERE nfts.block_number < $7 OR (nfts.block_number = $7 AND nfts.tx_index < $8)
`

type UpsertNFTV2Params struct {
	Token       string
	Identifier  string
	Owner       string
	TokenUri    sql.NullString
	Metadata    pqtype.NullRawMessage
	IsBurned    bool
	BlockNumber string
	TxIndex     int64
}

func (q *Queries) UpsertNFTV2(ctx context.Context, arg UpsertNFTV2Params) error {
	_, err := q.db.ExecContext(ctx, upsertNFTV2,
		arg.Token,
		arg.Identifier,
		arg.Owner,
		arg.TokenUri,
		arg.Metadata,
		arg.IsBurned,
		arg.BlockNumber,
		arg.TxIndex,
	)
	return err
}
