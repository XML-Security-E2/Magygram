package mongodb

import (
	"context"
	"errors"
	"github.com/beevik/guid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"post-service/domain/model"
	"post-service/domain/repository"
	"post-service/logger"
)

type postRepository struct {
	Col *mongo.Collection
	locationCol *mongo.Collection
	tagCol *mongo.Collection
}

func NewPostRepository(Col *mongo.Collection, locationCol *mongo.Collection, tagCol *mongo.Collection) repository.PostRepository {
	return &postRepository{Col, locationCol, tagCol}
}

func (r *postRepository) Create(ctx context.Context, post *model.Post) (*mongo.InsertOneResult, error) {
	r.InsertLocation(ctx, post.Location)
	for _, tag := range post.Tags {
		r.InsertTag(ctx, tag)
	}
	return r.Col.InsertOne(ctx, post)
}

func (r *postRepository) InsertLocation(ctx context.Context, name string) error {
	location := model.Location{guid.New().String(), name}
	_, err := r.locationCol.InsertOne(ctx, location)
	return err
}

func (r *postRepository) InsertTag(ctx context.Context, tag model.Tag) error {
	_, err := r.tagCol.InsertOne(ctx, tag)
	return err
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

func (r *postRepository) GetByID(ctx context.Context, id string) (*model.Post, error) {

	var post = model.Post{}
	err := r.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"post_id" : id}).Warn("Invalid post id")
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
		{"post_type" , post.ContentType},
		{"tags" , post.Tags},
		{"hashTags" , post.HashTags},
		{"media" , post.Media},
		{"user_info" , post.UserInfo},
		{"liked_by" , post.LikedBy},
		{"disliked_by" , post.DislikedBy},
		{"comments" , post.Comments}}}})
}

func (r *postRepository) GetPostsForUser(ctx context.Context, userId string) ([]*model.Post, error) {
	cursor, err := r.Col.Find(ctx, bson.M{"user_info.id": userId})
	var results []*model.Post

	if err != nil {
		defer cursor.Close(ctx)
		return nil, err
	} else {
		for cursor.Next(ctx) {
			var result model.Post

			_ = cursor.Decode(&result)
			results = append(results, &result)

		}
	}
	return results, nil
}

func (r *postRepository) GetPostsThatContainHashTag(ctx context.Context, hashTag string) ([]*model.Post, error) {
	var posts []*model.Post

	cursor, err := r.Col.Find(ctx, bson.M{"hashTags": bson.M{"$regex": hashTag, "$options": "i"}})

	if err != nil {
		return nil, err
	} else {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var record model.Post
			err := cursor.Decode(&record)
			posts = append(posts,&record)

			if err != nil {
				return nil, err
			}
		}
	}

	return posts, nil
}

func (r *postRepository) GetPostsByHashTag(ctx context.Context, hashTag string) ([]*model.Post, error) {
	var posts []*model.Post

	var arr []string
	arr=append(arr, hashTag)

	cursor, err := r.Col.Find(ctx, bson.M{"hashTags": bson.M{"$in": arr }})

	if err != nil {
		return nil, err
	} else {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var record model.Post
			err := cursor.Decode(&record)
			posts = append(posts,&record)

			if err != nil {
				return nil, err
			}
		}
	}

	return posts, nil
}

func (r *postRepository) GetPostsThatContainLocation(ctx context.Context, location string) ([]*model.Post, error) {
	var posts []*model.Post

	cursor, err := r.Col.Find(ctx, bson.M{"location": bson.M{"$regex": location, "$options": "i"}})
	if err != nil {
		return nil, err
	} else {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var record model.Post
			err := cursor.Decode(&record)
			posts = append(posts, &record)
			if err != nil {
				return nil, err
			}
		}
	}

	return posts, nil
}

func (r *postRepository) GetPostsByLocation(ctx context.Context, location string) ([]*model.Post, error) {
	var posts []*model.Post

	cursor, err := r.Col.Find(ctx, bson.M{"location": location})

	if err != nil {
		return nil, err
	} else {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var record model.Post
			err := cursor.Decode(&record)
			posts = append(posts,&record)

			if err != nil {
				return nil, err
			}
		}
	}

	return posts, nil
}

func (r *postRepository) GetPostsByPostIdArray(ctx context.Context, ids []string) ([]*model.Post, error) {
	cursor, err := r.Col.Find(ctx,bson.M{"_id": bson.M{"$in": ids}})
	var results []*model.Post

	if err != nil {
		defer cursor.Close(ctx)
		return nil, err
	} else {
		for cursor.Next(ctx) {
			var result model.Post

			_ = cursor.Decode(&result)
			results = append(results, &result)

		}
	}
	return results, nil
}