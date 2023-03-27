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
    n.token, n.identifier, n.owner, n.metadata, n.is_burned
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

const getNFTValidConsiderations = `-- name: GetNFTValidConsiderations :many
SELECT
    selected_nft.block_number,
    selected_nft.token,
    selected_nft.identifier,
    selected_nft.owner,
    selected_nft.metadata ->> 'image' AS image,
    selected_nft.metadata ->> 'name' AS name,
    selected_nft.metadata ->> 'description' AS description,
    selected_nft.metadata AS metadata,
    ci.order_hash,
    ci.item_type,
    ci.start_amount AS start_price,
    ci.end_amount AS end_price,
    o.start_time AS start_time,
    o.end_time AS end_time
FROM (
         SELECT token, identifier, owner, metadata, is_burned, block_number, tx_index FROM nfts WHERE nfts.token ILIKE $1 AND nfts.identifier = $2
     ) selected_nft
         LEFT JOIN "offer_items" oi ON oi.token ILIKE selected_nft.token AND oi.identifier = selected_nft.identifier
         LEFT JOIN "consideration_items" ci ON ci.order_hash ILIKE oi.order_hash
         LEFT JOIN (
    SELECT order_hash, offerer, zone, recipient, order_type, zone_hash, salt, start_time, end_time, signature, is_cancelled, is_validated, is_fulfilled FROM orders WHERE orders.is_fulfilled = FALSE AND orders.is_cancelled = FALSE
) o ON oi.order_hash ILIKE o.order_hash
`

type GetNFTValidConsiderationsParams struct {
	Token      string
	Identifier string
}

type GetNFTValidConsiderationsRow struct {
	BlockNumber string
	Token       string
	Identifier  string
	Owner       string
	Image       interface{}
	Name        interface{}
	Description interface{}
	Metadata    pqtype.NullRawMessage
	OrderHash   sql.NullString
	ItemType    sql.NullInt32
	StartPrice  sql.NullString
	EndPrice    sql.NullString
	StartTime   sql.NullString
	EndTime     sql.NullString
}

func (q *Queries) GetNFTValidConsiderations(ctx context.Context, arg GetNFTValidConsiderationsParams) ([]GetNFTValidConsiderationsRow, error) {
	rows, err := q.db.QueryContext(ctx, getNFTValidConsiderations, arg.Token, arg.Identifier)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetNFTValidConsiderationsRow{}
	for rows.Next() {
		var i GetNFTValidConsiderationsRow
		if err := rows.Scan(
			&i.BlockNumber,
			&i.Token,
			&i.Identifier,
			&i.Owner,
			&i.Image,
			&i.Name,
			&i.Description,
			&i.Metadata,
			&i.OrderHash,
			&i.ItemType,
			&i.StartPrice,
			&i.EndPrice,
			&i.StartTime,
			&i.EndTime,
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
SELECT
    paged_nfts.block_number,
    paged_nfts.token,
    paged_nfts.identifier,
    paged_nfts.owner,
    paged_nfts.metadata -> 'image' AS image,
    paged_nfts.metadata -> 'name' AS name,
    paged_nfts.metadata -> 'description' AS description,
    ci.order_hash,
    ci.item_type,
    ci.start_amount AS start_price,
    ci.end_amount AS end_price,
    o.start_time AS start_time,
    o.end_time AS end_time
FROM (
         SELECT token, identifier, owner, metadata, is_burned, block_number, tx_index FROM nfts
         WHERE nfts.is_burned = FALSE
         AND (nfts.owner ILIKE $1 OR $1 IS NULL)
         AND (nfts.token ILIKE $2 OR $2 IS NULL)
         LIMIT $4
         OFFSET $3
     ) AS paged_nfts
         LEFT JOIN offer_items oi
                   ON paged_nfts.token ILIKE oi.token AND paged_nfts.identifier = oi.identifier
         LEFT JOIN consideration_items ci ON oi.order_hash ILIKE ci.order_hash
         LEFT JOIN (
    SELECT order_hash, offerer, zone, recipient, order_type, zone_hash, salt, start_time, end_time, signature, is_cancelled, is_validated, is_fulfilled FROM orders WHERE orders.is_fulfilled = FALSE AND orders.is_cancelled = FALSE
) o ON oi.order_hash ILIKE o.order_hash
ORDER BY paged_nfts.block_number, paged_nfts.tx_index, ci.id, paged_nfts.token, paged_nfts.identifier
`

type GetNFTsWithPricesPaginatedParams struct {
	Owner  sql.NullString
	Token  sql.NullString
	Offset int32
	Limit  int32
}

type GetNFTsWithPricesPaginatedRow struct {
	BlockNumber string
	Token       string
	Identifier  string
	Owner       string
	Image       interface{}
	Name        interface{}
	Description interface{}
	OrderHash   sql.NullString
	ItemType    sql.NullInt32
	StartPrice  sql.NullString
	EndPrice    sql.NullString
	StartTime   sql.NullString
	EndTime     sql.NullString
}

func (q *Queries) GetNFTsWithPricesPaginated(ctx context.Context, arg GetNFTsWithPricesPaginatedParams) ([]GetNFTsWithPricesPaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, getNFTsWithPricesPaginated,
		arg.Owner,
		arg.Token,
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
			&i.BlockNumber,
			&i.Token,
			&i.Identifier,
			&i.Owner,
			&i.Image,
			&i.Name,
			&i.Description,
			&i.OrderHash,
			&i.ItemType,
			&i.StartPrice,
			&i.EndPrice,
			&i.StartTime,
			&i.EndTime,
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

const upsertNFTV2 = `-- name: UpsertNFTV2 :exec
INSERT INTO "nfts" (token, identifier, owner, metadata, is_burned)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (token, identifier) DO UPDATE SET
                                              owner = $3,
                                              metadata = $4,
                                              is_burned = $5
WHERE nfts.block_number < $6 OR (nfts.block_number = $6 AND nfts.tx_index < $7)
`

type UpsertNFTV2Params struct {
	Token       string
	Identifier  string
	Owner       string
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
		arg.Metadata,
		arg.IsBurned,
		arg.BlockNumber,
		arg.TxIndex,
	)
	return err
}
