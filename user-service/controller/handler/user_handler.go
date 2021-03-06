package handler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"user-service/domain/model"
	service_contracts "user-service/domain/service-contracts"
	"user-service/domain/service-contracts/exceptions"
	"user-service/logger"
	"user-service/tracer"

	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
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
	SearchForInfluencerByUsername(c echo.Context) error
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
	RegisterAgentByAdmin(c echo.Context) error
	GetLoggedAgentInfo(c echo.Context) error
	GetLoggedUserTargetGroup(c echo.Context) error
	GetLoggedAgentInfoById(c echo.Context) error
}

var (
	ErrWrongCredentials = echo.NewHTTPError(http.StatusUnauthorized, "username or password is invalid")
)

type userHandler struct {
	UserService service_contracts.UserService
	tracer      opentracing.Tracer
	closer      io.Closer
}

func NewUserHandler(u service_contracts.UserService) UserHandler {
	tracer, closer := tracer.Init("user-service")
	opentracing.SetGlobalTracer(tracer)
	return &userHandler{
		UserService: u,
		tracer:      tracer,
		closer:      closer,
	}
}

func (u *userHandler) CloseTracer() error {
	return u.closer.Close()
}

func (h *userHandler) GetUsersInfo(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetUsersInfo", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get users info at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	user, err := h.UserService.GetUsersInfo(ctx, userId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) GetUsersNotificationsSettings(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetUsersNotificationsSettings", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get users notification settings at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	settings, err := h.UserService.GetUsersNotificationsSettings(ctx, bearer, userId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, settings)
}

func (h userHandler) DeleteUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerDeleteUser", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling delete user at %s\n", c.Path())),
	)

	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	err := h.UserService.DeleteUser(ctx, postId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) ChangeUsersNotificationsSettings(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerChangeUsersNotificationsSettings", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling change users notification settings at %s\n", c.Path())),
	)

	userId := c.Param("userId")
	setReq := &model.SettingsRequest{}
	if err := c.Bind(setReq); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")
	err := h.UserService.ChangeUsersNotificationsSettings(ctx, bearer, setReq, userId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) GetUsersForPostNotification(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetUsersForPostNotification", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get users post notification at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	users, err := h.UserService.GetUsersForPostNotification(ctx, userId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetUsersForStoryNotification(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetUsersForStoryNotification", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get users for story notification at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	users, err := h.UserService.GetUsersForStoryNotification(ctx, userId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) CheckIfPostInteractionNotificationEnabled(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerCheckIfPostInteractionNotificationEnabled", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling check if post interaction notification enabled at %s\n", c.Path())),
	)

	userId := c.Param("userId")
	fromId := c.Param("fromId")

	interactionType := c.Param("interactionType")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	check, err := h.UserService.CheckIfPostInteractionNotificationEnabled(ctx, userId, fromId, interactionType)
	if err != nil {
		tracer.LogError(span, err)
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
	span := tracer.StartSpanFromRequest("UserHandlerEditUser", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling edit user at %s\n", c.Path())),
	)

	userId := c.Param("userId")
	userRequest := &model.EditUserRequest{}
	if err := c.Bind(userRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	updatedId, err := h.UserService.EditUser(ctx, bearer, userId, userRequest)
	if err != nil {
		tracer.LogError(span, err)
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
	span := tracer.StartSpanFromRequest("UserHandlerEditUsersNotifications", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling edit users notifications at %s\n", c.Path())),
	)

	notificationReq := &model.NotificationSettingsUpdateReq{}
	if err := c.Bind(notificationReq); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	err := h.UserService.EditUsersNotifications(ctx, bearer, notificationReq)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) EditUsersPrivacySettings(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerEditUsersPrivacySettings", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling edit users privacy settings at %s\n", c.Path())),
	)

	privacySettingsReq := &model.PrivacySettings{}
	if err := c.Bind(privacySettingsReq); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	err := h.UserService.EditUsersPrivacySettings(ctx, bearer, privacySettingsReq)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) EditUserImage(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerEditUserImage", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling edit user image at %s\n", c.Path())),
	)

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
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	url, err := h.UserService.EditUserImage(ctx, bearer, userId, headers)
	if err != nil {
		tracer.LogError(span, err)
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
	span := tracer.StartSpanFromRequest("UserHandlerRegisterUser", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling register user at %s\n", c.Path())),
	)

	userRequest := &model.UserRequest{}
	if err := c.Bind(userRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bufer, err := h.UserService.RegisterUser(ctx, userRequest)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	yter := bytes.NewReader(bufer)

	return c.Stream(http.StatusCreated, "image/png", yter)
}

func (h *userHandler) ActivateUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerActivateUser", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling activate user at %s\n", c.Path())),
	)

	activationId := c.Param("activationId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	activated, err := h.UserService.ActivateUser(ctx, activationId)
	if err != nil || activated == false {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not be activated.")
	}

	if os.Getenv("IS_PRODUCTION") == "true" {
		return c.Redirect(http.StatusMovedPermanently, "http://localhost:3000/#/login") //c.JSON(http.StatusNoContent, activationId)
	} else {
		return c.Redirect(http.StatusMovedPermanently, "https://localhost:3000/#/login") //c.JSON(http.StatusNoContent, activationId)
	}
	//return c.Redirect(http.StatusMovedPermanently, "https://localhost:3000/#/login")//c.JSON(http.StatusNoContent, activationId)
}

