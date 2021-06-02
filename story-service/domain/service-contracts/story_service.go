package service_contracts

import (
	"context"
	"mime/multipart"
	"story-service/domain/model"
)

type StoryService interface {
	CreatePost(ctx context.Context, bearer string, storyContent *multipart.FileHeader) (string, error)
	GetStoriesForStoryline(ctx context.Context, bearer string) ([]*model.StoryInfoResponse , error)
	GetStoriesForUser(ctx context.Context, userId string) (*model.StoryResponse , error)
}