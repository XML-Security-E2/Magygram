package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
	"post-service/domain/model"
	"post-service/domain/service-contracts"
	"post-service/domain/service-contracts/exceptions"
	"post-service/logger"
)


type PostHandler interface {
	CreatePost(c echo.Context) error
	GetPostsForTimeline(c echo.Context) error
	LikePost(c echo.Context) error
	UnlikePost(c echo.Context) error
	DislikePost(c echo.Context) error
	UndislikePost(c echo.Context) error
	GetPostsFirstImage(c echo.Context) error
	AddComment(c echo.Context) error
	EditPost(c echo.Context) error
	GetUsersPosts(c echo.Context) error
	GetUsersPostsCount(c echo.Context) error
	GetPostById(c echo.Context) error
	SearchPostsByHashTagByGuest(c echo.Context) error
	GetPostForGuestLineByHashTag(c echo.Context) error
	GetPostForUserTimelineByHashTag(c echo.Context) error
	SearchPostsByLocation(c echo.Context) error
	GetPostForGuestTimelineByLocation(c echo.Context) error
	GetPostForUserTimelineByLocation(c echo.Context) error
	GetPostByIdForGuest(c echo.Context) error
	LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	GetLikedPosts(c echo.Context) error
	GetDislikedPosts(c echo.Context) error
	DeletePost(c echo.Context) error
	EditPostOwnerInfo(c echo.Context) error
	EditLikedByInfo(c echo.Context) error
	EditDislikedByInfo(c echo.Context) error
	EditCommentedByInfo(c echo.Context) error
}

type postHandler struct {
	PostService service_contracts.PostService
}

func (p postHandler) DeletePost(c echo.Context) error {
	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := p.PostService.DeletePost(ctx, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func NewPostHandler(p service_contracts.PostService) PostHandler {
	return &postHandler{p}
}

func (p postHandler) LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.LoggingEntry = logger.Logger.WithFields(logrus.Fields{"request_ip": c.RealIP()})
		return next(c)
	}
}

