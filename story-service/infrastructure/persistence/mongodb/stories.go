package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"story-service/domain/model"
	"story-service/domain/repository"
)

type storyRepository struct {
	Col *mongo.Collection
}


func NewStoryRepository(Col *mongo.Collection) repository.StoryRepository {
	return &storyRepository{Col}
}

func (s storyRepository) Create(ctx context.Context, story *model.Story) (*mongo.InsertOneResult, error) {
	return s.Col.InsertOne(ctx, story)
}