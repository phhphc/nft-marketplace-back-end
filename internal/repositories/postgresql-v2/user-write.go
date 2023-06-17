package postgresql

import (
	"context"
	"fmt"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2/gen"
)

func (r *PostgresqlRepository) InsertUser(
	ctx context.Context,
	user *entities.User,
) (*entities.User, error) {
	arg := gen.InsertUserParams{
		PublicAddress: user.Address,
		Nonce:         user.Nonce,
	}

	row, err := r.queries.InsertUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	for _, role := range user.Roles {
		arg := gen.InsertUserRoleParams{
			Address: user.Address,
			RoleID:  int32(role.Id),
		}
		_, err := r.queries.InsertUserRole(ctx, arg)
		if err != nil {
			if err.Error() == "pq: duplicate key value violates unique constraint \"user_role_pkey\"" {
				continue
			}
		}
	}
	return &entities.User{
		Address: row.PublicAddress,
		Nonce:   row.Nonce,
	}, nil
}

func (r *PostgresqlRepository) UpdateUserBlockState(
	ctx context.Context,
	address string,
	isBlock bool,
) error {
	user, err := r.FindOneUser(ctx, address)
	if err != nil {
		return err
	}

	for _, role := range user.Roles {
		if role.Id == 1 {
			return fmt.Errorf("can not block admin")
		}
	}

	arg := gen.UpdateUserBlockStateParams{
		PublicAddress: address,
		IsBlock:       isBlock,
	}
	_, err = r.queries.UpdateUserBlockState(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresqlRepository) InsertUserRole(
	ctx context.Context,
	address string, roleId int32) (*entities.Role, error) {
	if roleId == 1 {
		return nil, fmt.Errorf("can not insert role admin")
	}
	arg := gen.InsertUserRoleParams{
		Address: address,
		RoleID:  roleId,
	}
	role, err := r.queries.InsertUserRole(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &entities.Role{
		Id: int(role.RoleID),
	}, nil
}

func (r *PostgresqlRepository) DeleteUserRole(
	ctx context.Context,
	address string, roleId int32) error {
	if roleId == 1 {
		return fmt.Errorf("can not delete role admin")
	}

	arg := gen.DeleteUserRoleParams{
		Address: address,
		RoleID:  roleId,
	}
	err := r.queries.DeleteUserRole(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresqlRepository) TransferAdminRole(
	ctx context.Context,
	maker string, taker string) (*entities.Role, error) {
	arg1 := gen.InsertUserRoleParams{
		Address: taker,
		RoleID:  1,
	}
	role, err := r.queries.InsertUserRole(ctx, arg1)
	if err != nil {
		return nil, err
	}

	arg2 := gen.DeleteUserRoleParams{
		Address: maker,
		RoleID:  1,
	}
	err = r.queries.DeleteUserRole(ctx, arg2)
	if err != nil {
		return nil, err
	}

	return &entities.Role{
		Id: int(role.RoleID),
	}, nil
}

func (r *PostgresqlRepository) UpdateUserNonce(
	ctx context.Context,
	address string,
	nonce string,
) (*entities.User, error) {
	// Check if the user is in the database
	arg := gen.UpdateNonceParams{
		Nonce:         nonce,
		PublicAddress: address,
	}
	res, err := r.queries.UpdateNonce(
		ctx,
		arg,
	)
	if err != nil {
		return nil, err
	}
	user := entities.User{
		Address: res.PublicAddress,
		Nonce:   res.Nonce,
	}
	return &user, nil
}
