package service_contracts

import (
	"context"
	"mime/multipart"
	"user-service/domain/model"
)

type UserService interface {
	RegisterUser(ctx context.Context, user *model.UserRequest) ([]byte, error)
	EditUser(ctx context.Context, bearer string, userId string, user *model.EditUserRequest) (string, error)
	EditUserImage(ctx context.Context, bearer string, userId string, userImage []*multipart.FileHeader) (string, error)
	ActivateUser(ctx context.Context, activationId string) (bool, error)
	ResetPassword(ctx context.Context, userEmail string) (bool, error)
	ResetPasswordActivation(ctx context.Context, resetPasswordId string) (bool, error)
	ChangeNewPassword(ctx context.Context, changePasswordRequest *model.ChangeNewPasswordRequest) (bool, error)
	ResendActivationLink(ctx context.Context, activateLinkRequest *model.ActivateLinkRequest) (bool, error)
	GetUserEmailIfUserExist(ctx context.Context, userId string) (*model.User, error)
	GetUserById(ctx context.Context, userId string) (*model.User, error)
	GetLoggedUserInfo(ctx context.Context, bearer string) (*model.UserInfo, error)
	GetUserProfileById(ctx context.Context, bearer string, userId string) (*model.UserProfileResponse, error)
	GetFollowedUsers(ctx context.Context, bearer string, userId string) ([]*model.UserFollowingResponse, error)
	GetFollowingUsers(ctx context.Context, bearer string, userId string) ([]*model.UserFollowingResponse, error)
	GetFollowRequests(ctx context.Context, bearer string) ([]*model.UserFollowingResponse, error)
	FollowUser(ctx context.Context, bearer string, userId string) (bool, error)
	UnfollowUser(ctx context.Context, bearer string, userId string) error
	MuteUser(ctx context.Context, bearer string, userId string) error
	UnmuteUser(ctx context.Context, bearer string, userId string) error
	BlockUser(ctx context.Context, bearer string, userId string) error
	UnblockUser(ctx context.Context, bearer string, userId string) error
	SearchForUsersByUsername(ctx context.Context, username string, bearer string) ([]model.User, error)
	SearchForInfluencerByUsername(ctx context.Context, username string, bearer string) ([]model.User, error)
	SearchForUsersByUsernameByGuest(ctx context.Context, username string) ([]model.User, error)
	AcceptFollowRequest(ctx context.Context, bearer string, userId string) error
	UpdateLikedPost(ctx context.Context, bearer string, postId string) error
	UpdateDislikedPost(ctx context.Context, bearer string, postId string) error
	AddComment(ctx context.Context, bearer string, postId string) error
	GetUserLikedPost(ctx context.Context, bearer string) ([]string, error)
	GetUserDislikedPost(ctx context.Context, bearer string) ([]string, error)
	EditUsersPrivacySettings(ctx context.Context, bearer string, privacySettingsReq *model.PrivacySettings) error
	EditUsersNotifications(ctx context.Context, bearer string, notificationReq *model.NotificationSettingsUpdateReq) error

	GetUsersForPostNotification(ctx context.Context, userId string) ([]*model.UserInfo, error)
	GetUsersForStoryNotification(ctx context.Context, userId string) ([]*model.UserInfo, error)
	CheckIfPostInteractionNotificationEnabled(ctx context.Context, userId string, userFromId string, interactionType string) (bool, error)
	VerifyUser(ctx context.Context, dto *model.VerifyAccountDTO) error
	CheckIfUserVerified(ctx context.Context, bearer string) (bool, error)

	GetUsersNotificationsSettings(ctx context.Context, bearer string, userId string) (*model.SettingsRequest, error)
	ChangeUsersNotificationsSettings(ctx context.Context, bearer string, settingsReq *model.SettingsRequest, userId string) error
	DeleteUser(ctx context.Context, requestId string) error
	GetFollowRecommendation(ctx context.Context, bearer string) (*model.FollowRecommendationResponse, error)
	RegisterAgent(ctx context.Context, agentRegistrationDTO *model.AgentRegistrationDTO) (string, error)

	CheckIfUserVerifiedById(ctx context.Context, userId string) (bool, error)

	GetUsersInfo(ctx context.Context, userId string) (*model.UserInfo, error)
	RegisterAgentByAdmin(ctx context.Context, agentRequest *model.AgentRequest) (string, error)
	RedisConnection(finished chan bool)
}
