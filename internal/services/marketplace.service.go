package services

import "context"

type MarketplaceService interface {
	UpdateMarketplaceLastSyncBlock(ctx context.Context, block uint64) error
	GetMarketplaceLastSyncBlock(ctx context.Context) (uint64, error)
}

func (s *Services) UpdateMarketplaceLastSyncBlock(ctx context.Context, block uint64) error {
	err := s.repo.UpdateMarketplaceLastSyncBlock(ctx, int64(block))
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update block")
	}
	return err
}

func (s *Services) GetMarketplaceLastSyncBlock(ctx context.Context) (uint64, error) {
	lastSyncBlock, err := s.repo.GetMarketplaceLastSyncBlock(ctx)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error get last block")
	}
	return uint64(lastSyncBlock), err
}
