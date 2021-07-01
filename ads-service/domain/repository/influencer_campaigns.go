package repository

import (
	"ads-service/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type InfluencerCampaignRepository interface {
	Create(ctx context.Context, campaign *model.InfluencerCampaign) (*mongo.InsertOneResult, error)
	GetByID(ctx context.Context, id string) (*model.InfluencerCampaign, error)
	GetAllByParentID(ctx context.Context, id string) ([]*model.InfluencerCampaign, error)
}
