package service

import (
	"context"
	"github.com/go-playground/validator"
	_ "net/http"
	"post-service/domain/model"
	"post-service/domain/repository"
	"post-service/domain/service-contracts"
	"post-service/service/intercomm"
)



type postService struct {
	repository.PostRepository
	intercomm.MediaClient
	intercomm.UserClient
}


func NewPostService(r repository.PostRepository, ic intercomm.MediaClient, uc intercomm.UserClient) service_contracts.PostService {
	return &postService{r , ic, uc}
}

func (p postService) CreatePost(ctx context.Context, bearer string, postRequest *model.PostRequest) (string, error) {

	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil { return "", err}

	media, err := p.MediaClient.SaveMedia(postRequest.Media)
	if err != nil { return "", err}

	post, err := model.NewPost(postRequest, *userInfo, "REGULAR", media)

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
