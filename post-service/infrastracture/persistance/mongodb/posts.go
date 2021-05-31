package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
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

func (r *postRepository) GetAll(ctx context.Context) ([]*model.Post, error) {

	cursor, err := r.Col.Find(context.TODO(), bson.D{})
	var results []*model.Post

	if err != nil {
		defer cursor.Close(ctx)
	} else {
		for cursor.Next(ctx) {
			var result model.Post

			err := cursor.Decode(&result)
			results = append(results, &result)

			if err != nil {
				os.Exit(1)
			}
		}
	}
	return results, nil
}