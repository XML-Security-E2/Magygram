package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"story-service/domain/model"
	"story-service/domain/service-contracts"
	"story-service/domain/service-contracts/exceptions/expired"
	"story-service/domain/service-contracts/exceptions/unauthorized"
	"story-service/logger"
	"story-service/tracer"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type StoryHandler interface {
	CreateStory(c echo.Context) error
	CreateStoryCampaign(c echo.Context) error
	GetStoriesForStoryline(c echo.Context) error
	GetStoryForAdmin(c echo.Context) error
	GetStoriesForUser(c echo.Context) error
	GetAllUserStories(c echo.Context) error
	VisitedStoryByUser(c echo.Context) error
	GetStoryHighlight(c echo.Context) error
	HaveActiveStoriesLoggedUser(c echo.Context) error
	LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	DeleteStory(c echo.Context) error
	UpdateUserInfo(c echo.Context) error
	GetStoryForUserMessage(c echo.Context) error
	GetUserStoryCampaign(c echo.Context) error
	GetStoryMediaAndWebsiteByIds(c echo.Context) error
}

type storyHandler struct {
	StoryService service_contracts.StoryService
	tracer       opentracing.Tracer
	closer       io.Closer
}

func NewStoryHandler(p service_contracts.StoryService) StoryHandler {
	tracer, closer := tracer.Init("post-service")
	opentracing.SetGlobalTracer(tracer)
	return &storyHandler{p, tracer, closer}
}

func (p storyHandler) GetUserStoryCampaign(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerGetUserStoryCampaign", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get user story campaign at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	stories, err := p.StoryService.GetAllUserStoryCampaigns(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stories)
}

func (p storyHandler) GetStoryForAdmin(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerGetStoryForAdmin", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get story for admin at %s\n", c.Path())),
	)

	storyId := c.Param("storyId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	post, err := p.StoryService.GetStoryForAdmin(ctx, storyId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}
func (p storyHandler) DeleteStory(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerDeleteStory", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling delete story at %s\n", c.Path())),
	)

	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	err := p.StoryService.DeleteStory(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p storyHandler) LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.LoggingEntry = logger.Logger.WithFields(logrus.Fields{"request_ip": c.RealIP()})
		return next(c)
	}
}

func (p storyHandler) GetStoryMediaAndWebsiteByIds(c echo.Context) error {
	request := &model.FollowedUsersResponse{}

	if err := c.Bind(request); err != nil {
		return err
	}
	fmt.Println(len(request.Users))
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	retVal, err := p.StoryService.GetStoryMediaAndWebsiteByIds(ctx, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, retVal)
}

func (p storyHandler) CreateStoryCampaign(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerCreateStoryCampaign", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create story campaign at %s\n", c.Path())),
	)

	headers, err := c.FormFile("images")
	tagsString := c.FormValue("tags")
	var tags []model.Tag
	json.Unmarshal([]byte(tagsString), &tags)

	minD := c.FormValue("minDisplays")
	minDisplays, _ := strconv.Atoi(minD)
	frequency := c.FormValue("frequency")
	minA := c.FormValue("minAge")
	minAge, _ := strconv.Atoi(minA)
	maxA := c.FormValue("maxAge")
	maxAge, _ := strconv.Atoi(maxA)
	gender := c.FormValue("gender")

	dateF := c.FormValue("startDate")
	dateFr, _ := strconv.ParseInt(dateF, 10, 64)
	dateFrom := time.Unix(0, dateFr*int64(time.Millisecond))

	dateT := c.FormValue("endDate")
	dateTt, _ := strconv.ParseInt(dateT, 10, 64)
	dateTo := time.Unix(0, dateTt*int64(time.Millisecond))

	exposeD := c.FormValue("exposeOnceDate")
	exposeDa, _ := strconv.ParseInt(exposeD, 10, 64)
	exposeDate := time.Unix(0, exposeDa*int64(time.Millisecond))

	displayT := c.FormValue("displayTime")
	displayTime, _ := strconv.Atoi(displayT)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	campaignRequest := &model.CampaignRequest{
		MinDisplaysForRepeatedly: minDisplays,
		Frequency:                model.CampaignFrequency(frequency),
		TargetGroup: model.TargetGroup{
			MinAge: minAge,
			MaxAge: maxAge,
			Gender: model.GenderType(gender),
		},
		DateFrom:       dateFrom,
		DateTo:         dateTo,
		Type:           "STORY",
		DisplayTime:    displayTime,
		ContentId:      "",
		ExposeOnceDate: exposeDate,
	}

	storyId, err := p.StoryService.CreateStoryCampaign(ctx, bearer, headers, tags, campaignRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, storyId)
}

func (p storyHandler) CreateStory(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerCreateStory", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create story at %s\n", c.Path())),
	)

	headers, err := c.FormFile("images")
	tagsString := c.FormValue("tags")
	var tags []model.Tag
	json.Unmarshal([]byte(tagsString), &tags)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	storyId, err := p.StoryService.CreatePost(ctx, bearer, headers, tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, storyId)
}

func (p storyHandler) GetStoriesForStoryline(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerGetStoriesForStoryline", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get stories for storyline at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	stories, err := p.StoryService.GetStoriesForStoryline(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if stories == nil {
		return c.JSON(http.StatusOK, []model.StoryResponse{})
	}

	return c.JSON(http.StatusOK, stories)
}

func (p storyHandler) GetStoriesForUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerGetStoriesForUser", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get stories for users at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	stories, err := p.StoryService.GetStoriesForUser(ctx, userId, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stories)
}

func (p storyHandler) VisitedStoryByUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerVisitedStoryByUser", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling visited stories for users at %s\n", c.Path())),
	)

	storyId := c.Param("storyId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	err := p.StoryService.VisitedStoryByUser(ctx, storyId, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p storyHandler) GetStoryForUserMessage(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerGetStoryForUserMessage", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get story for user message at %s\n", c.Path())),
	)

	storyId := c.Param("storyId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	story, userInfo, err := p.StoryService.GetStoryForUserMessage(ctx, bearer, storyId)
	if err != nil {
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *unauthorized.UnauthorizedAccessError:
			return echo.NewHTTPError(http.StatusUnauthorized, userInfo)
		case *expired.StoryError:
			return echo.NewHTTPError(http.StatusForbidden, userInfo)
		}
	}

	return c.JSON(http.StatusOK, story)
}

func (p storyHandler) GetAllUserStories(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerGetAllUserStories", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get all user stories at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	stories, err := p.StoryService.GetAllUserStories(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if stories == nil {
		return c.JSON(http.StatusOK, []model.StoryResponse{})
	}
	fmt.Println("test3")

	return c.JSON(http.StatusOK, stories)
}

func (p storyHandler) GetStoryHighlight(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerGetStoryHighlight", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get story highlight at %s\n", c.Path())),
	)

	highRequest := &model.HighlightRequest{}
	if err := c.Bind(highRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	highlight, err := p.StoryService.GetStoryHighlight(ctx, bearer, highRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, highlight)
}

func (p storyHandler) HaveActiveStoriesLoggedUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerHaveActiveStoriesLoggedUser", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling have active stories logged user at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")
	retVal, err := p.StoryService.HaveActiveStoriesLoggedUser(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (p storyHandler) UpdateUserInfo(c echo.Context) error {
	span := tracer.StartSpanFromRequest("StoryHandlerUpdateUserInfo", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling update user info at %s\n", c.Path())),
	)

	userInfo := &model.UserInfo{}
	if err := c.Bind(userInfo); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	err := p.StoryService.EditStoryOwnerInfo(ctx, bearer, userInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
