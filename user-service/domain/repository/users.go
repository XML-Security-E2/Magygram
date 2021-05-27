package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*mongo.InsertOneResult, error)
	Update(ctx context.Context, user *model.User) (*mongo.UpdateResult, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	//GetAllRolesByUserId(userId string) ([]model.Role, error)
}