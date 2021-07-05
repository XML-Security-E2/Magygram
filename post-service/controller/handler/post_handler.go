package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"net/http"
	"post-service/domain/model"
	"post-service/domain/service-contracts"
	"post-service/domain/service-contracts/exceptions"
	"post-service/logger"
	"post-service/tracer"
	"strconv"
	"time"
)


type PostHandler interface {
	CreatePost(c echo.Context) error
	CreatePostCampaign(c echo.Context) error
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
	GetPostForMessagesById(c echo.Context) error
	GetUsersPostCampaigns(c echo.Context) error
}

type postHandler struct {
	PostService service_contracts.PostService
	tracer      opentracing.Tracer
	closer      io.Closer
}

func NewPostHandler(p service_contracts.PostService) PostHandler {
	tracer, closer := tracer.Init("post-service")
	opentracing.SetGlobalTracer(tracer)
	return &postHandler{p, tracer, closer}
}

func (p postHandler) DeletePost(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerDeletePost", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling delete post at %s\n", c.Path())),
	)

	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	err := p.PostService.DeletePost(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.LoggingEntry = logger.Logger.WithFields(logrus.Fields{"request_ip": c.RealIP()})
		return next(c)
	}
}

func (p postHandler) CreatePost(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerCreatePost", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create post at %s\n", c.Path())),
	)

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
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	postId, err := p.PostService.CreatePost(ctx, bearer, postRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, postId)
}

func (p postHandler) CreatePostCampaign(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerCreatePostCampaign", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create post campaign at %s\n", c.Path())),
	)

	location := c.FormValue("location")
	description := c.FormValue("description")
	tagsString := c.FormValue("tags")
	minD := c.FormValue("minDisplays")
	minDisplays, _ := strconv.Atoi(minD)
	frequency := c.FormValue("frequency")
	minA := c.FormValue("minAge")
	minAge, _ := strconv.Atoi(minA)
	maxA := c.FormValue("maxAge")
	maxAge, _ := strconv.Atoi(maxA)
	gender := c.FormValue("gender")

	dateF := c.FormValue("startDate")
	dateFr, _ := strconv.ParseInt(dateF, 10, 64)
	dateFrom := time.Unix(0, dateFr * int64(time.Millisecond))

	dateT := c.FormValue("endDate")
	dateTt, _ := strconv.ParseInt(dateT, 10, 64)
	dateTo := time.Unix(0, dateTt * int64(time.Millisecond))

	exposeD := c.FormValue("exposeOnceDate")
	exposeDa, _ := strconv.ParseInt(exposeD, 10, 64)
	exposeDate := time.Unix(0, exposeDa * int64(time.Millisecond))

	displayT := c.FormValue("displayTime")
	displayTime, _ := strconv.Atoi(displayT)

	mpf, _ := c.MultipartForm()
	var tags []model.Tag
	json.Unmarshal([]byte(tagsString), &tags)

	var headers []*multipart.FileHeader
	for _, v := range mpf.File {
		headers = append(headers, v[0])
	}

	postRequest := &model.PostRequest{
		Description:              description,
		Location:                 location,
		Media:                    headers,
		Tags:                     tags,
	}

	campaignRequest := &model.CampaignRequest{
		MinDisplaysForRepeatedly: minDisplays,
		Frequency:                model.CampaignFrequency(frequency),
		TargetGroup:              model.TargetGroup{
			MinAge: minAge,
			MaxAge: maxAge,
			Gender: model.GenderType(gender),
		},
		DateFrom:                 dateFrom,
		DateTo:                   dateTo,
		Type: "POST",
		DisplayTime: displayTime,
		ContentId: "",
		ExposeOnceDate: exposeDate,
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	postId, err := p.PostService.CreatePostCampaign(ctx, bearer, postRequest, campaignRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, postId)
}

func (p postHandler) GetUsersPostCampaigns(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerGetUsersPostCampaigns", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get users post campaigns at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	posts, err := p.PostService.GetUserPostCampaigns(ctx,bearer)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if posts == nil {
		return c.JSON(http.StatusOK, []model.PostProfileResponse{})
	}

	fmt.Println(posts[0].Id)
	return c.JSON(http.StatusOK, posts)
}

func (p postHandler) GetPostsForTimeline(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerGetPostsForTimeline", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get posts for timeline at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
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
	span := tracer.StartSpanFromRequest("PostHandlerLikePost", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling like post at %s\n", c.Path())),
	)

	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	err := p.PostService.LikePost(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) UnlikePost(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerUnlikePost", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling unlike post at %s\n", c.Path())),
	)

	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	err := p.PostService.UnlikePost(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) DislikePost(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerDislikePost", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling dislike post at %s\n", c.Path())),
	)

	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	err := p.PostService.DislikePost(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) UndislikePost(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerUndislikePost", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling undislike post at %s\n", c.Path())),
	)

	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")
	err := p.PostService.UndislikePost(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) GetPostsFirstImage(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerGetPostsFirstImage", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get post first image at %s\n", c.Path())),
	)

	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	postImage, err := p.PostService.GetPostsFirstImage(ctx, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, postImage)
}

