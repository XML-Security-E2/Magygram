package mongodb

import (
	"ads-service/domain/model"
	"ads-service/domain/repository"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
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

func (c campaignUpdateRequestsRepository) Update(ctx context.Context, request *model.CampaignUpdateRequest) (*mongo.UpdateResult, error) {
	return c.Col.UpdateOne(ctx, bson.M{"_id":  request.Id},bson.D{{"$set", bson.D{
		{"campaign_update_status" , request.CampaignUpdateStatus}}}})}

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

func (c campaignUpdateRequestsRepository) GetAllPending(ctx context.Context) ([]*model.CampaignUpdateRequest, error) {
	cursor, err := c.Col.Find(ctx, bson.M{"campaign_update_status" : "PENDING",
										  "requested_date" : bson.M{"$lt" : primitive.NewDateTimeFromTime(time.Now().AddDate(0,0,-1))}})


	var results []*model.CampaignUpdateRequest

	if err != nil {
		if cursor != nil {
			defer cursor.Close(ctx)
		}
		return []*model.CampaignUpdateRequest{}, err
	} else {
		if cursor != nil {
			for cursor.Next(ctx) {
				var result model.CampaignUpdateRequest

				err := cursor.Decode(&result)
				results = append(results, &result)

				if err != nil {
					return []*model.CampaignUpdateRequest{}, err
				}
			}
		} else {
			return []*model.CampaignUpdateRequest{}, nil
		}
	}
	return results, nil
}


