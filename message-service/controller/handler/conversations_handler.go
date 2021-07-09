package handler

import (
	"context"
	"fmt"
	"io"
	"message-service/controller/hub"
	"message-service/domain/model"
	"message-service/domain/service-contracts/exceptions/denied"
	"message-service/tracer"
	"net/http"

	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
)

type ConversationHandler interface {
	SendMessage(c echo.Context) error
	GetAllConversationsForUser(c echo.Context) error
	GetAllMessagesFromUser(c echo.Context) error
	ViewMessages(c echo.Context) error
	ViewMediaMessages(c echo.Context) error
	GetAllMessagesFromUserFromRequest(c echo.Context) error
	GetAllMessageRequestsForUser(c echo.Context) error
	AcceptConversationRequest(c echo.Context) error
	DenyConversationRequest(c echo.Context) error
	DeleteConversationRequest(c echo.Context) error
	HandleNotifyMessagesWs(c echo.Context) error
	HandleMessagesWs(c echo.Context) error
}

type conversationHandler struct {
	ConversationService service_contracts.ConversationService
	Hub                 *hub.MessageHub
	NotifyHub           *hub.MessageNotificationsHub
	tracer              opentracing.Tracer
	closer              io.Closer
}

func NewConversationHandler(p service_contracts.ConversationService, h *hub.MessageHub, nh *hub.MessageNotificationsHub) ConversationHandler {
	tracer, closer := tracer.Init("message-service")
	opentracing.SetGlobalTracer(tracer)
	return &conversationHandler{p, h, nh, tracer, closer}
}

func (ch conversationHandler) ViewMediaMessages(c echo.Context) error {
	span := tracer.StartSpanFromRequest("ConversationHandlerViewMediaMessages", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling view media messages at %s\n", c.Path())),
	)

	conversationId := c.Param("conversationId")
	messageId := c.Param("messageId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	err := ch.ConversationService.ViewUserMediaMessages(ctx, bearer, conversationId, messageId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (ch conversationHandler) ViewMessages(c echo.Context) error {
	span := tracer.StartSpanFromRequest("ConversationHandlerViewMessages", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling view messages at %s\n", c.Path())),
	)

	conversationId := c.Param("conversationId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	userId, err := ch.ConversationService.ViewUsersMessages(ctx, bearer, conversationId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	notifications, _ := ch.ConversationService.GetAllNotViewedConversationsForUser(ctx, userId)

	ch.NotifyHub.Notify <- &hub.Notification{
		Count:    len(notifications),
		Receiver: userId,
	}

	return c.JSON(http.StatusOK, "")
}

func (ch conversationHandler) SendMessage(c echo.Context) error {
	span := tracer.StartSpanFromRequest("ConversationHandlerSendMessage", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling send message at %s\n", c.Path())),
	)

	messageTo := c.FormValue("messageTo")
	messageType := c.FormValue("messageType")
	text := c.FormValue("text")
	contentId := c.FormValue("contentId")

	media, _ := c.FormFile("media")

	messageRequest := &model.MessageSentRequest{
		MessageTo:   messageTo,
		MessageType: model.MessageType(messageType),
		Media:       media,
		Text:        text,
		ContentId:   contentId,
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	message, err := ch.ConversationService.SendMessage(ctx, bearer, messageRequest)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	notifications, _ := ch.ConversationService.GetAllNotViewedConversationsForUser(ctx, messageTo)

	ch.NotifyHub.Notify <- &hub.Notification{
		Count:    len(notifications),
		Receiver: messageTo,
	}

	ch.Hub.Broadcast <- message

	return c.JSON(http.StatusCreated, message)
}

func (ch conversationHandler) GetAllConversationsForUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("ConversationHandlerGetAllConversationsForUser", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get all conversations for user at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	conversations, err := ch.ConversationService.GetAllConversationsForUser(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, conversations)
}

func (ch conversationHandler) GetAllMessagesFromUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("ConversationHandlerGetAllMessagesFromUser", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get all messages for user at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	messages, err := ch.ConversationService.GetAllMessagesFromUser(ctx, bearer, userId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, messages)
}

func (ch conversationHandler) GetAllMessagesFromUserFromRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("ConversationHandlerGetAllMessagesFromUserFromRequest", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get all messages from user from request at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	messages, err := ch.ConversationService.GetAllMessagesFromUserFromRequest(ctx, bearer, userId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, messages)
}

func (ch conversationHandler) GetAllMessageRequestsForUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("ConversationHandlerGetAllMessageRequestsForUser", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get all message requests for user at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	messages, err := ch.ConversationService.GetAllMessageRequestsForUser(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, messages)
}

func (ch conversationHandler) AcceptConversationRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("ConversationHandlerAcceptConversationRequest", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling accept conversation request at %s\n", c.Path())),
	)

	requestId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	err := ch.ConversationService.AcceptConversationRequest(ctx, bearer, requestId)
	if err != nil {
		tracer.LogError(span, err)
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *denied.MessageRequestDeniedError:
			return echo.NewHTTPError(http.StatusForbidden, t.Error())
		}
	}

	return c.JSON(http.StatusOK, "")
}

func (ch conversationHandler) DenyConversationRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("ConversationHandlerDenyConversationRequest", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling deny conversation request at %s\n", c.Path())),
	)

	requestId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	err := ch.ConversationService.DenyConversationRequest(ctx, bearer, requestId)
	if err != nil {
		tracer.LogError(span, err)
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *denied.MessageRequestDeniedError:
			return echo.NewHTTPError(http.StatusForbidden, t.Error())
		}
	}

	return c.JSON(http.StatusOK, "")
}

func (ch conversationHandler) DeleteConversationRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("ConversationHandlerDeleteConversationRequest", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling delete conversation request at %s\n", c.Path())),
	)

	requestId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	err := ch.ConversationService.DeleteConversationRequest(ctx, bearer, requestId)
	if err != nil {
		tracer.LogError(span, err)
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *denied.MessageRequestDeniedError:
			return echo.NewHTTPError(http.StatusForbidden, t.Error())
		}
	}

	return c.JSON(http.StatusOK, "")
}

func (ch conversationHandler) HandleNotifyMessagesWs(c echo.Context) error {
	span := tracer.StartSpanFromRequest("ConversationHandlerHandleNotifyMessagesWs", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling notify messages at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	notifications, _ := ch.ConversationService.GetAllNotViewedConversationsForUser(ctx, userId)
	a := 0
	if notifications != nil {
		a = len(notifications)
	}
	hub.ServeMessageNotificationWs(ch.NotifyHub, c.Response().Writer, c.Request(), userId, a)

	return nil
}

func (ch conversationHandler) HandleMessagesWs(c echo.Context) error {
	span := tracer.StartSpanFromRequest("ConversationHandlerHandleMessagesWs", ch.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling messages at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	hub.ServeMessageWs(ch.Hub, c.Response().Writer, c.Request(), userId)

	return nil
}
