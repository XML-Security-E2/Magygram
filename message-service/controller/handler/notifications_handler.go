package handler

import (
	"context"
	"fmt"
	"io"
	"message-service/controller/hub"
	"message-service/domain/model"
	"message-service/tracer"
	"net/http"

	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
)

type NotificationHandler interface {
	CreateNotification(c echo.Context) error
	CreateNotifications(c echo.Context) error
	GetAllNotViewedNotificationsForUser(c echo.Context) error
	GetAllNotificationsForUser(c echo.Context) error
	HandleNotifyWs(c echo.Context) error
	ViewNotifications(c echo.Context) error
}

type notificationHandler struct {
	NotificationService service_contracts.NotificationService
	Hub                 *hub.NotifyHub
	tracer              opentracing.Tracer
	closer              io.Closer
}

func NewNotificationHandler(p service_contracts.NotificationService, h *hub.NotifyHub) NotificationHandler {
	tracer, closer := tracer.Init("message-service")
	opentracing.SetGlobalTracer(tracer)
	return &notificationHandler{p, h, tracer, closer}
}

func (n notificationHandler) CreateNotification(c echo.Context) error {
	span := tracer.StartSpanFromRequest("NotificationHandlerCreateNotification", n.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create notification at %s\n", c.Path())),
	)

	notificationRequest := &model.NotificationRequest{}
	if err := c.Bind(notificationRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	notify, err := n.NotificationService.CreatePostInteractionNotification(ctx, notificationRequest)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if notify {
		notifications, err := n.NotificationService.GetAllNotViewedForUser(ctx, notificationRequest.UserId)
		if err == nil {
			n.Hub.Notify <- &hub.Notification{
				Count:    len(notifications),
				Receiver: notificationRequest.UserId,
			}
		}
	}

	return c.JSON(http.StatusCreated, "")
}

func (n notificationHandler) CreateNotifications(c echo.Context) error {
	span := tracer.StartSpanFromRequest("NotificationHandlerCreateNotifications", n.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create notifications at %s\n", c.Path())),
	)

	notificationRequest := &model.NotificationRequest{}
	if err := c.Bind(notificationRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfos, err := n.NotificationService.CreatePostOrStoryNotification(ctx, notificationRequest)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for _, userInfo := range userInfos {
		notifications, err := n.NotificationService.GetAllNotViewedForUser(ctx, userInfo.Id)
		if err == nil {
			n.Hub.Notify <- &hub.Notification{
				Count:    len(notifications),
				Receiver: userInfo.Id,
			}
		}
	}

	return c.JSON(http.StatusCreated, "")
}

func (n notificationHandler) GetAllNotViewedNotificationsForUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("NotificationHandlerGetAllNotViewedNotificationsForUser", n.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get all not viewed notifications for user at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	retVal, err := n.NotificationService.GetAllNotViewedForLoggedUser(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, retVal)
}

func (n notificationHandler) HandleNotifyWs(c echo.Context) error {
	span := tracer.StartSpanFromRequest("NotificationHandlerHandleNotifyWs", n.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling notify at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	notifications, _ := n.NotificationService.GetAllNotViewedForUser(ctx, userId)
	a := 0
	if notifications != nil {
		a = len(notifications)
	}
	hub.ServeNotifyWs(n.Hub, c.Response().Writer, c.Request(), userId, a)

	return nil
}

func (n notificationHandler) GetAllNotificationsForUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("NotificationHandlerGetAllNotificationsForUser", n.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get all notifications for user at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	retVal, err := n.NotificationService.GetAllForUserSortedByTime(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, retVal)
}

func (n notificationHandler) ViewNotifications(c echo.Context) error {
	span := tracer.StartSpanFromRequest("NotificationHandlerViewNotifications", n.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling view notifications at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	err := n.NotificationService.ViewUsersNotifications(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "")
}
