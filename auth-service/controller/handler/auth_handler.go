package handler

import (
	"auth-service/conf"
	"auth-service/domain/model"
	"auth-service/domain/service-contracts"
	"auth-service/logger"
	"auth-service/tracer"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AuthHandler interface {
	LoginFirstStage(c echo.Context) error
	AdminCheck(c echo.Context) error
	AuthorizationSuccess(c echo.Context) error
	AuthorizationMiddleware() echo.MiddlewareFunc
	GetLoggedUserId(c echo.Context) error
	AuthLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	LoginSecondStage(c echo.Context) error
	GenerateNewAgentCampaignJWTToken(c echo.Context) error
	DeleteCampaignJWTToken(c echo.Context) error
	GetCampaignJWTToken(c echo.Context) error
}

type authHandler struct {
	AuthService service_contracts.AuthService
	tracer      opentracing.Tracer
	closer      io.Closer
}

func NewAuthHandler(a service_contracts.AuthService) AuthHandler {
	tracer, closer := tracer.Init("auth-service")
	opentracing.SetGlobalTracer(tracer)
	return &authHandler{a, tracer, closer}
}

func (a authHandler) AuthLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.LoggingEntry = logger.Logger.WithFields(logrus.Fields{"request_ip": c.RealIP()})
		return next(c)
	}
}

func (a authHandler) LoginFirstStage(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AuthHandlerLoginFirstStage", a.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling login first stage at %s\n", c.Path())),
	)

	loginRequest := &model.LoginRequest{}
	if err := c.Bind(loginRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	user, err := a.AuthService.AuthenticateUser(ctx, loginRequest)

	if err != nil && user == nil {
		tracer.LogError(span, err)
		return ErrWrongCredentials
	}

	if err != nil && user != nil {
		tracer.LogError(span, err)
		return c.JSON(http.StatusForbidden, map[string]string{
			"userId": user.Id,
		})
	}
	//vracati ako je ukljucen 2fa
	//return c.JSON(http.StatusOK, user.Id)

	expireTime := time.Now().Add(time.Hour).Unix() * 1000
	token, err := generateToken(user, expireTime)
	if err != nil {
		tracer.LogError(span, err)
		return ErrHttpGenericMessage
	}

	rolesString, _ := json.Marshal(user.Roles)
	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": token,
		"roles":       string(rolesString),
		"expireTime":  strconv.FormatInt(expireTime, 10),
	})
}

func (a authHandler) LoginSecondStage(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AuthHandlerLoginSecondStage", a.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling login second stage at %s\n", c.Path())),
	)

	loginRequest := &model.LoginTwoFactoryRequest{}
	if err := c.Bind(loginRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	user, err := a.AuthService.AuthenticateTwoFactoryUser(ctx, loginRequest)

	if err != nil {
		tracer.LogError(span, err)
		return err
	}

	if user == nil {
		tracer.LogError(span, err)
		return c.JSON(http.StatusForbidden, "")
	}

	expireTime := time.Now().Add(time.Hour).Unix() * 1000
	token, err := generateToken(user, expireTime)
	if err != nil {
		tracer.LogError(span, err)
		return ErrHttpGenericMessage
	}

	rolesString, _ := json.Marshal(user.Roles)
	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": token,
		"roles":       string(rolesString),
		"expireTime":  strconv.FormatInt(expireTime, 10),
	})
}

func generateToken(user *model.User, expireTime int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	rolesString, _ := json.Marshal(user.Roles)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["roles"] = string(rolesString)
	claims["id"] = user.Id
	claims["exp"] = expireTime

	return token.SignedString([]byte(conf.Current.Server.Secret))
}

func (a authHandler) AdminCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OKET")
}

func (a authHandler) AuthorizationSuccess(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func (a authHandler) AuthorizationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var allowedPermissions []string
			permissionsHeader := c.Request().Header.Get("X-permissions")
			log.Println(permissionsHeader)
			json.Unmarshal([]byte(permissionsHeader), &allowedPermissions)

			authStringHeader := c.Request().Header.Get("Authorization")
			if authStringHeader == "" {
				return ErrUnauthorized
			}
			authHeader := strings.Split(authStringHeader, "Bearer ")
			jwtToken := authHeader[1]

			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
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
					_ = next(c)
				}

				return ErrUnauthorized
			} else {
				return ErrUnauthorized
			}
		}
	}
}

func checkPermission(userRoles []model.Role, allowedPermissions []string) bool {
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

func (a authHandler) GetLoggedUserId(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AuthHandlerGetLoggedUserId", a.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get logged user id at %s\n", c.Path())),
	)

	authStringHeader := c.Request().Header.Get("Authorization")

	if authStringHeader == "" {
		return ErrUnauthorized
	}
	authHeader := strings.Split(authStringHeader, "Bearer ")
	jwtToken := authHeader[1]

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.Current.Server.Secret), nil
	})

	if err != nil {
		tracer.LogError(span, err)
		return ErrHttpGenericMessage
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, _ := claims["id"].(string)
		return c.JSON(http.StatusOK, userId)
	} else {
		return ErrUnauthorized
	}
}

func (a authHandler) GenerateNewAgentCampaignJWTToken(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AuthHandlerGenerateNewAgentCampaignJWTToken", a.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling generate new agent campaign jwt token at %s\n", c.Path())),
	)

	expireTime := time.Now().Add(time.Hour*8760).Unix() * 1000 // 1 year

	bearer := c.Request().Header.Get("Authorization")
	userId := getUserId(bearer)
	if userId == ""{
		return ErrHttpGenericMessage
	}

	token, err := generateAgentCampaignJWTToken(expireTime, userId)
	if err != nil {
		tracer.LogError(span, err)
		return ErrHttpGenericMessage
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}



	err = a.AuthService.UpdateAgentCampaignJWTToken(ctx, bearer, token)
	if err != nil {
		tracer.LogError(span, err)
		return ErrHttpGenericMessage
	}

	return c.JSON(http.StatusOK, token)
}

func getUserId(authStringHeader string) string {
	if authStringHeader == "" {
		return ""
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
		return ""
	}

	userId := ""
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, _ = claims["id"].(string)
	}
	return userId
}

func generateAgentCampaignJWTToken(expireTime int64, userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	var roles = []model.Role{{Name: "campaign_role", Permissions: []model.Permission{
		{"create_campaign"},
		{"get_agent_campaign"},
		{"get_monitoring_for_campaign"}}}}

	rolesString, _ := json.Marshal(roles)

	claims := token.Claims.(jwt.MapClaims)
	claims["roles"] = string(rolesString)
	claims["exp"] = expireTime
	claims["id"] = userId

	return token.SignedString([]byte(conf.Current.Server.Secret))
}

func (a authHandler) DeleteCampaignJWTToken(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AuthHandlerDeleteCampaignJWTToken", a.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling delete campaign jwt token at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	err := a.AuthService.DeleteCampaignJWTToken(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return ErrHttpGenericMessage
	}

	return c.JSON(http.StatusOK, "")
}

func (a authHandler) GetCampaignJWTToken(c echo.Context) error {
	span := tracer.StartSpanFromRequest("AuthHandlerGetCampaignJWTToken", a.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get campaign jwt token at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	jwtToken, err := a.AuthService.GetCampaignJWTToken(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return ErrHttpGenericMessage
	}

	return c.JSON(http.StatusOK, jwtToken)
}
