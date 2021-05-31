package interactor

import (
	"media-service/controller/handler"
	"media-service/service"
)

type Interactor interface {
	NewMediaService() service.MediaService
	NewMediaHandler() handler.MediaHandler
	NewAppHandler() handler.AppHandler
}

type interactor struct {
}

func NewInteractor() Interactor {
	return &interactor{}
}

type appHandler struct {
	handler.MediaHandler
	// embed all handler interfaces
}

func (i *interactor) NewMediaHandler() handler.MediaHandler {
	return handler.NewMediaHandler(i.NewMediaService())
}


func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.MediaHandler = i.NewMediaHandler()
	return appHandler
}

func (i *interactor) NewMediaService() service.MediaService {
	return service.NewMediaService()
}