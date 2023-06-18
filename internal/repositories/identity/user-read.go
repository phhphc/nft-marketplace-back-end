package identity

import (
	"context"
	"database/sql"
	"encoding/json"

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

type roleDb struct {
	Role   string `json:"role"`
	RoleId int    `json:"role_id"`
}

type userDb struct {
	Address string   `json:"address"`
	Nonce   string   `json:"nonce"`
	Roles   []roleDb `json:"roles"`
	IsBlock bool     `json:"is_block"`
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

	res, err := r.queries.GetUsers(ctx, arg)
	if err != nil {
		return nil, err
	}

	users := make([]*entities.User, len(res))
	for i, row := range res {
		var udb userDb
		err = json.Unmarshal(row, &udb)
		if err != nil {
			r.lg.Error().Caller().Err(err).Msg("error unmarshal")
			return nil, err
		}

		rdb := udb.Roles
		roles := make([]entities.Role, len(rdb))
		for i, r := range rdb {
			roles[i] = entities.Role{
				Id:   r.RoleId,
				Name: r.Role,
			}
		}

		users[i] = &entities.User{
			Address: udb.Address,
			Nonce:   udb.Nonce,
			IsBlock: udb.IsBlock,
			Roles:   roles,
		}
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
