package controllers

import (
	"github.com/labstack/echo/v4"
)

type AuthorizationController interface {
}

func (ctl *Controls) GetRolesByUser(c echo.Context) error {
	return nil
}

func (ctl *Controls) CreateRole(c echo.Context) error {
	return nil
}

func (ctl *Controls) SetUserRole(c echo.Context) error {
	return nil
}
