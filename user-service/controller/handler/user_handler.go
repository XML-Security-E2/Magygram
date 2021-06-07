package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"user-service/domain/model"
	"user-service/domain/service-contracts"
	"user-service/domain/service-contracts/exceptions"
)


type UserHandler interface {
	RegisterUser(c echo.Context) error
	EditUser(c echo.Context) error
	ActivateUser(c echo.Context) error
	ResetPasswordRequest(c echo.Context) error
	ResetPasswordActivation(c echo.Context) error
	ChangeNewPassword(c echo.Context) error
	ResendActivationLink(c echo.Context) error
	GetUserEmailIfUserExist(c echo.Context) error
	GetUserById(c echo.Context) error
	GetLoggedUserInfo(c echo.Context) error
	SearchForUsersByUsername(c echo.Context) error
	GetUserProfileById(c echo.Context) error
	GetFollowedUsers(c echo.Context) error
	GetFollowingUsers(c echo.Context) error
	FollowUser(c echo.Context) error
	UnollowUser(c echo.Context) error
	SearchForUsersByUsernameByGuest(c echo.Context) error
	IsUserPrivate(c echo.Context) error
}

var (
	ErrWrongCredentials = echo.NewHTTPError(http.StatusUnauthorized, "username or password is invalid")
)
type userHandler struct {
	UserService service_contracts.UserService
}


func NewUserHandler(u service_contracts.UserService) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) EditUser(c echo.Context) error {
	userId := c.Param("userId")
	userRequest := &model.EditUserRequest{}
	if err := c.Bind(userRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	updatedId, err := h.UserService.EditUser(ctx, bearer, userId, userRequest)
	if err != nil{
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *exceptions.UnauthorizedAccessError:
			return echo.NewHTTPError(http.StatusUnauthorized, t.Error())
		}
	}
	return c.JSON(http.StatusOK, updatedId)
}
func (h *userHandler) RegisterUser(c echo.Context) error {
	userRequest := &model.UserRequest{}
	if err := c.Bind(userRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userId, err := h.UserService.RegisterUser(ctx, userRequest)
	fmt.Println(userId)
	if err != nil {
		fmt.Println(err)

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userId)
}

func (h *userHandler) ActivateUser(c echo.Context) error {
	activationId := c.Param("activationId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	activated, err := h.UserService.ActivateUser(ctx, activationId)
	if err != nil || activated == false{
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not be activated.")
	}

	return c.Redirect(http.StatusMovedPermanently, "https://localhost:3000/#/login")//c.JSON(http.StatusNoContent, activationId)
}

func (h *userHandler) ResendActivationLink(c echo.Context) error {

	activateLinkRequest := &model.ActivateLinkRequest{}
	if err := c.Bind(activateLinkRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	_, err := h.UserService.ResendActivationLink(ctx, activateLinkRequest)

	if err != nil {
		return ErrWrongCredentials
	}

	return c.JSON(http.StatusOK, "Activation link has been send")
}

func (h *userHandler) ResetPasswordRequest(c echo.Context) error {
	resetPasswordRequest := &model.ResetPasswordRequest{}
	if err := c.Bind(resetPasswordRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()

	_, err := h.UserService.ResetPassword(ctx, resetPasswordRequest.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "Email has been send")
}

func (h *userHandler) ResetPasswordActivation(c echo.Context) error {

	resetPasswordId := c.Param("resetPasswordId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	activated, err := h.UserService.ResetPasswordActivation(ctx, resetPasswordId)
	if err != nil || activated == false{
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not reset password.")
	}

	return c.Redirect(http.StatusMovedPermanently, "https://localhost:3000/#/reset-password/" + resetPasswordId)//c.JSON(http.StatusNoContent, activationId)
}

func (h *userHandler) ChangeNewPassword(c echo.Context) error {
	changeNewPasswordRequest := &model.ChangeNewPasswordRequest{}
	if err := c.Bind(changeNewPasswordRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()

	successful, err := h.UserService.ChangeNewPassword(ctx, changeNewPasswordRequest)

	if err != nil || !successful {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Password has been changed")
}

func (h *userHandler) GetUserEmailIfUserExist(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserService.GetUserEmailIfUserExist(ctx, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"emailAddress": user.Email,
	})
}

func (h *userHandler) GetUserById(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserService.GetUserById(ctx, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	c.Response().Header().Set("Content-Type" , "text/javascript")
	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) SearchForUsersByUsername(c echo.Context) error {
	username := c.Param("username")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	users, err := h.UserService.SearchForUsersByUsername(ctx, username, bearer)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	c.Response().Header().Set("Content-Type" , "text/javascript")
	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) SearchForUsersByUsernameByGuest(c echo.Context) error {
	username := c.Param("username")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	users, err := h.UserService.SearchForUsersByUsernameByGuest(ctx, username)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	c.Response().Header().Set("Content-Type" , "text/javascript")
	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetLoggedUserInfo(c echo.Context) error {
	ctx := c.Request().Context()
	bearer := c.Request().Header.Get("Authorization")

	if ctx == nil {
		ctx = context.Background()
	}
	userInfo, err := h.UserService.GetLoggedUserInfo(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	return c.JSON(http.StatusOK, userInfo)
}

func (h *userHandler) GetUserProfileById(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	user, err := h.UserService.GetUserProfileById(ctx,bearer, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK,user)
}

func (h *userHandler) IsUserPrivate(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserService.GetUserById(ctx, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, user.IsPrivate)
}

func (h *userHandler) GetFollowedUsers(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	users, err := h.UserService.GetFollowedUsers(ctx, bearer, userId)
	if err != nil{
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *exceptions.UnauthorizedAccessError:
			return echo.NewHTTPError(http.StatusUnauthorized, t.Error())
		}
	}

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetFollowingUsers(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	users, err := h.UserService.GetFollowingUsers(ctx, bearer, userId)
	if err != nil{
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *exceptions.UnauthorizedAccessError:
			return echo.NewHTTPError(http.StatusUnauthorized, t.Error())
		}
	}

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) FollowUser(c echo.Context) error {
	userId := c.FormValue("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	err := h.UserService.FollowUser(ctx, bearer, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) UnollowUser(c echo.Context) error {
	userId := c.FormValue("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	err := h.UserService.UnfollowUser(ctx, bearer, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, "")
}