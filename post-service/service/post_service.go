package service

import (
	"context"
	"github.com/go-playground/validator"
	"post-service/domain/model"
	"post-service/domain/repository"
	"post-service/domain/service-contracts"
	"post-service/service/intercomm"
)



type postService struct {
	repository.PostRepository
	intercomm.MediaClient
}


func NewPostService(r repository.PostRepository, ic intercomm.MediaClient) service_contracts.PostService {
	return &postService{r , ic}
}

func (p postService) CreatePost(ctx context.Context, postRequest *model.PostRequest) (string, error) {
	media, err := p.MediaClient.SaveMedia(postRequest.Media)
	if err != nil { return "", err}

	post, err := model.NewPost(postRequest, model.UserInfo{
		Id:       "123131232112",
		Username: "nikola",
		ImageURL: "nikola.jpg",
	}, "REGULAR", media)

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
