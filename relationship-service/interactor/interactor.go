package interactor

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"relationship-service/controller/handler"
	"relationship-service/infrastructure/persistence/neo4jdb"
	"relationship-service/service"
	"relationship-service/service/intercomm"
)

type Interactor interface {
	NewFollowHandler() handler.FollowHandler
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	Driver neo4j.Driver
}

func NewInteractor(driver neo4j.Driver) Interactor {
	return &interactor{Driver: driver}
}

type appHandler struct {
	handler.FollowHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.FollowHandler = i.NewFollowHandler()
	return appHandler
}

func (i *interactor) NewFollowHandler() handler.FollowHandler {
	return handler.NewFollowHandler(service.NewFollowService(neo4jdb.NewFollowRepository(i.Driver), intercomm.NewUserClient()))
}