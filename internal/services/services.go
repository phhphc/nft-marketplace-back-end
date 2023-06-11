package services

import (
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/pkg/asyncQueue"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

func New(repo postgresql.Querier, redisUrl string, redisPass string) *Services {
	return &Services{
		lg:    *log.GetLogger(),
		repo:  repo,
		asynq: asyncQueue.New(redisUrl, redisPass),
	}
}

func (s *Services) Close() error {
	return s.asynq.Close()
}

type Services struct {
	lg    log.Logger
	repo  postgresql.Querier
	asynq asyncQueue.AsyncQueue
}
