package services

import (
	"context"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type RoleService interface {
	GetRoles(ctx context.Context) ([]*entities.Role, error)
}

func (s *Services) GetRoles(ctx context.Context) ([]*entities.Role, error) {
	res, err := s.repo.GetAllRoles(ctx)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("service cannot get all roles")
		return nil, err
	}
	var roles []*entities.Role
	for _, role := range res {
		roles = append(roles, &entities.Role{
			Id:   int(role.ID),
			Name: role.Name,
		})
	}
	return roles, nil
}
