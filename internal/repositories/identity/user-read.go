package identity

import (
	"context"
	"database/sql"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/identity/gen"
)

func (r *IdentityRepository) FindOneUser(
	ctx context.Context,
	address string,
) (*entities.User, error) {
	res, err := r.queries.GetUserByAddress(ctx, address)
	r.lg.Debug().Err(err).Msg("error")
	if err != nil {
		return nil, err
	}

	dbr, err := r.queries.GetUserRoles(
		ctx,
		address,
	)

	roles := make([]entities.Role, len(dbr))
	for i, r := range dbr {
		roles[i] = entities.Role{
			Id:   int(r.ID),
			Name: r.Name,
		}
	}

	user := &entities.User{
		Address: res.PublicAddress,
		Nonce:   res.Nonce,
		Roles:   roles,
		IsBlock: res.IsBlock,
	}

	r.lg.Debug().Caller().Err(err).Interface("user", user).Interface("dbr", dbr).Msg("error")

	return user, nil
}

func (r *IdentityRepository) FindUser(
	ctx context.Context,
	isBlock *bool,
	role string,
	offset int32,
	limit int32,
) ([]*entities.User, error) {

	arg := gen.GetUsersParams{
		Role:   sql.NullString{String: role, Valid: role != ""},
		Offset: offset,
		Limit:  limit,
	}

	if isBlock != nil {
		arg.IsBlock = sql.NullBool{Bool: *isBlock, Valid: true}
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

func (r *IdentityRepository) FindRoles(
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
