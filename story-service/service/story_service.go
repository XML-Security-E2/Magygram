package service

import (
	"context"
	"fmt"
	"github.com/go-playground/validator"
	"mime/multipart"
	"story-service/domain/model"
	"story-service/domain/repository"
	"story-service/domain/service-contracts"
	"story-service/service/intercomm"
)

type storyService struct {
	repository.StoryRepository
	intercomm.MediaClient
	intercomm.UserClient
}

func NewStoryService(r repository.StoryRepository, ic intercomm.MediaClient, uc intercomm.UserClient) service_contracts.StoryService {
	return &storyService{r , ic, uc}
}

func (p storyService) CreatePost(ctx context.Context, bearer string, file *multipart.FileHeader) (string, error) {
	fmt.Println("udje2")
	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil { return "", err}
	fmt.Println("udje3")

	media, err := p.MediaClient.SaveMedia(file)
	if err != nil { return "", err}
	fmt.Println("udje4")

	post, err := model.NewStory(*userInfo, "REGULAR", media)
	fmt.Println("udje5")

	if err != nil { return "", err}

	if err := validator.New().Struct(post); err!= nil {
		return "", err
	}
	fmt.Println("udje6")

	result, err := p.StoryRepository.Create(ctx, post)
	fmt.Println("udje7")

	if err != nil { return "", err}
	if postId, ok := result.InsertedID.(string); ok {
		if err != nil { return "", err}
		return postId, nil
	}
	fmt.Println("udje8")

	return "", err
}

func (p storyService) GetStoriesForStoryline(ctx context.Context, bearer string) ([]*model.StoryResponse, error) {
	result, err := p.StoryRepository.GetAll(ctx)

	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return nil, err
	}

	retVal := mapStoriesToResponseStoriesDTO(result, userInfo.Id)

	return retVal, nil
}

func mapStoriesToResponseStoriesDTO(result []*model.Story, id string) []*model.StoryResponse {
	var retVal []*model.StoryResponse

	for _, story := range result {
		res, err := model.NewStoryResponse(story)

		if err != nil { return nil}

		retVal = append(retVal, res)
	}

	return retVal
}