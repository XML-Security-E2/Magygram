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

func (i influencerCampaignRepository) Update(ctx context.Context, campaign *model.InfluencerCampaign) (*mongo.UpdateResult, error) {
	return i.Col.UpdateOne(ctx, bson.M{"_id":  campaign.Id},bson.D{{"$set", bson.D{
															{"seen_by" , campaign.SeenBy},
															{"daily_seen_by", campaign.DailySeenBy},
															{"website_click_count", campaign.WebsiteClickCount}}}})}

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

func (i influencerCampaignRepository) GetAllByOwnerID(ctx context.Context, id string, campaignType string) ([]*model.InfluencerCampaign, error) {
	cursor, err := i.Col.Find(ctx, bson.M{"owner_id": id, "campaign_type": campaignType})
	var results []*model.InfluencerCampaign

	if err != nil {
		if cursor != nil {
			defer cursor.Close(ctx)
		}
		return nil, err
	} else {
		if cursor != nil {
			for cursor.Next(ctx) {
			var result model.InfluencerCampaign

			_ = cursor.Decode(&result)
			results = append(results, &result)
			}
		} else {
			return nil, err
		}
	}
	return results, nil
}

func (i influencerCampaignRepository) GetByContentIDAndType(ctx context.Context, contentId string, campaignType string) (*model.InfluencerCampaign, error) {
	var campaign = model.InfluencerCampaign{}
	err := i.Col.FindOne(ctx, bson.M{"content_id": contentId,
		"campaign_type": campaignType}).Decode(&campaign)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &campaign, nil
}