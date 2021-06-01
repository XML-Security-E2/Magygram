package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"user-service/domain/model"
	"user-service/domain/service-contracts"
)

type CollectionsHandler interface {
	CreateCollection(c echo.Context) error
	AddPostToCollection(c echo.Context) error
}

type collectionsHandler struct {
	CollectionsService service_contracts.CollectionsService
}

func NewCollectionsHandler(u service_contracts.CollectionsService) CollectionsHandler {
	return &collectionsHandler{u}
}

func (ch collectionsHandler) CreateCollection(c echo.Context) error {

	collectionName := c.FormValue("name")
	ctx := c.Request().Context()
	bearer := c.Request().Header.Get("Authorization")

	err := ch.CollectionsService.CreateCollection(ctx, bearer, collectionName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}

func (ch collectionsHandler) AddPostToCollection(c echo.Context) error {

	postCollectionRequest := &model.FavouritePostRequest{}
	if err := c.Bind(postCollectionRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	bearer := c.Request().Header.Get("Authorization")

	err := ch.CollectionsService.AddPostToCollection(ctx,bearer, postCollectionRequest)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}