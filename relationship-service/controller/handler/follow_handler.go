package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"relationship-service/domain/model"
	"relationship-service/service"
)

type FollowHandler interface {
	FollowRequest(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	ReturnFollowedUsers(ctx echo.Context) error
	ReturnFollowRequests(ctx echo.Context) error
	AcceptFollowRequest(ctx echo.Context) error
}

type followHandler struct {
	FollowService service.FollowService
}

func NewFollowHandler(f service.FollowService) FollowHandler {
	return &followHandler{f}
}

func (f *followHandler) FollowRequest(c echo.Context) error {
	followRequest := &model.FollowRequest{}
	if err := c.Bind(followRequest); err != nil {
		return err
	}

	followId, err := f.FollowService.FollowRequest(followRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, followId)
}

func (f *followHandler) AcceptFollowRequest(ctx echo.Context) error {
	followRequest := &model.FollowRequest{}
	if err := ctx.Bind(followRequest); err != nil {
		return err
	}

	err := f.FollowService.AcceptFollowRequest(followRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, true)
}

func (f *followHandler) CreateUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	if err := f.FollowService.CreateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, true)
}

func (f *followHandler) ReturnFollowedUsers(ctx echo.Context) error {
	user := &model.User{}
	if err := ctx.Bind(user); err != nil {
		return err
	}

	result, err := f.FollowService.ReturnFollowedUsers(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (f *followHandler) ReturnFollowRequests(ctx echo.Context) error {
	user := &model.User{}
	if err := ctx.Bind(user); err != nil {
		return err
	}

	result, err := f.FollowService.ReturnFollowRequests(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, result)
}