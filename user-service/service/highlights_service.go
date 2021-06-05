package service

import (
	"context"
	"errors"
	"fmt"
	"user-service/domain/model"
	"user-service/domain/repository"
	"user-service/domain/service-contracts"
	"user-service/service/intercomm"
)

type highlightsService struct {
	repository.UserRepository
	intercomm.AuthClient
	intercomm.StoryClient
}

func NewHighlightsService(r repository.UserRepository, ic 	intercomm.AuthClient, pc intercomm.StoryClient) service_contracts.HighlightsService {
	return &highlightsService{r, ic, pc}
}


func (h highlightsService) CreateHighlights(ctx context.Context, bearer string, highlights *model.HighlightRequest) (*model.HighlightProfileResponse,error) {
	userId, err := h.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	user, err := h.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	if _, ok := user.HighlightsStory[highlights.Name]; ok {
		return nil, errors.New(fmt.Sprintf("highlights with name %s already exist", highlights.Name))
	}

	user.HighlightsStory[highlights.Name] = model.HighlightImageWithMedia{}

	highlightsResp, err := h.StoryClient.GetStoryHighlightIfValid(bearer, highlights)
	if err != nil {
		return nil, err
	}

	user.HighlightsStory[highlights.Name] = *highlightsResp

	_, err = h.UserRepository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &model.HighlightProfileResponse{
		Name: highlights.Name,
		Url:  highlightsResp.Url,
	},nil
}


func (h highlightsService) GetProfileHighlights(ctx context.Context, bearer string) ([]*model.HighlightProfileResponse, error) {
	userId, err := h.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	user, err := h.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	var response []*model.HighlightProfileResponse

	for highlightName, _ := range user.HighlightsStory {
		response = append(response, &model.HighlightProfileResponse{
			Name: highlightName,
			Url:  user.HighlightsStory[highlightName].Url,
		})
	}

	return response, nil
}

func (h highlightsService) GetProfileHighlightsByHighlightName(ctx context.Context, bearer string, name string) (*model.HighlightImageWithMedia, error) {
	userId, err := h.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	user, err := h.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	if _, ok := user.HighlightsStory[name]; !ok {
		return nil, errors.New(fmt.Sprintf("highlights with name %s already not exist", name))
	}
	retVal := &model.HighlightImageWithMedia{
		Url:   user.HighlightsStory[name].Url,
		Media: user.HighlightsStory[name].Media,
	}
	return retVal, nil
}