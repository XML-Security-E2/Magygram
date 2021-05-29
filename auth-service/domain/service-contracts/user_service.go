package service_contracts

import (
	"auth-service/domain/model"
	"context"
)

type UserService interface {
	RegisterUser(ctx context.Context, user *model.UserRequest) (string, error)
	ActivateUser(ctx context.Context, userId string) (bool, error)
	DeactivateUser(ctx context.Context, userEmail string) (bool, error)
	ResetPassword(ctx context.Context, changePasswordRequest *model.ChangeNewPasswordRequest) (bool, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserById(ctx context.Context, userId string) (*model.User, error)
	GetAllRolesByUserId(ctx context.Context, userId string) ([]model.Role, error)
}