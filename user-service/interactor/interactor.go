package interactor

import (
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/controller/handler"
	"user-service/domain/repository"
	"user-service/domain/service-contracts"
	"user-service/infrastructure/persistence/mongodb"
	"user-service/service"
	"user-service/service/intercomm"
)

type Interactor interface {
	NewUserRepository() repository.UserRepository
	NewAccountActivationRepository() repository.AccountActivationRepository
	NewUserService() service_contracts.UserService
	NewAccountActivationService() service_contracts.AccountActivationService
	NewAuthClient() intercomm.AuthClient
	NewUserHandler() handler.UserHandler
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	UserCol *mongo.Collection
	AccCol *mongo.Collection
	ResPwdCol *mongo.Collection
}

func NewInteractor(UserCol *mongo.Collection, AccCol *mongo.Collection, ResPwdCol *mongo.Collection) Interactor {
	return &interactor{UserCol, AccCol,  ResPwdCol}
}

type appHandler struct {
	handler.UserHandler
	// embed all handler interfaces
}

func (i *interactor) NewAuthClient() intercomm.AuthClient {
	return intercomm.NewAuthClient()
}


func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	return appHandler
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	return mongodb.NewUserRepository(i.UserCol)
}

func (i *interactor) NewAccountActivationRepository() repository.AccountActivationRepository {
	return mongodb.NewAccountActivationRepository(i.AccCol)
}

func (i *interactor) NewAccountResetPasswordRepository() repository.ResetPasswordRepository {
	return mongodb.NewResetPasswordRepository(i.ResPwdCol)
}

func (i *interactor) NewUserService() service_contracts.UserService {
	return service.NewAuthService(i.NewUserRepository(), i.NewAccountActivationService(), i.NewAuthClient(),i.NewResetPasswordService())
}

func (i *interactor) NewAccountActivationService() service_contracts.AccountActivationService {
	return service.NewAccountActivationService(i.NewAccountActivationRepository())
}

func (i *interactor) NewResetPasswordService() service_contracts.ResetPasswordService {
	return service.NewResetPasswordService(i.NewAccountResetPasswordRepository())
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserService())
}