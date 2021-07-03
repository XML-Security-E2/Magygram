package router

import (
	"auth-service/controller/handler"
	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/users", h.RegisterUser, h.UserLoggingMiddleware)
	e.GET("/api/users/activate/:userId", h.ActivateUser, h.UserLoggingMiddleware)
	e.POST("/api/users/reset-password", h.ResetPassword, h.UserLoggingMiddleware)
	e.POST("/api/auth/login/firststage", h.LoginFirstStage, h.AuthLoggingMiddleware)
	e.POST("/api/auth/login/secondstage", h.LoginSecondStage, h.AuthLoggingMiddleware)
	e.GET("/api/auth/logged-user", h.GetLoggedUserId, h.AuthLoggingMiddleware)
	e.GET("/api/auth/admin-check", h.AdminCheck, h.AuthLoggingMiddleware)
	e.GET("/api/auth/has-role", h.AuthorizationSuccess, h.AuthorizationMiddleware())
	e.POST("/api/users/agent", h.RegisterAgent, h.UserLoggingMiddleware)
	e.GET("/api/auth/generate-campaign-jwt-token", h.GenerateNewAgentCampaignJWTToken, h.AuthLoggingMiddleware)
	e.DELETE("/api/auth/delete-campaign-jwt-token", h.DeleteCampaignJWTToken, h.AuthLoggingMiddleware)
	e.GET("/api/auth/get-campaign-jwt-token", h.GetCampaignJWTToken, h.AuthLoggingMiddleware)
}