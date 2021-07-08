package router

import (
	"request-service/controller/handler"

	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/requests/verification", h.CreateVerificationRequest)
	e.POST("/api/report", h.CreateReportRequest)
	e.POST("/api/requests/campaign", h.CreateCampaignRequest)
	e.GET("/api/report", h.GetReportRequests)
	e.GET("/api/requests/campaign/:requestId/get", h.GetCampaignRequests)
	e.GET("/api/requests/verification", h.GetVerificationRequests)
	e.PUT("/api/report/:requestId/delete", h.DeleteReportRequest)
	e.PUT("/api/requests/campaign/:requestId/delete", h.DeleteCampaignRequest)
	e.PUT("/api/requests/verification/:requestId/approve", h.ApproveVerificationRequest)
	e.PUT("/api/requests/verification/:requestId/reject", h.RejectVerificationRequest)
	e.GET("/api/requests/verification/has-pending-request", h.HasUserPendingRequest)
	e.POST("/api/requests/agent-registration", h.CreateAgentRegistrationRequest)
	e.GET("/api/requests/agent-registration", h.GetAgentRegistrationRequests)
	e.PUT("/api/requests/agent-registration/:requestId/approve", h.ApproveAgentRegistrationRequest)
	e.PUT("/api/requests/agent-registration/:requestId/reject", h.RejectAgentRegistrationRequest)
	e.GET("/api/requests/metrics", echo.WrapHandler(promhttp.Handler()))
}
