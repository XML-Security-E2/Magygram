package repository

import (
	"context"
	"message-service/domain/model"
)

type NotificationRepository interface {
	Create(ctx context.Context, notification *model.Notification) error
	GetAllForUser(ctx context.Context, userId string, limit int64) ([]*model.Notification, error)
	ViewNotifications(ctx context.Context, userId string) error
	GetAllNotViewedForUser(ctx context.Context, userId string, limit int64) ([]*model.Notification, error)
	GetAllForUserSortedByTime(ctx context.Context, userId string, limit int64) ([]*model.Notification, error)
}