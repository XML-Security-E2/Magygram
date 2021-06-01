package service_contracts

import (
	"context"
	"post-service/domain/model"
)

type PostService interface {
	CreatePost(ctx context.Context, bearer string,  post *model.PostRequest) (string, error)
	GetPostsFirstImage(ctx context.Context, postId string) (*model.Media, error)
}