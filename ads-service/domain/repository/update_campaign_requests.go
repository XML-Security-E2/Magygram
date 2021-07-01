package repository

import (
	"ads-service/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type CampaignUpdateRequestsRepository interface {
	Create(ctx context.Context, campaignUpdateReq *model.CampaignUpdateRequest) (*mongo.InsertOneResult, error)
	GetByID(ctx context.Context, id string) (*model.CampaignUpdateRequest, error)
	GetPendingByCampaignId(ctx context.Context, campaignId string) (*model.CampaignUpdateRequest, error)
}
