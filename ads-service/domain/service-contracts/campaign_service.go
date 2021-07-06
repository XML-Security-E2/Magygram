package service_contracts

import (
	"ads-service/domain/model"
	"context"
)

type CampaignService interface {
	CreateCampaign(ctx context.Context, bearer string, campaignRequest *model.CampaignRequest) (string , error)
	CreateInfluencerCampaign(ctx context.Context, bearer string, campaignRequest *model.InfluencerCampaignCreateRequest) (string , error)
	UpdateCampaignRequest(ctx context.Context, bearer string, campaignRequest *model.CampaignUpdateRequestDTO) (string , error)
	GetAllActiveAgentsPostCampaigns(ctx context.Context, bearer string) ([]string, error)
	GetAllActiveAgentsStoryCampaigns(ctx context.Context,bearer string) ([]string, error)
	GetCampaignById(ctx context.Context, bearer string, campaignIds string) (*model.Campaign, error)

	GetCampaignByPostId(ctx context.Context, bearer string, contentId string) (*model.CampaignRetreiveRequest, error)
	GetCampaignByStoryId(ctx context.Context, bearer string, contentId string) (*model.CampaignRetreiveRequest, error)

	DeleteCampaignByPostId(ctx context.Context, bearer string, contentId string) error
	DeleteCampaignByStoryId(ctx context.Context, bearer string, contentId string) error

	GetUnseenPostIdsCampaignsForUser(ctx context.Context, bearer string, count int) ([]string, error)
	GetUnseenStoryIdsCampaignsForUser(ctx context.Context, bearer string, count int) ([]string, error)

	ClickOnStoryCampaignWebsite(ctx context.Context, contentId string) error
	ClickOnPostCampaignWebsite(ctx context.Context, contentId string) error

	GetPostCampaignStatistic(ctx context.Context, bearer string) ([]*model.CampaignStatisticResponse, error)
	GetStoryCampaignStatistic(ctx context.Context, bearer string) ([]*model.CampaignStatisticResponse, error)
	UpdateCampaignVisitor(ctx context.Context, bearer string, id string, campaignType string) error
}
