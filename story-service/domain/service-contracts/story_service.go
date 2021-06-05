package service_contracts

import (
	"context"
	"mime/multipart"
	"story-service/domain/model"
)

type StoryService interface {
	CreatePost(ctx context.Context, bearer string, storyContent *multipart.FileHeader) (string, error)
	GetStoriesForStoryline(ctx context.Context, bearer string) ([]*model.StoryInfoResponse , error)
	GetStoriesForUser(ctx context.Context, userId string, bearer string) (*model.StoryResponse , error)
	GetAllUserStories(ctx context.Context, bearer string) ([]*model.UsersStoryResponse , error)
	VisitedStoryByUser(ctx context.Context, storyId string, bearer string) error
	GetStoryHighlight(ctx context.Context, bearer string, request *model.HighlightRequest) (*model.HighlightImageWithMedia , error)
}