package repository

import (
	"context"
	"magyAgent/domain/model"
)

type AccountActivationRepository interface {
	Create(ctx context.Context, user *model.AccountActivation) (*model.AccountActivation, error)
	GetById(ctx context.Context, id string) (*model.AccountActivation, error)
	Update(ctx context.Context, a *model.AccountActivation) (*model.AccountActivation, error)
}
