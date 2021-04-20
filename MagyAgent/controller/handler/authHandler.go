package handler

import (
	"context"
	"github.com/labstack/echo"
	"magyAgent/domain/model"
	service_contracts "magyAgent/domain/service-contracts"
	"net/http"
)

type AuthHandler interface {
	RegisterUser(c echo.Context) error
}

type authHandler struct {
	AuthService service_contracts.AuthService
}

func NewAuthHandler(a service_contracts.AuthService) AuthHandler {
	return &authHandler{a}
}

func (h *authHandler) RegisterUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.AuthService.RegisterUser(ctx, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not Create.")
	}

	return c.JSON(http.StatusCreated, user.Id)
}