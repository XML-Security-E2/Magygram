package repository

import (
	"ads-service/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type CampaignRepository interface {
	Create(ctx context.Context, campaign *model.Campaign) (*mongo.InsertOneResult, error)
	GetAll(ctx context.Context) ([]*model.Campaign, error)
	Update(ctx context.Context, post *model.Campaign) (*mongo.UpdateResult, error)
	GetByID(ctx context.Context, id string) (*model.Campaign, error)
}