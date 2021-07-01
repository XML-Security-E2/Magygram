package service

import (
	"ads-service/domain/model"
	"ads-service/domain/repository"
	"ads-service/domain/service-contracts"
	"ads-service/service/intercomm"
	"context"
	"errors"
	"github.com/go-playground/validator"
)

type campaignService struct {
	repository.CampaignRepository
	repository.InfluencerCampaignRepository
	intercomm.AuthClient
}

func NewCampaignService(r repository.CampaignRepository, ic repository.InfluencerCampaignRepository, ac intercomm.AuthClient) service_contracts.CampaignService {
	return &campaignService{r , ic, ac}
}

func (c campaignService) CreateCampaign(ctx context.Context, bearer string, campaignRequest *model.CampaignRequest) (string, error) {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return "", err
	}

	campaign, err := model.NewCampaign(campaignRequest, loggedId)
	if err != nil {
		return "", err
	}

	_, err = c.CampaignRepository.Create(ctx, campaign)
	if err != nil {
		return "", err
	}

	return campaign.Id, nil
}

func (c campaignService) CreateInfluencerCampaign(ctx context.Context, bearer string, campaignRequest *model.InfluencerCampaignRequest) (string, error) {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return "", err
	}

	//mozda posle nece trebati provera
	_, err = c.CampaignRepository.GetByID(ctx, campaignRequest.ParentCampaignId)
	if err != nil {
		return "", errors.New("invalid parent campaign id")
	}

	campaign, err := model.NewInfluencerCampaign(campaignRequest, loggedId)
	if err != nil {
		return "", err
	}

	if err = validator.New().Struct(campaign); err!= nil {
		return "", err
	}

	_, err = c.InfluencerCampaignRepository.Create(ctx, campaign)
	if err != nil {
		return "", err
	}

	return campaign.Id, nil
}
