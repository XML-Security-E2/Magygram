package service_contracts

import (
	"context"
	"post-service/domain/model"
)

type PostService interface {
	GetPostsForTimeline(ctx context.Context) ([]*model.Post , error)
	CreatePost(ctx context.Context, bearer string,  post *model.PostRequest) (string, error)
}