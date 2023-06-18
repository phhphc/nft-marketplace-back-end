// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: nft-read.sql

package gen

import (
	"context"
	"database/sql"
	"encoding/json"
)

const getNft = `-- name: GetNft :one
SELECT token, identifier, owner, token_uri, metadata, is_burned, is_hidden, block_number, tx_index
FROM "nfts"
WHERE "token" = $1
  AND "identifier" = $2
`

type GetNftParams struct {
	Token      string
	Identifier string
}

func (q *Queries) GetNft(ctx context.Context, arg GetNftParams) (Nft, error) {
	row := q.queryRow(ctx, q.getNftStmt, getNft, arg.Token, arg.Identifier)
	var i Nft
	err := row.Scan(
		&i.Token,
		&i.Identifier,
		&i.Owner,
		&i.TokenUri,
		&i.Metadata,
		&i.IsBurned,
		&i.IsHidden,
		&i.BlockNumber,
		&i.TxIndex,
	)
	return i, err
}

const listNftWithListing = `-- name: ListNftWithListing :many
SELECT json_build_object(
               'token', n.token,
               'identifier', n.identifier::VARCHAR,
               'owner', n.owner,
               'metadata', n.metadata,
               'is_hidden', n.is_hidden,
               'listing',
               (SELECT json_agg(
                               json_build_object(
                                       'order_hash', l.order_hash,
                                       'item_type', l.item_type,
                                       'start_time', l.start_time::VARCHAR,
                                       'end_time', l.end_time::VARCHAR,
                                       'start_price', l.start_price::VARCHAR,
                                       'end_price', l.end_price::VARCHAR
                                   )
                           )
                FROM (SELECT o.order_hash,
                             ci.item_type,
                             o.start_time         AS start_time,
                             o.end_time           AS end_time,
                             SUM(ci.start_amount) AS start_price,
                             SUM(ci.end_amount)   AS end_price
                      FROM orders o
                               JOIN offer_items oi on o.order_hash = oi.order_hash
                               JOIN consideration_items ci on o.order_hash = ci.order_hash
                      WHERE o.order_hash NOT IN (SELECT DISTINCT c.order_hash
                                                 FROM consideration_items c
                                                 WHERE c.item_type != $1)
                        AND o.is_fulfilled = FALSE
                        AND o.is_cancelled = FALSE
                        AND o.is_invalid = FALSE
                        AND o.start_time <= $2
                        AND o.end_time > $2
                        AND oi.token ILIKE n.token
                        AND oi.identifier = n.identifier
                      GROUP BY o.order_hash,
                               ci.item_type,
                               o.start_time,
                               o.end_time
                      LIMIT $3) as l)
           )
FROM nfts n
WHERE n."is_burned" = FALSE
  AND n."is_hidden" = COALESCE($4, n."is_hidden")
  AND n."owner" ILIKE COALESCE($5, n."owner")
  AND n."token" ILIKE COALESCE($6, n."token")
  AND n."identifier" = COALESCE($7, n."identifier")
LIMIT $9 OFFSET $8
`

type ListNftWithListingParams struct {
	ItemType     int32
	Now          sql.NullString
	LimitListing int32
	IsHidden     sql.NullBool
	Owner        sql.NullString
	Token        sql.NullString
	Identifier   sql.NullString
	OffsetNft    int32
	LimitNft     int32
}

func (q *Queries) ListNftWithListing(ctx context.Context, arg ListNftWithListingParams) ([]json.RawMessage, error) {
	rows, err := q.query(ctx, q.listNftWithListingStmt, listNftWithListing,
		arg.ItemType,
		arg.Now,
		arg.LimitListing,
		arg.IsHidden,
		arg.Owner,
		arg.Token,
		arg.Identifier,
		arg.OffsetNft,
		arg.LimitNft,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []json.RawMessage{}
	for rows.Next() {
		var json_build_object json.RawMessage
		if err := rows.Scan(&json_build_object); err != nil {
			return nil, err
		}
		items = append(items, json_build_object)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
