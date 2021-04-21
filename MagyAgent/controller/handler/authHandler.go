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
	ActivateUser(c echo.Context) error
}

type authHandler struct {
	AuthService service_contracts.AuthService
}

func NewAuthHandler(a service_contracts.AuthService) AuthHandler {
	return &authHandler{a}
}

func (h *authHandler) RegisterUser(c echo.Context) error {
	userRequest := &model.UserRequest{}
	if err := c.Bind(userRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.AuthService.RegisterUser(ctx, userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not Create.")
	}

	return c.JSON(http.StatusCreated, user.Id)
}

func (h *authHandler) ActivateUser(c echo.Context) error {
	activationId := c.Param("activationId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	activated, err := h.AuthService.ActivateUser(ctx, activationId)
	if err != nil || activated == false{
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not be activated.")
	}

	return c.Redirect(http.StatusMovedPermanently, "https://localhost:3000/#/login")//c.JSON(http.StatusNoContent, activationId)
}