package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	_ "html/template"
	"io"
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
	ResetPasswordRequest(c echo.Context) error
	AdminCheck(c echo.Context) error
	OtherCheck(c echo.Context) error
	AuthorizationMiddleware(allowedPermissions ...string) echo.MiddlewareFunc
	ResetPasswordActivation(c echo.Context) error
	ChangeNewPassword(c echo.Context) error
	ResendActivationLink(c echo.Context) error
	GetUserEmailIfUserExist(c echo.Context) error
}

var (
	ErrHttpGenericMessage = echo.NewHTTPError(http.StatusInternalServerError, "something went wrong, please try again later")
	ErrWrongCredentials = echo.NewHTTPError(http.StatusUnauthorized, "username or password is invalid")
	ErrUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	ErrBlockedUser = echo.NewHTTPError(http.StatusForbidden, "User is not activated")
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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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

func HTMLEscape(w io.Writer, b []byte) {

}

type Todo struct {
	Name        string
	Password string
}
func (h *authHandler) LoginUser(c echo.Context) error {

	/*loginRequest := &model.LoginRequest{}
	fmt.Println("sakjefskjfs" + loginRequest.Email)
	td := Todo{"Test templates", "Let's test a template to see the magic."}


	t, err := template.New("todos").Parse("You have a task named \"{{ .loginRequest.Email}}\" with description: \"{{ .loginRequest.Password}}\"")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}

	if err := c.Bind(loginRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	user, err := h.AuthService.AuthenticateUser(ctx, loginRequest)

	if err != nil && user==nil {
		return ErrWrongCredentials
	}

	if err != nil && user != nil {
		return ErrBlockedUser
	}

	token, err := generateToken(user)
	if err != nil {
		return ErrHttpGenericMessage
	}

	rolesString, _ := json.Marshal(user.Roles)
	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": token,
		"roles" : string(rolesString),
	})

	loginRequest := &model.LoginRequest{}
	if err := c.Bind(loginRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	user, err := h.AuthService.AuthenticateUser(ctx, loginRequest)

	td := Todo{user.Email, user.Password}
	t, err := template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Password}}\"")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}



	token, err := generateToken(user)
	if err != nil {
		return ErrHttpGenericMessage
	}

	rolesString, _ := json.Marshal(user.Roles)
	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": token,
		"roles" : string(rolesString),
	})*/
	loginRequest := &model.LoginRequest{}
	if err := c.Bind(loginRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	user, err := h.AuthService.AuthenticateUser(ctx, loginRequest)

	if err != nil && user==nil {
		return ErrWrongCredentials
	}

	if err != nil && user != nil {
		return c.JSON(http.StatusForbidden, map[string]string{
			"userId" : user.Id,
		})
	}

	token, err := generateToken(user)
	if err != nil {
		return ErrHttpGenericMessage
	}

	rolesString, _ := json.Marshal(user.Roles)
	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": token,
		"roles" : string(rolesString),
	})
}

func (h *authHandler) ResendActivationLink(c echo.Context) error {

	activateLinkRequest := &model.ActivateLinkRequest{}
	if err := c.Bind(activateLinkRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	_, err := h.AuthService.ResendActivationLink(ctx, activateLinkRequest)

	if err != nil {
		return ErrWrongCredentials
	}

	return c.JSON(http.StatusOK, "Activation link has been send")
}

func (h *authHandler) ResetPasswordRequest(c echo.Context) error {
	resetPasswordRequest := &model.ResetPasswordRequest{}
	if err := c.Bind(resetPasswordRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()

	_, err := h.AuthService.ResetPassword(ctx, resetPasswordRequest.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "Email has been send")
}
func generateToken(user *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	rolesString, _ := json.Marshal(user.Roles)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["surname"] = user.Surname
	claims["roles"] = string(rolesString)
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	return token.SignedString([]byte(conf.Current.Server.Secret))
}

func (h *authHandler) AdminCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OKET")
}

func (h *authHandler) OtherCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OKET")
}

func (h *authHandler) AuthorizationMiddleware(allowedPermissions ...string) echo.MiddlewareFunc {
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

func (h *authHandler) ResetPasswordActivation(c echo.Context) error {
	resetPasswordId := c.Param("resetPasswordId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	activated, err := h.AuthService.ResetPasswordActivation(ctx, resetPasswordId)
	if err != nil || activated == false{
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not reset password.")
	}

	return c.Redirect(http.StatusMovedPermanently, "https://localhost:3000/#/reset-password/" + resetPasswordId)//c.JSON(http.StatusNoContent, activationId)
}

func (h *authHandler) ChangeNewPassword(c echo.Context) error {
	changeNewPasswordRequest := &model.ChangeNewPasswordRequest{}
	if err := c.Bind(changeNewPasswordRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()

	_, err := h.AuthService.ChangeNewPassword(ctx, changeNewPasswordRequest)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "Password has been changed")
}

func (h *authHandler) GetUserEmailIfUserExist(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.AuthService.GetUserEmailIfUserExist(ctx, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"emailAddress": user.Email,
	})
}