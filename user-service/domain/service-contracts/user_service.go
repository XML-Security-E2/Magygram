package service_contracts

import (
	"context"
	"user-service/domain/model"
)

type UserService interface {
	RegisterUser(ctx context.Context, user *model.UserRequest) (string, error)
	EditUser(ctx context.Context, user *model.EditUserRequest) (string, error)
	ActivateUser(ctx context.Context, activationId string) (bool, error)
	ResetPassword(ctx context.Context, userEmail string) (bool, error)
	ResetPasswordActivation(ctx context.Context, resetPasswordId string) (bool, error)
	ChangeNewPassword(ctx context.Context, changePasswordRequest *model.ChangeNewPasswordRequest) (bool, error)
	ResendActivationLink(ctx context.Context, activateLinkRequest *model.ActivateLinkRequest) (bool, error)
	GetUserEmailIfUserExist(ctx context.Context, userId string) (*model.User, error)
	GetUserById(ctx context.Context, userId string) (*model.User, error)
	GetLoggedUserInfo(ctx context.Context, bearer string) (*model.UserInfo, error)
	SearchForUsersByUsername(ctx context.Context, username string, bearer string) ([]model.User, error)
	SearchForUsersByUsernameByGuest(ctx context.Context, username string) ([]model.User, error)
}
