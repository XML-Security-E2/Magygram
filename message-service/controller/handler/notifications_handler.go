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
	GetAllNotViewedNotificationsForUser(c echo.Context) error
	HandleNotifyWs(c echo.Context) error
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

	err := n.NotificationService.Create(ctx, notificationRequest)

	notifications, err := n.NotificationService.GetAllNotViewedForUser(ctx, notificationRequest.UserId)

	fmt.Println(len(notifications))
	n.Hub.Notify <- &hub.Notification{
		Count:    len(notifications),
		Receiver: notificationRequest.UserId,
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "")
}

func (n notificationHandler) GetAllNotViewedNotificationsForUser(c echo.Context) error {
	panic("implement me")
}

func (n notificationHandler) HandleNotifyWs(c echo.Context) error {
	fmt.Println("USAO")
	userId := c.Param("userId")

	hub.ServeNotifyWs(n.Hub, c.Response().Writer, c.Request(), userId)
	return nil
}