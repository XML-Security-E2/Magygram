package handler

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"magyAgent/conf"
	"magyAgent/domain/model"
	service_contracts "magyAgent/domain/service-contracts"
	"net/http"
	"strings"
	"time"
)

type AuthHandler interface {
	RegisterUser(c echo.Context) error
	ActivateUser(c echo.Context) error
	LoginUser(c echo.Context) error
	AdminCheck(c echo.Context) error
	OtherCheck(c echo.Context) error
	AuthorizationMiddleware(roles ...string) echo.MiddlewareFunc
}

var (
	// ErrHttpGenericMessage that is returned in general case, details should be logged in such case
	ErrHttpGenericMessage = echo.NewHTTPError(http.StatusInternalServerError, "something went wrong, please try again later")

	// ErrWrongCredentials indicates that login attempt failed because of incorrect login or password
	ErrWrongCredentials = echo.NewHTTPError(http.StatusUnauthorized, "username or password is invalid")

	ErrUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")

)

type authHandler struct {
	AuthService service_contracts.AuthService
}

func NewAuthHandler(a service_contracts.AuthService) AuthHandler {
	return &authHandler{a}
}

func (h *authHandler) RegisterUser(c echo.Context) error {
	userRequest := &model.UserRequest{}
	if err := c.Bind(userRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.AuthService.RegisterUser(ctx, userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not Create.")
	}

	return c.JSON(http.StatusCreated, user.Id)
}

func (h *authHandler) ActivateUser(c echo.Context) error {
	activationId := c.Param("activationId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	activated, err := h.AuthService.ActivateUser(ctx, activationId)
	if err != nil || activated == false{
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not be activated.")
	}

	return c.Redirect(http.StatusMovedPermanently, "https://localhost:3000/#/login")//c.JSON(http.StatusNoContent, activationId)
}

func (h *authHandler) LoginUser(c echo.Context) error {

	loginRequest := &model.LoginRequest{}
	if err := c.Bind(loginRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	user, err := h.AuthService.AuthenticateUser(ctx, loginRequest)
	if err != nil || user == nil {
		return ErrWrongCredentials
	}

	token, err := generateToken(user)
	if err != nil {
		return ErrHttpGenericMessage
	}

	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": token,
		"role" : user.Role,
	})
}

func generateToken(user *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["surname"] = user.Surname
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	return token.SignedString([]byte(conf.Current.Server.Secret))
}

func (h *authHandler) AdminCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OKET")
}

func (h *authHandler) OtherCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OKET")
}

func (h *authHandler) AuthorizationMiddleware(roles ...string) echo.MiddlewareFunc {
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
				tokenRole, _ := claims["role"].(string)
				for _, role := range roles {
					if tokenRole == role {
						next(c)
					}
				}
				return ErrUnauthorized
			} else{
				return ErrUnauthorized
			}
		}
	}
}
