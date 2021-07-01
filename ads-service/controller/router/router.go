package router

import (
	"ads-service/controller/handler"
	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/ads/campaign", h.CreateCampaign)
	e.POST("/api/ads/campaign/influencer", h.CreateInfluencerCampaign)
}