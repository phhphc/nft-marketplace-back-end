// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: search.sql

package postgresql

import (
	"context"
	"database/sql"
)

const fullTextSearch = `-- name: FullTextSearch :many
SELECT paged_nfts.block_number,
       paged_nfts.token,
       paged_nfts.identifier,
       paged_nfts.owner,
       paged_nfts.is_hidden,
       paged_nfts.metadata -> 'image'       AS image,
       paged_nfts.metadata -> 'name'        AS name,
       paged_nfts.metadata -> 'description' AS description,
       ci.order_hash,
       ci.item_type,
       ci.start_amount                      AS start_price,
       ci.end_amount                        AS end_price,
       o.start_time                         AS start_time,
       o.end_time                           AS end_time
FROM (SELECT token, identifier, owner, metadata, is_burned, is_hidden, block_number, tx_index
      FROM nfts, plainto_tsquery('simple', $1) AS q
      WHERE nfts.is_burned = FALSE
        AND nfts."is_hidden" = COALESCE($2, "nfts"."is_hidden")
        AND (nfts.owner ILIKE $3 OR $3 IS NULL)
        AND (nfts.token ILIKE $4 OR $4 IS NULL)
        AND (nfts.tsv @@ q OR $1 IS NULL)
      LIMIT $6 OFFSET $5) AS paged_nfts
         LEFT JOIN offer_items oi
                   ON paged_nfts.token ILIKE oi.token AND paged_nfts.identifier = oi.identifier
         LEFT JOIN consideration_items ci ON oi.order_hash ILIKE ci.order_hash
         LEFT JOIN (SELECT order_hash, offerer, recipient, salt, start_time, end_time, signature, is_cancelled, is_validated, is_fulfilled, is_invalid
                    FROM orders
                    WHERE orders.is_fulfilled = FALSE
                      AND orders.is_cancelled = FALSE
                      AND orders.is_invalid = FALSE
                      AND orders.start_time <= round(EXTRACT(EPOCH FROM now()))
                      AND orders.end_time >= round(EXTRACT(EPOCH FROM now()))) o
                   ON oi.order_hash ILIKE o.order_hash
ORDER BY paged_nfts.block_number, paged_nfts.tx_index, ci.id, paged_nfts.token, paged_nfts.identifier
`

type FullTextSearchParams struct {
	Q        sql.NullString
	IsHidden sql.NullBool
	Owner    sql.NullString
	Token    sql.NullString
	Offset   int32
	Limit    int32
}

type FullTextSearchRow struct {
	BlockNumber string
	Token       string
	Identifier  string
	Owner       string
	IsHidden    bool
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

func (q *Queries) FullTextSearch(ctx context.Context, arg FullTextSearchParams) ([]FullTextSearchRow, error) {
	rows, err := q.db.QueryContext(ctx, fullTextSearch,
		arg.Q,
		arg.IsHidden,
		arg.Owner,
		arg.Token,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FullTextSearchRow{}
	for rows.Next() {
		var i FullTextSearchRow
		if err := rows.Scan(
			&i.BlockNumber,
			&i.Token,
			&i.Identifier,
			&i.Owner,
			&i.IsHidden,
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
