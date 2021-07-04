package handler

import (
	"ads-service/domain/model"
	"ads-service/domain/service-contracts"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

type CampaignHandler interface {
	CreateCampaign(c echo.Context) error
	CreateInfluencerCampaign(c echo.Context) error
	UpdateCampaignRequest(c echo.Context) error
	GetAllActiveAgentsPostCampaigns(c echo.Context) error
	GetAllActiveAgentsStoryCampaigns(c echo.Context) error
	GetCampaignById(c echo.Context) error
	GetCampaignByPostId(c echo.Context) error
	GetCampaignByStoryId(c echo.Context) error
	DeleteCampaignByPostId(c echo.Context) error
	DeleteCampaignByStory(c echo.Context) error
	GetPostCampaignSuggestion(c echo.Context) error
	GetStoryCampaignSuggestion(c echo.Context) error
	ClickOnStoryCampaignWebsite(c echo.Context) error
	ClickOnPostCampaignWebsite(c echo.Context) error
	GetPostCampaignStatistic(c echo.Context) error
	GetStoryCampaignStatistic(c echo.Context) error
	UpdatePostCampaignVisitor(c echo.Context) error
	UpdateStoryCampaignVisitor(c echo.Context) error
}

type campaignHandler struct {
	CampaignService service_contracts.CampaignService
}


func NewCampaignHandler(p service_contracts.CampaignService) CampaignHandler {
	return &campaignHandler{p}
}

func (ch campaignHandler) GetPostCampaignStatistic(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	campaigns, err := ch.CampaignService.GetPostCampaignStatistic(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaigns)
}

func (ch campaignHandler) GetStoryCampaignStatistic(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	campaigns, err := ch.CampaignService.GetStoryCampaignStatistic(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaigns)
}


func (ch campaignHandler) ClickOnStoryCampaignWebsite(c echo.Context) error {
	contentId := c.Param("contentId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := ch.CampaignService.ClickOnStoryCampaignWebsite(ctx, contentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (ch campaignHandler) ClickOnPostCampaignWebsite(c echo.Context) error {
	contentId := c.Param("contentId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := ch.CampaignService.ClickOnPostCampaignWebsite(ctx, contentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")}


func (ch campaignHandler) DeleteCampaignByPostId(c echo.Context) error {
	contentId := c.Param("contentId")
	fmt.Println("USO")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	err := ch.CampaignService.DeleteCampaignByPostId(ctx, bearer, contentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")}

func (ch campaignHandler) DeleteCampaignByStory(c echo.Context) error {
	contentId := c.Param("contentId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	err := ch.CampaignService.DeleteCampaignByStoryId(ctx, bearer, contentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (ch campaignHandler) GetPostCampaignSuggestion(c echo.Context) error {
	count := c.Param("count")
	countI, _ := strconv.Atoi(count)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	campaign, err := ch.CampaignService.GetUnseenPostIdsCampaignsForUser(ctx, bearer, countI)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaign)
}

func (ch campaignHandler) GetStoryCampaignSuggestion(c echo.Context) error {
	count := c.Param("count")
	countI, _ := strconv.Atoi(count)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	campaign, err := ch.CampaignService.GetUnseenStoryIdsCampaignsForUser(ctx, bearer, countI)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaign)
}

func (ch campaignHandler) GetCampaignByPostId(c echo.Context) error {
	contentId := c.Param("contentId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	campaign, err := ch.CampaignService.GetCampaignByPostId(ctx, bearer, contentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaign)}

func (ch campaignHandler) GetCampaignByStoryId(c echo.Context) error {
	contentId := c.Param("contentId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	campaign, err := ch.CampaignService.GetCampaignByStoryId(ctx, bearer, contentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaign)
}


func (ch campaignHandler) GetCampaignById(c echo.Context) error {
	campaignId := c.Param("campaignId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	campaign, err := ch.CampaignService.GetCampaignById(ctx, bearer, campaignId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaign)
}

func (ch campaignHandler) GetAllActiveAgentsStoryCampaigns(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	campaigns, err := ch.CampaignService.GetAllActiveAgentsStoryCampaigns(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaigns)
}

func (ch campaignHandler) GetAllActiveAgentsPostCampaigns(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	campaigns, err := ch.CampaignService.GetAllActiveAgentsPostCampaigns(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaigns)
}

func (ch campaignHandler) UpdateCampaignRequest(c echo.Context) error {
	campaignRequest := &model.CampaignUpdateRequestTimeDTO{}
	if err := c.Bind(campaignRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	campaignReq := &model.CampaignUpdateRequestDTO{
		CampaignId:               campaignRequest.CampaignId,
		MinDisplaysForRepeatedly: campaignRequest.MinDisplaysForRepeatedly,
		TargetGroup:              campaignRequest.TargetGroup,
		DateFrom:                 time.Unix(0, campaignRequest.DateFrom * int64(time.Millisecond)),
		DateTo:                   time.Unix(0, campaignRequest.DateTo * int64(time.Millisecond)),
	}

	bearer := c.Request().Header.Get("Authorization")
	campaignId, err := ch.CampaignService.UpdateCampaignRequest(ctx, bearer, campaignReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaignId)
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
	campaignRequest := &model.InfluencerCampaignCreateRequest{}
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

	return c.JSON(http.StatusCreated, campaignId)
}

func (ch campaignHandler) UpdatePostCampaignVisitor(c echo.Context) error {
	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := ch.CampaignService.UpdateCampaignVisitor(ctx, bearer, postId, "POST")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (ch campaignHandler) UpdateStoryCampaignVisitor(c echo.Context) error {
	storyId := c.Param("storyId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := ch.CampaignService.UpdateCampaignVisitor(ctx, bearer, storyId, "STORY")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}