package service

import (
	"context"
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
	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil { return "", err}

	media, err := p.MediaClient.SaveMedia(file)
	if err != nil { return "", err}

	post, err := model.NewStory(*userInfo, "REGULAR", media)

	if err != nil { return "", err}

	if err := validator.New().Struct(post); err!= nil {
		return "", err
	}

	result, err := p.StoryRepository.Create(ctx, post)

	if err != nil { return "", err}
	if postId, ok := result.InsertedID.(string); ok {
		if err != nil { return "", err}
		return postId, nil
	}

	return "", err
}

func (p storyService) GetStoriesForStoryline(ctx context.Context, bearer string) ([]*model.StoryResponse, error) {
	//result, err := p.StoryRepository.GetAll(ctx)

	//userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	//if err != nil {
	//	return nil, err
	//}

	var retVal []*model.StoryResponse
	story1 := model.StoryResponse{Id: "123",
		Media: model.Media{
			Url: "http://lorempixel.com/1000/600/nature/2/",
			MediaType: "IMAGE",
		},
		UserInfo: model.UserInfo{
			Id: "123",
			Username: "nikolakolovic",
			ImageURL: "http://lorempixel.com/1000/600/nature/2/",
		},
		ContentType: "REGULAR",
	}
	retVal=append(retVal, &story1)
	story2 := model.StoryResponse{Id: "123",
		Media: model.Media{
			Url: "http://lorempixel.com/1000/600/nature/1/",
			MediaType: "IMAGE",
		},
		UserInfo: model.UserInfo{
			Id: "123",
			Username: "nkl",
			ImageURL: "assets/images/profiles/profile-1.jpg",
		},
		ContentType: "REGULAR",
	}
	retVal=append(retVal, &story2)

	//retVal := mapStoriesToResponseStoriesDTO(result, userInfo.Id)

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