func (p postHandler) AddComment(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerAddComment", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling add comment to post at %s\n", c.Path())),
	)

	commentRequest := &model.CommentRequest{}
	if err := c.Bind(commentRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	retVal, err := p.PostService.AddComment(ctx, commentRequest.PostId, commentRequest.Content, bearer, commentRequest.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, retVal)
}

func (p postHandler) EditPost(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerEditPost", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling edit post at %s\n", c.Path())),
	)

	editRequest := &model.PostEditRequest{}
	if err := c.Bind(editRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

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
	span := tracer.StartSpanFromRequest("PostHandlerGetUsersPosts", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get users posts at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

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
	span := tracer.StartSpanFromRequest("PostHandlerGetUsersPostsCount", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get users posts count at %s\n", c.Path())),
	)

	userId := c.Param("userId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	postsCount, err := p.PostService.GetUsersPostsCount(ctx, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, postsCount)
}


func (p postHandler) GetPostForUserTimelineByHashTag(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerGetPostForUserTimelineByHashTag", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get post for user timeline by hashtag at %s\n", c.Path())),
	)

	hashTag := c.Param("value")

	ctx := c.Request().Context()
	if ctx == nil {
	ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
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

func (p postHandler) GetPostForMessagesById(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerGetPostForMessagesById", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get post for messages by id at %s\n", c.Path())),
	)

	postId := c.Param("postId")
	fmt.Println(postId)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	post, userInfo, err := p.PostService.GetPostForMessagesById(ctx, bearer, postId)
	if err != nil{
		switch t := err.(type) {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, t.Error())
		case *exceptions.UnauthorizedAccessError:
			return echo.NewHTTPError(http.StatusUnauthorized, userInfo)
		}
	}

	return c.JSON(http.StatusOK, post)}

func (p postHandler) GetPostById(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerGetPostById", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get post by id at %s\n", c.Path())),
	)

	postId := c.Param("postId")
	fmt.Println(postId)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	post, err := p.PostService.GetPostById(ctx, bearer, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}

func (p postHandler) SearchPostsByHashTagByGuest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerSearchPostsByHashTagByGuest", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling search posts by hashtag by guest at %s\n", c.Path())),
	)

	hashTag := c.Param("value")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	hashTagsInfo, err := p.PostService.SearchForPostsByHashTagByGuest(ctx, hashTag)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	c.Response().Header().Set("Content-Type" , "text/javascript")
	return c.JSON(http.StatusOK, hashTagsInfo)
}

func (p postHandler) GetPostForGuestLineByHashTag(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerGetPostForGuestLineByHashTag", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get post for guestline by hashtag at %s\n", c.Path())),
	)

	hashTag := c.Param("value")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

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
	span := tracer.StartSpanFromRequest("PostHandlerSearchPostsByLocation", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling search posts by location at %s\n", c.Path())),
	)

	hashTag := c.Param("value")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	hashTagsInfo, err := p.PostService.SearchPostsByLocation(ctx, hashTag)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't find any users")
	}

	c.Response().Header().Set("Content-Type" , "text/javascript")
	return c.JSON(http.StatusOK, hashTagsInfo)
}

func (p postHandler) GetPostForGuestTimelineByLocation(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerGetPostForGuestTimelineByLocation", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get post for guest timeline by location at %s\n", c.Path())),
	)

	location := c.Param("value")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

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
	span := tracer.StartSpanFromRequest("PostHandlerGetPostForUserTimelineByLocation", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get post for user timeline by location at %s\n", c.Path())),
	)

	location := c.Param("value")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
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
	span := tracer.StartSpanFromRequest("PostHandlerGetPostByIdForGuest", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get post by id for guest at %s\n", c.Path())),
	)

	postId := c.Param("postId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	post, err := p.PostService.GetPostByIdForGuest(ctx, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}

func (p postHandler) GetLikedPosts(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerGetLikedPosts", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get liked posts at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

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
	span := tracer.StartSpanFromRequest("PostHandlerGetDislikedPosts", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get disliked posts at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

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
	span := tracer.StartSpanFromRequest("PostHandlerEditPostOwnerInfo", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling edit post owner info at %s\n", c.Path())),
	)

	userInfo := &model.UserInfo{}
	if err := c.Bind(userInfo); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	err := p.PostService.EditPostOwnerInfo(ctx, bearer, userInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) EditLikedByInfo(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerEditLikedByInfo", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling edit liked by info at %s\n", c.Path())),
	)

	userInfo := &model.UserInfoEdit{}
	if err := c.Bind(userInfo); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	err := p.PostService.EditLikedByInfo(ctx, bearer, userInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) EditDislikedByInfo(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerEditDislikedByInfo", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling edit disliked by info at %s\n", c.Path())),
	)

	userInfo := &model.UserInfoEdit{}
	if err := c.Bind(userInfo); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	err := p.PostService.EditDislikedByInfo(ctx, bearer, userInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (p postHandler) EditCommentedByInfo(c echo.Context) error {
	span := tracer.StartSpanFromRequest("PostHandlerEditCommentedByInfo", p.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling edit commented by info at %s\n", c.Path())),
	)

	userInfo := &model.UserInfoEdit{}
	if err := c.Bind(userInfo); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	err := p.PostService.EditCommentedByInfo(ctx, bearer, userInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}