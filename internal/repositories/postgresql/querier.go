// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
)

type Querier interface {
	DeleteProfile(ctx context.Context, address string) error
	FullTextSearch(ctx context.Context, arg FullTextSearchParams) ([]FullTextSearchRow, error)
	GetCategoryByName(ctx context.Context, name string) (Category, error)
	GetCollection(ctx context.Context, arg GetCollectionParams) ([]GetCollectionRow, error)
	GetCollectionLastSyncBlock(ctx context.Context, token string) (int64, error)
	GetCollectionWithCategory(ctx context.Context, arg GetCollectionWithCategoryParams) ([]GetCollectionWithCategoryRow, error)
	GetEvent(ctx context.Context, arg GetEventParams) ([]GetEventRow, error)
	GetExpiredOrder(ctx context.Context) ([]GetExpiredOrderRow, error)
	GetListValidNFT(ctx context.Context, arg GetListValidNFTParams) ([]GetListValidNFTRow, error)
	GetMarketplaceLastSyncBlock(ctx context.Context) (int64, error)
	GetNFTValidConsiderations(ctx context.Context, arg GetNFTValidConsiderationsParams) ([]GetNFTValidConsiderationsRow, error)
	GetNFTsWithPricesPaginated(ctx context.Context, arg GetNFTsWithPricesPaginatedParams) ([]GetNFTsWithPricesPaginatedRow, error)
	GetNotification(ctx context.Context, address sql.NullString) ([]GetNotificationRow, error)
	GetOffer(ctx context.Context, arg GetOfferParams) ([]GetOfferRow, error)
	GetOrder(ctx context.Context, arg GetOrderParams) ([]json.RawMessage, error)
	GetProfile(ctx context.Context, address string) (Profile, error)
	InsertCategory(ctx context.Context, name string) (Category, error)
	InsertCollection(ctx context.Context, arg InsertCollectionParams) (Collection, error)
	InsertEvent(ctx context.Context, arg InsertEventParams) (Event, error)
	InsertNotification(ctx context.Context, arg InsertNotificationParams) (Notification, error)
	InsertOrder(ctx context.Context, arg InsertOrderParams) error
	InsertOrderConsiderationItem(ctx context.Context, arg InsertOrderConsiderationItemParams) error
	InsertOrderOfferItem(ctx context.Context, arg InsertOrderOfferItemParams) error
	MarkOrderInvalid(ctx context.Context, arg MarkOrderInvalidParams) error
	UpdateCollectionLastSyncBlock(ctx context.Context, arg UpdateCollectionLastSyncBlockParams) error
	UpdateMarketplaceLastSyncBlock(ctx context.Context, lastSyncBlock int64) error
	UpdateNft(ctx context.Context, arg UpdateNftParams) error
	UpdateNftMetadata(ctx context.Context, arg UpdateNftMetadataParams) error
	UpdateNftStatus(ctx context.Context, arg UpdateNftStatusParams) (Nft, error)
	UpdateNotification(ctx context.Context, arg UpdateNotificationParams) (Notification, error)
	UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) error
	UpsertNFTV2(ctx context.Context, arg UpsertNFTV2Params) error
	UpsertProfile(ctx context.Context, arg UpsertProfileParams) (Profile, error)
}

var _ Querier = (*Queries)(nil)
