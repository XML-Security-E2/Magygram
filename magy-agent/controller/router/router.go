package router

import (
	"github.com/labstack/echo"
	"magyAgent/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	//users
	e.POST("/api/agent/users", h.RegisterUser)
	e.GET("/api/agent/users/activate/:activationId", h.ActivateUser)
	e.POST("/api/agent/users/login", h.LoginUser)
	e.POST("/api/agent/users/reset-password-link-request", h.ResetPasswordRequest)
	e.GET("/api/agent/users/reset-password/:resetPasswordId", h.ResetPasswordActivation)
	e.POST("/api/agent/users/reset-password", h.ChangeNewPassword)
	e.GET("/api/agent/admin-check", h.AdminCheck, h.AuthorizationMiddleware("execute_admin_check"))
	e.GET("/api/agent/other-check", h.OtherCheck, h.AuthorizationMiddleware("execute_admin_agent_check"))
	e.POST("/api/agent/users/resend-activation-link", h.ResendActivationLink)
	e.GET("/api/agent/users/check-existence/:userId", h.GetUserEmailIfUserExist)
	e.GET("/api/agent/users/:userId", h.GetUserById)

	e.POST("/api/agent/products", h.CreateProduct)
	e.PUT("/api/agent/products/:productId", h.UpdateProduct)
	e.PUT("/api/agent/products/:productId/image", h.UpdateProductImage)
	e.GET("/api/agent/products/:productId", h.GetProductById)
	e.GET("/api/agent/products", h.GetAllProducts)
	e.DELETE("/api/agent/products/:productId", h.DeleteProductById)

	e.Static("/api/agent/media", "./files")

}
