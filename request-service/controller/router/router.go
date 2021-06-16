package router

import (
	"github.com/labstack/echo"
	"request-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/requests/verification", h.CreateVerificationRequest)
	e.POST("/api/report", h.CreateReportRequest)
	e.GET("/api/requests/verification", h.GetVerificationRequests)
	e.PUT("/api/requests/verification/:requestId/approve", h.ApproveVerificationRequest)
	e.PUT("/api/requests/verification/:requestId/reject", h.RejectVerificationRequest)

}