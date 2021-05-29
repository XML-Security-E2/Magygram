package repository

import (
	"auth-service/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*mongo.InsertOneResult, error)
	Update(ctx context.Context, user *model.User) (*mongo.UpdateResult, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetAllRolesByUserId(ctx context.Context, userId string) ([]model.Role, error)
}