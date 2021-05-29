package service_contracts

import (
	"context"
	"user-service/domain/model"
)

type ResetPasswordService interface {
	Create(ctx context.Context, userId string) (string, error)
	UseAccountReset(ctx context.Context, id string) (string, error)
	IsActivationValid(accActivation *model.ResetPassword) bool
	GetValidActivationById(ctx context.Context, id string) (*model.ResetPassword, error)
}