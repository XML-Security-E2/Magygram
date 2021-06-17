package service_contracts

import (
	"context"
	"message-service/domain/model"
)

type NotificationService interface {
	CreatePostInteractionNotification(ctx context.Context, notificationReq *model.NotificationRequest) (bool, error)
	CreatePostOrStoryNotification(ctx context.Context, notificationReq *model.NotificationRequest) ([]*model.UserInfo, error)
	GetAllForUser(ctx context.Context, bearer string) ([]*model.Notification, error)
	GetAllForUserSortedByTime(ctx context.Context, bearer string) ([]*model.Notification, error)
	ViewUsersNotifications(ctx context.Context, bearer string) error
	GetAllNotViewedForLoggedUser(ctx context.Context, bearer string) ([]*model.Notification, error)
	GetAllNotViewedForUser(ctx context.Context, userId string) ([]*model.Notification, error)
}