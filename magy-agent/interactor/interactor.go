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
	NewLoginEventRepository() repository.LoginEventRepository
	NewAccountActivationRepository() repository.AccountActivationRepository
	NewAuthService() service_contracts.AuthService
	NewAccountActivationService() service_contracts.AccountActivationService
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

func (i *interactor) NewLoginEventRepository() repository.LoginEventRepository {
return pgsql.NewLoginEventRepository(i.Conn)
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	return pgsql.NewUserRepository(i.Conn)
}

func (i *interactor) NewAccountActivationRepository() repository.AccountActivationRepository {
	return pgsql.NewAccountActivationRepository(i.Conn)
}

func (i *interactor) NewAccountResetPasswordRepository() repository.AccountResetPasswordRepository {
	return pgsql.NewAccountResetPasswordRepository(i.Conn)
}

func (i *interactor) NewAuthService() service_contracts.AuthService {
	return service.NewAuthService(i.NewUserRepository(), i.NewAccountActivationService(), i.NewLoginEventRepository(),i.NewAccountResetPasswordService())
}

func (i *interactor) NewAccountActivationService() service_contracts.AccountActivationService {
	return service.NewAccountActivationService(i.NewAccountActivationRepository())
}

func (i *interactor) NewAccountResetPasswordService() service_contracts.AccountResetPasswordService {
	return service.NewAccountResetPasswordService(i.NewAccountResetPasswordRepository())
}

func (i *interactor) NewAuthHandler() handler.AuthHandler {
	return handler.NewAuthHandler(i.NewAuthService())
}