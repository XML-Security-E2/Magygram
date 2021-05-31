package service

import (
	"context"
	"github.com/go-playground/validator"
	"post-service/domain/model"
	"post-service/domain/repository"
	"post-service/domain/service-contracts"
)



type postService struct {
	repository.PostRepository
}


func NewPostService(r repository.PostRepository) service_contracts.PostService {
	return &postService{r }
}

func (p postService) CreatePost(ctx context.Context, postRequest *model.PostRequest) (string, error) {
	post, err := model.NewPost(postRequest)

	if err != nil { return "", err}

	if err := validator.New().Struct(post); err!= nil {
		return "", err
	}

	result, err := p.PostRepository.Create(ctx, post)

	if err != nil { return "", err}
	if postId, ok := result.InsertedID.(string); ok {
		if err != nil { return "", err}
		return postId, nil
	}

	return "", err
}

func (p postService) GetPostsForTimeline(ctx context.Context) ([]*model.Post, error) {
	result, err := p.PostRepository.GetAll(ctx)

	if err != nil { return nil, err}

	return result, nil
}

