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

func (p storyService) GetStoriesForStoryline(ctx context.Context, bearer string) ([]*model.StoryInfoResponse, error) {
	//TODO: napraviti getStory za usera koji eliminise njegove storije a onda izbrisati iz mapStories proveru
	result, err := p.StoryRepository.GetAll(ctx)

	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return nil, err
	}

	retVal := mapStoriesToResponseStoriesInfoDTO(result, userInfo.Id)

	return retVal, nil
}

func mapStoriesToResponseStoriesInfoDTO(result []*model.Story, id string) []*model.StoryInfoResponse {
	var retVal []*model.StoryInfoResponse

	for _, story := range result {
		if story.UserInfo.Id!=id && !hasUserCreatedSomeStory(retVal, story.UserInfo.Id) {
			res, err := model.NewStoryInfoResponse(story)

			if err != nil { return nil}

			retVal = append(retVal, res)
		}
	}

	return retVal
}

func hasUserCreatedSomeStory(result []*model.StoryInfoResponse, id string) bool {
	for _, storyInfo := range result {
		if storyInfo.UserInfo.Id==id {
			return true
		}
	}
	return false
}

func (p storyService) GetStoriesForUser(ctx context.Context, userId string) (*model.StoryResponse, error) {
	result, err := p.StoryRepository.GetStoriesForUser(ctx, userId)

	if err != nil {
		return nil, err
	}
	var media []model.Media
	media = mapStoriesToMediaArray(result)

	res, err := model.NewStoryResponse(result[0], media)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func mapStoriesToMediaArray(result []*model.Story) []model.Media {
	var retVal []model.Media

	for _, story := range result {
		retVal = append(retVal, story.Media)
	}

	return retVal
}
