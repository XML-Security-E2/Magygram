package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"message-service/controller/hub"
	"message-service/domain/model"
	"message-service/domain/service-contracts"
	"message-service/service/intercomm"
	"net/http"
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
	Hub *hub.NotifyHub
	intercomm.AuthClient
}

func NewNotificationHandler(p service_contracts.NotificationService, h *hub.NotifyHub, ac intercomm.AuthClient) NotificationHandler {
	return &notificationHandler{p, h, ac}
}

func (n notificationHandler) CreateNotification(c echo.Context) error {
	notificationRequest := &model.NotificationRequest{}
	if err := c.Bind(notificationRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	notify, err := n.NotificationService.CreatePostInteractionNotification(ctx, notificationRequest)
	if err != nil {
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
	notificationRequest := &model.NotificationRequest{}
	if err := c.Bind(notificationRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userInfos, err := n.NotificationService.CreatePostOrStoryNotification(ctx, notificationRequest)
	if err != nil {
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

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	retVal, err := n.NotificationService.GetAllNotViewedForLoggedUser(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, retVal)
}

func (n notificationHandler) HandleNotifyWs(c echo.Context) error {
	fmt.Println("USAO")
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	notifications, _ := n.NotificationService.GetAllNotViewedForUser(ctx, userId)
	a := 0
	if notifications != nil {
		a = len(notifications)
	}
	hub.ServeNotifyWs(n.Hub, c.Response().Writer, c.Request(), userId, a)

	return nil
}

func (n notificationHandler) GetAllNotificationsForUser(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	retVal, err := n.NotificationService.GetAllForUserSortedByTime(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, retVal)
}

func (n notificationHandler) ViewNotifications(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	err := n.NotificationService.ViewUsersNotifications(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "")
}