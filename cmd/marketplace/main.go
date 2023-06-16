package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	httpApi "github.com/phhphc/nft-marketplace-back-end/api/http-api"
	"github.com/phhphc/nft-marketplace-back-end/configs"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers"
	postgresqlV1 "github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2"
	"github.com/phhphc/nft-marketplace-back-end/internal/services"
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

	wg := sync.WaitGroup{}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	postgresql, err := postgresql.NewPostgresqlRepository(ctx, cfg.PostgreUri)
	if err != nil {
		lg.Panic().Caller().Err(err).Msg("error")
	}
	defer postgresql.Close()

	var repositoryV1 postgresqlV1.Querier = postgresqlV1.New(postgreClient.Database)

	var service services.Servicer = services.New(
		repositoryV1,
		cfg.RedisUrl,
		cfg.RedisPass,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
		postgresql,
	)
	var controller controllers.Controller = controllers.New(service)

	controller.InitMarketplaceSettings()
	controller.InitAdmin()

	wg.Add(1)
	go func() {
		defer wg.Done()
		httpServer := httpApi.NewHttpServer(controller)
		address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
		httpServer.Run(ctx, address)
	}()

	wg.Wait()
}
