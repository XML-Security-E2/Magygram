package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *postRepository) GetByID(ctx context.Context, id string) (*model.Post, error) {

	var post = model.Post{}
	err := r.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &post, nil
}