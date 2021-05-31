package interactor

import (
	"go.mongodb.org/mongo-driver/mongo"
	"post-service/controller/handler"
	"post-service/domain/repository"
	"post-service/domain/service-contracts"
	"post-service/infrastracture/persistance/mongodb"
	"post-service/service"
	"post-service/service/intercomm"
)

type Interactor interface {
	NewPostRepository() repository.PostRepository
	NewPostService() service_contracts.PostService
	NewPostHandler() handler.PostHandler
	NewMediaClient() intercomm.MediaClient
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	PostCol *mongo.Collection
}



func NewInteractor(PostCol *mongo.Collection) Interactor {
	return &interactor{PostCol}
}

type appHandler struct {
	handler.PostHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.PostHandler = i.NewPostHandler()
	return appHandler
}

func (i *interactor) NewMediaClient() intercomm.MediaClient {
	return intercomm.NewMediaClient()
}

func (i *interactor) NewPostRepository() repository.PostRepository {
	return mongodb.NewPostRepository(i.PostCol)
}

func (i *interactor) NewPostService() service_contracts.PostService {
	return service.NewPostService(i.NewPostRepository(), i.NewMediaClient())
}

func (i *interactor) NewPostHandler() handler.PostHandler {
	return handler.NewPostHandler(i.NewPostService())
}