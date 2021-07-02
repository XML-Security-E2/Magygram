package repository

import (
	"ads-service/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type CampaignRepository interface {
	Create(ctx context.Context, campaign *model.Campaign) (*mongo.InsertOneResult, error)
	DeleteByID(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*model.Campaign, error)
	Update(ctx context.Context, post *model.Campaign) (*mongo.UpdateResult, error)
	GetByID(ctx context.Context, id string) (*model.Campaign, error)
	GetAllFutureByOwnerIDAndType(ctx context.Context, ownerId string, campaignType string) ([]*model.Campaign, error)
	GetFutureByContentIDAndType(ctx context.Context, contentId string, campaignType string) (*model.Campaign, error)
	GetByContentIDAndType(ctx context.Context, contentId string, campaignType string) (*model.Campaign, error)
}