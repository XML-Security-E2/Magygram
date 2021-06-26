package interactor

import (
	"github.com/go-redis/redis/v8"
	"message-service/controller/handler"
	"message-service/controller/hub"
	"message-service/domain/repository"
	"message-service/domain/service-contracts"
	"message-service/infrastructure/persistance/redisdb"
	"message-service/service"
	"message-service/service/intercomm"
)

type Interactor interface {
	NewNotificationRepository() repository.NotificationRepository
	NewConversationRepository() repository.ConversationRepository
	NewNotificationService() service_contracts.NotificationService
	NewConversationService() service_contracts.ConversationService
	NewNotificationHandler() handler.NotificationHandler
	NewConversationHandler() handler.ConversationHandler
	NewAuthClient() intercomm.AuthClient
	NewUserClient() intercomm.UserClient
	NewMediaClient() intercomm.MediaClient
	NewRelationshipClient() intercomm.RelationshipClient
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	Db *redis.Client
	Hub *hub.NotifyHub
	MHub *hub.MessageHub
	MNHub *hub.MessageNotificationsHub
}

func NewInteractor(db *redis.Client, hub *hub.NotifyHub, mhub *hub.MessageHub, mnhub *hub.MessageNotificationsHub) Interactor {
	return &interactor{db, hub, mhub, mnhub}
}

type appHandler struct {
	handler.NotificationHandler
	handler.ConversationHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.NotificationHandler = i.NewNotificationHandler()
	appHandler.ConversationHandler = i.NewConversationHandler()
	return appHandler
}

func (i *interactor) NewMediaClient() intercomm.MediaClient {
	return intercomm.NewMediaClient()
}

func (i *interactor) NewRelationshipClient() intercomm.RelationshipClient {
	return intercomm.NewRelationshipClient()
}

func (i *interactor) NewConversationRepository() repository.ConversationRepository {
	return redisdb.NewConversationRepository(i.Db)
}

func (i *interactor) NewConversationService() service_contracts.ConversationService {
	return service.NewConversationService(i.NewConversationRepository(), i.NewAuthClient(), i.NewUserClient(), i.NewMediaClient(), i.NewRelationshipClient())
}

func (i *interactor) NewConversationHandler() handler.ConversationHandler {
	return handler.NewConversationHandler(i.NewConversationService(), i.MHub, i.MNHub)
}

func (i *interactor) NewUserClient() intercomm.UserClient {
	return intercomm.NewUserClient()
}

func (i *interactor) NewAuthClient() intercomm.AuthClient {
	return intercomm.NewAuthClient()
}

func (i *interactor) NewNotificationRepository() repository.NotificationRepository {
	return redisdb.NewNotificationRepository(i.Db)
}

func (i *interactor) NewNotificationService() service_contracts.NotificationService {
	return service.NewNotificationService(i.NewNotificationRepository(), i.NewAuthClient(), i.NewUserClient())
}

func (i *interactor) NewNotificationHandler() handler.NotificationHandler {
	return handler.NewNotificationHandler(i.NewNotificationService(), i.Hub)
}