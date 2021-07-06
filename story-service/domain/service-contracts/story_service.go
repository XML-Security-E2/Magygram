package service_contracts

import (
	"context"
	"mime/multipart"
	"story-service/domain/model"
)

type StoryService interface {
	CreatePost(ctx context.Context, bearer string, storyContent *multipart.FileHeader, tags []model.Tag) (string, error)
	CreateStoryCampaign(ctx context.Context, bearer string, storyContent *multipart.FileHeader, tags []model.Tag, campaignReq *model.CampaignRequest) (string, error)
	GetAllUserStoryCampaigns(ctx context.Context, bearer string) ([]*model.UsersStoryResponseWithUserInfo , error)

	GetStoriesForStoryline(ctx context.Context, bearer string) ([]*model.StoryInfoResponse , error)
	GetStoriesForUser(ctx context.Context, userId string, bearer string) (*model.StoryResponse , error)
	GetAllUserStories(ctx context.Context, bearer string) ([]*model.UsersStoryResponse , error)
	VisitedStoryByUser(ctx context.Context, storyId string, bearer string) error
	GetStoryHighlight(ctx context.Context, bearer string, request *model.HighlightRequest) (*model.HighlightImageWithMedia , error)
	HaveActiveStoriesLoggedUser(ctx context.Context, bearer string) (bool, error)
	DeleteStory(ctx context.Context, bearer string, requestId string) error
	EditStoryOwnerInfo(ctx context.Context, bearer string, userInfo *model.UserInfo) error
	GetStoryForUserMessage(ctx context.Context, bearer string, storyId string) (*model.UsersStoryResponseWithUserInfo, *model.UserInfo, error)
	GetStoryForAdmin(ctx context.Context, storyId string) (*model.StoryResponseForAdmin, error)

	GetStoryMediaAndWebsiteByIds(ctx context.Context, ids *model.FollowedUsersResponse) ([]*model.IdMediaWebsiteResponse, error)
}