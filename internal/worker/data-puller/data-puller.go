package dataPuller

import (
	"context"
	"sync"

	"github.com/phhphc/nft-marketplace-back-end/internal/services"
	"github.com/phhphc/nft-marketplace-back-end/pkg/clients"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

type DataPuller interface {
	Run(ctx context.Context) error
}

type worker struct {
	lg        *log.Logger
	Service   services.Servicer
	ethClient *clients.EthClient
}

func NewDataPuller(service services.Servicer, ethClient *clients.EthClient) (DataPuller, error) {
	return &worker{
		lg: log.GetLogger(),

		Service:   service,
		ethClient: ethClient,
	}, nil
}

func (w *worker) Run(ctx context.Context) error {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go w.pullErc721Metadata(ctx, &wg)

	wg.Wait()
	return nil
}
