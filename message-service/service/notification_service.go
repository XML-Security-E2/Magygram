package service

import (
	"context"
	"fmt"
	"message-service/domain/model"
	"message-service/domain/repository"
	"message-service/domain/service-contracts"
	"message-service/service/intercomm"
)

var (
	limit int64 = 100
)

type notificationService struct {
	repository.NotificationRepository
	intercomm.AuthClient
}

func NewNotificationService(r repository.NotificationRepository, ac intercomm.AuthClient) service_contracts.NotificationService {
	return &notificationService{r, ac}
}

func (n notificationService) Create(ctx context.Context, notificationReq *model.NotificationRequest) error {
	notification := model.NewNotification(notificationReq)
	fmt.Println(notification.Id)
	return n.NotificationRepository.Create(ctx, notification)
}

func (n notificationService) GetAllForUser(ctx context.Context, bearer string) ([]*model.Notification, error) {
	userId, err := n.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	return n.NotificationRepository.GetAllForUser(ctx, userId, limit)
}

func (n notificationService) GetAllNotViewedForUser(ctx context.Context, userId string) ([]*model.Notification, error) {
	return n.NotificationRepository.GetAllNotViewedForUser(ctx, userId, limit)
}

func (n notificationService) GetAllNotViewedForLoggedUser(ctx context.Context, bearer string) ([]*model.Notification, error) {
	userId, err := n.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	return n.NotificationRepository.GetAllNotViewedForUser(ctx, userId, limit)}

func (n notificationService) ViewUsersNotifications(ctx context.Context, bearer string) error {
	userId, err := n.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	return n.NotificationRepository.ViewNotifications(ctx, userId)
}