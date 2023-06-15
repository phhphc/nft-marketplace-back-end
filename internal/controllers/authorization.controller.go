package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/configs"
)

type AuthorizationController interface {
	InitAdmin() error
}

func (ctl *Controls) GetRolesByUser(c echo.Context) error {
	return nil
}

func (ctl *Controls) GetRoles(c echo.Context) error {
	return nil
}

func (ctl *Controls) CreateRole(c echo.Context) error {
	return nil
}

func (ctl *Controls) SetUserRole(c echo.Context) error {
	return nil
}

func (ctl *Controls) InitAdmin() error {
	cfg, err := configs.LoadConfig()
	if err != nil {
		return err
	}
	adminAddress := cfg.MarketplaceAdmin
	// Add admin
	_, err = ctl.service.InitAdmin(context.Background(), adminAddress)
	if err != nil {
		ctl.lg.Error().Err(err).Caller().Msg("error init admin")
		return err
	}
	return nil
}
