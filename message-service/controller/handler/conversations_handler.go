package handler

import (
	"context"
	"github.com/labstack/echo"
	"message-service/controller/hub"
	"message-service/domain/model"
	"message-service/domain/service-contracts"
	"message-service/domain/service-contracts/exceptions/denied"
	"net/http"
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

}

type conversationHandler struct {
	ConversationService service_contracts.ConversationService
	Hub *hub.MessageHub
}

func NewConversationHandler(p service_contracts.ConversationService, h *hub.MessageHub) ConversationHandler {
	return &conversationHandler{p, h}
}

func (ch conversationHandler) ViewMediaMessages(c echo.Context) error {
	conversationId := c.Param("conversationId")
	messageId := c.Param("messageId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := ch.ConversationService.ViewUserMediaMessages(ctx, bearer, conversationId, messageId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")}

func (ch conversationHandler) ViewMessages(c echo.Context) error {
	conversationId := c.Param("conversationId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := ch.ConversationService.ViewUsersMessages(ctx, bearer, conversationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (ch conversationHandler) SendMessage(c echo.Context) error {
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

	bearer := c.Request().Header.Get("Authorization")

	message, err := ch.ConversationService.SendMessage(ctx, bearer, messageRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, message)
}

func (ch conversationHandler) GetAllConversationsForUser(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	conversations, err := ch.ConversationService.GetAllConversationsForUser(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, conversations)
}

func (ch conversationHandler) GetAllMessagesFromUser(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	messages, err := ch.ConversationService.GetAllMessagesFromUser(ctx, bearer, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, messages)
}

func (ch conversationHandler) GetAllMessagesFromUserFromRequest(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	messages, err := ch.ConversationService.GetAllMessagesFromUserFromRequest(ctx, bearer, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, messages)
}

func (ch conversationHandler) GetAllMessageRequestsForUser(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	messages, err := ch.ConversationService.GetAllMessageRequestsForUser(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, messages)
}

func (ch conversationHandler) AcceptConversationRequest(c echo.Context) error {
	requestId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := ch.ConversationService.AcceptConversationRequest(ctx, bearer, requestId)
	if err != nil{
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
	requestId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := ch.ConversationService.DenyConversationRequest(ctx, bearer, requestId)
	if err != nil{
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
	requestId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := ch.ConversationService.DeleteConversationRequest(ctx, bearer, requestId)
	if err != nil{
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *denied.MessageRequestDeniedError:
			return echo.NewHTTPError(http.StatusForbidden, t.Error())
		}
	}

	return c.JSON(http.StatusOK, "")
}