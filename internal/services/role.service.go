package services

import (
	"context"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type RoleService interface {
	GetRoles(ctx context.Context) ([]*entities.Role, error)
}

func (s *Services) GetRoles(ctx context.Context) ([]*entities.Role, error) {
	return s.userReader.FindRoles(
		ctx,
	)

}
