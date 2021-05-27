package repository

import (
	"context"
	"magyAgent/domain/model"
)

type AccountResetPasswordRepository interface {
	Create(ctx context.Context, user *model.AccountResetPassword) (*model.AccountResetPassword, error)
	GetById(ctx context.Context, id string) (*model.AccountResetPassword, error)
	Update(ctx context.Context, a *model.AccountResetPassword) (*model.AccountResetPassword, error)
}
