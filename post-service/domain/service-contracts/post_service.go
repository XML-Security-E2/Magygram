package service_contracts

import (
	"context"
	"post-service/domain/model"
)

type PostService interface {
	CreatePost(ctx context.Context, post *model.PostRequest) (string, error)
	GetPostsForTimeline(ctx context.Context) ([]*model.Post , error)
}