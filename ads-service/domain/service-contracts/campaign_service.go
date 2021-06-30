package service_contracts

import (
	"ads-service/domain/model"
	"context"
)

type CampaignService interface {
	CreateCampaign(ctx context.Context, bearer string) (*model.Campaign , error)
}
