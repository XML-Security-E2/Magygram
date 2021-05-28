package router

import (
	"github.com/labstack/echo"
	"user-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	//users
	e.POST("/users", h.RegisterUser)
	e.GET("/users/activate/:activationId", h.ActivateUser)
	e.POST("/users/login", h.LoginUser)
	e.POST("/users/reset-password-link-request", h.ResetPasswordRequest)
	e.GET("/users/reset-password/:resetPasswordId", h.ResetPasswordActivation)
	e.POST("/users/reset-password", h.ChangeNewPassword)
	e.GET("/admin-check", h.AdminCheck, h.AuthorizationMiddleware("execute_admin_check"))
	e.GET("/other-check", h.OtherCheck, h.AuthorizationMiddleware("execute_admin_agent_check"))
	e.POST("/users/resend-activation-link", h.ResendActivationLink)
	e.GET("/users/check-existence/:userId", h.GetUserEmailIfUserExist)
	e.GET("/users/:userId", h.GetUserById)
}