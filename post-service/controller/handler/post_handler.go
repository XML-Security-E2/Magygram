package handler

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"post-service/conf"
	"post-service/domain/model"
	"post-service/domain/service-contracts"
	"strings"
)


type PostHandler interface {
	CreatePost(c echo.Context) error
	GetPostsForTimeline(c echo.Context) error
}

type postHandler struct {
	PostService service_contracts.PostService
}

func NewPostHandler(p service_contracts.PostService) PostHandler {
	return &postHandler{p}
}

func (p postHandler) CreatePost(c echo.Context) error {
	postRequest := &model.PostRequest{}
	if err := c.Bind(postRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	postId, err := p.PostService.CreatePost(ctx, postRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, postId)
}

func (p postHandler) GetPostsForTimeline(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	//var userId= GetUserIdFromJWTToken(c)
	//fmt.Println(userId)
	posts, err := p.PostService.GetPostsForTimeline(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}

func GetUserIdFromJWTToken(c echo.Context) string{
	authStringHeader := c.Request().Header.Get("Authorization")
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

	claims, _ := token.Claims.(jwt.MapClaims)
	id, _ := claims["id"].(string)
	return id
}

