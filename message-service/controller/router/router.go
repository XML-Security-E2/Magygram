package router

import (
	"github.com/labstack/echo"
	"message-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/notifications", h.CreateNotification)
	e.GET("/api/notifications", h.GetAllNotViewedNotificationsForUser)

	e.Any("/ws/notify/:userId", h.HandleNotifyWs)
}