package router

import (
	"github.com/labstack/echo"
	"user-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	//users
	e.POST("/api/users", h.RegisterUser, h.LoggingMiddleware)
	e.PUT("/api/users/:userId", h.EditUser, h.LoggingMiddleware)
	e.PUT("/api/users/:userId/image", h.EditUserImage, h.LoggingMiddleware)

	e.GET("/api/users/logged", h.GetLoggedUserInfo)
	e.GET("/api/users/activate/:activationId", h.ActivateUser, h.LoggingMiddleware)
	e.POST("/api/users/reset-password-link-request", h.ResetPasswordRequest, h.LoggingMiddleware)
	e.GET("/api/users/reset-password/:resetPasswordId", h.ResetPasswordActivation, h.LoggingMiddleware)
	e.POST("/api/users/reset-password", h.ChangeNewPassword, h.LoggingMiddleware)
	e.POST("/api/users/resend-activation-link", h.ResendActivationLink, h.LoggingMiddleware)
	e.GET("/api/users/check-existence/:userId", h.GetUserEmailIfUserExist)
	e.GET("/api/users/:userId", h.GetUserById)

	e.GET("/api/users/:userId/is-private", h.IsUserPrivate)
	e.GET("/api/users/:userId/followed", h.GetFollowedUsers)
	e.GET("/api/users/:userId/following", h.GetFollowingUsers)
	e.GET("/api/users/follow-requests", h.GetFollowRequests)
	e.POST("/api/users/follow-requests/:userId/accept", h.AcceptFollowRequest, h.LoggingMiddleware)
	e.POST("/api/users/follow", h.FollowUser, h.LoggingMiddleware)
	e.POST("/api/users/unfollow", h.UnollowUser, h.LoggingMiddleware)

	e.GET("/api/users/:userId/profile", h.GetUserProfileById)

	e.GET("/api/users/search/:username", h.SearchForUsersByUsername)
	e.GET("/api/users/search/:username/user", h.SearchForUsersByUsername)
	e.GET("/api/users/search/:username/guest", h.SearchForUsersByUsernameByGuest)

	e.POST("/api/users/highlights", h.CreateHighlights, h.LoggingMiddleware)
	e.GET("/api/users/:userId/highlights", h.GetProfileHighlights)
	e.GET("/api/users/:userId/highlights/:name", h.GetProfileHighlightsByHighlightName)


	e.POST("/api/users/collections", h.CreateCollection, h.LoggingMiddleware)
	e.POST("/api/users/collections/posts", h.AddPostToCollection, h.LoggingMiddleware)
	e.GET("/api/users/collections/:collectionName/posts", h.GetCollectionPosts)
	e.DELETE("/api/users/collections/posts/:postId", h.DeleteFromCollection, h.LoggingMiddleware)
	e.GET("/api/users/collections/except-default", h.GetUsersCollectionsExceptDefault)
	e.GET("/api/users/collections", h.GetUserCollections)
	e.POST("/api/users/collections/check-favourites", h.CheckIfPostInFavourites, h.LoggingMiddleware)
}