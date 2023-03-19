package services

import (
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

func New(repo postgresql.Querier) *Services {
	return &Services{
		lg:   *log.GetLogger(),
		repo: repo,
	}
}

type Services struct {
	lg   log.Logger
	repo postgresql.Querier
}