func (p postHandler) CreatePost(c echo.Context) error {

	location := c.FormValue("location")
	description := c.FormValue("description")
	tagsString := c.FormValue("tags")

	mpf, _ := c.MultipartForm()
	var tags []model.Tag
	json.Unmarshal([]byte(tagsString), &tags)

	var headers []*multipart.FileHeader
	for _, v := range mpf.File {
		headers = append(headers, v[0])
	}

	postRequest := &model.PostRequest{
		Description: description,
		Location:    location,
		Media:       headers,
		Tags:        tags,
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")

	postId, err := p.PostService.CreatePost(ctx, bearer, postRequest)
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

	bearer := c.Request().Header.Get("Authorization")
	posts, err := p.PostService.GetPostsForTimeline(ctx,bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if posts==nil{
		return c.JSON(http.StatusOK, []model.PostResponse{})
	}
	return c.JSON(http.StatusOK, posts)
}

func (p postHandler) LikePost(c echo.Context) error {

	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	err := p.PostService.LikePost(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) UnlikePost(c echo.Context) error {
	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	err := p.PostService.UnlikePost(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) DislikePost(c echo.Context) error {
	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	err := p.PostService.DislikePost(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) UndislikePost(c echo.Context) error {
	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	err := p.PostService.UndislikePost(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) GetPostsFirstImage(c echo.Context) error {
	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	postImage, err := p.PostService.GetPostsFirstImage(ctx, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, postImage)
}

func (p postHandler) AddComment(c echo.Context) error {
	commentRequest := &model.CommentRequest{}
	if err := c.Bind(commentRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	retVal, err := p.PostService.AddComment(ctx, commentRequest.PostId, commentRequest.Content, bearer, commentRequest.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, retVal)
}

func (p postHandler) EditPost(c echo.Context) error {
	editRequest := &model.PostEditRequest{}
	if err := c.Bind(editRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := p.PostService.EditPost(ctx, bearer, editRequest)
	if err != nil{
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *exceptions.UnauthorizedAccessError:
			return echo.NewHTTPError(http.StatusUnauthorized, t.Error())
		}
	}
	return c.JSON(http.StatusOK, "")
}

func (p postHandler) GetUsersPosts(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	posts, err := p.PostService.GetUsersPosts(ctx, bearer, userId)
	if err != nil{
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *exceptions.UnauthorizedAccessError:
			return echo.NewHTTPError(http.StatusUnauthorized, t.Error())
		}
	}

	return c.JSON(http.StatusOK, posts)
}

func (p postHandler) GetUsersPostsCount(c echo.Context) error {
	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	postsCount, err := p.PostService.GetUsersPostsCount(ctx, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, postsCount)
}


func (p postHandler) GetPostForUserTimelineByHashTag(c echo.Context) error {
	hashTag := c.Param("value")

	ctx := c.Request().Context()
	if ctx == nil {
	ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	hashTagsPosts, err := p.PostService.GetPostForUserTimelineByHashTag(ctx, hashTag,bearer)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	if hashTagsPosts == nil{
		return c.JSON(http.StatusOK, []*model.GuestTimelinePostResponse{})
	}
	c.Response().Header().Set("Content-Type" , "text/javascript")
		return c.JSON(http.StatusOK, hashTagsPosts)
}

func (p postHandler) GetPostById(c echo.Context) error {
	postId := c.Param("postId")
	fmt.Println(postId)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	post, err := p.PostService.GetPostById(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}

func (p postHandler) SearchPostsByHashTagByGuest(c echo.Context) error {
	hashTag := c.Param("value")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	hashTagsInfo, err := p.PostService.SearchForPostsByHashTagByGuest(ctx, hashTag)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	c.Response().Header().Set("Content-Type" , "text/javascript")
	return c.JSON(http.StatusOK, hashTagsInfo)
}

func (p postHandler) GetPostForGuestLineByHashTag(c echo.Context) error {
	hashTag := c.Param("value")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	hashTagsPosts, err := p.PostService.GetPostsByHashTagForGuest(ctx, hashTag)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	if hashTagsPosts == nil{
		return c.JSON(http.StatusOK, []*model.GuestTimelinePostResponse{})
	}

	c.Response().Header().Set("Content-Type" , "text/javascript")
	return c.JSON(http.StatusOK, hashTagsPosts)
}

func (p postHandler) SearchPostsByLocation(c echo.Context) error {
	hashTag := c.Param("value")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	hashTagsInfo, err := p.PostService.SearchPostsByLocation(ctx, hashTag)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	c.Response().Header().Set("Content-Type" , "text/javascript")
	return c.JSON(http.StatusOK, hashTagsInfo)
}

func (p postHandler) GetPostForGuestTimelineByLocation(c echo.Context) error {
	location := c.Param("value")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	locationPosts, err := p.PostService.GetPostForGuestTimelineByLocation(ctx, location)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	if locationPosts == nil{
		return c.JSON(http.StatusOK, []*model.GuestTimelinePostResponse{})
	}

	c.Response().Header().Set("Content-Type" , "text/javascript")
	return c.JSON(http.StatusOK, locationPosts)
}

func (p postHandler) GetPostForUserTimelineByLocation(c echo.Context) error {
	location := c.Param("value")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	locationPosts, err := p.PostService.GetPostForUserTimelineByLocation(ctx, location,bearer)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	if locationPosts == nil{
		return c.JSON(http.StatusOK, []*model.GuestTimelinePostResponse{})
	}

	c.Response().Header().Set("Content-Type" , "text/javascript")
	return c.JSON(http.StatusOK, locationPosts)
}

func (p postHandler) GetPostByIdForGuest(c echo.Context) error {
	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	post, err := p.PostService.GetPostByIdForGuest(ctx, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}

func (p postHandler) GetLikedPosts(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	posts, err := p.PostService.GetUserLikedPosts(ctx, bearer)
	if err != nil{
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *exceptions.UnauthorizedAccessError:
			return echo.NewHTTPError(http.StatusUnauthorized, t.Error())
		}
	}

	return c.JSON(http.StatusOK, posts)
}

func (p postHandler) GetDislikedPosts(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	posts, err := p.PostService.GetUserDislikedPosts(ctx, bearer)
	if err != nil{
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *exceptions.UnauthorizedAccessError:
			return echo.NewHTTPError(http.StatusUnauthorized, t.Error())
		}
	}

	return c.JSON(http.StatusOK, posts)
}

func (p postHandler) EditPostOwnerInfo(c echo.Context) error {
	userInfo := &model.UserInfo{}
	if err := c.Bind(userInfo); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := p.PostService.EditPostOwnerInfo(ctx, bearer, userInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) EditLikedByInfo(c echo.Context) error {

	userInfo := &model.UserInfoEdit{}
	if err := c.Bind(userInfo); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := p.PostService.EditLikedByInfo(ctx, bearer, userInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) EditDislikedByInfo(c echo.Context) error {

	userInfo := &model.UserInfoEdit{}
	if err := c.Bind(userInfo); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := p.PostService.EditDislikedByInfo(ctx, bearer, userInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) EditCommentedByInfo(c echo.Context) error {
	userInfo := &model.UserInfoEdit{}
	if err := c.Bind(userInfo); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	err := p.PostService.EditCommentedByInfo(ctx, bearer, userInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}