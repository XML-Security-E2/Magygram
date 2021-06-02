package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"story-service/domain/service-contracts"
)

type StoryHandler interface {
	CreateStory(c echo.Context) error
	GetStoriesForStoryline(c echo.Context) error
	GetStoriesForUser(c echo.Context) error
}

type storyHandler struct {
	StoryService service_contracts.StoryService
}


func NewStoryHandler(p service_contracts.StoryService) StoryHandler {
	return &storyHandler{p}
}

func (p storyHandler) CreateStory(c echo.Context) error {
	fmt.Println("UDJE")
	headers, err := c.FormFile("images")
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

func (p storyHandler) GetStoriesForStoryline(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	stories, err := p.StoryService.GetStoriesForStoryline(ctx,bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stories)
}

func (p storyHandler) GetStoriesForUser(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	stories, err := p.StoryService.GetStoriesForUser(ctx,userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stories)
}