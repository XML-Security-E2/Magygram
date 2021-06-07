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
	NewAuthClient() intercomm.AuthClient
	NewUserClient() intercomm.UserClient
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	PostCol *mongo.Collection
	locationCol *mongo.Collection
	tagCol *mongo.Collection
}

func NewInteractor(PostCol *mongo.Collection, locationCol *mongo.Collection, tagCol *mongo.Collection) Interactor {
	return &interactor{PostCol, locationCol, tagCol}
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

func (i *interactor) NewAuthClient() intercomm.AuthClient {
	return intercomm.NewAuthClient()
}

func (i *interactor) NewUserClient() intercomm.UserClient {
	return intercomm.NewUserClient()
}

func (i *interactor) NewRelationshipClient() intercomm.RelationshipClient {
	return intercomm.NewRelationshipClient()
}

func (i *interactor) NewPostRepository() repository.PostRepository {
	return mongodb.NewPostRepository(i.PostCol, i.locationCol, i.tagCol)
}

func (i *interactor) NewPostService() service_contracts.PostService {
	return service.NewPostService(i.NewPostRepository(), i.NewMediaClient(), i.NewUserClient(), i.NewRelationshipClient(), i.NewAuthClient())
}

func (i *interactor) NewPostHandler() handler.PostHandler {
	return handler.NewPostHandler(i.NewPostService())
}