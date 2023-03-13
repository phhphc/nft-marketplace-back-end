package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/phhphc/nft-marketplace-back-end/configs"
	chainListener "github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/interfaces/chain-listener"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/interfaces/repository"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/services"
	"github.com/phhphc/nft-marketplace-back-end/pkg/clients"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

func main() {
	lg := log.GetLogger()

	cfg, err := configs.LoadConfig()
	if err != nil {
		lg.Fatal().Err(err).Caller().Msg("error load config")
	}

	if cfg.Env == "Dev" {
		log.SetPrettyLogging()
	}

	postgreClient, err := clients.NewPostgreClient(cfg.PostgreUri)
	if err != nil {
		lg.Fatal().Err(err).Caller().Msg("error create postgre client")
	}
	defer postgreClient.Disconnect()

	lg.Info().Caller().Str("chain url", cfg.ChainUrl).Msg("Create new eth client")
	ethClient, err := clients.NewEthClient(cfg.ChainUrl)
	if err != nil {
		lg.Fatal().Err(err).Caller().Msg("error create eth client")
	}
	defer ethClient.Disconnect()

	wg := sync.WaitGroup{}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	orderRepository := repository.NewRepository(postgreClient.Database)
	orderService := services.NewOrderService(orderRepository)

	wg.Add(1)
	go func() {
		defer wg.Done()

		lg.Info().Caller().Msg("Start chain watcher")
		chainListener, err := chainListener.NewChainListener(postgreClient, ethClient, orderService, cfg.MarkeplaceAddr)
		if err != nil {
			lg.Fatal().Err(err).Caller().Msg("error create chain listener")
		}
		chainListener.Run(ctx)
	}()

	wg.Wait()
}
