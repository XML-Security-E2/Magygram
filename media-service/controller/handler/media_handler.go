package handler

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"media-service/logger"
	"media-service/service"
	"mime/multipart"
	"net/http"
)

type MediaHandler interface {
	SaveMedia(c echo.Context) error
}

type mediaHandler struct {
	MediaService service.MediaService
}

func NewMediaHandler(m service.MediaService) MediaHandler {
	return &mediaHandler{m}
}

func (m mediaHandler) SaveMedia(c echo.Context) error {
	logger.LoggingEntry = logger.Logger.WithFields(logrus.Fields{"request_ip": c.RealIP()})

	mpf, _ := c.MultipartForm()

	var headers []*multipart.FileHeader
	for _, v := range mpf.File {
		headers = append(headers, v[0])
	}

	ctx := c.Request().Context()

	media, err := m.MediaService.SaveMedia(ctx, headers)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, media)
}