package router

import (
	"github.com/labstack/echo"
	"request-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/requests", h.CreateVerificationRequest)
	e.POST("/api/report", h.CreateReportRequest)
}