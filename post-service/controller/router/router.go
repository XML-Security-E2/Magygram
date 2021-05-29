package router

import (
	"github.com/labstack/echo"
	"post-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/posts", h.CreatePost)
}