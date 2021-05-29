package interactor

import (
	"auth-service/controller/handler"
	"auth-service/domain/repository"
	"auth-service/domain/service-contracts"
	"auth-service/infrastructure/persistence/mongodb"
	"auth-service/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type Interactor interface {
	NewUserRepository() repository.UserRepository
	NewLoginEventRepository() repository.LoginEventRepository
	NewUserService() service_contracts.UserService
	NewAuthService() service_contracts.AuthService
	NewUserHandler() handler.UserHandler
	NewAuthHandler() handler.AuthHandler
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	UserCol *mongo.Collection
	LogECol *mongo.Collection
}

func NewInteractor(UserCol *mongo.Collection, LogECol *mongo.Collection) Interactor {
	return &interactor{UserCol, LogECol}
}

type appHandler struct {
	handler.UserHandler
	handler.AuthHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	appHandler.AuthHandler = i.NewAuthHandler()
	return appHandler
}

func (i *interactor) NewLoginEventRepository() repository.LoginEventRepository {
	return mongodb.NewLoginEventRepository(i.LogECol)
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	return mongodb.NewUserRepository(i.UserCol)
}



func (i *interactor) NewUserService() service_contracts.UserService {
	return service.NewUserService(i.NewUserRepository(), i.NewLoginEventRepository())
}

func (i *interactor) NewAuthService() service_contracts.AuthService {
	return service.NewAuthService(i.NewLoginEventRepository(), i.NewUserService())
}


func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserService())
}

func (i *interactor) NewAuthHandler() handler.AuthHandler {
	return handler.NewAuthHandler(i.NewAuthService())
}