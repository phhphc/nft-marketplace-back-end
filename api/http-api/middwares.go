package httpApi

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

func RequestLogger() echo.MiddlewareFunc {
	lg := log.GetLogger()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			latency := time.Since(start)
			byteIn, _ := strconv.ParseInt(req.Header.Get(echo.HeaderContentLength), 10, 64)

			lg.Info().
				Str("remote_ip", c.RealIP()).
				Str("host", req.Host).
				Str("method", req.Method).
				Str("uri", req.RequestURI).
				Str("user_agent", req.UserAgent()).
				Int("status", res.Status).
				Err(err).
				Dur("latency", latency).
				Str("latency_human", latency.String()).
				Int64("bytes_in", byteIn).
				Int64("bytes_out", res.Size).
				Msg("request income")
			return nil
		}
	}
}

func HTTPErrorHandler(err error, c echo.Context) {

	var code int
	var res models.Response
	if he, ok := err.(*dto.HTTPError); ok {
		code = he.Code
		res.Error = he.Err.Error()
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		res.Error = he.Message
	} else {
		code = http.StatusInternalServerError
		res.Error = "internal server error"
	}

	c.JSON(code, res)
}
