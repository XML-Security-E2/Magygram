package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"relationship-service/conf"
)

var (
	ErrUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
)

func NewMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(RequestsMiddleware)
}

func RequestsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Host == conf.Current.Server.Name {
			return next(c)
		} else {
			handshake := c.Request().Header.Get(conf.Current.Server.Handshake)

			if !validSecret(handshake) {
				return ErrUnauthorized
			} else {
				return next(c)
			}
		}
	}
}

func validSecret(secretRequest string) bool {
	byteHash := []byte(secretRequest)
	plainPwd := []byte(conf.Current.Server.Secret)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
