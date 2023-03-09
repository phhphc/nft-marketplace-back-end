package httpApi

import (
	"context"
	orderControllers "github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/interfaces/controllers"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/interfaces/repository"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/services"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers"
	"github.com/phhphc/nft-marketplace-back-end/pkg/clients"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

type HttpServer interface {
	Run(ctx context.Context, address string) error
}

type httpServer struct {
	lg   *log.Logger
	Echo *echo.Echo
}

func NewHttpServer(postgreClient *clients.PostgreClient) HttpServer {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(RequestLogger())

	e.HTTPErrorHandler = HTTPErrorHandler
	e.Validator = NewValidator()

	nftController := controllers.NewNftController(postgreClient.Database)

	nftRoute := e.Group("/api/v0.1/nft")
	nftRoute.GET("/:contract_addr/:token_id", nftController.GetNft)
	nftRoute.GET("", nftController.GetNftsOfCollection)

	ver1group := e.Group("/api/v0.1")
	orderRepository := repository.NewRepository(postgreClient.Database)
	orderService := services.NewOrderService(orderRepository)
	orderControllers.NewOrderController(ver1group, orderService)

	return &httpServer{
		lg:   log.GetLogger(),
		Echo: e,
	}
}

func (s *httpServer) Run(ctx context.Context, address string) error {
	go func() {
		<-ctx.Done()

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer shutdownCancel()
		err := s.Echo.Shutdown(shutdownCtx)
		if err != nil {
			s.lg.Fatal().Err(err).Caller().Msg("error shutdown server")
		}
	}()

	err := s.Echo.Start(address)
	if err != nil && err != http.ErrServerClosed {
		s.lg.Fatal().Err(err).Caller().Msg("error echo server")
	}
	return err
}
