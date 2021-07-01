package interactor

import (
	"ads-service/controller/handler"
	"ads-service/domain/repository"
	service_contracts "ads-service/domain/service-contracts"
	"ads-service/infrastructure/persistence/mongodb"
	"ads-service/service"
	"ads-service/service/intercomm"
	"go.mongodb.org/mongo-driver/mongo"
)

type Interactor interface {
	NewCampaignRepository() repository.CampaignRepository
	NewInfluencerCampaignRepository() repository.InfluencerCampaignRepository
	NewCampaignUpdateRequestsRepository() repository.CampaignUpdateRequestsRepository
	NewCampaignService() service_contracts.CampaignService
	NewCampaignHandler() handler.CampaignHandler
	NewAuthClient() intercomm.AuthClient
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	campaignCol *mongo.Collection
	influencerCampaignCol *mongo.Collection
	updateReqCampaignCol *mongo.Collection
}

func NewInteractor(campaignCol *mongo.Collection, influencerCampaignCol *mongo.Collection, updateReqCampaignCol *mongo.Collection) Interactor {
	return &interactor{campaignCol, influencerCampaignCol, updateReqCampaignCol}
}

type appHandler struct {
	handler.CampaignHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.CampaignHandler = i.NewCampaignHandler()
	return appHandler
}

func (i *interactor) NewAuthClient() intercomm.AuthClient {
	return intercomm.NewAuthClient()
}

func (i *interactor) NewCampaignUpdateRequestsRepository() repository.CampaignUpdateRequestsRepository {
	return mongodb.NewCampaignUpdateRequestsRepository(i.updateReqCampaignCol)
}

func (i *interactor) NewCampaignRepository() repository.CampaignRepository {
	return mongodb.NewCampaignRepository(i.campaignCol)
}

func (i *interactor) NewInfluencerCampaignRepository() repository.InfluencerCampaignRepository {
	return mongodb.NewInfluencerCampaignRepository(i.influencerCampaignCol)
}

func (i *interactor) NewCampaignService() service_contracts.CampaignService {
	return service.NewCampaignService(i.NewCampaignRepository(), i.NewInfluencerCampaignRepository(), i.NewCampaignUpdateRequestsRepository(), i.NewAuthClient())
}

func (i *interactor) NewCampaignHandler() handler.CampaignHandler {
	return handler.NewCampaignHandler(i.NewCampaignService())
}
