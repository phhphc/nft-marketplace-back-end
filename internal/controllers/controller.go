package controllers

import "github.com/labstack/echo/v4"

type Controller interface {
	PostOrder(c echo.Context) error
	GetOrderHash(c echo.Context) error
	GetOrder(c echo.Context) error

	PostCollection(c echo.Context) error
	GetCollection(c echo.Context) error

	GetNFTsWithListings(c echo.Context) error
	GetNFTWithListings(c echo.Context) error
}

var _ Controller = (*Controls)(nil)
