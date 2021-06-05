package interactor

import (
	"go.mongodb.org/mongo-driver/mongo"
	"story-service/controller/handler"
	"story-service/domain/repository"
	service_contracts "story-service/domain/service-contracts"
	"story-service/infrastructure/persistence/mongodb"
	"story-service/service"
	"story-service/service/intercomm"
)

type Interactor interface {
	NewStoryRepository() repository.StoryRepository
	NewStoryService() service_contracts.StoryService
	NewStoryHandler() handler.StoryHandler
	NewMediaClient() intercomm.MediaClient
	NewUserClient() intercomm.UserClient
	NewAuthClient() intercomm.AuthClient
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	StoryCol *mongo.Collection
}


func NewInteractor(StoryCol *mongo.Collection) Interactor {
	return &interactor{StoryCol}
}

type appHandler struct {
	handler.StoryHandler
}

func (i *interactor) NewAuthClient() intercomm.AuthClient {
	return intercomm.NewAuthClient()
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.StoryHandler = i.NewStoryHandler()
	return appHandler
}

func (i *interactor) NewMediaClient() intercomm.MediaClient {
	return intercomm.NewMediaClient()
}

func (i *interactor) NewUserClient() intercomm.UserClient {
	return intercomm.NewUserClient()
}

func (i *interactor) NewStoryRepository() repository.StoryRepository {
	return mongodb.NewStoryRepository(i.StoryCol)
}

func (i *interactor) NewStoryService() service_contracts.StoryService {
	return service.NewStoryService(i.NewStoryRepository(), i.NewMediaClient(), i.NewUserClient(), i.NewAuthClient())
}

func (i *interactor) NewStoryHandler() handler.StoryHandler {
	return handler.NewStoryHandler(i.NewStoryService())
}