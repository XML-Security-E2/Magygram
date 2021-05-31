package service_contracts

import (
	"context"
	"post-service/domain/model"
)

type PostService interface {
	GetPostsForTimeline(ctx context.Context, bearer string) ([]*model.PostResponse , error)
	CreatePost(ctx context.Context, bearer string,  post *model.PostRequest) (string, error)
}