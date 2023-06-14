package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phhphc/nft-marketplace-back-end/configs"
)

var (
	IsLoggedIn  = isLoggedIn()
	IsAdmin     = isAdmin
	IsModerator = isModerator
	IsUser      = isUser
)

type AddressRequired struct {
	Address string `query:"address" validate:"required,eth_addr"`
}

func isLoggedIn() echo.MiddlewareFunc {
	cfg, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}
	secret := []byte(cfg.JwtSecret)
	jwt := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: secret,
	})
	return jwt
}

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		roles := claims["roles"].([]string)
		for _, role := range roles {
			if role == "admin" {
				return next(c)
			}
		}
		return echo.ErrUnauthorized
	}
}

func isModerator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		roles := claims["roles"].([]string)
		for _, role := range roles {
			if role == "moderator" {
				return next(c)
			}
		}
		return echo.ErrUnauthorized
	}
}

func isUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		roles := claims["roles"].([]string)
		for _, role := range roles {
			if role == "user" {
				return next(c)
			}
		}
		return echo.ErrUnauthorized
	}
}
