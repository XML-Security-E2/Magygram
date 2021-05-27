package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/domain/model"
)

type LoginEventRepository interface {
	Create(ctx context.Context, event *model.LoginEvent) (*mongo.InsertOneResult, error)
	GetLastByUserEmail(ctx context.Context, email string) (*model.LoginEvent, error)
}
