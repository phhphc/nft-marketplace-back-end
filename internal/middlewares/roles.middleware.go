package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phhphc/nft-marketplace-back-end/configs"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
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
		Claims:     &entities.Claims{},
	})
	return jwt
}

func OrMiddleware(handlers ...echo.MiddlewareFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, handler := range handlers {
				if err := handler(next)(c); err == nil {
					return nil
				}
			}
			return echo.ErrUnauthorized
		}
	}
}

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*entities.Claims)
		roles := claims.Roles
		for _, role := range roles {
			if role.Name == "admin" {
				return next(c)
			}
		}
		return echo.ErrUnauthorized
	}
}

func isModerator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*entities.Claims)
		roles := claims.Roles
		for _, role := range roles {
			if role.Name == "moderator" {
				return next(c)
			}
		}
		return echo.ErrUnauthorized
	}
}

func isUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*entities.Claims)
		roles := claims.Roles
		for _, role := range roles {
			if role.Name == "user" {
				return next(c)
			}
		}
		return echo.ErrUnauthorized
	}
}
