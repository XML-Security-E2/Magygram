package handler

import (
	"ads-service/domain/service-contracts"
	"github.com/labstack/echo"
)

type CampaignHandler interface {
	CreateCampaign(c echo.Context) error
}

type campaignHandler struct {
	CampaignService service_contracts.CampaignService
}

func NewCampaignHandler(p service_contracts.CampaignService) CampaignHandler {
	return &campaignHandler{p}
}

func (ch campaignHandler) CreateCampaign(c echo.Context) error {
	panic("implement me")
}
