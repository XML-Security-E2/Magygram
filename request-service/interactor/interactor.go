package interactor

import (
	"go.mongodb.org/mongo-driver/mongo"
	"request-service/controller/handler"
	"request-service/domain/repository"
	"request-service/domain/service-contracts"
	"request-service/infrastructure/persistance/mongodb"
	"request-service/service"
	"request-service/service/intercomm"
)

type Interactor interface {
	NewVerificationRequestRepository() repository.VerificationRequestsRepository
	NewCampaignRequestsRepository() repository.CampaignRequestsRepository
	NewVerificationRequestService() service_contracts.VerificationRequestService
	NewAgentRegistrationRequestService() service_contracts.AgentRegistrationRequestService
	NewVerificationRequestHandler() handler.VerificationRequestHandler
	NewAgentRegistrationRequestHandler() handler.AgentRegistrationRequestHandler
	NewAgentRegistrationRequestRepository() repository.AgentRegistrationRequests
	NewAppHandler() handler.AppHandler
	NewMediaClient() intercomm.MediaClient
	NewAuthClient() intercomm.AuthClient
	NewUserClient() intercomm.UserClient
}

type interactor struct {
	VerificationRequestCol *mongo.Collection
	ReportRequestCol *mongo.Collection
	AgentRegistrationRequestCol *mongo.Collection
	CampaignContentCol *mongo.Collection
}


func NewInteractor(VerificationRequestCol *mongo.Collection, ReportRequestCol *mongo.Collection, AgentRegistrationRequestCol *mongo.Collection,CampaignContentCol *mongo.Collection) Interactor {
	return &interactor{VerificationRequestCol, ReportRequestCol, AgentRegistrationRequestCol,CampaignContentCol}
}

type appHandler struct {
	handler.VerificationRequestHandler
	handler.AgentRegistrationRequestHandler
}

func (i *interactor) NewMediaClient() intercomm.MediaClient {
	return intercomm.NewMediaClient()
}

func (i *interactor) NewAuthClient() intercomm.AuthClient {
	return intercomm.NewAuthClient()
}

func (i *interactor) NewUserClient() intercomm.UserClient {
	return intercomm.NewUserClient()
}

func (i interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.VerificationRequestHandler = i.NewVerificationRequestHandler()
	appHandler.AgentRegistrationRequestHandler = i.NewAgentRegistrationRequestHandler()
	return appHandler
}

func (i interactor) NewVerificationRequestRepository() repository.VerificationRequestsRepository {
	return mongodb.NewVerificatioRequestsRepository(i.VerificationRequestCol)
}
func (i interactor) NewCampaignRequestsRepository() repository.CampaignRequestsRepository {
	return mongodb.NewCampaignRequestsRepository(i.CampaignContentCol)
}

func (i interactor) NewReportRequestRepository() repository.ReportRequestsRepository {
	return mongodb.NewReportRequestsRepository(i.ReportRequestCol)
}

func (i interactor) NewAgentRegistrationRequestRepository() repository.AgentRegistrationRequests {
	return mongodb.NewAgentRegistrationRequestsRepository(i.AgentRegistrationRequestCol)
}

func (i interactor) NewVerificationRequestService() service_contracts.VerificationRequestService {
	return service.NewVerificationServiceService(i.NewVerificationRequestRepository(),i.NewReportRequestRepository(),i.NewCampaignRequestsRepository(),i.NewMediaClient(),i.NewAuthClient(),i.NewUserClient())
}

func (i *interactor) NewAgentRegistrationRequestService() service_contracts.AgentRegistrationRequestService {
	return service.NewAgentRegistrationRequestService(i.NewAgentRegistrationRequestRepository(),i.NewAuthClient(),i.NewUserClient())
}

func (i interactor) NewVerificationRequestHandler() handler.VerificationRequestHandler {
	return handler.NewVerificationRequestHandler(i.NewVerificationRequestService())
}

func (i interactor) NewAgentRegistrationRequestHandler() handler.AgentRegistrationRequestHandler {
	return handler.NewAgentRegistrationRequestHandler(i.NewAgentRegistrationRequestService())
}


