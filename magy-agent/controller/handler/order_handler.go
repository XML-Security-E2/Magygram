package handler

import (
	"context"
	"github.com/labstack/echo"
	"magyAgent/domain/model"
	"magyAgent/domain/service-contracts"
	"net/http"
)

type OrderHandler interface {
	CreateOrder(c echo.Context) error
}

type orderHandler struct {
	OrderService service_contracts.OrderService
}
func NewOrderHandler(a service_contracts.OrderService) OrderHandler {
	return &orderHandler{a}
}

func (o orderHandler) CreateOrder(c echo.Context) error {
	orderReq := &model.OrderRequest{}
	if err := c.Bind(orderReq); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	order, err := o.OrderService.CreateOrder(ctx, orderReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusCreated, order.Id)
}
