package repository

import (
	"context"
	"magyAgent/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
}
