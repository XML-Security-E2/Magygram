package router

import (
	"github.com/labstack/echo"
	"story-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/story", h.CreateStory)
}