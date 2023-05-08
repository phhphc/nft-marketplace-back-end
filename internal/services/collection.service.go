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

type CollectionService interface {
	CreateCollection(ctx context.Context, collection entities.Collection) (entities.Collection, error)
	GetListCollection(ctx context.Context, query entities.Collection, offset int, limit int) ([]entities.Collection, error)
	GetListCollectionWithCategory(ctx context.Context, categogy string, offset int, limit int) ([]entities.Collection, error)
	UpdateCollectionLastSyncBlock(ctx context.Context, token common.Address, block uint64) error
	GetCollectionLastSyncBlock(ctx context.Context, token common.Address) (uint64, error)
}

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

	payload := models.NewCollectionTask{
		Address: collection.Token,
	}
	bs, err := json.Marshal(payload)
	if err != nil {
		s.lg.Panic().Caller().Err(err).Msg("error")
	}
	err = s.EmitTask(ctx, models.TaskNewCollection, bs)
	s.lg.Info().Caller().Str("token", collection.Token.Hex()).Msg("emit task new collection")
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

func (s *Services) UpdateCollectionLastSyncBlock(ctx context.Context, token common.Address, block uint64) error {
	err := s.repo.UpdateCollectionLastSyncBlock(ctx, postgresql.UpdateCollectionLastSyncBlockParams{
		Token:         token.Hex(),
		LastSyncBlock: int64(block),
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update")
	}
	return nil
}

func (s *Services) GetCollectionLastSyncBlock(ctx context.Context, token common.Address) (uint64, error) {
	block, err := s.repo.GetCollectionLastSyncBlock(ctx, token.Hex())
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot get last sync block")
	}

	return uint64(block), err
}
