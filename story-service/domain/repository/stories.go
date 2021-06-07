package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"story-service/domain/model"
)

type StoryRepository interface {
	Create(ctx context.Context, story *model.Story) (*mongo.InsertOneResult, error)
	GetAll(ctx context.Context) ([]*model.Story, error)
	GetStoriesForUser(ctx context.Context, userId string) ([]*model.Story, error)
	GetByID(ctx context.Context, storyId string) (*model.Story, error)
	Update(ctx context.Context, story *model.Story) (*mongo.UpdateResult, error)
	GetActiveStoriesForUser(ctx context.Context, userId string) ([]*model.Story, error)
}