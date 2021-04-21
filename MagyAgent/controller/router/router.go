package router

import (
	"github.com/labstack/echo"
	"magyAgent/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/users", h.RegisterUser)
	e.GET("/users/activate/:activationId", h.ActivateUser)
	e.POST("/users/login", h.LoginUser)
	e.GET("/admin-check", h.AdminCheck, h.AuthorizationMiddleware("admin"))
	e.GET("/other-check", h.OtherCheck, h.AuthorizationMiddleware("user", "admin"))
}
