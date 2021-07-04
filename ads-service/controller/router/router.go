package router

import (
	"ads-service/controller/handler"
	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/ads/campaign", h.CreateCampaign)
	e.GET("/api/ads/campaign/:campaignId", h.GetCampaignById)
	e.GET("/api/ads/campaign/post/:contentId", h.GetCampaignByPostId)
	e.GET("/api/ads/campaign/story/:contentId", h.GetCampaignByStoryId)

	e.DELETE("/api/ads/campaign/post/:contentId", h.DeleteCampaignByPostId)
	e.DELETE("/api/ads/campaign/story/:contentId", h.DeleteCampaignByStory)

	e.GET("/api/ads/campaign/post/suggestion/:count", h.GetPostCampaignSuggestion)
	e.GET("/api/ads/campaign/story/suggestion/:count", h.GetStoryCampaignSuggestion)

	e.GET("/api/ads/campaign/post", h.GetAllActiveAgentsPostCampaigns)
	e.GET("/api/ads/campaign/story", h.GetAllActiveAgentsStoryCampaigns)

	e.POST("/api/ads/campaign/influencer", h.CreateInfluencerCampaign)

	e.PUT("/api/ads/campaign", h.UpdateCampaignRequest)

	e.PUT("/api/ads/campaign/post/visited/:postId", h.UpdatePostCampaignVisitor)
	e.PUT("/api/ads/campaign/story/visited/:storyId", h.UpdateStoryCampaignVisitor)

}