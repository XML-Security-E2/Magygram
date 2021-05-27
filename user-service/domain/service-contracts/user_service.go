package service_contracts

import (
	"context"
	"user-service/domain/model"
)

type UserService interface {
	RegisterUser(ctx context.Context, user *model.UserRequest) (*model.User, error)
	ActivateUser(ctx context.Context, activationId string) (bool, error)
	DeactivateUser(ctx context.Context, userEmail string) (bool, error)
	AuthenticateUser(ctx context.Context, loginRequest *model.LoginRequest) (*model.User, error)
	HasUserPermission(desiredPermission string, userId string) (bool, error)
	HandleLoginEventAndAccountActivation(ctx context.Context, userEmail string, successful bool, eventType string)
	ResetPassword(ctx context.Context, userEmail string) (bool, error)
	ResetPasswordActivation(ctx context.Context, resetPasswordId string) (bool, error)
	ChangeNewPassword(ctx context.Context, changePasswordRequest *model.ChangeNewPasswordRequest) (bool, error)
	ResendActivationLink(ctx context.Context, activateLinkRequest *model.ActivateLinkRequest) (bool, error)
	GetUserEmailIfUserExist(ctx context.Context, userId string) (*model.User, error)
	GetUserById(ctx context.Context, userId string) (*model.User, error)
}
