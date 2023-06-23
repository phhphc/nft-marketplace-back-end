package services

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
)

type CollectionService interface {
	CreateCollection(ctx context.Context, collection entities.Collection) (entities.Collection, error)
	GetListCollection(ctx context.Context, query entities.Collection, offset int, limit int) ([]entities.Collection, error)
	GetListCollectionWithCategory(ctx context.Context, categogy string, offset int, limit int) ([]entities.Collection, error)
	UpdateCollectionLastSyncBlock(ctx context.Context, token common.Address, block uint64) error
	GetCollectionLastSyncBlock(ctx context.Context, token common.Address) (uint64, error)
}

func (s *Services) CreateCollection(ctx context.Context, collection entities.Collection) (ec entities.Collection, err error) {
	ec, err = s.collectionWriter.CreateCollection(
		ctx,
		collection,
	)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error insert")
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

	return
}

func (s *Services) GetListCollection(ctx context.Context, query entities.Collection, offset int, limit int) (ecs []entities.Collection, err error) {
	ecs, err = s.collectionReader.FindCollection(
		ctx,
		query,
		offset,
		limit,
	)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot get list collection")
		return
	}

	return
}

func (s *Services) GetListCollectionWithCategory(ctx context.Context, categogy string, offset int, limit int) (ecs []entities.Collection, err error) {
	ecs, err = s.collectionReader.FindCollection(
		ctx,
		entities.Collection{
			Category: categogy,
		},
		offset,
		limit,
	)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot get list collection")
		return
	}

	return
}

func (s *Services) UpdateCollectionLastSyncBlock(ctx context.Context, token common.Address, block uint64) error {
	err := s.collectionWriter.UpdateCollectionLastSyncBlock(
		ctx,
		token,
		block,
	)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update")
		return err
	}
	return nil
}

func (s *Services) GetCollectionLastSyncBlock(ctx context.Context, token common.Address) (uint64, error) {
	block, err := s.collectionReader.GetCollectionLastSyncBlock(ctx, token)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot get last sync block")
	}

	return uint64(block), err
}
