package router

import (
	"github.com/labstack/echo"
	"relationship-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/relationship/follow", h.FollowRequest)
	e.POST("/api/user", h.CreateUser)
	e.GET("/api/relationship/followed-users", h.ReturnFollowedUsers)
	e.GET("/api/relationship/follow-requests", h.ReturnFollowRequests)
}