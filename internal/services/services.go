package services

import (
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/pkg/asyncQueue"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

func New(repo postgresql.Querier) *Services {
	return &Services{
		lg:     *log.GetLogger(),
		repo:   repo,
		asynq: 	asyncQueue.New("165.232.160.106:6379", "12345"),
	}
}

type Services struct {
	lg     log.Logger
	repo   postgresql.Querier
	asynq  asyncQueue.AsyncQueue
}
