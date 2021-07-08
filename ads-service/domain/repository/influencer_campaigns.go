package repository

import (
	"ads-service/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type InfluencerCampaignRepository interface {
	Create(ctx context.Context, campaign *model.InfluencerCampaign) (*mongo.InsertOneResult, error)
	Update(ctx context.Context, post *model.InfluencerCampaign) (*mongo.UpdateResult, error)

	GetByID(ctx context.Context, id string) (*model.InfluencerCampaign, error)
	GetAllByOwnerID(ctx context.Context, id string, campaignType string) ([]*model.InfluencerCampaign, error)

	GetByContentIDAndType(ctx context.Context, contentId string, campaignType string) (*model.InfluencerCampaign, error)

}
