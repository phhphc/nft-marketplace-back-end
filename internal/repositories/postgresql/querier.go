// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package postgresql

import (
	"context"
)

type Querier interface {
	DeleteProfile(ctx context.Context, address string) error
	DeleteUserRole(ctx context.Context, arg DeleteUserRoleParams) error
	FullTextSearch(ctx context.Context, arg FullTextSearchParams) ([]FullTextSearchRow, error)
	GetAllRoles(ctx context.Context) ([]Role, error)
	GetMarketplaceLastSyncBlock(ctx context.Context) (int64, error)
	GetMarketplaceSettings(ctx context.Context, arg GetMarketplaceSettingsParams) (GetMarketplaceSettingsRow, error)
	GetNotification(ctx context.Context, arg GetNotificationParams) ([]GetNotificationRow, error)
	GetOffer(ctx context.Context, arg GetOfferParams) ([]GetOfferRow, error)
	GetProfile(ctx context.Context, address string) (Profile, error)
	GetUserByAddress(ctx context.Context, publicAddress string) (User, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUsersRow, error)
	GetValidMarketplaceSettings(ctx context.Context, marketplace string) (GetValidMarketplaceSettingsRow, error)
	InsertMarketplaceSettings(ctx context.Context, arg InsertMarketplaceSettingsParams) (MarketplaceSetting, error)
	InsertNotification(ctx context.Context, arg InsertNotificationParams) (Notification, error)
	InsertUser(ctx context.Context, arg InsertUserParams) (User, error)
	InsertUserRole(ctx context.Context, arg InsertUserRoleParams) (UserRole, error)
	UpdateMarketplaceLastSyncBlock(ctx context.Context, lastSyncBlock int64) error
	UpdateNonce(ctx context.Context, arg UpdateNonceParams) (User, error)
	UpdateNotification(ctx context.Context, arg UpdateNotificationParams) (Notification, error)
	UpdateUserBlockState(ctx context.Context, arg UpdateUserBlockStateParams) (User, error)
	UpsertProfile(ctx context.Context, arg UpsertProfileParams) (Profile, error)
}

var _ Querier = (*Queries)(nil)
