// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package postgresql

import (
	"context"
)

type Querier interface {
	GetListNft(ctx context.Context, arg GetListNftParams) ([]GetListNftRow, error)
	GetNftDetail(ctx context.Context, arg GetNftDetailParams) (GetNftDetailRow, error)
	UpsertListing(ctx context.Context, arg UpsertListingParams) error
	UpsertNft(ctx context.Context, arg UpsertNftParams) error
}

var _ Querier = (*Queries)(nil)