func (h *userHandler) ResendActivationLink(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerResendActivationLink", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling resend activation link at %s\n", c.Path())),
	)

	activateLinkRequest := &model.ActivateLinkRequest{}
	if err := c.Bind(activateLinkRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	ctx = tracer.ContextWithSpan(ctx, span)
	_, err := h.UserService.ResendActivationLink(ctx, activateLinkRequest)

	if err != nil {
		tracer.LogError(span, err)
		return ErrWrongCredentials
	}

	return c.JSON(http.StatusOK, "Activation link has been send")
}

func (h *userHandler) ResetPasswordRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerResetPasswordRequest", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling reset password request at %s\n", c.Path())),
	)

	resetPasswordRequest := &model.ResetPasswordRequest{}
	if err := c.Bind(resetPasswordRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	ctx = tracer.ContextWithSpan(ctx, span)

	_, err := h.UserService.ResetPassword(ctx, resetPasswordRequest.Email)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "Email has been send")
}

func (h *userHandler) ResetPasswordActivation(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerResetPasswordActivation", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling reset password activation at %s\n", c.Path())),
	)

	resetPasswordId := c.Param("resetPasswordId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	activated, err := h.UserService.ResetPasswordActivation(ctx, resetPasswordId)
	if err != nil || activated == false {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not reset password.")
	}

	return c.Redirect(http.StatusMovedPermanently, "https://localhost:3000/#/reset-password/"+resetPasswordId) //c.JSON(http.StatusNoContent, activationId)
}

func (h *userHandler) ChangeNewPassword(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerChangeNewPassword", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling change new password at %s\n", c.Path())),
	)

	changeNewPasswordRequest := &model.ChangeNewPasswordRequest{}
	if err := c.Bind(changeNewPasswordRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	ctx = tracer.ContextWithSpan(ctx, span)

	successful, err := h.UserService.ChangeNewPassword(ctx, changeNewPasswordRequest)

	if err != nil || !successful {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Password has been changed")
}

func (h *userHandler) GetUserEmailIfUserExist(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetUserEmailIfUserExist", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get user email if user exists at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	user, err := h.UserService.GetUserEmailIfUserExist(ctx, userId)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"emailAddress": user.Email,
	})
}

func (h *userHandler) GetUserById(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetUserById", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get user by id at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	user, err := h.UserService.GetUserById(ctx, userId)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	c.Response().Header().Set("Content-Type", "text/javascript")
	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) SearchForUsersByUsername(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerSearchForUsersByUsername", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling search for users by username at %s\n", c.Path())),
	)

	username := c.Param("username")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	users, err := h.UserService.SearchForUsersByUsername(ctx, username, bearer)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	fmt.Println(len(users))
	c.Response().Header().Set("Content-Type", "text/javascript")
	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) SearchForInfluencerByUsername(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerSearchForInfluencerByUsername", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling search for influencers by username at %s\n", c.Path())),
	)

	username := c.Param("username")
	fmt.Println(username)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	users, err := h.UserService.SearchForInfluencerByUsername(ctx, username, bearer)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	c.Response().Header().Set("Content-Type", "text/javascript")
	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) SearchForUsersByUsernameByGuest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerSearchForUsersByUsernameByGuest", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling search for influencers by username by guest at %s\n", c.Path())),
	)

	username := c.Param("username")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	users, err := h.UserService.SearchForUsersByUsernameByGuest(ctx, username)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	c.Response().Header().Set("Content-Type", "text/javascript")
	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetLoggedUserInfo(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetLoggedUserInfo", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get logged user info at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	bearer := c.Request().Header.Get("Authorization")
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	userInfo, err := h.UserService.GetLoggedUserInfo(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	return c.JSON(http.StatusOK, userInfo)
}

func (h *userHandler) GetLoggedUserTargetGroup(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetLoggedUserTargetGroup", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get logged user target group at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	userInfo, err := h.UserService.GetLoggedUserTargetGroup(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	return c.JSON(http.StatusOK, userInfo)
}


func (u *userHandler) GetLoggedAgentInfoById(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}
	userInfo, err := u.UserService.GetLoggedAgentInfoById(ctx, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unauthorized")
	}

	return c.JSON(http.StatusOK, userInfo)
}

func (h *userHandler) GetLoggedAgentInfo(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetLoggedAgentInfo", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get logged agent info at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	if ctx == nil {
		ctx = context.Background()
	}
	userInfo, err := h.UserService.GetLoggedAgentInfo(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	return c.JSON(http.StatusOK, userInfo)
}

