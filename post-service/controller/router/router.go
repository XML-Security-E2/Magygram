package router

import (
	"github.com/labstack/echo"
	"post-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/posts", h.CreatePost)
	e.PUT("/api/posts", h.EditPost)
	e.GET("/api/posts", h.GetPostsForTimeline)
	e.GET("/api/posts/id/:postId", h.GetPostById)
	e.GET("/api/posts/:userId/count", h.GetUsersPostsCount)
	e.GET("/api/posts/:userId", h.GetUsersPosts)
	e.PUT("/api/posts/:postId/like", h.LikePost)
	e.PUT("/api/posts/:postId/unlike", h.UnlikePost)
	e.PUT("/api/posts/:postId/dislike", h.DislikePost)
	e.PUT("/api/posts/:postId/undislike", h.UndislikePost)
	e.GET("/api/posts/:postId/image", h.GetPostsFirstImage)
	e.PUT("/api/posts/comments", h.AddComment)
	e.GET("/api/posts/hashtag-search/:value/guest", h.SearchPostsByHashTagByGuest)
	e.GET("/api/posts/hashtag/:value/guest", h.GetPostForGuestLineByHashTag)
	e.GET("/api/posts/hashtag/:value/user", h.GetPostForUserTimelineByHashTag)

}