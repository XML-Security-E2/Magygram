package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"post-service/domain/model"
	"post-service/domain/repository"
)

type postRepository struct {
	Col *mongo.Collection
}

func NewPostRepository(Col *mongo.Collection) repository.PostRepository {
	return &postRepository{Col}
}

func (r *postRepository) Create(ctx context.Context, post *model.Post) (*mongo.InsertOneResult, error) {
	return r.Col.InsertOne(ctx, post)
}