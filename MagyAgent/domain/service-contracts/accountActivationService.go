package service_contracts

import (
	"context"
	"magyAgent/domain/model"
)

type AccountActivationService interface {
	Create(ctx context.Context, userId string) (*model.AccountActivation, error)
	IsActivationValid(accActivation *model.AccountActivation) bool
	GetValidActivationById(ctx context.Context, id string) (*model.AccountActivation, error)
}
