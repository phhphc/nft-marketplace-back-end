package services

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/tabbed/pqtype"
)

func (s *Services) CreateCollection(ctx context.Context, collection entities.Collection) (ec entities.Collection, err error) {
	//	TODO: use transaction

	category, err := s.repo.GetCategoryByName(ctx, collection.Category)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot find category")
		return
	}

	metadata, err := json.Marshal(collection.Metadata)
	if err != nil {
		s.lg.Panic().Caller().Err(err).Msg("cannot unmashal metadata")
		return
	}

	dbCollection, err := s.repo.InsertCollection(ctx, postgresql.InsertCollectionParams{
		Token:       collection.Token.Hex(),
		Owner:       collection.Owner.Hex(),
		Name:        collection.Name,
		Description: collection.Description,
		Metadata: pqtype.NullRawMessage{
			RawMessage: metadata,
			Valid:      len(metadata) > 0,
		},
		Category: category.ID,
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot create collection")
		return
	}

	// TODO: later
	err = s.EmitEvent(ctx, models.EventNewCollection, collection.Token[:])
	if err != nil {
		s.lg.Panic().Caller().Err(err).Msg("error")
	}

	ec.Token = common.HexToAddress(dbCollection.Token)
	ec.Owner = common.HexToAddress(dbCollection.Owner)
	ec.Name = dbCollection.Name
	ec.Description = dbCollection.Description
	ec.Category = category.Name
	ec.CreatedAt = dbCollection.CreatedAt.Time

	if dbCollection.Metadata.Valid {
		var metadata map[string]any
		if err = json.Unmarshal(dbCollection.Metadata.RawMessage, &metadata); err != nil {
			if err != nil {
				s.lg.Panic().Caller().Err(err).Msg("error")
			}
		}
		ec.Metadata = metadata
	}
	return
}

func (s *Services) GetListCollection(ctx context.Context, query entities.Collection, offset int, limit int) (ec []entities.Collection, err error) {
	params := postgresql.GetCollectionParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	}

	if query.Token != (common.Address{}) {
		params.Token = sql.NullString{
			String: query.Token.Hex(),
			Valid:  true,
		}
	}
	if query.Owner != (common.Address{}) {
		params.Owner = sql.NullString{
			String: query.Owner.Hex(),
			Valid:  true,
		}
	}
	if len(query.Name) > 0 {
		params.Name = sql.NullString{
			String: query.Name,
			Valid:  true,
		}
	}

	cs, err := s.repo.GetCollection(ctx, params)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot get list collection")
		return
	}

	for _, c := range cs {
		e := entities.Collection{
			Token:       common.HexToAddress(c.Token),
			Name:        c.Name,
			Description: c.Description,
			Owner:       common.HexToAddress(c.Owner),
			Category:    c.Category,
			CreatedAt:   c.CreatedAt.Time,
		}
		if c.Metadata.Valid {
			var metadata map[string]any
			if err = json.Unmarshal(c.Metadata.RawMessage, &metadata); err != nil {
				if err != nil {
					s.lg.Panic().Caller().Err(err).Msg("error")
				}
			}
			e.Metadata = metadata
		}

		ec = append(ec, e)
	}

	return
}

func (s *Services) GetListCollectionWithCategory(ctx context.Context, categogy string, offset int, limit int) (ec []entities.Collection, err error) {
	params := postgresql.GetCollectionWithCategoryParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	}

	if len(categogy) > 0 {
		params.Category = sql.NullString{
			String: categogy,
			Valid:  true,
		}
	}

	cs, err := s.repo.GetCollectionWithCategory(ctx, params)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot get list collection with category")
		return
	}

	for _, c := range cs {
		e := entities.Collection{
			Token:       common.HexToAddress(c.Token),
			Name:        c.Name,
			Description: c.Description,
			Owner:       common.HexToAddress(c.Owner),
			Category:    c.Category,
			CreatedAt:   c.CreatedAt.Time,
		}
		if c.Metadata.Valid {
			var metadata map[string]any
			if err = json.Unmarshal(c.Metadata.RawMessage, &metadata); err != nil {
				if err != nil {
					s.lg.Panic().Caller().Err(err).Msg("error")
				}
			}
			e.Metadata = metadata
		}

		ec = append(ec, e)
	}

	return
}