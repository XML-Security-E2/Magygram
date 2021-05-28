package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/domain/model"
)

type ResetPasswordRepository interface {
	Create(ctx context.Context, user *model.ResetPassword) (*mongo.InsertOneResult, error)
	GetById(ctx context.Context, id string) (*model.ResetPassword, error)
	Update(ctx context.Context, a *model.ResetPassword) (*mongo.UpdateResult, error)
}