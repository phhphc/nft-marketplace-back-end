package httpApi

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

type HttpServer interface {
	Run(ctx context.Context, address string) error
}

type httpServer struct {
	lg   *log.Logger
	Echo *echo.Echo
}

func NewHttpServer(
	controller controllers.Controller,
) HttpServer {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(RequestLogger())

	e.HTTPErrorHandler = HTTPErrorHandler
	e.Validator = NewValidator()

	s := httpServer{
		lg:   log.GetLogger(),
		Echo: e,
	}
	s.applyRoutes(controller)

	return &s
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
