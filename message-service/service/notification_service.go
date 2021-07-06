package service

import (
	"context"
	"fmt"
	"math"
	"message-service/domain/model"
	"message-service/domain/repository"
	"message-service/domain/service-contracts"
	"message-service/service/intercomm"
	"message-service/tracer"
)

var (
	limit int64 = math.MaxInt64
)

type notificationService struct {
	repository.NotificationRepository
	intercomm.AuthClient
	intercomm.UserClient
}

func NewNotificationService(r repository.NotificationRepository, ac intercomm.AuthClient, uc intercomm.UserClient) service_contracts.NotificationService {
	return &notificationService{r, ac, uc}
}

func (n notificationService) CreatePostInteractionNotification(ctx context.Context, notificationReq *model.NotificationRequest) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "NotificationServiceCreatePostInteractionNotification")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	if notificationReq.Type == model.Liked {
		notify, err := n.UserClient.CheckIfPostInteractionNotificationEnabled(ctx, notificationReq.UserId, notificationReq.UserFromId, "like")
		if err != nil {
			return false, err
		}

		if notify {
			notification := model.NewNotification(notificationReq)
			return true, n.NotificationRepository.Create(ctx, notification)
		}
	} else if notificationReq.Type == model.Disliked {
		notify, err := n.UserClient.CheckIfPostInteractionNotificationEnabled(ctx, notificationReq.UserId, notificationReq.UserFromId, "dislike")
		if err != nil {
			return false, err
		}

		if notify {
			notification := model.NewNotification(notificationReq)
			return true, n.NotificationRepository.Create(ctx, notification)
		}
	} else if notificationReq.Type == model.Commented {
		notify, err := n.UserClient.CheckIfPostInteractionNotificationEnabled(ctx, notificationReq.UserId, notificationReq.UserFromId, "comment")
		if err != nil {
			return false, err
		}

		if notify {
			notification := model.NewNotification(notificationReq)
			return true, n.NotificationRepository.Create(ctx, notification)
		}
	} else if notificationReq.Type == model.Followed {
		notify, err := n.UserClient.CheckIfPostInteractionNotificationEnabled(ctx, notificationReq.UserId, notificationReq.UserFromId, "follow")
		if err != nil {
			return false, err
		}

		if notify {
			notification := model.NewNotification(notificationReq)
			return true, n.NotificationRepository.Create(ctx, notification)
		}
	} else if notificationReq.Type == model.FollowRequest {
		notify, err := n.UserClient.CheckIfPostInteractionNotificationEnabled(ctx, notificationReq.UserId, notificationReq.UserFromId, "follow-request")
		if err != nil {
			return false, err
		}

		if notify {
			notification := model.NewNotification(notificationReq)
			return true, n.NotificationRepository.Create(ctx, notification)
		}
	} else if notificationReq.Type == model.AcceptedFollowRequest {
		notify, err := n.UserClient.CheckIfPostInteractionNotificationEnabled(ctx, notificationReq.UserId, notificationReq.UserFromId, "accepted-follow-request")
		if err != nil {
			return false, err
		}

		if notify {
			notification := model.NewNotification(notificationReq)
			return true, n.NotificationRepository.Create(ctx, notification)
		}
	}

	return false, nil
}


func (n notificationService) CreatePostOrStoryNotification(ctx context.Context, notificationReq *model.NotificationRequest) ([]*model.UserInfo, error) {
	span := tracer.StartSpanFromContext(ctx, "NotificationServiceCreatePostOrStoryNotification")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	if notificationReq.Type == model.PublishedPost {
		userInfos, err := n.UserClient.GetUsersForPostNotification(ctx, notificationReq.UserId)
		fmt.Println(len(userInfos))
		fmt.Println(notificationReq.UserId)

		if err != nil {
			return []*model.UserInfo{}, err
		}
		for _, userInfo := range userInfos {
			notification := model.NewNotification(&model.NotificationRequest{
				Username:  notificationReq.Username,
				UserId:    userInfo.Id,
				NotifyUrl: "TODO",
				ImageUrl:  notificationReq.ImageUrl,
				Type:      model.PublishedPost,
			})

			_ = n.NotificationRepository.Create(ctx, notification)
		}
		return userInfos, nil
	} else if notificationReq.Type == model.PublishedStory {
		userInfos, err := n.UserClient.GetUsersForStoryNotification(ctx, notificationReq.UserId)
		if err != nil {
			return []*model.UserInfo{}, err
		}
		for _, userInfo := range userInfos {
			notification := model.NewNotification(&model.NotificationRequest{
				Username:  notificationReq.Username,
				UserId:    userInfo.Id,
				NotifyUrl: "TODO",
				ImageUrl:  notificationReq.ImageUrl,
				Type:      model.PublishedStory,
			})

			_ = n.NotificationRepository.Create(ctx, notification)
		}
		return userInfos, nil
	}
	return []*model.UserInfo{}, nil
}


func (n notificationService) GetAllForUser(ctx context.Context, bearer string) ([]*model.Notification, error) {
	span := tracer.StartSpanFromContext(ctx, "NotificationServiceGetAllForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := n.AuthClient.GetLoggedUserId(ctx,
		bearer)
	if err != nil {
		return nil, err
	}

	return n.NotificationRepository.GetAllForUser(ctx, userId, limit)
}

func (n notificationService) GetAllForUserSortedByTime(ctx context.Context, bearer string) ([]*model.Notification, error) {
	span := tracer.StartSpanFromContext(ctx, "NotificationServiceGetAllForUserSortedByTime")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := n.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return nil, err
	}

	return n.NotificationRepository.GetAllForUserSortedByTime(ctx, userId, limit)
}

func (n notificationService) GetAllNotViewedForUser(ctx context.Context, userId string) ([]*model.Notification, error) {
	span := tracer.StartSpanFromContext(ctx, "NotificationServiceGetAllNotViewedForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	return n.NotificationRepository.GetAllNotViewedForUser(ctx, userId, limit)
}

func (n notificationService) GetAllNotViewedForLoggedUser(ctx context.Context, bearer string) ([]*model.Notification, error) {
	span := tracer.StartSpanFromContext(ctx, "NotificationServiceGetAllNotViewedForLoggedUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := n.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return nil, err
	}

	return n.NotificationRepository.GetAllNotViewedForUser(ctx, userId, limit)}

func (n notificationService) ViewUsersNotifications(ctx context.Context, bearer string) error {
	span := tracer.StartSpanFromContext(ctx, "NotificationServiceViewUsersNotifications")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := n.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return err
	}

	return n.NotificationRepository.ViewNotifications(ctx, userId)
}