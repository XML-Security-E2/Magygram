package router

import (
	"github.com/labstack/echo"
	"user-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	//users
	e.POST("/api/users", h.RegisterUser)
	e.GET("/api/users/logged", h.GetLoggedUserInfo)
	e.GET("/api/users/activate/:activationId", h.ActivateUser)
	e.POST("/api/users/reset-password-link-request", h.ResetPasswordRequest)
	e.GET("/api/users/reset-password/:resetPasswordId", h.ResetPasswordActivation)
	e.POST("/api/users/reset-password", h.ChangeNewPassword)
	e.POST("/api/users/resend-activation-link", h.ResendActivationLink)
	e.GET("/api/users/check-existence/:userId", h.GetUserEmailIfUserExist)
	e.GET("/api/users/:userId", h.GetUserById)

	e.POST("/api/users/collections", h.CreateCollection)
	e.POST("/api/users/collections/posts", h.AddPostToCollection)

}