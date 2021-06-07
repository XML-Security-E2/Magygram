package router

import (
	"github.com/labstack/echo"
	"story-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/story", h.CreateStory)
	e.POST("/api/story/highlights", h.GetStoryHighlight)
	e.GET("/api/story", h.GetStoriesForStoryline)
	e.GET("/api/story/:userId", h.GetStoriesForUser)
	e.GET("/api/story/user", h.GetAllUserStories)
	e.PUT("/api/story/:storyId/visited", h.VisitedStoryByUser)
	e.GET("/api/story/activestories", h.HaveActiveStoriesLoggedUser)

}