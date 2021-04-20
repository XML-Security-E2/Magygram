package interactor

import (
	"gorm.io/gorm"
	"magyAgent/controller/handler"
	"magyAgent/domain/repository"
	service_contracts "magyAgent/domain/service-contracts"
	"magyAgent/infrastructure/persistence/pgsql"
	"magyAgent/service"
)

type Interactor interface {
	NewUserRepository() repository.UserRepository
	NewAuthService() service_contracts.AuthService
	NewAuthHandler() handler.AuthHandler
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	Conn *gorm.DB
}

func NewInteractor(Conn *gorm.DB) Interactor {
	return &interactor{Conn}
}

type appHandler struct {
	handler.AuthHandler
	// embed all handler interfaces
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.AuthHandler = i.NewAuthHandler()
	return appHandler
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	return pgsql.NewUserRepository(i.Conn)
}

func (i *interactor) NewAuthService() service_contracts.AuthService {
	return service.NewAuthService(i.NewUserRepository())
}

func (i *interactor) NewAuthHandler() handler.AuthHandler {
	return handler.NewAuthHandler(i.NewAuthService())
}