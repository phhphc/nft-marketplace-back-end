package infrastructure

import (
	"context"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type UserReader interface {
	FindOneUser(
		ctx context.Context,
		address string,
	) (*entities.User, error)

	FindUser(
		ctx context.Context,
		isBlock bool,
		role string,
		offset int32,
		limit int32,
	) ([]*entities.User, error)

	FindRoles(
		ctx context.Context,
	) ([]*entities.Role, error)
}
