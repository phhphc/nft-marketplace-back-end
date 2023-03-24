// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package postgresql

import (
	"context"
)

type Querier interface {
	GetListValidNFT(ctx context.Context, arg GetListValidNFTParams) ([]GetListValidNFTRow, error)
	GetNFTsWithPricesPaginated(ctx context.Context, arg GetNFTsWithPricesPaginatedParams) ([]GetNFTsWithPricesPaginatedRow, error)
	GetValidNFT(ctx context.Context, arg GetValidNFTParams) (GetValidNFTRow, error)
	InsertOrder(ctx context.Context, arg InsertOrderParams) error
	InsertOrderConsiderationItem(ctx context.Context, arg InsertOrderConsiderationItemParams) error
	InsertOrderOfferItem(ctx context.Context, arg InsertOrderOfferItemParams) error
	UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) (string, error)
	UpsertNFTV2(ctx context.Context, arg UpsertNFTV2Params) error
}

var _ Querier = (*Queries)(nil)
