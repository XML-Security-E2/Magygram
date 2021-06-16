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
	NewNotificationService() service_contracts.NotificationService
	NewNotificationHandler() handler.NotificationHandler
	NewAuthClient() intercomm.AuthClient
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	Db *redis.Client
	Hub *hub.NotifyHub
}

func NewInteractor(db *redis.Client, hub *hub.NotifyHub) Interactor {
	return &interactor{db, hub}
}

type appHandler struct {
	handler.NotificationHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.NotificationHandler = i.NewNotificationHandler()
	return appHandler
}

func (i *interactor) NewAuthClient() intercomm.AuthClient {
	return intercomm.NewAuthClient()
}


func (i *interactor) NewNotificationRepository() repository.NotificationRepository {
	return redisdb.NewNotificationRepository(i.Db)
}

func (i *interactor) NewNotificationService() service_contracts.NotificationService {
	return service.NewNotificationService(i.NewNotificationRepository(), i.NewAuthClient())
}

func (i *interactor) NewNotificationHandler() handler.NotificationHandler {
	return handler.NewNotificationHandler(i.NewNotificationService(), i.Hub, i.NewAuthClient())
}