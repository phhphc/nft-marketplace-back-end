package services

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
)

func (s *Services) EmitTask(ctx context.Context, event models.EnumTask, value []byte) error {
	s.lg.Debug().Caller().Str("event", string(event)).Bytes("value", value).Msg("emit task")
	err := s.asynq.DistributeTask(string(event), value)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("err emit task")
	}
	return err
}

func (s *Services) SubcribeTask(ctx context.Context, event models.EnumTask, handler asynq.HandlerFunc) error {
	err := s.asynq.ProcessTask(string(event), handler)

	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("err subcribe task")
	}
	return err
}
