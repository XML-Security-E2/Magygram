package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"mime/multipart"
	"net/http"
	"post-service/domain/model"
	"post-service/domain/service-contracts"
)


type PostHandler interface {
	CreatePost(c echo.Context) error

}

type postHandler struct {
	PostService service_contracts.PostService
}

func NewPostHandler(p service_contracts.PostService) PostHandler {
	return &postHandler{p}
}

func (p postHandler) CreatePost(c echo.Context) error {

	location := c.FormValue("location")
	description := c.FormValue("description")
	tags := c.FormValue("tags")

	fmt.Println(location)
	mpf, _ := c.MultipartForm()
	var headers []*multipart.FileHeader
	for _, v := range mpf.File {
		headers = append(headers, v[0])
	}

	fmt.Println(len(headers))
	postRequest := &model.PostRequest{
		Description: description,
		Location:    location,
		Media:       headers,
		Tags:        []string{tags},
	}

	if err := c.Bind(postRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	postId, err := p.PostService.CreatePost(ctx, postRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, postId)
}


