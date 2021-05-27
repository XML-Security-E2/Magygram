package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/domain/model"
)

type AccountActivationRepository interface {
	Create(ctx context.Context, user *model.AccountActivation)  (*mongo.InsertOneResult, error)
	GetById(ctx context.Context, id string) (*model.AccountActivation, error)
	Update(ctx context.Context, a *model.AccountActivation) (*mongo.UpdateResult, error)
}

