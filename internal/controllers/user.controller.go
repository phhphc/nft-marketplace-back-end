package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers/dto"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type UserController interface {
	GetUsers(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateBlockState(c echo.Context) error
	CreateUserRole(c echo.Context) error
	DeleteUserRole(c echo.Context) error
}

type GetUsersReq struct {
	Offset  int32  `query:"offset" validate:"gte=0"`
	Limit   int32  `query:"limit" validate:"gte=0,lte=100"`
	IsBlock bool   `query:"is_block"`
	Role    string `query:"role"`
}

type GetUsersRes struct {
	Users  []*entities.User `json:"users"`
	Offset int32            `json:"offset"`
	Limit  int32            `json:"limit"`
}

type GetUserReq struct {
	Address string `param:"address" validate:"required,eth_addr"`
}

type GetUserRes struct {
	User *entities.User `json:"user"`
}

type UpdateBlockStateReq struct {
	Address string `param:"address" validate:"required,eth_addr"`
	IsBlock bool   `json:"is_block"`
}

type UpdateBlockStateRes struct {
	IsSuccess bool `json:"is_success"`
}

type CreateUserRoleReq struct {
	Address string `json:"address" validate:"required,eth_addr"`
	RoleId  int32  `json:"role_id" validate:"required"`
}
type CreateUserRoleRes struct {
	IsSuccess bool   `json:"is_success"`
	Address   string `json:"address"`
	RoleId    int32  `json:"role_id"`
}
type DeleteUserRoleReq struct {
	Address string `json:"address" validate:"required,eth_addr"`
	RoleId  int32  `json:"role_id" validate:"required"`
}
type DeleteUserRoleRes struct {
	IsSuccess bool   `json:"is_success"`
	Address   string `json:"address"`
}

func (ctl *Controls) GetUsers(c echo.Context) error {
	var req GetUsersReq
	if err := c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err := c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	users, err := ctl.service.GetUsers(c.Request().Context(), req.IsBlock, req.Role, req.Offset, req.Limit)
	if err != nil {
		return dto.NewHTTPError(400, err)
	}
	return c.JSON(200, dto.Response{
		Data: GetUsersRes{
			Users:  users,
			Offset: req.Offset,
			Limit:  req.Limit,
		},
		IsSuccess: true,
	})
}

func (ctl *Controls) GetUser(c echo.Context) error {
	var req GetUserReq
	if err := c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err := c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	user, err := ctl.service.GetUserByAddress(c.Request().Context(), req.Address)
	if err != nil {
		return dto.NewHTTPError(400, err)
	}
	return c.JSON(200, dto.Response{
		Data: GetUserRes{
			User: user,
		},
		IsSuccess: true,
	})
}

func (ctl *Controls) UpdateBlockState(c echo.Context) error {
	var req UpdateBlockStateReq
	if err := c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err := c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	err := ctl.service.UpdateUserBlockState(c.Request().Context(), req.Address, req.IsBlock)
	if err != nil {
		return dto.NewHTTPError(400, err)
	}
	return c.JSON(200, dto.Response{
		Data: UpdateBlockStateRes{
			IsSuccess: true,
		},
		IsSuccess: true,
	})
}

func (ctl *Controls) CreateUserRole(c echo.Context) error {
	var req CreateUserRoleReq
	if err := c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err := c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	role, err := ctl.service.InsertUserRole(c.Request().Context(), req.Address, req.RoleId)
	if err != nil {
		return dto.NewHTTPError(400, err)
	}
	return c.JSON(200, dto.Response{
		Data: CreateUserRoleRes{
			IsSuccess: true,
			Address:   req.Address,
			RoleId:    int32(role.Id),
		},
		IsSuccess: true,
	})
}

func (ctl *Controls) DeleteUserRole(c echo.Context) error {
	var req DeleteUserRoleReq
	if err := c.Bind(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}
	if err := c.Validate(&req); err != nil {
		return dto.NewHTTPError(400, err)
	}

	err := ctl.service.DeleteUserRole(c.Request().Context(), req.Address, req.RoleId)
	if err != nil {
		return dto.NewHTTPError(400, err)
	}
	return c.JSON(200, dto.Response{
		Data: DeleteUserRoleRes{
			IsSuccess: true,
			Address:   req.Address,
		},
		IsSuccess: true,
	})
}
