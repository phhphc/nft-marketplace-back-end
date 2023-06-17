package services

import (
	"context"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type UserService interface {
	GetUserByAddress(ctx context.Context, address string) (*entities.User, error)
	GetUsers(ctx context.Context, isBlock bool, role string, offset int32, limit int32) ([]*entities.User, error)
	UpdateUserBlockState(ctx context.Context, address string, isBlock bool) error
	InsertUserRole(ctx context.Context, address string, roleId int32) (*entities.Role, error)
	DeleteUserRole(ctx context.Context, address string, roleID int32) error
}

func (s *Services) GetUserByAddress(ctx context.Context, address string) (*entities.User, error) {
	return s.userReader.FindOneUser(
		ctx,
		address,
	)
}

func (s *Services) GetUsers(ctx context.Context, isBlock bool, role string, offset int32, limit int32) ([]*entities.User, error) {
	return s.userReader.FindUser(
		ctx,
		isBlock,
		role,
		offset,
		limit,
	)
}

func (s *Services) InsertUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	return s.userWriter.InsertUser(
		ctx,
		user,
	)
}

func (s *Services) UpdateUserBlockState(ctx context.Context, address string, isBlock bool) error {
	return s.userWriter.UpdateUserBlockState(
		ctx,
		address,
		isBlock,
	)
}

func (s *Services) InsertUserRole(ctx context.Context, address string, roleId int32) (*entities.Role, error) {
	return s.userWriter.InsertUserRole(
		ctx,
		address,
		roleId,
	)
}

func (s *Services) DeleteUserRole(ctx context.Context, address string, roleId int32) error {
	return s.userWriter.DeleteUserRole(
		ctx,
		address,
		roleId,
	)
}

func (s *Services) TransferAdminRole(ctx context.Context, maker string, taker string) (*entities.Role, error) {
	return s.userWriter.TransferAdminRole(
		ctx,
		maker,
		taker,
	)
}
