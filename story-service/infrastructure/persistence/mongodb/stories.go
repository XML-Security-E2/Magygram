package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
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

func (s storyRepository) GetAll(ctx context.Context) ([]*model.Story, error) {
	cursor, err := s.Col.Find(context.TODO(), bson.D{})
	var results []*model.Story

	if err != nil {
		defer cursor.Close(ctx)
	} else {
		for cursor.Next(ctx) {
			var result model.Story

			err := cursor.Decode(&result)
			results = append(results, &result)

			if err != nil {
				os.Exit(1)
			}
		}
	}
	return results, nil
}