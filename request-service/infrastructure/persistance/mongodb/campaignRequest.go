package mongodb

import (
"context"
"errors"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo"
"os"
"request-service/domain/model"
"request-service/domain/repository"
)

type campaignRequest struct {
	Col *mongo.Collection
}

func (v *campaignRequest) GetAllRequests(ctx context.Context, requestId string) ([]*model.CampaignRequest, error) {
	cursor, err := v.Col.Find(context.TODO(), bson.M{"deleted": false, "_id":  requestId})
	var results []*model.CampaignRequest

	if err != nil {
		defer cursor.Close(ctx)
	} else {
		for cursor.Next(ctx) {
			var result model.CampaignRequest

			err := cursor.Decode(&result)
			results = append(results, &result)

			if err != nil {
				os.Exit(1)
			}
		}
	}
	return results, nil
}
func (v campaignRequest) DeleteRequest(ctx context.Context, request *model.CampaignRequest) (*mongo.UpdateResult, error) {
	return v.Col.UpdateOne(ctx, bson.M{"_id":  request.Id},bson.D{{"$set", bson.D{
		{"deleted" , true}}}})
}

func (v *campaignRequest) GetRequestById(ctx context.Context, requestId string) (*model.CampaignRequest, error) {
	var post = model.CampaignRequest{}
	err := v.Col.FindOne(ctx, bson.M{"_id": requestId}).Decode(&post)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &post, nil
}

func (v campaignRequest) CreateRequest(ctx context.Context, request *model.CampaignRequest) (*mongo.InsertOneResult, error) {
	return v.Col.InsertOne(ctx, request)
}

func NewCampaignRequestsRepository(Col *mongo.Collection) repository.CampaignRequestsRepository {
	return &campaignRequest{Col}
}

