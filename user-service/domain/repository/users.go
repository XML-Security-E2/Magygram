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
	SearchForUsersByUsername(ctx context.Context, username string, loggedUserId string) ([]model.User, error)
	SearchForUsersByUsernameByGuest(ctx context.Context, username string) ([]model.User, error)
	DeleteUser(ctx context.Context, request *model.User) (*mongo.UpdateResult, error)
}