package repository

import (
	"context"
	"magyAgent/domain/model"
)

type LoginEventRepository interface {
	Create(ctx context.Context, event *model.LoginEvent) (*model.LoginEvent, error)
	GetLastByUserEmail(ctx context.Context, email string) (*model.LoginEvent, error)
}