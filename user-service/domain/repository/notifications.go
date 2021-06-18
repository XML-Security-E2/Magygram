package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/domain/model"
)

type NotificationRulesRepository interface {
	Create(ctx context.Context, notificationRule *model.PostStoryNotifications) (*mongo.InsertOneResult, error)
	GetNotifiersForStory(ctx context.Context, userId string) ([]string, error)
	GetNotifiersForPost(ctx context.Context, userId string) ([]string, error)
	GetRuleForUser(ctx context.Context, userId string, userFromId string) (*model.PostStoryNotifications, error)
	Update(ctx context.Context, notificationRule *model.PostStoryNotifications) (*mongo.UpdateResult, error)
}
