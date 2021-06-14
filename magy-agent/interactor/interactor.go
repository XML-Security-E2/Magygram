package interactor

import (
	"gorm.io/gorm"
	"magyAgent/controller/handler"
	"magyAgent/domain/repository"
	"magyAgent/domain/service-contracts"
	"magyAgent/infrastructure/persistence/pgsql"
	"magyAgent/service"
)

type Interactor interface {
	NewUserRepository() repository.UserRepository
	NewProductRepository() repository.ProductRepository
	NewLoginEventRepository() repository.LoginEventRepository
	NewAccountActivationRepository() repository.AccountActivationRepository
	NewAuthService() service_contracts.AuthService
	NewProductService() service_contracts.ProductService
	NewAccountActivationService() service_contracts.AccountActivationService
	NewAuthHandler() handler.AuthHandler
	NewProductHandler() handler.ProductHandler
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
	handler.ProductHandler
	// embed all handler interfaces
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.AuthHandler = i.NewAuthHandler()
	appHandler.ProductHandler = i.NewProductHandler()

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

func (i *interactor) NewProductHandler() handler.ProductHandler {
	return handler.NewProductHandler(i.NewProductService())
}

func (i *interactor) NewProductRepository() repository.ProductRepository {
	return pgsql.NewProductRepository(i.Conn)
}

func (i *interactor) NewProductService() service_contracts.ProductService {
	return service.NewProductService(i.NewProductRepository())
}