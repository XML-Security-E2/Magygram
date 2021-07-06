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
	"story-service/domain/service-contracts/exceptions/expired"
	"story-service/domain/service-contracts/exceptions/unauthorized"
	"story-service/logger"
	"story-service/service/intercomm"
	"story-service/tracer"
	"time"
)

type storyService struct {
	repository.StoryRepository
	intercomm.MediaClient
	intercomm.UserClient
	intercomm.AuthClient
	intercomm.RelationshipClient
	intercomm.MessageClient
	intercomm.AdsClient
}

func (p storyService) DeleteStory(ctx context.Context, bearer string, requestId string) error {
	span := tracer.StartSpanFromContext(ctx, "StoryServiceDeleteStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	retVal, err := p.AuthClient.HasRole(ctx, bearer,"delete_story")
	if err != nil{
		return errors.New("auth service not found")
	}

	request, err := p.StoryRepository.GetByID(ctx, requestId)
	if err!=nil {
		return errors.New("story not found")
	}

	if !retVal {
		userId, err := p.AuthClient.GetLoggedUserId(ctx, bearer)
		if err != nil {
			return err
		}
		if request.UserInfo.Id != userId {
			return errors.New("user not authorized for story delete")
		}
	}

	err = p.AdsClient.DeleteCampaign(ctx, bearer, request.Id)
	if err != nil {
		return err
	}

	request.IsDeleted = true

	_, err = p.StoryRepository.DeleteStory(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func NewStoryService(r repository.StoryRepository, ic intercomm.MediaClient, uc intercomm.UserClient, ac intercomm.AuthClient, rc intercomm.RelationshipClient, mc intercomm.MessageClient, adscli 	intercomm.AdsClient,
						) service_contracts.StoryService {
	return &storyService{r , ic, uc,ac,rc, mc, adscli}
}

func (p storyService) CreatePost(ctx context.Context, bearer string, file *multipart.FileHeader, tags []model.Tag) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "StoryServiceCreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
	if err != nil { return "", err}

	media, err := p.MediaClient.SaveMedia(ctx, file)
	if err != nil { return "", err}

	post, err := model.NewStory(*userInfo, "REGULAR", media, tags, "")
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

	err = p.MessageClient.CreateNotifications(ctx, &intercomm.NotificationRequest{
		Username:  userInfo.Username,
		UserId:    userInfo.Id,
		UserFromId:userInfo.Id,
		NotifyUrl: "TODO",
		ImageUrl:  post.UserInfo.ImageURL,
		Type:      intercomm.PublishedStory,
	})
	if err != nil {
		return "", err
	}


	if postId, ok := result.InsertedID.(string); ok {
		logger.LoggingEntry.WithFields(logrus.Fields{"story_id": post.Id, "user_id" : userInfo.Id}).Info("Story created")
		return postId, nil
	}

	return "", err
}

func (p storyService) CreateStoryCampaign(ctx context.Context, bearer string, file *multipart.FileHeader, tags []model.Tag, campaignReq *model.CampaignRequest) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "StoryServiceCreateStoryCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedAgentInfo(ctx, bearer)
	if err != nil { return "", err}

	media, err := p.MediaClient.SaveMedia(ctx, file)
	if err != nil { return "", err}

	post, err := model.NewStory(model.UserInfo{
		Id:       userInfo.Id,
		Username: userInfo.Username,
		ImageURL: userInfo.ImageURL,
	}, "CAMPAIGN", media, tags, userInfo.Website)
	if err != nil {
		return "", err}

	if err := validator.New().Struct(post); err!= nil {
		return "", err
	}

	campaignReq.ContentId = post.Id

	err = p.AdsClient.CreatePostCampaign(ctx, bearer, campaignReq)
	if err != nil {
		return "", err
	}

	result, err := p.StoryRepository.Create(ctx, post)
	if err != nil {
		return "", err}

	err = p.MessageClient.CreateNotifications(ctx, &intercomm.NotificationRequest{
		Username:  userInfo.Username,
		UserId:    userInfo.Id,
		UserFromId:userInfo.Id,
		NotifyUrl: "TODO",
		ImageUrl:  post.UserInfo.ImageURL,
		Type:      intercomm.PublishedStory,
	})
	if err != nil {
		return "", err
	}


	if postId, ok := result.InsertedID.(string); ok {
		return postId, nil
	}

	return "", err
}

func (p storyService) GetAllUserStoryCampaigns(ctx context.Context, bearer string) ([]*model.UsersStoryResponseWithUserInfo, error) {
	span := tracer.StartSpanFromContext(ctx, "StoryServiceGetAllUserStoryCampaigns")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	storyIds, err := p.AdsClient.GetAllActiveAgentsStoryCampaigns(ctx, bearer)
	if err != nil {
		return nil, err
	}
	var stories []*model.UsersStoryResponseWithUserInfo

	for _, storyId := range storyIds {
		story, err := p.StoryRepository.GetByID(ctx, storyId)
		if err != nil {
			return nil, err
		}

		stories = append(stories, &model.UsersStoryResponseWithUserInfo{
			Id: story.Id,
			ContentType: story.ContentType,
			Media:      story.Media,
			DateTime:    "",
			UserInfo:   story.UserInfo,
			Website:    story.Website,
		})
	}

	return stories, nil
}

func (p storyService) GetStoriesForStoryline(ctx context.Context, bearer string) ([]*model.StoryInfoResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "StoryServiceGetStoriesForStoryline")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)
	//TODO 1: napraviti getStory za usera koji eliminise njegove storije a onda izbrisati iz mapStories proveru
	log.Println("test")
	var stories []*model.Story
	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
	if err != nil {
		return nil, err
	}

	var followedUsers model.FollowedUsersResponse
	followedUsers, err = p.RelationshipClient.GetUnmutedFollowedUsers(ctx, userInfo.Id)
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

