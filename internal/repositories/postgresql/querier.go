// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package postgresql

import (
	"context"
	"encoding/json"
)

type Querier interface {
	DeleteProfile(ctx context.Context, address string) error
	DeleteUserRole(ctx context.Context, arg DeleteUserRoleParams) error
	FullTextSearch(ctx context.Context, arg FullTextSearchParams) ([]FullTextSearchRow, error)
	GetAllRoles(ctx context.Context) ([]Role, error)
	GetCategoryByName(ctx context.Context, name string) (Category, error)
	GetCollection(ctx context.Context, arg GetCollectionParams) ([]GetCollectionRow, error)
	GetCollectionLastSyncBlock(ctx context.Context, token string) (int64, error)
	GetCollectionWithCategory(ctx context.Context, arg GetCollectionWithCategoryParams) ([]GetCollectionWithCategoryRow, error)
	GetEvent(ctx context.Context, arg GetEventParams) ([]GetEventRow, error)
	GetExpiredOrder(ctx context.Context) ([]GetExpiredOrderRow, error)
	GetMarketplaceLastSyncBlock(ctx context.Context) (int64, error)
	GetMarketplaceSettings(ctx context.Context, marketplace string) (MarketplaceSetting, error)
	GetNotification(ctx context.Context, arg GetNotificationParams) ([]GetNotificationRow, error)
	GetOffer(ctx context.Context, arg GetOfferParams) ([]GetOfferRow, error)
	GetOrder(ctx context.Context, arg GetOrderParams) ([]json.RawMessage, error)
	GetProfile(ctx context.Context, address string) (Profile, error)
	GetUserByAddress(ctx context.Context, publicAddress string) (User, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUsersRow, error)
	InsertCategory(ctx context.Context, name string) (Category, error)
	InsertCollection(ctx context.Context, arg InsertCollectionParams) (Collection, error)
	InsertEvent(ctx context.Context, arg InsertEventParams) (Event, error)
	InsertNotification(ctx context.Context, arg InsertNotificationParams) (Notification, error)
	InsertOrder(ctx context.Context, arg InsertOrderParams) error
	InsertOrderConsiderationItem(ctx context.Context, arg InsertOrderConsiderationItemParams) error
	InsertOrderOfferItem(ctx context.Context, arg InsertOrderOfferItemParams) error
	InsertUser(ctx context.Context, arg InsertUserParams) (User, error)
	InsertUserRole(ctx context.Context, arg InsertUserRoleParams) (UserRole, error)
	MarkOrderInvalid(ctx context.Context, arg MarkOrderInvalidParams) error
	UpdateCollectionLastSyncBlock(ctx context.Context, arg UpdateCollectionLastSyncBlockParams) error
	UpdateMarketplaceLastSyncBlock(ctx context.Context, lastSyncBlock int64) error
	//
	// -- name: GetValidMarketplaceSettings :one
	// SELECT ms.id, ms.marketplace, ms.admin, ms.signer, ms.royalty, ms.sighash, ms.signature, ms.created_at
	// FROM "marketplace_settings" ms
	// WHERE ms.marketplace = sqlc.arg('marketplace')
	// AND ms.signature IS NOT NULL OR ms.id = 1
	// ORDER BY ms.id DESC
	// LIMIT 1;
	// -- name: InsertMarketplaceSettings :one
	// INSERT INTO "marketplace_settings" ("marketplace", "admin", "signer", "royalty", "typed_data", "sighash", "signature", "created_at")
	// VALUES (sqlc.arg('marketplace'), sqlc.arg('admin'), sqlc.arg('signer'), sqlc.arg('royalty'), sqlc.arg('typed_data'), sqlc.arg('sighash'), sqlc.arg('signature'), sqlc.arg('created_at'))
	// RETURNING *;
	UpdateMarketplaceSettings(ctx context.Context, arg UpdateMarketplaceSettingsParams) (MarketplaceSetting, error)
	UpdateNonce(ctx context.Context, arg UpdateNonceParams) (User, error)
	UpdateNotification(ctx context.Context, arg UpdateNotificationParams) (Notification, error)
	UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) error
	UpdateUserBlockState(ctx context.Context, arg UpdateUserBlockStateParams) (User, error)
	UpsertProfile(ctx context.Context, arg UpsertProfileParams) (Profile, error)
}

var _ Querier = (*Queries)(nil)
