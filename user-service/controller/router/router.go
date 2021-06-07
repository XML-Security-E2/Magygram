package router

import (
	"github.com/labstack/echo"
	"user-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	//users
	e.POST("/api/users", h.RegisterUser)
	e.PUT("/api/users/:userId", h.EditUser)
	e.PUT("/api/users/:userId/image", h.EditUserImage)

	e.GET("/api/users/logged", h.GetLoggedUserInfo)
	e.GET("/api/users/activate/:activationId", h.ActivateUser)
	e.POST("/api/users/reset-password-link-request", h.ResetPasswordRequest)
	e.GET("/api/users/reset-password/:resetPasswordId", h.ResetPasswordActivation)
	e.POST("/api/users/reset-password", h.ChangeNewPassword)
	e.POST("/api/users/resend-activation-link", h.ResendActivationLink)
	e.GET("/api/users/check-existence/:userId", h.GetUserEmailIfUserExist)
	e.GET("/api/users/:userId", h.GetUserById)

	e.GET("/api/users/:userId/is-private", h.IsUserPrivate)
	e.GET("/api/users/:userId/followed", h.GetFollowedUsers)
	e.GET("/api/users/:userId/following", h.GetFollowingUsers)
	e.GET("/api/users/follow-requests", h.GetFollowRequests)
	e.POST("/api/users/follow-requests/:userId/accept", h.AcceptFollowRequest)
	e.POST("/api/users/follow", h.FollowUser)
	e.POST("/api/users/unfollow", h.UnollowUser)

	e.GET("/api/users/:userId/profile", h.GetUserProfileById)

	e.GET("/api/users/search/:username", h.SearchForUsersByUsername)
	e.GET("/api/users/search/:username/user", h.SearchForUsersByUsername)
	e.GET("/api/users/search/:username/guest", h.SearchForUsersByUsernameByGuest)

	e.POST("/api/users/highlights", h.CreateHighlights)
	e.GET("/api/users/:userId/highlights", h.GetProfileHighlights)
	e.GET("/api/users/:userId/highlights/:name", h.GetProfileHighlightsByHighlightName)


	e.POST("/api/users/collections", h.CreateCollection)
	e.POST("/api/users/collections/posts", h.AddPostToCollection)
	e.GET("/api/users/collections/:collectionName/posts", h.GetCollectionPosts)
	e.DELETE("/api/users/collections/posts/:postId", h.DeleteFromCollection)
	e.GET("/api/users/collections/except-default", h.GetUsersCollectionsExceptDefault)
	e.GET("/api/users/collections", h.GetUserCollections)
	e.POST("/api/users/collections/check-favourites", h.CheckIfPostInFavourites)
}