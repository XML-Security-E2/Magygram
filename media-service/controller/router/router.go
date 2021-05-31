package router

import (
	"github.com/labstack/echo"
	"media-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/media", h.SaveMedia)
	e.GET("/api/media/:mediaId", h.GetMedia)
}
