package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/phhphc/nft-marketplace-back-end/configs"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/identity"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/internal/services"
	dataPuller "github.com/phhphc/nft-marketplace-back-end/internal/worker/data-puller"
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

	wg := sync.WaitGroup{}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	postgresql, err := postgresql.NewPostgresqlRepository(ctx, cfg.PostgreUri)
	if err != nil {
		lg.Panic().Caller().Err(err).Msg("error")
	}
	defer postgresql.Close()

	identity, err := identity.NewIdentityRepository(ctx, cfg.PostgreIdentityUri)
	if err != nil {
		lg.Panic().Caller().Err(err).Msg("error")
	}
	defer postgresql.Close()
	service := services.New(
		cfg.RedisUrl,
		cfg.RedisPass,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
		identity,
		identity,
		identity,
		identity,
	)
	defer func() {
		err := service.Close()
		if err != nil {
			lg.Error().Caller().Err(err).Msg("fail to close")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		lg.Info().Caller().Msg("Start data puller")
		dataPuller, err := dataPuller.NewDataPuller(service)
		if err != nil {
			lg.Fatal().Err(err).Caller().Msg("error create data puller")
		}
		dataPuller.Run(ctx)
	}()

	wg.Wait()
}
