package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"post-service/domain/model"
)

type PostRepository interface {
	Create(ctx context.Context, post *model.Post) (*mongo.InsertOneResult, error)
}