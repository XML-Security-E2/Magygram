package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"relationship-service/domain/model"
	"relationship-service/logger"
	"relationship-service/service"
)

type FollowHandler interface {
	FollowRequest(ctx echo.Context) error
	Unfollow(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	IsUserFollowed(ctx echo.Context) error
	ReturnFollowedUsers(ctx echo.Context) error
	ReturnUnmutedFollowedUsers(ctx echo.Context) error
	ReturnFollowingUsers(ctx echo.Context) error
	ReturnFollowRequests(ctx echo.Context) error
	AcceptFollowRequest(ctx echo.Context) error
	ReturnFollowRequestsForUser(ctx echo.Context) error
	LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	Mute(ctx echo.Context) error
	Unmute(ctx echo.Context) error
	IsMuted(ctx echo.Context) error
	ReturnRecommendedUsers(ctx echo.Context) error
}

type followHandler struct {
	FollowService service.FollowService
}

func NewFollowHandler(f service.FollowService) FollowHandler {
	return &followHandler{f}
}

func (f followHandler) LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.LoggingEntry = logger.Logger.WithFields(logrus.Fields{"request_ip": c.RealIP()})
		return next(c)
	}
}

func (f followHandler) Mute(ctx echo.Context) error {
	mute := &model.Mute{}
	if err := ctx.Bind(mute); err != nil {
		return err
	}

	if err := f.FollowService.Mute(mute); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "User successfully muted")
}

func (f followHandler) Unmute(ctx echo.Context) error {
	mute := &model.Mute{}
	if err := ctx.Bind(mute); err != nil {
		return err
	}

	if err := f.FollowService.Unmute(mute); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "User successfully muted")
}

func (f *followHandler) FollowRequest(c echo.Context) error {
	followRequest := &model.FollowRequest{}
	if err := c.Bind(followRequest); err != nil {
		return err
	}

	sentRequest, err := f.FollowService.FollowRequest(followRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, sentRequest)
}

func (f *followHandler) Unfollow(c echo.Context) error {
	followRequest := &model.FollowRequest{}
	if err := c.Bind(followRequest); err != nil {
		return err
	}

	err := f.FollowService.Unfollow(followRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (f *followHandler) AcceptFollowRequest(c echo.Context) error {
	userId := c.Param("userId")
	fmt.Println(userId)
	bearer := c.Request().Header.Get("Authorization")
	err := f.FollowService.AcceptFollowRequest(bearer, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, true)
}

func (f *followHandler) CreateUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	if err := f.FollowService.CreateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, true)
}

func (f *followHandler) ReturnFollowedUsers(ctx echo.Context) error {
	user := &model.User{Id: ctx.Param("userId")}

	result, err := f.FollowService.ReturnFollowedUsers(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (f *followHandler) ReturnUnmutedFollowedUsers(ctx echo.Context) error {
	user := &model.User{Id: ctx.Param("userId")}

	result, err := f.FollowService.ReturnUnmutedFollowedUsers(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (f *followHandler) ReturnFollowingUsers(ctx echo.Context) error {
	user := &model.User{Id: ctx.Param("userId")}

	result, err := f.FollowService.ReturnFollowingUsers(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, result)}

func (f *followHandler) ReturnFollowRequests(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	result, err := f.FollowService.ReturnFollowRequests(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (f *followHandler) IsUserFollowed(ctx echo.Context) error {
	followRequest := &model.FollowRequest{}
	if err := ctx.Bind(followRequest); err != nil {
		return err
	}

	exists, err := f.FollowService.IsUserFollowed(followRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, exists)
}

func (f *followHandler) IsMuted(ctx echo.Context) error {
	mute := &model.Mute{}
	if err := ctx.Bind(mute); err != nil {
		return err
	}

	exists, err := f.FollowService.IsMuted(mute)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, exists)
}

func (f *followHandler) ReturnFollowRequestsForUser(c echo.Context) error {
	objectId := c.Param("objectId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	exists, err := f.FollowService.ReturnFollowRequestsForUser(bearer, objectId)
	fmt.Println(exists)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, exists)
}

func (f *followHandler) ReturnRecommendedUsers(ctx echo.Context) error {
	user := &model.User{Id: ctx.Param("userId")}

	result, err := f.FollowService.ReturnRecommendedUsers(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, result)
}