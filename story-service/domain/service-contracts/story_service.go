package service_contracts

import (
	"context"
	"mime/multipart"
)

type StoryService interface {
	CreatePost(ctx context.Context, bearer string, storyContent *multipart.FileHeader) (string, error)
}