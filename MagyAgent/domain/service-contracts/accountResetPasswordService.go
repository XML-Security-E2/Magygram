package service_contracts

import (
	"context"
	"magyAgent/domain/model"
)

type AccountResetPasswordService interface {
	Create(ctx context.Context, userId string) (*model.AccountResetPassword, error)
	UseAccountReset(ctx context.Context, id string) (*model.AccountResetPassword, error)
	IsActivationValid(accActivation *model.AccountResetPassword) bool
	GetValidActivationById(ctx context.Context, id string) (*model.AccountResetPassword, error)
}
