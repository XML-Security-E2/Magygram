package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"user-service/domain/model"
	"user-service/domain/service-contracts"
	"user-service/logger"
)

type CollectionsHandler interface {
	CreateCollection(c echo.Context) error
	AddPostToCollection(c echo.Context) error
	GetUsersCollections(c echo.Context) error
	GetUsersCollectionsExceptDefault(c echo.Context) error
	CheckIfPostInFavourites(c echo.Context) error
	DeleteFromCollection(c echo.Context) error
	GetUserCollections(c echo.Context) error
	GetCollectionPosts(c echo.Context) error
	CollectionsLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type collectionsHandler struct {
	CollectionsService service_contracts.CollectionsService
}

func NewCollectionsHandler(u service_contracts.CollectionsService) CollectionsHandler {
	return &collectionsHandler{u}
}


func (h *collectionsHandler) CollectionsLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.LoggingEntry = logger.Logger.WithFields(logrus.Fields{"request_ip": c.RealIP()})
		return next(c)
	}
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

func (ch collectionsHandler) GetUsersCollections(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	collection, err := ch.CollectionsService.GetUsersCollections(ctx,bearer, "")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, collection)
}

func (ch collectionsHandler) GetUsersCollectionsExceptDefault(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	collection, err := ch.CollectionsService.GetUsersCollections(ctx,bearer, model.DefaultCollection)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, collection)
}

func (ch collectionsHandler) CheckIfPostInFavourites(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	postIds := &[]string{}

	if err := c.Bind(postIds); err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	postsFavFlags, err := ch.CollectionsService.CheckIfPostsInFavourites(ctx,bearer, postIds)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, postsFavFlags)
}

func (ch collectionsHandler) DeleteFromCollection(c echo.Context) error {

	postId := c.Param("postId")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")

	err := ch.CollectionsService.DeletePostFromCollections(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (ch collectionsHandler) GetUserCollections(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	collection, err := ch.CollectionsService.GetUsersCollections(ctx,bearer, "")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, collection)
}

func (ch collectionsHandler) GetCollectionPosts(c echo.Context) error {
	collectionName := c.Param("collectionName")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")

	posts, err := ch.CollectionsService.GetCollectionPosts(ctx, bearer, collectionName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}