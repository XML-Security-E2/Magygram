package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"media-service/service"
	"mime/multipart"
	"net/http"
)

type MediaHandler interface {
	SaveMedia(c echo.Context) error
	GetMedia(c echo.Context) error
}

type mediaHandler struct {
	MediaService service.MediaService
}

func NewMediaHandler(m service.MediaService) MediaHandler {
	return &mediaHandler{m}
}

func (m mediaHandler) SaveMedia(c echo.Context) error {
	fmt.Println("USAO")

	mpf, _ := c.MultipartForm()

	var headers []*multipart.FileHeader
	for k, v := range mpf.File {
		fmt.Printf("Key[%s] ", k)
		headers = append(headers, v[0])
	}

	ctx := c.Request().Context()

	mediaIds, err := m.MediaService.SaveMedia(ctx, headers)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, mediaIds)
}

func (m mediaHandler) GetMedia(c echo.Context) error {
	panic("implement me")
}
