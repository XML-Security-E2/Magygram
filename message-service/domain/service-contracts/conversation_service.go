package service_contracts

import (
	"context"
	"message-service/domain/model"
)

type ConversationService interface {
	SendMessage(ctx context.Context, bearer string, messageRequest *model.MessageSentRequest) (*model.MessageSendResponse, error)
	GetAllConversationsForUser(ctx context.Context, bearer string) ([]*model.ConversationResponse, error)
	GetAllMessagesFromUser(ctx context.Context, bearer string, userId string) (*model.MessagesResponse, error)
	ViewUsersMessages(ctx context.Context, bearer string, userId string) error
	ViewUserMediaMessages(ctx context.Context, bearer string, conversationId string, messageId string) error
	AcceptConversationRequest(ctx context.Context, bearer string, requestId string) error
	DenyConversationRequest(ctx context.Context, bearer string, requestId string) error
	DeleteConversationRequest(ctx context.Context, bearer string, requestId string) error

	GetAllMessageRequestsForUser(ctx context.Context, bearer string) ([]*model.ConversationResponse, error)
	GetAllMessagesFromUserFromRequest(ctx context.Context, bearer string, userId string) (*model.MessagesResponse, error)
}
