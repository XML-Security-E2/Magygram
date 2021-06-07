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
	NewCollectionsService() service_contracts.CollectionsService
	NewUserService() service_contracts.UserService
	NewAccountActivationService() service_contracts.AccountActivationService
	NewAuthClient() intercomm.AuthClient
	NewStoryClient() intercomm.StoryClient
	NewHighlightsHandler() handler.HighlightsHandler
	NewHighlightsService() service_contracts.HighlightsService
	NewPostClient() intercomm.PostClient
	NewMediaClient() intercomm.MediaClient
	NewUserHandler() handler.UserHandler
	NewCollectionsHandler() handler.CollectionsHandler
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
	handler.CollectionsHandler
	handler.HighlightsHandler
	// embed all handler interfaces
}

func (i *interactor) NewAuthClient() intercomm.AuthClient {
	return intercomm.NewAuthClient()
}

func (i *interactor) NewMediaClient() intercomm.MediaClient {
	return intercomm.NewMediaClient()
}

func (i *interactor) NewRelationshipClient() intercomm.RelationshipClient {
	return intercomm.NewRelationshipClient()
}

func (i *interactor) NewPostClient() intercomm.PostClient {
	return intercomm.NewPostClient()
}

func (i *interactor) NewStoryClient() intercomm.StoryClient {
	return intercomm.NewStoryClient()
}
func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	appHandler.CollectionsHandler = i.NewCollectionsHandler()
	appHandler.HighlightsHandler = i.NewHighlightsHandler()
	return appHandler
}

func (i *interactor) NewCollectionsHandler() handler.CollectionsHandler {
	return handler.NewCollectionsHandler(i.NewCollectionsService())
}

func (i *interactor) NewHighlightsHandler() handler.HighlightsHandler {
	return handler.NewHighlightsHandler(i.NewHighlightsService())
}

func (i *interactor) NewHighlightsService() service_contracts.HighlightsService {
	return service.NewHighlightsService(i.NewUserRepository(), i.NewAuthClient(), i.NewStoryClient(), i.NewRelationshipClient())
}

func (i *interactor) NewCollectionsService() service_contracts.CollectionsService {
	return service.NewCollectionsService(i.NewUserRepository(), i.NewAuthClient(), i.NewPostClient())
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
	return service.NewAuthService(i.NewUserRepository(), i.NewAccountActivationService(), i.NewAuthClient(),i.NewResetPasswordService(), i.NewRelationshipClient(), i.NewPostClient(), i.NewMediaClient())
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