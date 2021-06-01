package service_contracts

import (
	"context"
	"post-service/domain/model"
)

type PostService interface {
	GetPostsForTimeline(ctx context.Context, bearer string) ([]*model.PostResponse , error)
	CreatePost(ctx context.Context, bearer string,  post *model.PostRequest) (string, error)
	LikePost(ctx context.Context, bearer string,  postId string) error
	UnlikePost(ctx context.Context, bearer string,  postId string) error
	DislikePost(ctx context.Context, bearer string,  postId string) error
	UndislikePost(ctx context.Context, bearer string,  postId string) error
}