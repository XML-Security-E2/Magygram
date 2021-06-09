package router

import (
	"auth-service/controller/handler"
	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/users", h.RegisterUser, h.UserLoggingMiddleware)
	e.GET("/api/users/activate/:userId", h.ActivateUser, h.UserLoggingMiddleware)
	e.POST("/api/users/reset-password", h.ResetPassword, h.UserLoggingMiddleware)
	e.POST("/api/auth/login", h.LoginUser, h.AuthLoggingMiddleware)
	e.GET("/api/auth/logged-user", h.GetLoggedUserId, h.AuthLoggingMiddleware)
	e.GET("/api/auth/admin-check", h.AdminCheck, h.AuthLoggingMiddleware)
	e.GET("/api/auth/has-role", h.AuthorizationSuccess, h.AuthorizationMiddleware())
}