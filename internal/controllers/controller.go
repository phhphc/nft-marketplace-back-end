package controllers

import "github.com/labstack/echo/v4"

type Controller interface {
	PostOrder(c echo.Context) error
	GetOrderHash(c echo.Context) error
	GetOrder(c echo.Context) error

	PostCollection(c echo.Context) error

	GetNftsOfCollection(c echo.Context) error
}

var _ Controller = (*Controls)(nil)