package handler

import (
	"ads-service/domain/model"
	"ads-service/domain/service-contracts"
	"ads-service/tracer"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
	"io"
	"net/http"
	"strconv"
	"time"
)

type CampaignHandler interface {
	CreateCampaign(c echo.Context) error
	CreateInfluencerCampaign(c echo.Context) error
	CreateInfluencerCampaignProduct(c echo.Context) error
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
	GetCampaignByPostIdInfulencer(c echo.Context) error
	CreateCampaignFromAgentApi(c echo.Context) error
	GetCampaignStatisticsFromAgentApi(c echo.Context) error
}

type campaignHandler struct {
	CampaignService service_contracts.CampaignService
	tracer      opentracing.Tracer
	closer      io.Closer
}

func NewCampaignHandler(p service_contracts.CampaignService) CampaignHandler {
	tracer, closer := tracer.Init("ads-service")
	opentracing.SetGlobalTracer(tracer)
	return &campaignHandler{p, tracer, closer}
}

func (ch campaignHandler) GetCampaignStatisticsFromAgentApi(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerGetCampaignStatisticsFromAgentApi", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get campaign statistics from agent api at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	stats, err := ch.CampaignService.GetCampaignStatisticsFromAgentApi(ctx, bearer)
	fmt.Println(bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stats)
}

func (ch campaignHandler) CreateCampaignFromAgentApi(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerCreateCampaignFromAgentApi", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create campaign from agent api at %s\n", c.Path())),
	)

	minD := c.FormValue("minDisplaysForRepeatedly")
	minDisplays, _ := strconv.Atoi(minD)

	frequency := c.FormValue("frequency")

	minA := c.FormValue("minAge")
	minAge, _ := strconv.Atoi(minA)
	maxA := c.FormValue("maxAge")
	maxAge, _ := strconv.Atoi(maxA)
	gender := c.FormValue("gender")

	dateF := c.FormValue("dateFrom")
	dateFr, _ := strconv.ParseInt(dateF, 10, 64)
	dateFrom := time.Unix(0, dateFr * int64(time.Millisecond))

	dateT := c.FormValue("dateTo")
	dateTt, _ := strconv.ParseInt(dateT, 10, 64)
	dateTo := time.Unix(0, dateTt * int64(time.Millisecond))

	exposeD := c.FormValue("exposeOnceDate")
	exposeDa, _ := strconv.ParseInt(exposeD, 10, 64)
	exposeDate := time.Unix(0, exposeDa * int64(time.Millisecond))

	displayT := c.FormValue("displayTime")
	displayTime, _ := strconv.Atoi(displayT)

	campaignType := c.FormValue("campaignType")

	headers, err := c.FormFile("images")

	if err != nil {
		fmt.Println(err.Error())
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	req := &model.CampaignApiRequest{
		MinDisplaysForRepeatedly: minDisplays,
		Frequency:                model.CampaignFrequency(frequency),
		TargetGroup:              model.TargetGroup{
			MinAge: minAge,
			MaxAge: maxAge,
			Gender: model.GenderType(gender),
		},
		DisplayTime:              displayTime,
		DateFrom:                 dateFrom,
		DateTo:                   dateTo,
		ExposeOnceDate:           exposeDate,
		Type:                     model.ContentType(campaignType),
		Media:                    headers,
	}

	err = ch.CampaignService.CreateCampaignFromAgentApi(ctx, bearer, req)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}

func (ch campaignHandler) GetPostCampaignStatistic(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerGetPostCampaignStatistic", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get posts campaign statistics at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	campaigns, err := ch.CampaignService.GetPostCampaignStatistic(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaigns)
}

func (ch campaignHandler) GetStoryCampaignStatistic(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerGetStoryCampaignStatistic", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get stories campaign statistics at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")
	campaigns, err := ch.CampaignService.GetStoryCampaignStatistic(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaigns)
}


func (ch campaignHandler) ClickOnStoryCampaignWebsite(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerClickOnStoryCampaignWebsite", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling click on story campaign website at %s\n", c.Path())),
	)

	contentId := c.Param("contentId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	err := ch.CampaignService.ClickOnStoryCampaignWebsite(ctx, contentId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (ch campaignHandler) ClickOnPostCampaignWebsite(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerClickOnPostCampaignWebsite", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling click on post campaign website at %s\n", c.Path())),
	)

	contentId := c.Param("contentId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	err := ch.CampaignService.ClickOnPostCampaignWebsite(ctx, contentId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")}


