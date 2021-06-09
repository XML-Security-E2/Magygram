package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"post-service/domain/model"
)

type PostRepository interface {
	Create(ctx context.Context, post *model.Post) (*mongo.InsertOneResult, error)
	GetAll(ctx context.Context) ([]*model.Post, error)
	Update(ctx context.Context, post *model.Post) (*mongo.UpdateResult, error)
	GetByID(ctx context.Context, id string) (*model.Post, error)
	GetPostsForUser(ctx context.Context, userId string) ([]*model.Post, error)
	GetPostsThatContainHashTag(ctx context.Context, hashTag string) ([]*model.Post, error)
	GetPostsByHashTag(ctx context.Context, hashTag string) ([]*model.Post, error)
	GetPostsThatContainLocation(ctx context.Context, location string) ([]*model.Post, error)
	GetPostsByLocation(ctx context.Context, location string)([]*model.Post, error)
}