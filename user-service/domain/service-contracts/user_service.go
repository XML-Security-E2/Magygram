package service_contracts

import (
	"context"
	"mime/multipart"
	"net/http"
	"user-service/domain/model"
)

type UserService interface {
	RegisterUser(ctx context.Context, user *model.UserRequest) (*http.Response, error)
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
	GetUserProfileById(ctx context.Context,bearer string, userId string) (*model.UserProfileResponse, error)
	GetFollowedUsers(ctx context.Context, bearer string, userId string) ([]*model.UserFollowingResponse, error)
	GetFollowingUsers(ctx context.Context, bearer string, userId string) ([]*model.UserFollowingResponse, error)
	GetFollowRequests(ctx context.Context, bearer string) ([]*model.UserFollowingResponse, error)
	FollowUser(ctx context.Context, bearer string, userId string) (bool, error)
	UnfollowUser(ctx context.Context, bearer string, userId string) error
	SearchForUsersByUsername(ctx context.Context, username string, bearer string) ([]model.User, error)
	SearchForUsersByUsernameByGuest(ctx context.Context, username string) ([]model.User, error)
	AcceptFollowRequest(ctx context.Context, bearer string, userId string) error
	UpdateLikedPost(ctx context.Context, bearer string, postId string) error
	UpdateDislikedPost(ctx context.Context, bearer string, postId string) error
	GetUserLikedPost(ctx context.Context, bearer string) ([]string,error)
	GetUserDislikedPost(ctx context.Context, bearer string) ([]string,error)
}
