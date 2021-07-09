package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"relationship-service/domain/model"
	"relationship-service/logger"
	"relationship-service/service"
	"relationship-service/tracer"

	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
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
	tracer        opentracing.Tracer
	closer        io.Closer
}

const name = "relationship-service"

func NewFollowHandler(f service.FollowService) FollowHandler {
	tracer, closer := tracer.Init(name)
	opentracing.SetGlobalTracer(tracer)
	return &followHandler{f, tracer, closer}
}

func (f followHandler) LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.LoggingEntry = logger.Logger.WithFields(logrus.Fields{"request_ip": c.RealIP()})
		return next(c)
	}
}

func (f followHandler) Mute(ctx echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerMute", f.tracer, ctx.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling mute user at %s\n", ctx.Path())),
	)

	c := ctx.Request().Context()
	if c == nil {
		c = context.Background()
	}
	c = tracer.ContextWithSpan(ctx.Request().Context(), span)

	mute := &model.Mute{}
	if err := ctx.Bind(mute); err != nil {
		tracer.LogError(span, err)
		return err
	}

	if err := f.FollowService.Mute(c, mute); err != nil {
		tracer.LogError(span, err)
		return err
	}

	return ctx.JSON(http.StatusOK, "User successfully muted")
}

func (f followHandler) Unmute(ctx echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerUnmute", f.tracer, ctx.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling unmute user at %s\n", ctx.Path())),
	)

	c := ctx.Request().Context()
	if c == nil {
		c = context.Background()
	}
	c = tracer.ContextWithSpan(ctx.Request().Context(), span)

	mute := &model.Mute{}
	if err := ctx.Bind(mute); err != nil {
		tracer.LogError(span, err)
		return err
	}

	if err := f.FollowService.Unmute(c, mute); err != nil {
		tracer.LogError(span, err)
		return err
	}

	return ctx.JSON(http.StatusOK, "User successfully muted")
}

func (f *followHandler) FollowRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerFollowRequest", f.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling follow requst at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(c.Request().Context(), span)

	followRequest := &model.FollowRequest{}
	if err := c.Bind(followRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	sentRequest, err := f.FollowService.FollowRequest(ctx, followRequest)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, sentRequest)
}

func (f *followHandler) Unfollow(c echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerUnfollow", f.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling unfollow user at %s\n", c.Path())),
	)

	followRequest := &model.FollowRequest{}
	if err := c.Bind(followRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}
	ctx := tracer.ContextWithSpan(c.Request().Context(), span)
	err := f.FollowService.Unfollow(ctx, followRequest)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (f *followHandler) AcceptFollowRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerAcceptFollowRequest", f.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling accept follow requst at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(c.Request().Context(), span)

	userId := c.Param("userId")
	fmt.Println(userId)
	bearer := c.Request().Header.Get("Authorization")
	err := f.FollowService.AcceptFollowRequest(ctx, bearer, userId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, true)
}

func (f *followHandler) CreateUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerCreateUser", f.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create user at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(c.Request().Context(), span)

	user := &model.User{}
	if err := c.Bind(user); err != nil {
		tracer.LogError(span, err)
		return err
	}
	if err := f.FollowService.CreateUser(ctx, user); err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, true)
}

func (f *followHandler) ReturnFollowedUsers(ctx echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerReturnFollowedUsers", f.tracer, ctx.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling return followed users at %s\n", ctx.Path())),
	)
	c := tracer.ContextWithSpan(ctx.Request().Context(), span)

	user := &model.User{Id: ctx.Param("userId")}

	result, err := f.FollowService.ReturnFollowedUsers(c, user)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (f *followHandler) ReturnUnmutedFollowedUsers(ctx echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerReturnUnmutedFollowedUsers", f.tracer, ctx.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling return unmuted followed users at %s\n", ctx.Path())),
	)

	c := ctx.Request().Context()
	if c == nil {
		c = context.Background()
	}
	c = tracer.ContextWithSpan(ctx.Request().Context(), span)

	user := &model.User{Id: ctx.Param("userId")}

	result, err := f.FollowService.ReturnUnmutedFollowedUsers(c, user)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (f *followHandler) ReturnFollowingUsers(ctx echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerReturnFollowingUsers", f.tracer, ctx.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling return following users at %s\n", ctx.Path())),
	)
	c := tracer.ContextWithSpan(ctx.Request().Context(), span)

	user := &model.User{Id: ctx.Param("userId")}

	result, err := f.FollowService.ReturnFollowingUsers(c, user)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, result)
}

func (f *followHandler) ReturnFollowRequests(c echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerReturnFollowRequests", f.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling return follow requsts at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(c.Request().Context(), span)
	bearer := c.Request().Header.Get("Authorization")

	result, err := f.FollowService.ReturnFollowRequests(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (f *followHandler) IsUserFollowed(ctx echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerIsUserFollowed", f.tracer, ctx.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling is user followed at %s\n", ctx.Path())),
	)

	c := ctx.Request().Context()
	if c == nil {
		c = context.Background()
	}
	c = tracer.ContextWithSpan(ctx.Request().Context(), span)

	followRequest := &model.FollowRequest{}
	if err := ctx.Bind(followRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	exists, err := f.FollowService.IsUserFollowed(c, followRequest)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, exists)
}

func (f *followHandler) IsMuted(ctx echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerIsMuted", f.tracer, ctx.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling is user muted at %s\n", ctx.Path())),
	)

	c := ctx.Request().Context()
	if c == nil {
		c = context.Background()
	}
	c = tracer.ContextWithSpan(ctx.Request().Context(), span)

	mute := &model.Mute{}
	if err := ctx.Bind(mute); err != nil {
		tracer.LogError(span, err)
		return err
	}

	exists, err := f.FollowService.IsMuted(c, mute)
	if err != nil {
		tracer.LogError(span, err)
		return err
	}

	return ctx.JSON(http.StatusOK, exists)
}

func (f *followHandler) ReturnFollowRequestsForUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerReturnFollowRequestsForUser", f.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling return follow requests for user at %s\n", c.Path())),
	)

	objectId := c.Param("objectId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(c.Request().Context(), span)
	bearer := c.Request().Header.Get("Authorization")

	exists, err := f.FollowService.ReturnFollowRequestsForUser(ctx, bearer, objectId)
	fmt.Println(exists)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, exists)
}

func (f *followHandler) ReturnRecommendedUsers(ctx echo.Context) error {
	span := tracer.StartSpanFromRequest("FollowHandlerReturnRecommendedUsers", f.tracer, ctx.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling return recommended users at %s\n", ctx.Path())),
	)

	c := ctx.Request().Context()
	if c == nil {
		c = context.Background()
	}
	c = tracer.ContextWithSpan(ctx.Request().Context(), span)

	user := &model.User{Id: ctx.Param("userId")}

	result, err := f.FollowService.ReturnRecommendedUsers(c, user)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, result)
}
