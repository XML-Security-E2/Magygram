package router

import (
	"github.com/labstack/echo"
	"request-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/requests/verification", h.CreateVerificationRequest)
	e.POST("/api/report", h.CreateReportRequest)
	e.GET("/api/report", h.GetReportRequests)
	e.GET("/api/requests/verification", h.GetVerificationRequests)
	e.PUT("/api/report/:requestId/delete", h.DeleteReportRequest)
	e.PUT("/api/requests/verification/:requestId/approve", h.ApproveVerificationRequest)
	e.PUT("/api/requests/verification/:requestId/reject", h.RejectVerificationRequest)
	e.GET("/api/requests/verification/has-pending-request", h.HasUserPendingRequest)
	e.POST("/api/requests/agent-registration", h.CreateAgentRegistrationRequest)
	e.GET("/api/requests/agent-registration", h.GetAgentRegistrationRequests)
	e.PUT("/api/requests/agent-registration/:requestId/approve", h.ApproveAgentRegistrationRequest)
	e.PUT("/api/requests/agent-registration/:requestId/reject", h.RejectAgentRegistrationRequest)
}