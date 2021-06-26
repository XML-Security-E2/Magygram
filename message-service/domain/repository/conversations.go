package repository

import (
	"context"
	"message-service/domain/model"
)

type ConversationRepository interface {
	CreateConversation(ctx context.Context, conversation *model.Conversation) error
	GetAllForUser(ctx context.Context, userId string, limit int64) ([]*model.Conversation, error)
	GetConversationForUser(ctx context.Context, loggedId string, userId string, limit int64) (*model.Conversation, error)
	GetAllMessagesFromUser(ctx context.Context, loggedId string, userId string, limit int64) ([]model.Message, error)
	Update(ctx context.Context, conversation *model.Conversation) error
	CreateMessageRequest(ctx context.Context, request *model.MessageRequest) error

	CreateConversationRequest(ctx context.Context, request *model.ConversationRequest) error
	GetConversationRequestFromUser(ctx context.Context, loggedId string, userId string, limit int64) (*model.ConversationRequest, error)
	GetAllConversationRequestsForUser(ctx context.Context, loggedId string, limit int64) ([]*model.ConversationRequest, error)
	UpdateConversationRequest(ctx context.Context, conversation *model.ConversationRequest) error
	DeleteConversationRequest(ctx context.Context, fromId string, toId string, requestId string) error
	GetConversationRequestById(ctx context.Context, requestId string) (*model.ConversationRequest, error)

	ViewUsersMessages(ctx context.Context, userId string,  conversationId string) error
	ViewUserMediaMessage(ctx context.Context, userId string, conversationId string, messageId string) error
}
