package router

import (
	"github.com/labstack/echo"
	"story-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/story", h.CreateStory)
	e.GET("/api/story", h.GetStoriesForStoryline)
	e.GET("/api/story/:userId", h.GetStoriesForUser)
}