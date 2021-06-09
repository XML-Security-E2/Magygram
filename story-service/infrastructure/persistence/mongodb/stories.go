package mongodb

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"story-service/domain/model"
	"story-service/domain/repository"
	"story-service/logger"
	"time"
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
	cursor, err := s.Col.Find(context.TODO(), bson.M{"created_time": bson.M{"$gte": primitive.NewObjectIDFromTimestamp(time.Now().AddDate(0,0,-1))}})

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


func (s storyRepository) GetStoriesForUser(ctx context.Context, userId string) ([]*model.Story, error) {
	cursor, err := s.Col.Find(context.TODO(), bson.M{"user_info.id": userId})
	var results []*model.Story

	if err != nil {
		defer cursor.Close(ctx)
		return nil, err
	} else {
		for cursor.Next(ctx) {
			var result model.Story

			_ = cursor.Decode(&result)
			results = append(results, &result)

		}
	}
	return results, nil
}


func (s storyRepository) GetActiveStoriesForUser(ctx context.Context, userId string) ([]*model.Story, error) {
	cursor, err := s.Col.Find(context.TODO(), bson.M{"user_info.id": userId , "created_time" : bson.M{"$gte" : primitive.NewDateTimeFromTime(time.Now().AddDate(0,0,-1))}})
	log.Println(primitive.NewObjectIDFromTimestamp(time.Now().AddDate(0,0,-1)))
	var results []*model.Story

	if err != nil {
		defer cursor.Close(ctx)
		return nil, err
	} else {
		for cursor.Next(ctx) {
			var result model.Story

			_ = cursor.Decode(&result)
			results = append(results, &result)
		}
	}
	return results, nil
}

func (s storyRepository) GetByID(ctx context.Context, storyId string) (*model.Story, error) {
	var story = model.Story{}

	err := s.Col.FindOne(ctx, bson.M{"_id": storyId}).Decode(&story)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"story_id" : storyId}).Warn("Invalid story id")
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &story, nil
}

func (s storyRepository) Update(ctx context.Context, story *model.Story) (*mongo.UpdateResult, error) {
	return s.Col.UpdateOne(ctx, bson.M{"_id":  story.Id},bson.D{{"$set", bson.D{{"content_type" , story.ContentType},
		{"media" , story.Media},
		{"user_info" , story.UserInfo},
		{"visited_by" , story.VisitedBy}}}})
}