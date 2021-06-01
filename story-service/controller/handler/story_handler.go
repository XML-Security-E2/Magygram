package handler

import (
	"context"
	"github.com/labstack/echo"
	"net/http"
	"story-service/domain/service-contracts"
)

type StoryHandler interface {
	CreateStory(c echo.Context) error

}

type storyHandler struct {
	StoryService service_contracts.StoryService
}

func NewStoryHandler(p service_contracts.StoryService) StoryHandler {
	return &storyHandler{p}
}

func (p storyHandler) CreateStory(c echo.Context) error {

	headers, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")

	storyId, err := p.StoryService.CreatePost(ctx, bearer, headers)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, storyId)
}