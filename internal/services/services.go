package services

import (
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/pkg/broker"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

func New(repo postgresql.Querier) *Services {
	return &Services{
		lg:     *log.GetLogger(),
		repo:   repo,
		broker: broker.New("165.232.160.106:9092"),
	}
}

type Services struct {
	lg     log.Logger
	repo   postgresql.Querier
	broker broker.Broker
}
