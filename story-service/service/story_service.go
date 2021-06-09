package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"log"
	"mime/multipart"
	"story-service/domain/model"
	"story-service/domain/repository"
	"story-service/domain/service-contracts"
	"story-service/logger"
	"story-service/service/intercomm"
)

type storyService struct {
	repository.StoryRepository
	intercomm.MediaClient
	intercomm.UserClient
	intercomm.AuthClient
	intercomm.RelationshipClient
}

func NewStoryService(r repository.StoryRepository, ic intercomm.MediaClient, uc intercomm.UserClient, ac intercomm.AuthClient, rc intercomm.RelationshipClient) service_contracts.StoryService {
	return &storyService{r , ic, uc,ac,rc}
}

func (p storyService) CreatePost(ctx context.Context, bearer string, file *multipart.FileHeader) (string, error) {
	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil { return "", err}

	media, err := p.MediaClient.SaveMedia(file)
	if err != nil { return "", err}

	post, err := model.NewStory(*userInfo, "REGULAR", media)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userInfo.Id}).Warn("Story creating validation failure")
		return "", err}

	if err := validator.New().Struct(post); err!= nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userInfo.Id}).Warn("Story creating validation failure")
		return "", err
	}

	result, err := p.StoryRepository.Create(ctx, post)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : userInfo.Id}).Error("Story database create failure")
		return "", err}

	if postId, ok := result.InsertedID.(string); ok {
		logger.LoggingEntry.WithFields(logrus.Fields{"story_id": post.Id, "user_id" : userInfo.Id}).Info("Story created")
		return postId, nil
	}

	return "", err
}

func (p storyService) GetStoriesForStoryline(ctx context.Context, bearer string) ([]*model.StoryInfoResponse, error) {
	//TODO 1: napraviti getStory za usera koji eliminise njegove storije a onda izbrisati iz mapStories proveru
	log.Println("test")
	var stories []*model.Story
	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return nil, err
	}

	var followedUsers model.FollowedUsersResponse
	followedUsers, err = p.RelationshipClient.GetFollowedUsers(userInfo.Id)
	if err != nil {
		return nil, err
	}

	for _, userId := range followedUsers.Users {
		var userStories []*model.Story
		userStories, _ = p.StoryRepository.GetActiveStoriesForUser(ctx,userId)
		stories= append(stories, userStories...)
	}

	storiesMap := makeStoriesMapFromArray(stories, userInfo)

    retVal := mapStoriesFromMapToResponseStoriesInfoDTO(storiesMap, userInfo.Id)

	retVal = sortFirstUnvisited(retVal)

	return retVal, nil
}

func sortFirstUnvisited(stories []*model.StoryInfoResponse) []*model.StoryInfoResponse {
	var visited []*model.StoryInfoResponse
	var unvisited []*model.StoryInfoResponse

	for _, story := range stories {
		if story.Visited==true{
			visited= append(visited, story)
		}else{
			unvisited= append(unvisited,story)
		}
	}

	return append(unvisited, visited...)
}

func mapStoriesFromMapToResponseStoriesInfoDTO(storiesMap map[string][]*model.Story, userId string) []*model.StoryInfoResponse {
	var retVal []*model.StoryInfoResponse

	for _, element := range storiesMap {
		visited, _ := hasUserVisitedStories(element, userId)
		res, err := model.NewStoryInfoResponse(element[0],visited)
		if err != nil { return nil}
		retVal = append(retVal, res)
	}
	return retVal

}
//TODO 3: mora userId da bude u svakom, ako u jednom nije visited je false
func hasUserVisitedStories(stories []*model.Story, id string) (bool, int) {
	for index, story := range stories{
		if !hasUserVisitStory(story, id){
			return false,index
		}
	}

	return true,0
}

func hasUserVisitStory(story *model.Story, id string) bool {
	for _, storyVisitor := range story.VisitedBy{
		if storyVisitor.Id==id{
			return true
		}
	}
	return false
}

