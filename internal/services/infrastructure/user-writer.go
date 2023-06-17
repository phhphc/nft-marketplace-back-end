package infrastructure

import (
	"context"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type UserWriter interface {
	InsertUser(
		ctx context.Context,
		user *entities.User,
	) (*entities.User, error)

	UpdateUserBlockState(
		ctx context.Context,
		address string,
		isBlock bool,
	) error

	InsertUserRole(
		ctx context.Context,
		address string,
		roleId int32,
	) (*entities.Role, error)

	DeleteUserRole(
		ctx context.Context,
		address string,
		roleId int32,
	) error

	TransferAdminRole(
		ctx context.Context,
		maker string,
		taker string,
	) (*entities.Role, error)

	UpdateUserNonce(
		ctx context.Context,
		address string,
		nonce string,
	) (*entities.User, error)
}
