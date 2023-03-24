package services

import (
	"context"

	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/segmentio/kafka-go"
)

func (s *Services) EmitEvent(ctx context.Context, event models.EnumEvent, value []byte, key []byte) error {
	producer := s.broker.Producer()

	err := producer.WriteMessages(ctx, kafka.Message{
		Topic: string(event),
		Key:   key,
		Value: value,
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("err")
	}
	return err
}

func (s *Services) SubcribeEvent(ctx context.Context, event models.EnumEvent, ch chan<- models.AppEvent) (func(), <-chan error) {
	consumer := s.broker.Consumer(string(event), "app_service")

	cCtx, cancel := context.WithCancel(ctx)
	errCh := make(chan error)

	go func() {
		for {
			m, err := consumer.ReadMessage(cCtx)
			if err != nil {
				if err != context.Canceled {
					s.lg.Error().Caller().Err(err).Msg("err")
					errCh <- err
				}
				break
			}

			ch <- models.AppEvent{
				Key:   m.Key,
				Value: m.Value,
			}
		}
		close(errCh)
	}()

	return cancel, errCh
}
