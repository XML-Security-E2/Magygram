package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"story-service/domain/model"
)

type StoryRepository interface {
	Create(ctx context.Context, story *model.Story) (*mongo.InsertOneResult, error)
}