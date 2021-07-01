package service_contracts

import (
	"ads-service/domain/model"
	"context"
)

type CampaignService interface {
	CreateCampaign(ctx context.Context, bearer string, campaignRequest *model.CampaignRequest) (string , error)
	CreateInfluencerCampaign(ctx context.Context, bearer string, campaignRequest *model.InfluencerCampaignRequest) (string , error)
}
