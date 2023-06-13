package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/phhphc/nft-marketplace-back-end/configs"
	postgresqlV1 "github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2"
	"github.com/phhphc/nft-marketplace-back-end/internal/services"
	chainListener "github.com/phhphc/nft-marketplace-back-end/internal/worker/chain-listener"
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

	postgresql, err := postgresql.NewPostgresqlRepository(ctx, cfg.PostgreUri)
	if err != nil {
		lg.Panic().Caller().Err(err).Msg("error")
	}
	defer postgresql.Close()
	repo := postgresqlV1.New(postgreClient.Database)
	service := services.New(repo, cfg.RedisUrl, cfg.RedisPass, postgresql)
	defer func() {
		err := service.Close()
		if err != nil {
			lg.Error().Caller().Err(err).Msg("fail to close")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		lg.Info().Caller().Msg("Start chain watcher")
		chainListener, err := chainListener.NewChainListener(service, ethClient, cfg.MarkeplaceAddr)
		if err != nil {
			lg.Fatal().Err(err).Caller().Msg("error create chain listener")
		}
		chainListener.Run(ctx)
	}()

	wg.Wait()
}
