package mongodb

import (
	"context"
	"errors"
	"fmt"
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

func (r *postRepository) GetOne(ctx context.Context, postId string) (*model.Post, error) {
	var post = model.Post{}
	fmt.Println(postId)
	err := r.Col.FindOne(ctx, bson.M{"_id": postId}).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) Update(ctx context.Context, post *model.Post) (*mongo.UpdateResult, error) {
	return r.Col.UpdateOne(ctx, bson.M{"_id":  post.Id},bson.D{{"$set", bson.D{
		{"description" , post.Description},
		{"location" , post.Location},
		{"post_type" , post.PostType},
		{"tags" , post.Tags},
		{"hashTags" , post.HashTags},
		{"media" , post.Media},
		{"user_info" , post.UserInfo},
		{"liked_by" , post.LikedBy},
		{"disliked_by" , post.DislikedBy},
		{"comments" , post.Comments}}}})
}