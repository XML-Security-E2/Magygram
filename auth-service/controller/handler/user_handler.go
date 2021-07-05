package handler

import (
	"auth-service/domain/model"
	"auth-service/domain/service-contracts"
	"auth-service/logger"
	"auth-service/tracer"
	"bytes"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)


type UserHandler interface {
	RegisterUser(c echo.Context) error
	ActivateUser(c echo.Context) error
	ResetPassword(c echo.Context) error
	UserLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	RegisterAgent(c echo.Context) error
}

var (
	ErrHttpGenericMessage = echo.NewHTTPError(http.StatusInternalServerError, "something went wrong, please try again later")
	ErrWrongCredentials = echo.NewHTTPError(http.StatusUnauthorized, "username or password is invalid")
	ErrUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
)

type userHandler struct {
	UserService service_contracts.UserService
	tracer      opentracing.Tracer
	closer      io.Closer
}

func NewUserHandler(u service_contracts.UserService) UserHandler {
	tracer, closer := tracer.Init("auth-service")
	opentracing.SetGlobalTracer(tracer)
	return &userHandler{u, tracer, closer}
}

func (u userHandler) UserLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.LoggingEntry = logger.Logger.WithFields(logrus.Fields{"request_ip": c.RealIP()})
		return next(c)
	}
}

func (u userHandler) RegisterUser(c echo.Context) error {
	fmt.Println(c.Request().Header.Get("X-Forwarded-For"))// capitalisation )
	fmt.Println(c.Request().Header.Get("proba-proba"))// capitalisation )

	span := tracer.StartSpanFromRequest("UserHandlerRegisterUser", u.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling register user at %s\n", c.Path())),
	)

	userRequest := &model.UserRequest{}
	if err := c.Bind(userRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	_, bufer, err := u.UserService.RegisterUser(ctx, userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	yter := bytes.NewReader(bufer)

	return c.Stream(http.StatusCreated,"image/png",yter)
}

func (u userHandler) ActivateUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerActivateUser", u.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling activate user at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	activated, err := u.UserService.ActivateUser(ctx, userId)
	if err != nil || activated == false{
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not be activated.")
	}

	return c.JSON(http.StatusOK, userId)
}

func (u userHandler) ResetPassword(c echo.Context) error {
	changeNewPasswordRequest := &model.PasswordChangeRequest{}
	if err := c.Bind(changeNewPasswordRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()

	successful, err := u.UserService.ResetPassword(ctx, changeNewPasswordRequest)

	if err != nil || !successful {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Password has been changed")
}


func (u userHandler) RegisterAgent(c echo.Context) error {
	userRequest := &model.UserRequest{}
	if err := c.Bind(userRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	_, bufer, err := u.UserService.RegisterAgent(ctx, userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	yter := bytes.NewReader(bufer)

	return c.Stream(http.StatusCreated,"image/png",yter)
}
