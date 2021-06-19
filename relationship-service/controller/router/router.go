package router

import (
	"github.com/labstack/echo"
	"relationship-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/relationship/follow", h.FollowRequest, h.LoggingMiddleware)
	e.POST("/api/relationship/unfollow", h.Unfollow, h.LoggingMiddleware)
	e.POST("/api/relationship/accept-follow-request/:userId", h.AcceptFollowRequest, h.LoggingMiddleware)
	e.POST("/api/relationship/user", h.CreateUser, h.LoggingMiddleware)
	e.GET("/api/relationship/followed-users/:userId", h.ReturnFollowedUsers, h.LoggingMiddleware)
	e.GET("/api/relationship/unmuted-followed-users/:userId", h.ReturnUnmutedFollowedUsers, h.LoggingMiddleware)
	e.GET("/api/relationship/following-users/:userId", h.ReturnFollowingUsers, h.LoggingMiddleware)
	e.POST("/api/relationship/is-user-followed", h.IsUserFollowed, h.LoggingMiddleware)
	e.GET("/api/relationship/follow-requests", h.ReturnFollowRequests, h.LoggingMiddleware)
	e.GET("/api/relationship/follow-requests/:objectId", h.ReturnFollowRequestsForUser, h.LoggingMiddleware)
	e.POST("/api/relationship/mute", h.Mute)
	e.POST("/api/relationship/unmute", h.Unmute)
	e.POST("/api/relationship/is-muted", h.IsMuted)
	e.GET("/api/relationship/recommended-users/:userId", h.ReturnRecommendedUsers)
}