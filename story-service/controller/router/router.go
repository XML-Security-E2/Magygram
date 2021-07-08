package router

import (
	"story-service/controller/handler"

	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/story", h.CreateStory, h.LoggingMiddleware)
	e.POST("/api/story/campaign", h.CreateStoryCampaign)
	e.POST("/api/story/campaign/agent", h.CreateStoryCampaignFromApi)
	e.GET("/api/story/campaign", h.GetUserStoryCampaign)

	e.POST("/api/story/highlights", h.GetStoryHighlight, h.LoggingMiddleware)
	e.GET("/api/story", h.GetStoriesForStoryline, h.LoggingMiddleware)
	e.GET("/api/story/:userId", h.GetStoriesForUser, h.LoggingMiddleware)
	e.GET("/api/story/id/:storyId", h.GetStoryForUserMessage, h.LoggingMiddleware)
	e.PUT("/api/story/campaign/influencer", h.CreateStoryCampaignInfluencer,h.LoggingMiddleware)

	e.GET("/api/story/user", h.GetAllUserStories, h.LoggingMiddleware)
	e.PUT("/api/story/:storyId/visited", h.VisitedStoryByUser, h.LoggingMiddleware)
	e.GET("/api/story/activestories", h.HaveActiveStoriesLoggedUser, h.LoggingMiddleware)
	e.PUT("/api/story/:requestId/delete", h.DeleteStory, h.LoggingMiddleware)
	e.PUT("/api/story/user-info", h.UpdateUserInfo)
	e.GET("/api/story/:storyId/getForAdmin", h.GetStoryForAdmin, h.LoggingMiddleware)

	e.POST("/api/story/media/first/preview", h.GetStoryMediaAndWebsiteByIds)
	e.GET("/api/story/metrics", echo.WrapHandler(promhttp.Handler()))

}
