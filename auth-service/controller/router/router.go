package router

import (
	"auth-service/controller/handler"
	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/users", h.RegisterUser)
	e.GET("/api/users/activate/:userId", h.ActivateUser)
	e.POST("/api/users/reset-password", h.ResetPassword)
	e.POST("/api/auth/login", h.LoginUser)
	e.GET("/api/auth/admin-check", h.AdminCheck, h.AuthorizationMiddleware("execute_admin_check"))
	e.GET("/api/auth/other-check", h.OtherCheck, h.AuthorizationMiddleware("execute_admin_agent_check"))
}