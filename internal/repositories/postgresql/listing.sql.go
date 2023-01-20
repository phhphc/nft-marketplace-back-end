// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: listing.sql

package postgresql

import (
	"context"
)

const upsertListing = `-- name: UpsertListing :exec
INSERT INTO "listings" (listing_id, collection, token_id, seller, price, status, block_number, tx_index)
VALUES ($1,$2,$3,$4,$5,$6, $7, $8)
ON CONFLICT (collection, token_id) DO UPDATE
SET seller=$4,price=$5,status=$6, block_number=$7, tx_index=$8
WHERE $7 > listings.block_number or ($7 = listings.block_number and $8 > listings.tx_index)
`

type UpsertListingParams struct {
	ListingID   string `json:"listing_id"`
	Collection  string `json:"collection"`
	TokenID     string `json:"token_id"`
	Seller      string `json:"seller"`
	Price       string `json:"price"`
	Status      string `json:"status"`
	BlockNumber string `json:"block_number"`
	TxIndex     int64  `json:"tx_index"`
}

func (q *Queries) UpsertListing(ctx context.Context, arg UpsertListingParams) error {
	_, err := q.db.ExecContext(ctx, upsertListing,
		arg.ListingID,
		arg.Collection,
		arg.TokenID,
		arg.Seller,
		arg.Price,
		arg.Status,
		arg.BlockNumber,
		arg.TxIndex,
	)
	return err
}
