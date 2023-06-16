package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
)

type UserService interface {
	GetUserByAddress(ctx context.Context, address string) (*entities.User, error)
	GetUsers(ctx context.Context, isBlock bool, role string, offset int32, limit int32) ([]*entities.User, error)
	UpdateUserBlockState(ctx context.Context, address string, isBlock bool) error
	InsertUserRole(ctx context.Context, address string, roleId int32) (*entities.Role, error)
	DeleteUserRole(ctx context.Context, address string, roleID int32) error
	//InitAdmin(ctx context.Context, address string) (*entities.User, error)
}

func (s *Services) GetUserByAddress(ctx context.Context, address string) (*entities.User, error) {
	// Holy query roles
	arg := postgresql.GetUsersParams{
		PublicAddress: sql.NullString{String: address, Valid: true},
		Offset:        0,
		Limit:         1,
	}

	rows, err := s.repo.GetUsers(ctx, arg)
	if err != nil {
		return nil, err
	}

	var user *entities.User
	// combines the roles of user
	for _, row := range rows {
		if user == nil {
			user = &entities.User{
				Address: row.PublicAddress,
				Nonce:   row.Nonce,
				Roles:   []entities.Role{{Id: int(row.RoleID.Int32), Name: row.Role.String}},
				IsBlock: row.IsBlock,
			}
		} else {
			role := entities.Role{Id: int(row.RoleID.Int32), Name: row.Role.String}
			user.Roles = append(user.Roles, role)
		}
	}

	return user, nil
}

func (s *Services) GetUsers(ctx context.Context, isBlock bool, role string, offset int32, limit int32) ([]*entities.User, error) {
	arg := postgresql.GetUsersParams{
		IsBlock: sql.NullBool{Bool: isBlock, Valid: true},
		Role:    sql.NullString{String: role, Valid: role != ""},
		Offset:  offset,
		Limit:   limit,
	}

	rows, err := s.repo.GetUsers(ctx, arg)
	if err != nil {
		return nil, err
	}

	var user2Address = make(map[string]*entities.User)
	for _, row := range rows {
		if user2Address[row.PublicAddress] == nil {
			user2Address[row.PublicAddress] = &entities.User{
				Address: row.PublicAddress,
				Nonce:   row.Nonce,
				Roles:   []entities.Role{{Id: int(row.RoleID.Int32), Name: row.Role.String}},
				IsBlock: row.IsBlock,
			}
		} else {
			role := entities.Role{Id: int(row.RoleID.Int32), Name: row.Role.String}
			user2Address[row.PublicAddress].Roles = append(user2Address[row.PublicAddress].Roles, role)
		}
	}

	var users []*entities.User
	for _, user := range user2Address {
		users = append(users, user)
	}

	return users, nil
}

func (s *Services) InsertUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	arg := postgresql.InsertUserParams{
		PublicAddress: user.Address,
		Nonce:         user.Nonce,
	}

	row, err := s.repo.InsertUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	for _, role := range user.Roles {
		arg := postgresql.InsertUserRoleParams{
			Address: user.Address,
			RoleID:  int32(role.Id),
		}
		_, err := s.repo.InsertUserRole(ctx, arg)
		if err != nil {
			if err.Error() == "pq: duplicate key value violates unique constraint \"user_roles_pkey\"" {
				continue
			}
		}
	}
	return &entities.User{
		Address: row.PublicAddress,
		Nonce:   row.Nonce,
	}, nil
}

func (s *Services) UpdateUserBlockState(ctx context.Context, address string, isBlock bool) error {
	user, err := s.GetUserByAddress(ctx, address)
	if err != nil {
		return err
	}

	for _, role := range user.Roles {
		if role.Id == 1 {
			return fmt.Errorf("can not block admin")
		}
	}

	arg := postgresql.UpdateUserBlockStateParams{
		PublicAddress: address,
		IsBlock:       isBlock,
	}
	_, err = s.repo.UpdateUserBlockState(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (s *Services) InsertUserRole(ctx context.Context, address string, roleId int32) (*entities.Role, error) {
	arg := postgresql.InsertUserRoleParams{
		Address: address,
		RoleID:  roleId,
	}
	role, err := s.repo.InsertUserRole(ctx, arg)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"user_roles_pkey\"" {
			return nil, fmt.Errorf("user already have role id %d", role.RoleID)
		}
		return nil, err
	}

	return &entities.Role{
		Id: int(role.RoleID),
	}, nil
}

func (s *Services) DeleteUserRole(ctx context.Context, address string, roleId int32) error {
	arg := postgresql.DeleteUserRoleParams{
		Address: address,
		RoleID:  roleId,
	}
	err := s.repo.DeleteUserRole(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (s *Services) TransferAdminRole(ctx context.Context, maker string, taker string) (*entities.Role, error) {
	arg1 := postgresql.InsertUserRoleParams{
		Address: taker,
		RoleID:  1,
	}
	role, err := s.repo.InsertUserRole(ctx, arg1)
	if err != nil {
		return nil, err
	}

	arg2 := postgresql.DeleteUserRoleParams{
		Address: maker,
		RoleID:  1,
	}
	err = s.repo.DeleteUserRole(ctx, arg2)
	if err != nil {
		return nil, err
	}

	return &entities.Role{
		Id: int(role.RoleID),
	}, nil
}
