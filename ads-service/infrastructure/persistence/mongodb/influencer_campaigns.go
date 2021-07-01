package mongodb

import (
	"ads-service/domain/model"
	"ads-service/domain/repository"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type influencerCampaignRepository struct {
	Col *mongo.Collection
}

func NewInfluencerCampaignRepository(Col *mongo.Collection) repository.InfluencerCampaignRepository {
	return &influencerCampaignRepository{Col}
}

func (i influencerCampaignRepository) Create(ctx context.Context, campaign *model.InfluencerCampaign) (*mongo.InsertOneResult, error) {
	return i.Col.InsertOne(ctx, campaign)
}

func (i influencerCampaignRepository) GetByID(ctx context.Context, id string) (*model.InfluencerCampaign, error) {
	var campaign = model.InfluencerCampaign{}
	err := i.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&campaign)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &campaign, nil
}

func (i influencerCampaignRepository) GetAllByParentID(ctx context.Context, id string) ([]*model.InfluencerCampaign, error) {
	cursor, err := i.Col.Find(ctx, bson.M{"parent_campaign_id": id})
	var results []*model.InfluencerCampaign

	if err != nil {
		defer cursor.Close(ctx)
		return nil, err
	} else {
		for cursor.Next(ctx) {
			var result model.InfluencerCampaign

			_ = cursor.Decode(&result)
			results = append(results, &result)
		}
	}
	return results, nil
}