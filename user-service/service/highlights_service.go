package service

import (
	"context"
	"errors"
	"fmt"
	"user-service/domain/model"
	"user-service/domain/repository"
	"user-service/domain/service-contracts"
	"user-service/domain/service-contracts/exceptions"
	"user-service/service/intercomm"
)

type highlightsService struct {
	repository.UserRepository
	intercomm.AuthClient
	intercomm.StoryClient
	intercomm.RelationshipClient
}

func NewHighlightsService(r repository.UserRepository, ic 	intercomm.AuthClient, pc intercomm.StoryClient, rc intercomm.RelationshipClient) service_contracts.HighlightsService {
	return &highlightsService{r, ic, pc, rc}
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


func (h highlightsService) GetProfileHighlights(ctx context.Context, bearer string, usrId string) ([]*model.HighlightProfileResponse, error) {


	owner, err := h.UserRepository.GetByID(ctx, usrId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	if !h.checkIfUserContentIsAccessible(bearer, owner) {
		return nil, &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
	}

	var response []*model.HighlightProfileResponse

	for highlightName, _ := range owner.HighlightsStory {
		response = append(response, &model.HighlightProfileResponse{
			Name: highlightName,
			Url:  owner.HighlightsStory[highlightName].Url,
		})
	}

	return response, nil
}

func (h highlightsService) checkIfUserContentIsAccessible(bearer string, owner *model.User) bool {
	if owner.IsPrivate {
		if bearer == "" {
			return false
		}
		loggedId, err := h.AuthClient.GetLoggedUserId(bearer)
		if err != nil {
			return false
		}

		if loggedId != owner.Id {
			followedUsers, err := h.RelationshipClient.GetFollowedUsers(loggedId)
			if err != nil {
				return false
			}

			for _, usrId := range followedUsers.Users {
				if owner.Id == usrId {
					return true
				}
			}
			return false
		}
	}

	return true
}

func (h highlightsService) GetProfileHighlightsByHighlightName(ctx context.Context, bearer string, name string, userId string) (*model.HighlightImageWithMedia, error) {

	owner, err := h.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	if !h.checkIfUserContentIsAccessible(bearer, owner) {
		return nil, &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
	}

	if _, ok := owner.HighlightsStory[name]; !ok {
		return nil, errors.New(fmt.Sprintf("highlights with name %s already not exist", name))
	}
	retVal := &model.HighlightImageWithMedia{
		Url:   owner.HighlightsStory[name].Url,
		Media: owner.HighlightsStory[name].Media,
	}
	return retVal, nil
}