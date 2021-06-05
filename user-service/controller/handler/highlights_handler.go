package handler

import (
	"context"
	"github.com/labstack/echo"
	"net/http"
	"user-service/domain/model"
	"user-service/domain/service-contracts"
)

type HighlightsHandler interface {
	CreateHighlights(c echo.Context) error
	GetProfileHighlights(c echo.Context) error
	GetProfileHighlightsByHighlightName(c echo.Context) error
}

type highlightsHandler struct {
	HighlightService service_contracts.HighlightsService
}

func NewHighlightsHandler(u service_contracts.HighlightsService) HighlightsHandler {
	return &highlightsHandler{u}
}

func (h highlightsHandler) CreateHighlights(c echo.Context) error {
	highRequest := &model.HighlightRequest{}

	if err := c.Bind(highRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")

	resp, err := h.HighlightService.CreateHighlights(ctx, bearer, highRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}


func (h highlightsHandler) GetProfileHighlights(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")

	response, err := h.HighlightService.GetProfileHighlights(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}


func (h highlightsHandler) GetProfileHighlightsByHighlightName(c echo.Context) error {
	name := c.Param("name")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")

	response, err := h.HighlightService.GetProfileHighlightsByHighlightName(ctx, bearer, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}