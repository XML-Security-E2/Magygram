package redisdb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"message-service/domain/model"
	"message-service/domain/repository"
)

type notificationRepository struct {
	Db *redis.Client
}

func NewNotificationRepository(Db *redis.Client) repository.NotificationRepository {
	return &notificationRepository{Db}
}

func (n notificationRepository) Create(ctx context.Context, notification *model.Notification) error {
	jsonString, err := json.Marshal(notification)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = n.Db.Set(ctx, notification.Id, jsonString, 0).Err()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (n notificationRepository) GetAllForUser(ctx context.Context, userId string, limit int64) ([]*model.Notification, error) {
	keys, _, err := n.Db.Scan(ctx, 0, fmt.Sprintf("%s/%s/*", model.Prefix, userId), limit).Result()

	var notifications []*model.Notification
	for _, key := range keys {
		val, err := n.Db.Get(ctx, key).Bytes()
		if err != nil {
			return nil, err
		}

		var temp *model.Notification
		json.Unmarshal(val, &temp)
		notifications = append(notifications, temp)
	}

	return notifications, err
}

func (n notificationRepository) GetAllNotViewedForUser(ctx context.Context, userId string, limit int64) ([]*model.Notification, error) {
	keys, _, err := n.Db.Scan(ctx, 0, fmt.Sprintf("%s/%s/*/false", model.Prefix, userId), limit).Result()

	var notifications []*model.Notification
	for _, key := range keys {
		val, err := n.Db.Get(ctx, key).Bytes()
		if err != nil {
			return nil, err
		}

		var temp *model.Notification
		json.Unmarshal(val, &temp)
		notifications = append(notifications, temp)
	}

	return notifications, err
}

func (n notificationRepository) ViewNotifications(ctx context.Context, userId string) error {
	keys, _, err := n.Db.Scan(ctx, 0, fmt.Sprintf("%s/%s/*/false", model.Prefix, userId), 1000).Result()

	for _, key := range keys {
		val, err := n.Db.Get(ctx, key).Bytes()
		if err != nil {
			return err
		}

		var temp *model.Notification
		json.Unmarshal(val, &temp)

		n.Db.Del(ctx, key).Err()
		if err != nil {
			return err
		}

		err = n.Db.Set(ctx, fmt.Sprintf("%s/%s/true", model.Prefix, userId), temp, 0).Err()
		if err != nil {
			return err
		}
	}

	return err
}