package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/domain/model"
	"user-service/domain/repository"
)

type notificationRulesRepository struct {
	Col *mongo.Collection
}

func NewNotificationRulesRepository(Col *mongo.Collection) repository.NotificationRulesRepository {
	return &notificationRulesRepository{Col}
}


func (n notificationRulesRepository) Create(ctx context.Context, notificationRule *model.PostStoryNotifications) (*mongo.InsertOneResult, error) {
	return n.Col.InsertOne(ctx, notificationRule)
}


func (n notificationRulesRepository) GetRuleForUser(ctx context.Context, userId string, userFromId string) (*model.PostStoryNotifications, error) {
	var notiy = model.PostStoryNotifications{}
	err := n.Col.FindOne(ctx, bson.M{"userId": userId, "notificationsFromId": userFromId}).Decode(&notiy)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &notiy, nil
}

func (n notificationRulesRepository) GetNotifiersForStory(ctx context.Context, userId string) ([]string, error) {
	var notifierIds []string
	cursor, err := n.Col.Find(ctx, bson.M{"notificationsFromId": userId, "storyNotifications" : true})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []string{}, nil
		}
		return []string{}, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var notifyRule model.PostStoryNotifications
		if err = cursor.Decode(&notifyRule); err == nil {
			notifierIds = append(notifierIds, notifyRule.UserId)
		}
	}

	return notifierIds, nil
}

func (n notificationRulesRepository) GetNotifiersForPost(ctx context.Context, userId string) ([]string, error) {
	var notifierIds []string
	cursor, err := n.Col.Find(ctx, bson.M{"notificationsFromId": userId, "postNotifications" : true})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []string{}, nil
		}
		return []string{}, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var notifyRule model.PostStoryNotifications
		if err = cursor.Decode(&notifyRule); err == nil {
			notifierIds = append(notifierIds, notifyRule.UserId)
		}
	}

	return notifierIds, nil}

func (n notificationRulesRepository) Update(ctx context.Context, notificationRule *model.PostStoryNotifications) (*mongo.UpdateResult, error) {
	return n.Col.UpdateOne(ctx, bson.M{"_id":  notificationRule.Id},bson.D{{"$set", bson.D{{"postNotifications" , notificationRule.PostNotifications},
		{"storyNotifications" , notificationRule.StoryNotifications},
	}}})
}