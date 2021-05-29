package repository

import (
	"auth-service/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginEventRepository interface {
	Create(ctx context.Context, event *model.LoginEvent) (*mongo.InsertOneResult, error)
	GetLastByUserEmail(ctx context.Context, email string) (*model.LoginEvent, error)
}