func (h *userHandler) GetUserProfileById(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetUserProfileById", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get user profile by id at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	user, err := h.UserService.GetUserProfileById(ctx, bearer, userId)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) IsUserPrivate(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerIsUserPrivate", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling is user private at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	user, err := h.UserService.GetUserById(ctx, userId)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, user.IsPrivate)
}

func (h *userHandler) GetFollowedUsers(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetFollowedUsers", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get followed users at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	users, err := h.UserService.GetFollowedUsers(ctx, bearer, userId)
	if err != nil {
		tracer.LogError(span, err)
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
	span := tracer.StartSpanFromRequest("UserHandlerGetFollowingUsers", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get following users at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	users, err := h.UserService.GetFollowingUsers(ctx, bearer, userId)
	if err != nil {
		tracer.LogError(span, err)
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
	span := tracer.StartSpanFromRequest("UserHandlerFollowUser", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling follow user at %s\n", c.Path())),
	)

	userId := c.FormValue("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	followRequest, err := h.UserService.FollowUser(ctx, bearer, userId)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	if followRequest {
		return c.JSON(http.StatusCreated, "")
	}
	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) MuteUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerMuteUser", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling mute user at %s\n", c.Path())),
	)

	userId := c.FormValue("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	err := h.UserService.MuteUser(ctx, bearer, userId)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) UnmuteUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerUnmuteUser", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling unmute user at %s\n", c.Path())),
	)

	userId := c.FormValue("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	err := h.UserService.UnmuteUser(ctx, bearer, userId)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) BlockUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerBlockUser", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling block user at %s\n", c.Path())),
	)

	userId := c.FormValue("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	err := h.UserService.BlockUser(ctx, bearer, userId)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) UnblockUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerUnblockUser", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling unblock user at %s\n", c.Path())),
	)

	userId := c.FormValue("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	err := h.UserService.UnblockUser(ctx, bearer, userId)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) UnollowUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerUnfollowUser", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling unfollow user at %s\n", c.Path())),
	)

	userId := c.FormValue("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	err := h.UserService.UnfollowUser(ctx, bearer, userId)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) GetFollowRequests(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetFollowRequests", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get follow requests at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	reqs, err := h.UserService.GetFollowRequests(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
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
	span := tracer.StartSpanFromRequest("UserHandlerAcceptFollowRequest", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling accept follow request at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	err := h.UserService.AcceptFollowRequest(ctx, bearer, userId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) UpdateLikedPost(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerUpdateLikedPost", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling update liked post at %s\n", c.Path())),
	)

	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	err := h.UserService.UpdateLikedPost(ctx, bearer, postId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) UpdateDislikedPost(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerUpdateDisikedPost", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling update disliked post at %s\n", c.Path())),
	)

	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	err := h.UserService.UpdateDislikedPost(ctx, bearer, postId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) GetUserLikedPost(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetUserLikedPost", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get user liked post at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	retVal, err := h.UserService.GetUserLikedPost(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (h *userHandler) GetUserDislikedPost(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetUserDislikedPost", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get user disliked post at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	retVal, err := h.UserService.GetUserDislikedPost(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (h *userHandler) VerifyUser(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerVerifyUser", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling verify user at %s\n", c.Path())),
	)

	verifyAccountDTO := &model.VerifyAccountDTO{}
	if err := c.Bind(verifyAccountDTO); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	err := h.UserService.VerifyUser(ctx, verifyAccountDTO)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) CheckIfUserVerified(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerCheckIfUserVerified", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling check if user verified at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	result, err := h.UserService.CheckIfUserVerified(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *userHandler) GetFollowRecommendation(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerGetFollowRecommendation", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get follow recommendation at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	result, err := h.UserService.GetFollowRecommendation(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *userHandler) RegisterAgent(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerRegisterAgent", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling register agent at %s\n", c.Path())),
	)

	agentRegistrationDTO := &model.AgentRegistrationDTO{}
	if err := c.Bind(agentRegistrationDTO); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	result, err := h.UserService.RegisterAgent(ctx, agentRegistrationDTO)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *userHandler) AddComment(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerAddComment", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling add comment at %s\n", c.Path())),
	)
	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	if bearer == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	err := h.UserService.AddComment(ctx, bearer, postId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (h *userHandler) CheckIfUserVerifiedById(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerCheckIfUserVerifiedById", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling check if user verified by id at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	result, err := h.UserService.CheckIfUserVerifiedById(ctx, userId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *userHandler) RegisterAgentByAdmin(c echo.Context) error {
	span := tracer.StartSpanFromRequest("UserHandlerRegisterAgentByAdmin", h.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling register agent by admin at %s\n", c.Path())),
	)

	agentRequest := &model.AgentRequest{}
	if err := c.Bind(agentRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	resp, err := h.UserService.RegisterAgentByAdmin(ctx, agentRequest)

	if err != nil {
		fmt.Println(err)
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}
