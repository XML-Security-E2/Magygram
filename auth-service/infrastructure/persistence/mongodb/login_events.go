package mongodb

import (
	"auth-service/domain/model"
	"auth-service/domain/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type loginEventRepository struct {
	Col *mongo.Collection
}

func NewLoginEventRepository(Col *mongo.Collection) repository.LoginEventRepository {
	return &loginEventRepository{Col}
}

func (r *loginEventRepository) Create(ctx context.Context, event *model.LoginEvent) (*mongo.InsertOneResult, error) {
	return r.Col.InsertOne(ctx, event)
}

func (r *loginEventRepository) GetLastByUserEmail(ctx context.Context, email string) (*model.LoginEvent, error) {
	var events []model.LoginEvent

	opts := options.Find()
	opts.SetSort(bson.D{{"timestamp", -1}})
	sortCursor, err := r.Col.Find(ctx, bson.M{"userEmail": email}, opts)
	if err != nil {
		return nil, err
	}
	if err = sortCursor.All(ctx, &events); err != nil {
		return nil, err
	}
	return &events[0], err
}