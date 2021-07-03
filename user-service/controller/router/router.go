package router

import (
	"github.com/labstack/echo"
	"user-service/controller/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	//users
	e.POST("/api/users", h.RegisterUser, h.UserLoggingMiddleware)
	e.POST("/api/users/agent", h.RegisterAgent, h.UserLoggingMiddleware)
	e.PUT("/api/users/:userId", h.EditUser, h.UserLoggingMiddleware)
	e.PUT("/api/users/:userId/image", h.EditUserImage, h.UserLoggingMiddleware)
	e.PUT("/api/users/:userId/notifications", h.EditUsersNotifications)
	e.PUT("/api/users/:userId/privacy-settings", h.EditUsersPrivacySettings)

	e.GET("/api/users/logged/agent", h.GetLoggedAgentInfo)
	e.GET("/api/users/logged/target-group", h.GetLoggedUserTargetGroup)

	e.GET("/api/users/logged", h.GetLoggedUserInfo, h.UserLoggingMiddleware)
	e.GET("/api/users/info/:userId", h.GetUsersInfo)

	e.GET("/api/users/activate/:activationId", h.ActivateUser, h.UserLoggingMiddleware)
	e.POST("/api/users/reset-password-link-request", h.ResetPasswordRequest, h.UserLoggingMiddleware)
	e.GET("/api/users/reset-password/:resetPasswordId", h.ResetPasswordActivation, h.UserLoggingMiddleware)
	e.POST("/api/users/reset-password", h.ChangeNewPassword, h.UserLoggingMiddleware)
	e.POST("/api/users/resend-activation-link", h.ResendActivationLink, h.UserLoggingMiddleware)
	e.GET("/api/users/check-existence/:userId", h.GetUserEmailIfUserExist, h.UserLoggingMiddleware)
	e.GET("/api/users/:userId", h.GetUserById, h.UserLoggingMiddleware)

	e.GET("/api/users/:userId/is-private", h.IsUserPrivate, h.UserLoggingMiddleware)
	e.GET("/api/users/:userId/followed", h.GetFollowedUsers, h.UserLoggingMiddleware)
	e.GET("/api/users/:userId/following", h.GetFollowingUsers, h.UserLoggingMiddleware)
	e.GET("/api/users/follow-requests", h.GetFollowRequests, h.UserLoggingMiddleware)
	e.POST("/api/users/follow-requests/:userId/accept", h.AcceptFollowRequest, h.UserLoggingMiddleware)
	e.POST("/api/users/follow", h.FollowUser, h.UserLoggingMiddleware)
	e.POST("/api/users/unfollow", h.UnollowUser, h.UserLoggingMiddleware)
	e.POST("/api/users/mute", h.MuteUser, h.UserLoggingMiddleware)
	e.POST("/api/users/unmute", h.UnmuteUser, h.UserLoggingMiddleware)
	e.POST("/api/users/block", h.BlockUser, h.UserLoggingMiddleware)
	e.POST("/api/users/unblock", h.UnblockUser, h.UserLoggingMiddleware)

	e.GET("/api/users/:userId/profile", h.GetUserProfileById, h.UserLoggingMiddleware)

	e.GET("/api/users/search/:username", h.SearchForUsersByUsername, h.UserLoggingMiddleware)
	e.GET("/api/users/search/:username/user", h.SearchForUsersByUsername, h.UserLoggingMiddleware)
	e.GET("/api/users/search/:username/influencer", h.SearchForInfluencerByUsername, h.UserLoggingMiddleware)
	e.GET("/api/users/search/:username/guest", h.SearchForUsersByUsernameByGuest, h.UserLoggingMiddleware)

	e.POST("/api/users/highlights", h.CreateHighlights, h.HighlightsLoggingMiddleware)
	e.GET("/api/users/:userId/highlights", h.GetProfileHighlights, h.HighlightsLoggingMiddleware)
	e.GET("/api/users/:userId/highlights/:name", h.GetProfileHighlightsByHighlightName, h.HighlightsLoggingMiddleware)


	e.POST("/api/users/collections", h.CreateCollection, h.CollectionsLoggingMiddleware)
	e.POST("/api/users/collections/posts", h.AddPostToCollection, h.CollectionsLoggingMiddleware)
	e.GET("/api/users/collections/:collectionName/posts", h.GetCollectionPosts, h.CollectionsLoggingMiddleware)
	e.DELETE("/api/users/collections/posts/:postId", h.DeleteFromCollection, h.CollectionsLoggingMiddleware)
	e.GET("/api/users/collections/except-default", h.GetUsersCollectionsExceptDefault, h.CollectionsLoggingMiddleware)
	e.GET("/api/users/collections", h.GetUserCollections, h.CollectionsLoggingMiddleware)
	e.POST("/api/users/collections/check-favourites", h.CheckIfPostInFavourites, h.CollectionsLoggingMiddleware)

	e.PUT("api/users/post/like/:postId", h.UpdateLikedPost)
	e.PUT("api/users/post/comment/:postId", h.AddComment)

	e.PUT("api/users/post/dislike/:postId", h.UpdateDislikedPost)
	e.GET("api/users/post/liked", h.GetUserLikedPost)
	e.GET("api/users/post/disliked", h.GetUserDislikedPost)

	e.PUT("/api/users/verify", h.VerifyUser)
	e.GET("/api/users/isverified", h.CheckIfUserVerified)
	e.GET("/api/users/isverified/:userId", h.CheckIfUserVerifiedById)

	e.GET("api/users/:userId/notify/post", h.GetUsersForPostNotification)
	e.GET("api/users/:userId/notify/story", h.GetUsersForStoryNotification)

	e.GET("api/users/notifications/get/:userId", h.GetUsersNotificationsSettings)
	e.POST("api/users/notifications/settings/:userId", h.ChangeUsersNotificationsSettings)

	e.GET("api/users/:userId/:fromId/notify/:interactionType", h.CheckIfPostInteractionNotificationEnabled)

	e.PUT("/api/users/:requestId/delete", h.DeleteUser)

	e.GET("/api/users/follow-recommendation", h.GetFollowRecommendation)
	e.POST("/api/users/register-agent-by-admin", h.RegisterAgentByAdmin, h.UserLoggingMiddleware)

}