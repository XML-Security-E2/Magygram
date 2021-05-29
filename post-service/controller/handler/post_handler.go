package handler

import (
	"context"
	"github.com/labstack/echo"
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
	postRequest := &model.PostRequest{}
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