func (ch campaignHandler) DeleteCampaignByPostId(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerDeleteCampaignByPostId", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling delete campaign by post id at %s\n", c.Path())),
	)

	contentId := c.Param("contentId")
	fmt.Println("USO")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	err := ch.CampaignService.DeleteCampaignByPostId(ctx, bearer, contentId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")}

func (ch campaignHandler) DeleteCampaignByStory(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerDeleteCampaignByStory", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling delete campaign by story id at %s\n", c.Path())),
	)

	contentId := c.Param("contentId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	err := ch.CampaignService.DeleteCampaignByStoryId(ctx, bearer, contentId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (ch campaignHandler) GetPostCampaignSuggestion(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerGetPostCampaignSuggestion", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get post campaign suggestion at %s\n", c.Path())),
	)

	count := c.Param("count")
	countI, _ := strconv.Atoi(count)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	campaign, err := ch.CampaignService.GetUnseenPostIdsCampaignsForUser(ctx, bearer, countI)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("ODJE SI")
	fmt.Println(len(campaign))
	fmt.Println(count)

	return c.JSON(http.StatusOK, campaign)
}

func (ch campaignHandler) GetStoryCampaignSuggestion(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerGetStoryCampaignSuggestion", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get story campaign suggestion at %s\n", c.Path())),
	)

	count := c.Param("count")
	countI, _ := strconv.Atoi(count)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	campaign, err := ch.CampaignService.GetUnseenStoryIdsCampaignsForUser(ctx, bearer, countI)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaign)
}

func (ch campaignHandler) GetCampaignByPostIdInfulencer(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerGetCampaignByPostIdInfulencer", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get campaign by post id influencer at %s\n", c.Path())),
	)

	contentId := c.Param("contentId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	campaign, err := ch.CampaignService.GetCampaignByPostIdInfulencer(ctx, contentId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaign)}

func (ch campaignHandler) GetCampaignByPostId(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerGetCampaignByPostId", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get campaign by post id at %s\n", c.Path())),
	)

	contentId := c.Param("contentId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	campaign, err := ch.CampaignService.GetCampaignByPostId(ctx, bearer, contentId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaign)}

func (ch campaignHandler) GetCampaignByStoryId(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerGetCampaignByStoryId", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get campaign by story id at %s\n", c.Path())),
	)

	contentId := c.Param("contentId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	campaign, err := ch.CampaignService.GetCampaignByStoryId(ctx, bearer, contentId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaign)
}


func (ch campaignHandler) GetCampaignById(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerGetCampaignById", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get campaign by id at %s\n", c.Path())),
	)

	campaignId := c.Param("campaignId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	campaign, err := ch.CampaignService.GetCampaignById(ctx, bearer, campaignId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaign)
}

func (ch campaignHandler) GetAllActiveAgentsStoryCampaigns(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerGetAllActiveAgentsStoryCampaigns", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get all active agents story campaigns at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	campaigns, err := ch.CampaignService.GetAllActiveAgentsStoryCampaigns(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaigns)
}

func (ch campaignHandler) GetAllActiveAgentsPostCampaigns(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerGetAllActiveAgentsPostCampaigns", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get all active agents post campaigns at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	campaigns, err := ch.CampaignService.GetAllActiveAgentsPostCampaigns(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaigns)
}

func (ch campaignHandler) UpdateCampaignRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerUpdateCampaignRequest", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling update campaign request at %s\n", c.Path())),
	)

	campaignRequest := &model.CampaignUpdateRequestTimeDTO{}
	if err := c.Bind(campaignRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

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
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, campaignId)
}

func (ch campaignHandler) CreateCampaign(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerCreateCampaign", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create campaign at %s\n", c.Path())),
	)

	campaignRequest := &model.CampaignRequest{}
	if err := c.Bind(campaignRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	campaignId, err := ch.CampaignService.CreateCampaign(ctx, bearer, campaignRequest)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, campaignId)
}

func (ch campaignHandler) CreateInfluencerCampaign(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerCreateInfluencerCampaign", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create influencer campaign at %s\n", c.Path())),
	)

	campaignRequest := &model.InfluencerCampaignCreateRequest{}
	fmt.Println("USO1")
	if err := c.Bind(campaignRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	campaignId, err := ch.CampaignService.CreateInfluencerCampaign(ctx, bearer, campaignRequest)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, campaignId)
}
func (ch campaignHandler) CreateInfluencerCampaignProduct(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerCreateInfluencerCampaignProduct", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create influencer campaign product at %s\n", c.Path())),
	)

	campaignRequest := &model.InfluencerCampaignProductCreateRequest{}
	fmt.Println("USO1")
	if err := c.Bind(campaignRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	campaignId, err := ch.CampaignService.CreateCampaignForInfluencer(ctx, campaignRequest)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, campaignId)
}

func (ch campaignHandler) UpdatePostCampaignVisitor(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerUpdatePostCampaignVisitor", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling update post campaign visitor at %s\n", c.Path())),
	)

	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	err := ch.CampaignService.UpdateCampaignVisitor(ctx, bearer, postId, "POST")
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (ch campaignHandler) UpdateStoryCampaignVisitor(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AdsHandlerUpdateStoryCampaignVisitor", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling update story campaign visitor at %s\n", c.Path())),
	)

	storyId := c.Param("storyId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	err := ch.CampaignService.UpdateCampaignVisitor(ctx, bearer, storyId, "STORY")
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}