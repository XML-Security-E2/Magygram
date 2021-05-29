package router

import (
	"auth-service/controller/handler"
	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/users", h.RegisterUser)
	e.GET("/api/users/activate/:activationId", h.ActivateUser)
	e.POST("/api/users/login", h.LoginUser)
	e.POST("/api/users/reset-password", h.ResetPassword)
	e.GET("/api/admin-check", h.AdminCheck, h.AuthorizationMiddleware("execute_admin_check"))
	e.GET("/api/other-check", h.OtherCheck, h.AuthorizationMiddleware("execute_admin_agent_check"))
}