package router

import (
	"github.com/labstack/echo"
	"post-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/posts", h.CreatePost, h.LoggingMiddleware)
	e.PUT("/api/posts", h.EditPost, h.LoggingMiddleware)
	e.GET("/api/posts", h.GetPostsForTimeline, h.LoggingMiddleware)
	e.GET("/api/posts/id/:postId", h.GetPostById, h.LoggingMiddleware)
	e.GET("/api/posts/:userId/count", h.GetUsersPostsCount, h.LoggingMiddleware) // iz ms
	e.GET("/api/posts/:userId", h.GetUsersPosts, h.LoggingMiddleware)
	e.PUT("/api/posts/:postId/like", h.LikePost, h.LoggingMiddleware)
	e.PUT("/api/posts/:postId/unlike", h.UnlikePost, h.LoggingMiddleware)
	e.PUT("/api/posts/:postId/dislike", h.DislikePost, h.LoggingMiddleware)
	e.PUT("/api/posts/:postId/undislike", h.UndislikePost, h.LoggingMiddleware)
	e.GET("/api/posts/:postId/image", h.GetPostsFirstImage, h.LoggingMiddleware) // iz ms
	e.PUT("/api/posts/comments", h.AddComment, h.LoggingMiddleware)
	e.GET("/api/posts/hashtag-search/:value/guest", h.SearchPostsByHashTagByGuest, h.LoggingMiddleware)
	e.GET("/api/posts/hashtag/:value/guest", h.GetPostForGuestLineByHashTag, h.LoggingMiddleware)
	e.GET("/api/posts/hashtag/:value/user", h.GetPostForUserTimelineByHashTag, h.LoggingMiddleware)
	e.GET("/api/posts/location-search/:value/guest", h.SearchPostsByLocation, h.LoggingMiddleware)
	e.GET("/api/posts/location/:value/guest", h.GetPostForGuestTimelineByLocation, h.LoggingMiddleware)
	e.GET("/api/posts/location/:value/user", h.GetPostForUserTimelineByLocation, h.LoggingMiddleware)
	e.GET("/api/posts/id/:postId/guest", h.GetPostByIdForGuest, h.LoggingMiddleware)
	e.GET("/api/posts/likedposts", h.GetLikedPosts, h.LoggingMiddleware)
}