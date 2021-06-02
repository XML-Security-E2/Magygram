package router

import (
	"github.com/labstack/echo"
	"post-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/posts", h.CreatePost)
	e.GET("/api/posts", h.GetPostsForTimeline)
	e.PUT("/api/posts/:postId/like", h.LikePost)
	e.PUT("/api/posts/:postId/unlike", h.UnlikePost)
	e.PUT("/api/posts/:postId/dislike", h.DislikePost)
	e.PUT("/api/posts/:postId/undislike", h.UndislikePost)
	e.GET("/api/posts/:postId/image", h.GetPostsFirstImage)
	e.PUT("/api/posts/comments", h.AddComment)
}