package service_contracts

import (
	"context"
	"message-service/domain/model"
)

type NotificationService interface {
	Create(ctx context.Context, notification *model.NotificationRequest) error
	GetAllForUser(ctx context.Context, bearer string) ([]*model.Notification, error)
	ViewUsersNotifications(ctx context.Context, bearer string) error
	GetAllNotViewedForLoggedUser(ctx context.Context, bearer string) ([]*model.Notification, error)
	GetAllNotViewedForUser(ctx context.Context, userId string) ([]*model.Notification, error)
}