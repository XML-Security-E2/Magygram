package router

import (
	"github.com/labstack/echo"
	"magyAgent/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	//users
	e.POST("/api/users", h.RegisterUser)
	e.GET("/api/users/activate/:activationId", h.ActivateUser)
	e.POST("/api/users/login", h.LoginUser)
	e.POST("/api/users/reset-password-link-request", h.ResetPasswordRequest)
	e.GET("/api/users/reset-password/:resetPasswordId", h.ResetPasswordActivation)
	e.POST("/api/users/reset-password", h.ChangeNewPassword)
	e.GET("/api/admin-check", h.AdminCheck, h.AuthorizationMiddleware("execute_admin_check"))
	e.GET("/api/other-check", h.OtherCheck, h.AuthorizationMiddleware("execute_admin_agent_check"))
	e.POST("/api/users/resend-activation-link", h.ResendActivationLink)
	e.GET("/api/users/check-existence/:userId", h.GetUserEmailIfUserExist)
	e.GET("/api/users/:userId", h.GetUserById)

	e.POST("/api/products", h.CreateProduct)
	e.PUT("/api/products/:productId", h.UpdateProduct)
	e.PUT("/api/products/:productId/image", h.UpdateProductImage)
	e.GET("/api/products/:productId", h.GetProductById)
	e.GET("/api/products", h.GetAllProducts)
	e.DELETE("/api/products/:productId", h.DeleteProductById)

	e.Static("/api/media", "./files")

	e.POST("/api/orders", h.CreateOrder)

}
