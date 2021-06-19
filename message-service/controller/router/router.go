package router

import (
	"github.com/labstack/echo"
	"message-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/notifications", h.CreateNotification)
	e.POST("/api/notifications/multiple", h.CreateNotifications)
	e.GET("/api/notifications", h.GetAllNotificationsForUser)
	e.PUT("/api/notifications/view", h.ViewNotifications)

	e.Any("/ws/notify/:userId", h.HandleNotifyWs)
}