package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"story-service/domain/model"
	"story-service/domain/service-contracts"
	"story-service/domain/service-contracts/exceptions/expired"
	"story-service/domain/service-contracts/exceptions/unauthorized"
	"story-service/logger"
	"strconv"
	"time"
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
}

type storyHandler struct {
	StoryService service_contracts.StoryService
}

func (p storyHandler) GetUserStoryCampaign(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	stories, err := p.StoryService.GetAllUserStoryCampaigns(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stories)
}

func (p storyHandler) GetStoryForAdmin(c echo.Context) error {

		storyId := c.Param("storyId")

		ctx := c.Request().Context()
			if ctx == nil {
			ctx = context.Background()
		}
		post, err := p.StoryService.GetStoryForAdmin(ctx, storyId)
		if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

		return c.JSON(http.StatusOK, post)
}
func (p storyHandler) DeleteStory(c echo.Context) error {
	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
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

func NewStoryHandler(p service_contracts.StoryService) StoryHandler {
	return &storyHandler{p}
}

func (p storyHandler) CreateStoryCampaign(c echo.Context) error {
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
	dateFrom := time.Unix(0, dateFr * int64(time.Millisecond))

	dateT := c.FormValue("endDate")
	dateTt, _ := strconv.ParseInt(dateT, 10, 64)
	dateTo := time.Unix(0, dateTt * int64(time.Millisecond))

	exposeD := c.FormValue("exposeOnceDate")
	exposeDa, _ := strconv.ParseInt(exposeD, 10, 64)
	exposeDate := time.Unix(0, exposeDa * int64(time.Millisecond))

	displayT := c.FormValue("displayTime")
	displayTime, _ := strconv.Atoi(displayT)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")

	campaignRequest := &model.CampaignRequest{
		MinDisplaysForRepeatedly: minDisplays,
		Frequency:                model.CampaignFrequency(frequency),
		TargetGroup:              model.TargetGroup{
			MinAge: minAge,
			MaxAge: maxAge,
			Gender: model.GenderType(gender),
		},
		DateFrom:                 dateFrom,
		DateTo:                   dateTo,
		Type: "STORY",
		DisplayTime: displayTime,
		ContentId: "",
		ExposeOnceDate: exposeDate,
	}

	storyId, err := p.StoryService.CreateStoryCampaign(ctx, bearer, headers, tags, campaignRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, storyId)
}

func (p storyHandler) CreateStory(c echo.Context) error {
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
	bearer := c.Request().Header.Get("Authorization")

	storyId, err := p.StoryService.CreatePost(ctx, bearer, headers, tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, storyId)
}

func (p storyHandler) GetStoriesForStoryline(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	stories, err := p.StoryService.GetStoriesForStoryline(ctx,bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if stories==nil{
		return c.JSON(http.StatusOK, []model.StoryResponse{})
	}

	return c.JSON(http.StatusOK, stories)
}

func (p storyHandler) GetStoriesForUser(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	stories, err := p.StoryService.GetStoriesForUser(ctx,userId,bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stories)
}

func (p storyHandler) VisitedStoryByUser(c echo.Context) error {
	storyId := c.Param("storyId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	err := p.StoryService.VisitedStoryByUser(ctx,storyId,bearer);
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p storyHandler) GetStoryForUserMessage(c echo.Context) error {
	storyId := c.Param("storyId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	story, userInfo, err := p.StoryService.GetStoryForUserMessage(ctx, bearer, storyId)
	if err != nil{
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
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	stories, err := p.StoryService.GetAllUserStories(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if stories==nil{
		return c.JSON(http.StatusOK, []model.StoryResponse{})
	}
	fmt.Println("test3")

	return c.JSON(http.StatusOK, stories)
}

func (p storyHandler) GetStoryHighlight(c echo.Context) error {
	highRequest := &model.HighlightRequest{}
	if err := c.Bind(highRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	highlight, err := p.StoryService.GetStoryHighlight(ctx, bearer, highRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, highlight)
}


func (p storyHandler) HaveActiveStoriesLoggedUser(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	retVal, err := p.StoryService.HaveActiveStoriesLoggedUser(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (p storyHandler) UpdateUserInfo(c echo.Context) error {
	userInfo := &model.UserInfo{}
	if err := c.Bind(userInfo); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := p.StoryService.EditStoryOwnerInfo(ctx, bearer, userInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}