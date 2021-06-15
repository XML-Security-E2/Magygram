package interactor

import (
	"go.mongodb.org/mongo-driver/mongo"
	"request-service/controller/handler"
	"request-service/domain/repository"
	"request-service/domain/service-contracts"
	"request-service/infrastructure/persistance/mongodb"
	"request-service/service"
)

type Interactor interface {
	NewVerificationRequestRepository() repository.VerificationRequestsRepository
	NewVerificationRequestService() service_contracts.VerificationRequestService
	NewVerificationRequestHandler() handler.VerificationRequestHandler
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	VerificationRequestCol *mongo.Collection
	ReportRequestCol *mongo.Collection
}

func NewInteractor(VerificationRequestCol *mongo.Collection, ReportRequestCol *mongo.Collection) Interactor {
	return &interactor{VerificationRequestCol, ReportRequestCol}
}

type appHandler struct {
	handler.VerificationRequestHandler
}

func (i interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.VerificationRequestHandler = i.NewVerificationRequestHandler()
	return appHandler
}

func (i interactor) NewVerificationRequestRepository() repository.VerificationRequestsRepository {
	return mongodb.NewVerificatioRequestsRepository(i.VerificationRequestCol)
}

func (i interactor) NewReportRequestRepository() repository.ReportRequestsRepository {
	return mongodb.NewReportRequestsRepository(i.ReportRequestCol)
}

func (i interactor) NewVerificationRequestService() service_contracts.VerificationRequestService {
	return service.NewVerificationServiceService(i.NewVerificationRequestRepository(), i.NewReportRequestRepository())
}

func (i interactor) NewVerificationRequestHandler() handler.VerificationRequestHandler {
	return handler.NewVerificationRequestHandler(i.NewVerificationRequestService())
}



