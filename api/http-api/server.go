package httpApi

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/internal/services"
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

	// nftController := controllers.NewNftController(postgreClient.Database)

	// nftRoute := e.Group("/api/v0.1/nft")
	// nftRoute.GET("/:contract_addr/:token_id", nftController.GetNFT)
	// nftRoute.GET("", nftController.GetNFTsWithPrices)

	var repository postgresql.Querier = postgresql.New(postgreClient.Database)
	var service services.Servicer = services.New(repository)
	var controller controllers.Controller = controllers.New(service)

	nftRoute := e.Group("/api/v0.1/nft")
	nftRoute.GET("", controller.GetNFTsWithListings)
	nftRoute.PATCH("/:token/:identifier", controller.UpdateNftStatus)
	nftRoute.GET("/:token/:identifier", controller.GetNFTWithListings)

	orderRoute := e.Group("/api/v0.1/order")
	orderRoute.GET("", controller.GetOrder)
	orderRoute.POST("", controller.PostOrder)

	collectionRoute := e.Group("/api/v0.1/collection")
	collectionRoute.POST("", controller.PostCollection)
	collectionRoute.GET("", controller.GetCollection)
	collectionRoute.GET("/:category", controller.GetCollectionWithCategory)

	profileRoute := e.Group("/api/v0.1/profile")
	profileRoute.GET("/:address", controller.GetProfile)
	profileRoute.POST("", controller.PostProfile)
	profileRoute.GET("/offer", controller.GetOffer)

	eventRoute := e.Group("/api/v0.1/event")
	eventRoute.GET("", controller.GetEvent)

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
