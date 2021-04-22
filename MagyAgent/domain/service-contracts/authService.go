package service_contracts

import (
	"context"
	"magyAgent/domain/model"
)

type AuthService interface {
	RegisterUser(ctx context.Context, user *model.UserRequest) (*model.User, error)
	ActivateUser(ctx context.Context, activationId string) (bool, error)
	AuthenticateUser(ctx context.Context, loginRequest *model.LoginRequest) (*model.User, error)
	HasUserPermission(desiredPermission string, userId string) (bool, error)
}
