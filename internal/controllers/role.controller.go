package controllers

import "github.com/labstack/echo/v4"

type RoleController interface {
	GetRoles(c echo.Context) error
}

func (ctl *Controls) GetRoles(c echo.Context) error {
	roles, err := ctl.service.GetRoles(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(200, roles)
}
