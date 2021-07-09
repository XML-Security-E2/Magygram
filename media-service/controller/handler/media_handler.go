package handler

import (
	"context"
	"fmt"
	"io"
	"media-service/logger"
	"media-service/service"
	"media-service/tracer"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MediaHandler interface {
	SaveMedia(c echo.Context) error
}

type mediaHandler struct {
	MediaService service.MediaService
	tracer       opentracing.Tracer
	closer       io.Closer
}

func NewMediaHandler(m service.MediaService) MediaHandler {
	tracer, closer := tracer.Init("media-service")
	opentracing.SetGlobalTracer(tracer)
	return &mediaHandler{m, tracer, closer}
}

func (m mediaHandler) SaveMedia(c echo.Context) error {
	span := tracer.StartSpanFromRequest("MediaHandlerSaveMedia", m.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling save media at %s\n", c.Path())),
	)

	logger.LoggingEntry = logger.Logger.WithFields(logrus.Fields{"request_ip": c.RealIP()})

	mpf, _ := c.MultipartForm()

	var headers []*multipart.FileHeader
	for _, v := range mpf.File {
		headers = append(headers, v[0])
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	media, err := m.MediaService.SaveMedia(ctx, headers)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, media)
}