func (p storyService) GetStoryForUserMessage(ctx context.Context, bearer string, storyId string) (*model.UsersStoryResponseWithUserInfo, *model.UserInfo, error) {
	span := tracer.StartSpanFromContext(ctx, "StoryServiceGetStoryForUserMessage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	story, err := p.StoryRepository.GetByID(ctx, storyId)
	if err != nil {
		return nil, nil, err
	}

	if !p.checkIfUserContentIsAccessible(ctx, bearer, story.UserInfo.Id) {
		return nil, &story.UserInfo, &unauthorized.UnauthorizedAccessError{Msg: "User not authorized"}
	}

	if story.CreatedTime.AddDate(0,0,1).Before(time.Now()) {
		return nil, &story.UserInfo, &expired.StoryError{Msg: "Story expired"}
	}

	resp := &model.UsersStoryResponseWithUserInfo{
		Id: story.Id,
		ContentType: story.ContentType,
		Media:      story.Media,
		DateTime:    "",
		UserInfo:   story.UserInfo,
		Website:    story.Website,
	}
	return resp, nil, nil
}

func (p storyService) checkIfUserContentIsAccessible(ctx context.Context, bearer string, storyOwnerId string) bool {
	isPrivate, err := p.UserClient.IsUserPrivate(ctx, storyOwnerId)
	if err != nil {
		return false
	}

	if isPrivate {
		if bearer == "" {
			return false
		}
		userId, err := p.AuthClient.GetLoggedUserId(ctx, bearer)
		if err != nil {
			return false
		}
		if userId != storyOwnerId {
			followedUsers, err := p.RelationshipClient.GetFollowedUsers(ctx, userId)
			if err != nil {
				return false
			}
			for _, usrId := range followedUsers.Users {
				if storyOwnerId == usrId {
					return true
				}
			}
			return false
		}
	}

	return true
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
	span := tracer.StartSpanFromContext(ctx, "StoryServiceGetStoriesForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	result, err := p.StoryRepository.GetActiveStoriesForUser(ctx, userId)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"story_owner_id" : userId}).Warn("Error while getting user stories")
		return nil, err
	}
	fmt.Println(len(result))

	var media []model.MediaContent
	media = mapStoriesToMediaArray(result)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)

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

func (p storyService) GetStoryForAdmin(ctx context.Context, storyId string) (*model.StoryResponseForAdmin, error) {
	span := tracer.StartSpanFromContext(ctx, "StoryServiceGetStoryForAdmin")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	result, err := p.StoryRepository.GetByID(ctx, storyId)
	if err != nil {
		return nil, errors.New("invalid story id")
	}

	var media []model.MediaContent
	mediaContent := model.MediaContent{
		Url:         result.Media.Url,
		MediaType:   result.Media.MediaType,
		StoryId:     result.Id,
		Tags:        result.Tags,
		ContentType: result.ContentType,
		Website:     result.Website,
	}
	media = append(media, mediaContent)

	res, err := model.NewStoryResponseForAdmin(result, media, 0)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (p storyService) GetAllUserStories(ctx context.Context, bearer string) ([]*model.UsersStoryResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "StoryServiceGetAllUserStories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
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
	span := tracer.StartSpanFromContext(ctx, "StoryServiceGetStoryHighlight")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := p.AuthClient.GetLoggedUserId(ctx, bearer)
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
	span := tracer.StartSpanFromContext(ctx, "StoryServiceVisitedStoryByUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
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
			Url:         story.Media.Url,
			MediaType:   story.Media.MediaType,
			StoryId:     story.Id,
			Tags:        story.Tags,
			ContentType: story.ContentType,
			Website:     story.Website,
		}
		retVal = append(retVal, mediaContent)
	}

	return retVal
}

func (p storyService) HaveActiveStoriesLoggedUser(ctx context.Context, bearer string) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "StoryServiceHaveActiveStoriesLoggedUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
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

func (p storyService) EditStoryOwnerInfo(ctx context.Context, bearer string, userInfo *model.UserInfo) error {
	span := tracer.StartSpanFromContext(ctx, "StoryServiceEditStoryOwnerInfo")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := p.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return err
	}
	fmt.Println(userId)


	if userId != userInfo.Id {
		return errors.New("unauthorized edit")
	}

	stories, err := p.StoryRepository.GetActiveStoriesForUser(ctx, userId)
	if err != nil {
		return errors.New("invalid user id")
	}

	for _, userStory := range stories {
		userStory.UserInfo = *userInfo

		_, err = p.StoryRepository.Update(ctx, userStory)
		if err != nil {
			return err
		}
	}
	return nil
}
