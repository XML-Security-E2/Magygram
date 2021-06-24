package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
	"os"
	"user-service/domain/model"
	"user-service/domain/service-contracts"
	"user-service/domain/service-contracts/exceptions"
	"user-service/logger"
)


type UserHandler interface {
	RegisterUser(c echo.Context) error
	EditUser(c echo.Context) error
	EditUserImage(c echo.Context) error
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
	MuteUser(c echo.Context) error
	UnmuteUser(c echo.Context) error
	BlockUser(c echo.Context) error
	UnblockUser(c echo.Context) error
	SearchForUsersByUsernameByGuest(c echo.Context) error
	IsUserPrivate(c echo.Context) error
	GetFollowRequests(c echo.Context) error
	AcceptFollowRequest(c echo.Context) error
	UserLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	UpdateLikedPost(c echo.Context) error
	UpdateDislikedPost(c echo.Context) error
	GetUserLikedPost(c echo.Context) error
	GetUserDislikedPost(c echo.Context) error
	GetUsersForPostNotification(c echo.Context) error
	GetUsersForStoryNotification(c echo.Context) error
	CheckIfPostInteractionNotificationEnabled(c echo.Context) error
	EditUsersNotifications(c echo.Context) error
	EditUsersPrivacySettings(c echo.Context) error
	VerifyUser(c echo.Context) error
	CheckIfUserVerified(c echo.Context) error
	GetUsersNotificationsSettings(c echo.Context) error
	ChangeUsersNotificationsSettings(c echo.Context) error
	DeleteUser(c echo.Context) error
	GetFollowRecommendation(c echo.Context) error
	RegisterAgent(c echo.Context) error
	AddComment(c echo.Context) error
	CheckIfUserVerifiedById(c echo.Context) error
	GetUsersInfo(c echo.Context) error
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



func (h *userHandler) GetUsersInfo(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserService.GetUsersInfo(ctx, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) GetUsersNotificationsSettings(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	settings, err := h.UserService.GetUsersNotificationsSettings(ctx, bearer, userId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, settings)}

func (h userHandler) DeleteUser(c echo.Context) error {
	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := h.UserService.DeleteUser(ctx, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) ChangeUsersNotificationsSettings(c echo.Context) error {
	userId := c.Param("userId")
	setReq := &model.SettingsRequest{}
	if err := c.Bind(setReq); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	err := h.UserService.ChangeUsersNotificationsSettings(ctx, bearer, setReq, userId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) GetUsersForPostNotification(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	users, err := h.UserService.GetUsersForPostNotification(ctx, userId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK,users)
}

func (h *userHandler) GetUsersForStoryNotification(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	users, err := h.UserService.GetUsersForStoryNotification(ctx, userId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK,users)
}

func (h *userHandler) CheckIfPostInteractionNotificationEnabled(c echo.Context) error {
	userId := c.Param("userId")
	fromId := c.Param("fromId")

	interactionType := c.Param("interactionType")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	check, err := h.UserService.CheckIfPostInteractionNotificationEnabled(ctx, userId, fromId, interactionType)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, check)
}

func (h *userHandler) UserLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.LoggingEntry = logger.Logger.WithFields(logrus.Fields{"request_ip": c.RealIP()})
		return next(c)
	}
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

func (h *userHandler) EditUsersNotifications(c echo.Context) error {
	notificationReq := &model.NotificationSettingsUpdateReq{}
	if err := c.Bind(notificationReq); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	err := h.UserService.EditUsersNotifications(ctx, bearer, notificationReq)
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) EditUsersPrivacySettings(c echo.Context) error {
	privacySettingsReq := &model.PrivacySettings{}
	if err := c.Bind(privacySettingsReq); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	err := h.UserService.EditUsersPrivacySettings(ctx, bearer, privacySettingsReq)
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) EditUserImage(c echo.Context) error {
	userId := c.Param("userId")

	mpf, _ := c.MultipartForm()
	var headers []*multipart.FileHeader
	for _, v := range mpf.File {
		headers = append(headers, v[0])
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	url, err := h.UserService.EditUserImage(ctx, bearer, userId, headers)
	if err != nil{
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *exceptions.UnauthorizedAccessError:
			return echo.NewHTTPError(http.StatusUnauthorized, t.Error())
		}
	}
	return c.JSON(http.StatusOK, url)
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

	resp, err := h.UserService.RegisterUser(ctx, userRequest)

	if err != nil {
		fmt.Println(err)

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Stream(http.StatusCreated,"image/png",resp.Body)
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

	if os.Getenv("IS_PRODUCTION") == "true" {
		return c.Redirect(http.StatusMovedPermanently, "http://localhost:3000/#/login")//c.JSON(http.StatusNoContent, activationId)
	} else {
		return c.Redirect(http.StatusMovedPermanently, "https://localhost:3000/#/login")//c.JSON(http.StatusNoContent, activationId)
	}
	//return c.Redirect(http.StatusMovedPermanently, "https://localhost:3000/#/login")//c.JSON(http.StatusNoContent, activationId)
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

	fmt.Println("Profile 1")
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
	followRequest, err := h.UserService.FollowUser(ctx, bearer, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	if followRequest {
		return c.JSON(http.StatusCreated, "")
	}
	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) MuteUser(c echo.Context) error {
	userId := c.FormValue("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	err := h.UserService.MuteUser(ctx, bearer, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) UnmuteUser(c echo.Context) error {
	userId := c.FormValue("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	err := h.UserService.UnmuteUser(ctx, bearer, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) BlockUser(c echo.Context) error {
	userId := c.FormValue("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	err := h.UserService.BlockUser(ctx, bearer, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) UnblockUser(c echo.Context) error {
	userId := c.FormValue("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	err := h.UserService.UnblockUser(ctx, bearer, userId)

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

func (h *userHandler) GetFollowRequests(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	reqs, err := h.UserService.GetFollowRequests(ctx, bearer)
	if err != nil{
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *exceptions.UnauthorizedAccessError:
			return echo.NewHTTPError(http.StatusUnauthorized, t.Error())
		}
	}

	return c.JSON(http.StatusOK, reqs)
}

func (h *userHandler) AcceptFollowRequest(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")

	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	err := h.UserService.AcceptFollowRequest(ctx, bearer, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) UpdateLikedPost(c echo.Context) error {
	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	err := h.UserService.UpdateLikedPost(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) UpdateDislikedPost(c echo.Context) error {
	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	err := h.UserService.UpdateDislikedPost(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) GetUserLikedPost(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	retVal, err := h.UserService.GetUserLikedPost(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (h *userHandler) GetUserDislikedPost(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	retVal, err := h.UserService.GetUserDislikedPost(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (h *userHandler) VerifyUser(c echo.Context) error {
	verifyAccountDTO := &model.VerifyAccountDTO{}
	if err := c.Bind(verifyAccountDTO); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := h.UserService.VerifyUser(ctx, verifyAccountDTO)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) CheckIfUserVerified(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	result,err := h.UserService.CheckIfUserVerified(ctx,bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}


func (h *userHandler) GetFollowRecommendation(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	result,err := h.UserService.GetFollowRecommendation(ctx,bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *userHandler) RegisterAgent(c echo.Context) error {
	agentRegistrationDTO := &model.AgentRegistrationDTO{}
	if err := c.Bind(agentRegistrationDTO); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result,err := h.UserService.RegisterAgent(ctx,agentRegistrationDTO)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *userHandler) AddComment(c echo.Context) error {
	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	err := h.UserService.AddComment(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) CheckIfUserVerifiedById(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result,err := h.UserService.CheckIfUserVerifiedById(ctx,userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}