func makeStoriesMapFromArray(stories []*model.Story, userInfo *model.UserInfo) map[string][]*model.Story {
	elementMap := make(map[string][]*model.Story)
	for i := 0; i < len(stories); i +=1 {
		if stories[i].UserInfo.Id!=userInfo.Id { // TODO 2: eliminise svoje storije, to obrisati kada se odradi TODO1
			elementMap[stories[i].UserInfo.Id]=append(elementMap[stories[i].UserInfo.Id], stories[i])
		}
	}
	return elementMap
}


func (p storyService) GetStoriesForUser(ctx context.Context, userId string, bearer string) (*model.StoryResponse, error) {
	result, err := p.StoryRepository.GetActiveStoriesForUser(ctx, userId)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"story_owner_id" : userId}).Warn("Error while getting user stories")
		return nil, err
	}
	fmt.Println(len(result))

	var media []model.MediaContent
	media = mapStoriesToMediaArray(result)

	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)

	if err != nil {
		return nil, err
	}
	_, index := hasUserVisitedStories(result,userInfo.Id)

	res, err := model.NewStoryResponse(result[0], media, index)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (p storyService) GetAllUserStories(ctx context.Context, bearer string) ([]*model.UsersStoryResponse, error) {
	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return nil, err
	}

	var userStories []*model.UsersStoryResponse
	result, err := p.StoryRepository.GetStoriesForUser(ctx, userInfo.Id)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"story_owner_id" : userInfo.Id}).Warn("Error while getting user stories")
		return nil, err
	}

	for _, story := range result {
		userStories = append(userStories, &model.UsersStoryResponse{
			Id: story.Id,
			ContentType: story.ContentType,
			Media:      story.Media,
			DateTime:    "",
		})
	}

	return userStories, nil
}

func (p storyService) GetStoryHighlight(ctx context.Context, bearer string, request *model.HighlightRequest) (*model.HighlightImageWithMedia, error) {
	userId, err := p.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	highlights := &model.HighlightImageWithMedia{
		Url:   "",
		Media: []model.IdWithMedia{},
	}
	for _, storyId := range request.StoryIds {
		story, errs := p.StoryRepository.GetByID(ctx, storyId)

		if errs != nil {
			return nil, err
		}
		if story.UserInfo.Id != userId {
			logger.LoggingEntry.WithFields(logrus.Fields{"logged_user_id" : userId,
														 "story_owner_id" : story.UserInfo.Id,
														 "story_id" : storyId}).Warn("Unauthorized to use story as highlights")
			return nil, errors.New("desired stories cannot be in users highlights")
		}
		highlights.Media = append(highlights.Media, model.IdWithMedia{
			Id:    story.Id,
			Media: story.Media,
		})

		if story.Media.MediaType == "IMAGE" && highlights.Url == "" {
			highlights.Url = story.Media.Url
		}

	}

	return highlights, nil
}


func (p storyService) VisitedStoryByUser(ctx context.Context, storyId string, bearer string) error {
	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return err
	}

	story, err := p.StoryRepository.GetByID(ctx, storyId)
	if err != nil {
		return err
	}

	if !hasUserVisitStory(story,userInfo.Id){
		story.VisitedBy=append(story.VisitedBy, *userInfo)
	}

	p.StoryRepository.Update(ctx, story)
	if err != nil {
		return err
	}

	return nil
}

func mapStoriesToMediaArray(result []*model.Story) []model.MediaContent {
	var retVal []model.MediaContent

	for _, story := range result {
		mediaContent := model.MediaContent{
			Url: story.Media.Url,
			MediaType: story.Media.MediaType,
			StoryId: story.Id,
		}
		retVal = append(retVal, mediaContent)
	}

	return retVal
}

func (p storyService) HaveActiveStoriesLoggedUser(ctx context.Context, bearer string) (bool, error) {
	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return false,err
	}

	result, err := p.StoryRepository.GetActiveStoriesForUser(ctx, userInfo.Id)
	if err != nil {
		return false, err
	}

	if len(result)==0{
		return false,nil
	}

	return true, nil
}
