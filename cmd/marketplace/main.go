package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	httpApi "github.com/phhphc/nft-marketplace-back-end/api/http-api"
	"github.com/phhphc/nft-marketplace-back-end/configs"
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

	ethClient, err := clients.NewEthClient(cfg.ChainUrl)
	if err != nil {
		lg.Fatal().Err(err).Caller().Msg("error create eth client")
	}
	defer ethClient.Disconnect()

	wg := sync.WaitGroup{}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		httpServer := httpApi.NewHttpServer(postgreClient)
		address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
		httpServer.Run(ctx, address)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		chainListener, err := chainListener.NewChainListener(postgreClient, ethClient, cfg.MarkeplaceAddr)
		if err != nil {
			lg.Fatal().Err(err).Caller().Msg("error create chain listener")
		}
		chainListener.Run(ctx)
	}()

	wg.Wait()
}
