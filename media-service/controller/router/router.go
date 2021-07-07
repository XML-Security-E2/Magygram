package router

import (
	"media-service/controller/handler"

	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/media", h.SaveMedia)
	e.Static("/api/media", "./files")
	e.GET("/api/media/metrics", echo.WrapHandler(promhttp.Handler()))
}
