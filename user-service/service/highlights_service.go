package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"user-service/domain/model"
	"user-service/domain/repository"
	"user-service/domain/service-contracts"
	"user-service/domain/service-contracts/exceptions"
	"user-service/logger"
	"user-service/service/intercomm"
	"user-service/tracer"
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
	span := tracer.StartSpanFromContext(ctx, "UserServiceCreateHighlights")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := h.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return nil, err
	}

	user, err := h.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	if _, ok := user.HighlightsStory[highlights.Name]; ok {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": user.Id, "highlight_name" : highlights.Name}).Warn("Highlight name already exist")
		return nil, errors.New(fmt.Sprintf("highlights with name %s already exist", highlights.Name))
	}

	user.HighlightsStory[highlights.Name] = model.HighlightImageWithMedia{}

	highlightsResp, err := h.StoryClient.GetStoryHighlightIfValid(ctx, bearer, highlights)
	if err != nil {
		return nil, err
	}

	user.HighlightsStory[highlights.Name] = *highlightsResp

	_, err = h.UserRepository.Update(ctx, user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId,
													 "highlights_name" : highlights.Name,
													 "story_ids": highlights.StoryIds}).Error("User create highlights failure")

		return nil, err
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId,	"highlights_name" : highlights.Name}).Info("Highlights created")

	return &model.HighlightProfileResponse{
		Name: highlights.Name,
		Url:  highlightsResp.Url,
	},nil
}


func (h highlightsService) GetProfileHighlights(ctx context.Context, bearer string, usrId string) ([]*model.HighlightProfileResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "UserServiceGetProfileHighlights")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	owner, err := h.UserRepository.GetByID(ctx, usrId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	retVal, err := h.AuthClient.HasRole(ctx, bearer,"search_all_post_by_hashtag")
	if err!=nil{
		return nil, errors.New("auth service not found")
	}

	if !h.checkIfUserContentIsAccessible(ctx, bearer, owner) {
		if !retVal {
			logger.LoggingEntry.WithFields(logrus.Fields{"content_owner_id": owner.Id}).Warn("Unauthorized access")
			return nil, &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
		}
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

func (h highlightsService) checkIfUserContentIsAccessible(ctx context.Context, bearer string, owner *model.User) bool {
	if owner.IsPrivate {
		if bearer == "" {
			return false
		}
		loggedId, err := h.AuthClient.GetLoggedUserId(ctx, bearer)
		if err != nil {
			return false
		}

		if loggedId != owner.Id {
			followedUsers, err := h.RelationshipClient.GetFollowedUsers(ctx, loggedId)
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
	span := tracer.StartSpanFromContext(ctx, "UserServiceGetProfileHighlightsByHighlightName")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	owner, err := h.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	retValRole, err := h.AuthClient.HasRole(ctx, bearer,"view_profile_highlights")
	if err!=nil{
		return nil, errors.New("auth service not found")
	}

	if !h.checkIfUserContentIsAccessible(ctx, bearer, owner) {
		if !retValRole {
			logger.LoggingEntry.WithFields(logrus.Fields{"content_owner_id" : owner.Id}).Warn("Unauthorized access")
			return nil, &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
		}
	}

	if _, ok := owner.HighlightsStory[name]; !ok {
		logger.LoggingEntry.WithFields(logrus.Fields{"highlights_name" : name, "content_owner_id" : owner.Id}).Warn("Invalid highlights name")
		return nil, errors.New(fmt.Sprintf("highlights with name %s not exist", name))
	}

	retVal := &model.HighlightImageWithMedia{
		Url:   owner.HighlightsStory[name].Url,
		Media: owner.HighlightsStory[name].Media,
	}
	return retVal, nil
}