package router

import (
	"github.com/labstack/echo"
	"magyAgent/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/users", h.RegisterUser)
	e.GET("/users/activate/:activationId", h.ActivateUser)
	e.POST("/users/login", h.LoginUser)
	e.POST("/users/reset-password-link-request", h.ResetPassword)
	e.GET("/users/reset-password/:resetPasswordId", h.ResetPasswordActivation)
	e.POST("/users/reset-password", h.ChangeNewPassword)
	e.GET("/admin-check", h.AdminCheck, h.AuthorizationMiddleware("execute_admin_check"))
	e.GET("/other-check", h.OtherCheck, h.AuthorizationMiddleware("execute_admin_agent_check"))
}
