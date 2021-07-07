package router

import (
	"message-service/controller/handler"

	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/api/notifications", h.CreateNotification)
	e.POST("/api/notifications/multiple", h.CreateNotifications)
	e.GET("/api/notifications", h.GetAllNotificationsForUser)
	e.PUT("/api/notifications/view", h.ViewNotifications)

	e.Any("/ws/notify/:userId", h.HandleNotifyWs)
	e.Any("/ws/notify/messages/:userId", h.HandleNotifyMessagesWs)
	e.Any("/ws/messages/:userId", h.HandleMessagesWs)

	e.POST("/api/messages", h.SendMessage)
	e.GET("/api/messages/:userId", h.GetAllMessagesFromUser)
	e.GET("/api/conversations", h.GetAllConversationsForUser)
	e.PUT("/api/messages/:conversationId/view", h.ViewMessages)
	e.PUT("/api/messages/:conversationId/:messageId/view", h.ViewMediaMessages)

	e.GET("/api/messages/:userId/requests", h.GetAllMessagesFromUserFromRequest)
	e.GET("/api/messages/requests", h.GetAllMessageRequestsForUser)
	e.PUT("/api/conversations/request/:requestId/accept", h.AcceptConversationRequest)
	e.PUT("/api/conversations/request/:requestId/deny", h.DenyConversationRequest)
	e.DELETE("/api/conversations/request/:requestId", h.DeleteConversationRequest)
	e.GET("/api/messages/metrics", echo.WrapHandler(promhttp.Handler()))
}
