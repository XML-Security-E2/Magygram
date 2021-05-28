package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
	"time"
	"user-service/conf"
	"user-service/domain/model"
	service_contracts "user-service/domain/service-contracts"
)


type UserHandler interface {
	RegisterUser(c echo.Context) error
	ActivateUser(c echo.Context) error
	LoginUser(c echo.Context) error
	ResetPasswordRequest(c echo.Context) error
	AdminCheck(c echo.Context) error
	OtherCheck(c echo.Context) error
	AuthorizationMiddleware(allowedPermissions ...string) echo.MiddlewareFunc
	ResetPasswordActivation(c echo.Context) error
	ChangeNewPassword(c echo.Context) error
	ResendActivationLink(c echo.Context) error
	GetUserEmailIfUserExist(c echo.Context) error
	GetUserById(c echo.Context) error
}

var (
	ErrHttpGenericMessage = echo.NewHTTPError(http.StatusInternalServerError, "something went wrong, please try again later")
	ErrWrongCredentials = echo.NewHTTPError(http.StatusUnauthorized, "username or password is invalid")
	ErrUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	ErrBlockedUser = echo.NewHTTPError(http.StatusForbidden, "User is not activated")
)
type userHandler struct {
	UserService service_contracts.UserService
}

func NewUserHandler(u service_contracts.UserService) UserHandler {
	return &userHandler{u}
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
	if err != nil {
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

func (h *userHandler) LoginUser(c echo.Context) error {

	loginRequest := &model.LoginRequest{}
	if err := c.Bind(loginRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	user, err := h.UserService.AuthenticateUser(ctx, loginRequest)

	if err != nil && user==nil {
		return ErrWrongCredentials
	}

	if err != nil && user != nil {
		return c.JSON(http.StatusForbidden, map[string]string{
			"userId" : user.Id,
		})
	}
	expireTime := time.Now().Add(time.Hour).Unix() * 1000
	token, err := generateToken(user, expireTime)
	if err != nil {
		return ErrHttpGenericMessage
	}

	rolesString, _ := json.Marshal(user.Roles)
	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": token,
		"roles" : string(rolesString),
		"expireTime" : strconv.FormatInt(expireTime, 10) ,
	})
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
func generateToken(user *model.User, expireTime int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	rolesString, _ := json.Marshal(user.Roles)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["surname"] = user.Surname
	claims["roles"] = string(rolesString)
	claims["id"] = user.Id
	claims["exp"] = expireTime

	return token.SignedString([]byte(conf.Current.Server.Secret))
}

func (h *userHandler) AdminCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OKET")
}

func (h *userHandler) OtherCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OKET")
}

func (h *userHandler) AuthorizationMiddleware(allowedPermissions ...string) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func (c echo.Context) error {
			authStringHeader := c.Request().Header.Get("Authorization")
			if authStringHeader == "" {
				return ErrUnauthorized
			}
			authHeader := strings.Split(authStringHeader, "Bearer ")
			jwtToken := authHeader[1]

			token, err := jwt.Parse(jwtToken, func (token *jwt.Token) (interface{}, error){
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(conf.Current.Server.Secret), nil
			})

			if err != nil {
				return ErrHttpGenericMessage
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				rolesString, _ := claims["roles"].(string)
				var tokenRoles []model.Role

				if err := json.Unmarshal([]byte(rolesString), &tokenRoles); err != nil {
					return ErrUnauthorized
				}

				if checkPermission(tokenRoles, allowedPermissions) {
					next(c)
				}

				return ErrUnauthorized
			} else{
				return ErrUnauthorized
			}
		}
	}
}

func checkPermission(userRoles []model.Role, allowedPermissions []string) bool{
	for _, role := range userRoles {
		for _, permission := range role.Permissions {
			for _, allowedPermission := range allowedPermissions {
				if permission.Name == allowedPermission {
					return true
				}
			}
		}
	}
	return false
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