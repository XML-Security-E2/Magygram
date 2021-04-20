package service_contracts

import (
	"context"
	"magyAgent/domain/model"
)

type AuthService interface {
	RegisterUser(ctx context.Context, user *model.User) (*model.User, error)
}
