package redisdb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/beevik/guid"
	"github.com/go-redis/redis/v8"
	"math"
	"message-service/domain/model"
	"message-service/domain/repository"
	"sort"
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
	err = n.Db.Set(ctx, fmt.Sprintf("%s/%s/%s/false", model.Prefix, notification.UserId, notification.Id), jsonString, 0).Err()
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

func (n notificationRepository) GetAllForUserSortedByTime(ctx context.Context, userId string, limit int64) ([]*model.Notification, error) {
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

	sort.Slice(notifications, func(i, j int) bool {
		return notifications[i].Timestamp.After(notifications[j].Timestamp)
	})

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
	keys, _, err := n.Db.Scan(ctx, 0, fmt.Sprintf("%s/%s/*/false", model.Prefix, userId), math.MaxInt64).Result()

	for _, key := range keys {
		fmt.Println(key)
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

		newId := fmt.Sprintf("%s/%s/%s/true", model.Prefix, userId, guid.New().String())
		temp.Id = newId

		jsonString, _ := json.Marshal(temp)
		err = n.Db.Set(ctx, newId, jsonString, 0).Err()
		if err != nil {
			return err
		}
	}

	return err
}