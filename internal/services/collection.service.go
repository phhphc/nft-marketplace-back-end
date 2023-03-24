package services

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
)

func (s *Services) CreateCollection(ctx context.Context, collection entities.Collection) (ec entities.Collection, err error) {
	//	TODO: use transaction

	category, err := s.repo.GetCategoryByName(ctx, collection.Category)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot find category")
		return
	}

	dbCollection, err := s.repo.InsertCollection(ctx, postgresql.InsertCollectionParams{
		Token:       collection.Token.Hex(),
		Owner:       collection.Owner.Hex(),
		Name:        collection.Name,
		Description: collection.Description,
		Category:    category.ID,
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot create collection")
		return
	}

	// TODO: later
	err = s.EmitEvent(ctx, models.EventNewCollection, collection.Token[:], nil)
	if err != nil {
		s.lg.Panic().Caller().Err(err).Msg("error")
	}

	ec.Token = common.HexToAddress(dbCollection.Token)
	ec.Owner = common.HexToAddress(dbCollection.Owner)
	ec.Name = dbCollection.Name
	ec.Description = dbCollection.Description
	ec.Category = category.Name
	ec.CreatedAt = dbCollection.CreatedAt.Time
	return
}

func (s *Services) GetListCollection(ctx context.Context, offset int, limit int) (ec []entities.Collection, err error) {

	cs, err := s.repo.GetListCollection(ctx, postgresql.GetListCollectionParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot get list collection")
		return
	}

	for _, c := range cs {
		ec = append(ec, entities.Collection{
			Token:       common.HexToAddress(c.Token),
			Name:        c.Name,
			Description: c.Description,
			Owner:       common.HexToAddress(c.Owner),
			Category:    c.Category,
		})
	}

	return
}
