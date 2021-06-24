package service_contracts

import (
	"context"
	"message-service/domain/model"
)

type ConversationService interface {
	SendMessage(ctx context.Context, bearer string, messageRequest *model.MessageSentRequest) (*model.ConversationResponse, error)
	GetAllConversationsForUser(ctx context.Context, bearer string) ([]*model.ConversationResponse, error)
	GetAllMessagesFromUser(ctx context.Context, bearer string, userId string) (*model.MessagesResponse, error)
	ViewUsersMessages(ctx context.Context, bearer string, userId string) error
}
