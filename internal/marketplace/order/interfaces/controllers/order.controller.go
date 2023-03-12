package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/labstack/echo/v4"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/services"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
	"net/http"
)

type (
	Response struct {
		Data      interface{} `json:"data,omitempty""`
		IsSuccess bool        `json:"is_success"`
		Error     interface{} `json:"error,omitempty"`
	}
)

type OrderController interface {
	GetOrder(c echo.Context) error
}

type orderController struct {
	lg           *log.Logger
	orderService services.OrderService
}

func NewOrderController(group *echo.Group, orderService services.OrderService) *orderController {
	controller := &orderController{
		lg:           log.GetLogger(),
		orderService: orderService,
	}
	group.GET("/orders/offer", controller.GetOrderByOfferItem)
	// dummy controller for pseudo Get Listing, need rework
	group.GET("/orders/consideration", controller.GetOrderByConsiderationItem)
	group.GET("/orders/:order_hash", controller.GetOrder)

	group.POST("/orders", controller.CreateOrder)

	return controller
}

func (ctl *orderController) GetOrder(c echo.Context) error {
	orderHash := c.Param("order_hash")
	order, err := ctl.orderService.GetOrder(c.Request().Context(), orderHash)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("err")
		return err
	}

	orderJson, err := json.Marshal(order)
	if err != nil {
		return c.JSON(400, Response{
			Error:     err,
			IsSuccess: false,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Data:      orderJson,
		IsSuccess: true,
	})
}

func (ctl *orderController) GetOrderByOfferItem(c echo.Context) error {
	tokenAddress := c.QueryParam("token_address")
	tokenId := c.QueryParam("token_id")

	fmt.Println("tokenAddress: ", tokenAddress)
	tkId, err := hexutil.DecodeBig(tokenId)

	fmt.Println("tokenId: ", tkId)

	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("err")
		return err
	}

	orders, err := ctl.orderService.GetAllOrderByOfferItem(
		c.Request().Context(),
		common.HexToAddress(tokenAddress),
		tkId,
	)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("err")
		return err
	}

	return c.JSON(http.StatusOK, Response{
		Data:      orders,
		IsSuccess: true,
	})
}

func (ctl *orderController) GetOrderByConsiderationItem(c echo.Context) error {
	tokenAddress := c.QueryParam("token_address")
	tokenId := c.QueryParam("token_id")
	tkId, err := hexutil.DecodeBig(tokenId)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("err")
		return err
	}

	orders, err := ctl.orderService.GetAllOrderByConsiderationItem(
		c.Request().Context(),
		common.HexToAddress(tokenAddress),
		tkId,
	)
	if err != nil {
		ctl.lg.Error().Caller().Err(err).Msg("err")
		return err
	}

	return c.JSON(http.StatusOK, Response{
		Data:      orders,
		IsSuccess: true,
	})
}

func (ctl *orderController) CreateOrder(c echo.Context) error {
	var orderForm OrderForm
	err := c.Bind(&orderForm)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response{
			Error:     err,
			IsSuccess: false,
		})
	}

	order := orderForm.MapToDomainOrder()

	err = ctl.orderService.CreateOrder(c.Request().Context(), order)

	if err != nil {
		return c.JSON(400, Response{
			Data:      orderForm,
			Error:     err,
			IsSuccess: false,
		})
	}

	return c.JSON(http.StatusCreated, Response{
		Data:      order,
		IsSuccess: true,
	})
}
