package mongodb

import (
	"ads-service/domain/model"
	"ads-service/domain/repository"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type campaignUpdateRequestsRepository struct {
	Col *mongo.Collection
}

func NewCampaignUpdateRequestsRepository(Col *mongo.Collection) repository.CampaignUpdateRequestsRepository {
	return &campaignUpdateRequestsRepository{Col}
}

func (c campaignUpdateRequestsRepository) Create(ctx context.Context, campaignUpdateReq *model.CampaignUpdateRequest) (*mongo.InsertOneResult, error) {
	return c.Col.InsertOne(ctx, campaignUpdateReq)
}

func (c campaignUpdateRequestsRepository) GetByID(ctx context.Context, id string) (*model.CampaignUpdateRequest, error) {
	var campaign = model.CampaignUpdateRequest{}
	err := c.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&campaign)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &campaign, nil
}

func (c campaignUpdateRequestsRepository) GetPendingByCampaignId(ctx context.Context, campaignId string) (*model.CampaignUpdateRequest, error) {
	var campaign = model.CampaignUpdateRequest{}
	err := c.Col.FindOne(ctx, bson.M{"campaign_id": campaignId, "campaign_update_status" : "PENDING"}).Decode(&campaign)

	if err != nil {
		return nil, nil
	}

	return &campaign, nil
}


