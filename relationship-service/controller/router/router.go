package router

import (
	"github.com/labstack/echo"
	"relationship-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/relationship/follow", h.FollowRequest)
	e.POST("/api/relationship/unfollow", h.Unfollow)
	e.POST("/api/relationship/accept-follow-request", h.AcceptFollowRequest)
	e.POST("/api/relationship/user", h.CreateUser)
	e.GET("/api/relationship/followed-users/:userId", h.ReturnFollowedUsers)
	e.GET("/api/relationship/following-users/:userId", h.ReturnFollowingUsers)
	e.POST("/api/relationship/is-user-followed", h.IsUserFollowed)
	e.GET("/api/relationship/follow-requests", h.ReturnFollowRequests)
	e.GET("/api/relationship/follow-requests/:objectId", h.ReturnFollowRequestsForUser)

}