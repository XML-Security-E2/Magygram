package handler

import (
	"context"
	"github.com/labstack/echo"
	"message-service/controller/hub"
	"message-service/domain/model"
	"message-service/domain/service-contracts"
	"net/http"
)

type ConversationHandler interface {
	SendMessage(c echo.Context) error
	GetAllConversationsForUser(c echo.Context) error
	GetAllMessagesFromUser(c echo.Context) error
	ViewMessages(c echo.Context) error
	ViewMediaMessages(c echo.Context) error
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
	contentUrl := c.FormValue("contentUrl")

	media, _ := c.FormFile("media")

	messageRequest := &model.MessageSentRequest{
		MessageTo:   messageTo,
		MessageType: model.MessageType(messageType),
		Media:       media,
		Text:        text,
		ContentUrl:  contentUrl,
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