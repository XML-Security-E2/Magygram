package handler

import (
	"ads-service/domain/model"
	"ads-service/domain/service-contracts"
	"context"
	"github.com/labstack/echo"
	"net/http"
)

type CampaignHandler interface {
	CreateCampaign(c echo.Context) error
	CreateInfluencerCampaign(c echo.Context) error
}

type campaignHandler struct {
	CampaignService service_contracts.CampaignService
}

func NewCampaignHandler(p service_contracts.CampaignService) CampaignHandler {
	return &campaignHandler{p}
}

func (ch campaignHandler) CreateCampaign(c echo.Context) error {
	campaignRequest := &model.CampaignRequest{}
	if err := c.Bind(campaignRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	campaignId, err := ch.CampaignService.CreateCampaign(ctx, bearer, campaignRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, campaignId)
}

func (ch campaignHandler) CreateInfluencerCampaign(c echo.Context) error {
	campaignRequest := &model.InfluencerCampaignRequest{}
	if err := c.Bind(campaignRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	campaignId, err := ch.CampaignService.CreateInfluencerCampaign(ctx, bearer, campaignRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, campaignId)}
