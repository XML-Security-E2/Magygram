package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"request-service/domain/model"
)

type CampaignRequestsRepository interface {
	CreateRequest(ctx context.Context, user *model.CampaignRequest) (*mongo.InsertOneResult, error)
	GetAllRequests(ctx context.Context, requestId string) ([]*model.CampaignRequest, error)
	DeleteRequest(ctx context.Context, request *model.CampaignRequest) (*mongo.UpdateResult, error)
	GetRequestById(ctx context.Context, requestId string) (*model.CampaignRequest, error)

}

