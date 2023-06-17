package postgresql

import (
	"context"
	"database/sql"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2/gen"
)

func (r *PostgresqlRepository) FindOneUser(
	ctx context.Context,
	address string,
) (*entities.User, error) {
	// Holy query roles
	arg := gen.GetUsersParams{
		PublicAddress: sql.NullString{String: address, Valid: true},
		Offset:        0,
		Limit:         1,
	}

	rows, err := r.queries.GetUsers(ctx, arg)
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

func (r *PostgresqlRepository) FindUser(
	ctx context.Context,
	isBlock bool,
	role string,
	offset int32,
	limit int32,
) ([]*entities.User, error) {
	arg := gen.GetUsersParams{
		IsBlock: sql.NullBool{Bool: isBlock, Valid: true},
		Role:    sql.NullString{String: role, Valid: role != ""},
		Offset:  offset,
		Limit:   limit,
	}

	rows, err := r.queries.GetUsers(ctx, arg)
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

func (r *PostgresqlRepository) FindRoles(
	ctx context.Context,
) ([]*entities.Role, error) {
	res, err := r.queries.GetAllRoles(
		ctx,
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error find")
		return nil, err
	}

	rs := make([]*entities.Role, len(res))
	for i, row := range res {
		rs[i] = &entities.Role{
			Id:   int(row.ID),
			Name: row.Name,
		}
	}

	return rs, nil
}
