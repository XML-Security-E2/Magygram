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
	ViewUsersMessages(ctx context.Context, userId string,  conversationId string) error

}
