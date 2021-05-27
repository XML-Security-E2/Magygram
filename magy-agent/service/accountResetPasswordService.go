package service

import (
	"context"
	"errors"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
	"magyAgent/domain/service-contracts"
	"time"
)

type accountResetPasswordService struct {
	repository.AccountResetPasswordRepository
}

func NewAccountResetPasswordService(r repository.AccountResetPasswordRepository) service_contracts.AccountResetPasswordService {
	return &accountResetPasswordService{r}
}

func (a accountResetPasswordService) Create(ctx context.Context, userId string) (*model.AccountResetPassword, error) {
	return a.AccountResetPasswordRepository.Create(ctx, model.NewAccountResetPassword(userId))
}

func (a *accountResetPasswordService) UseAccountReset(ctx context.Context, id string) (*model.AccountResetPassword, error) {
	accActivation, err := a.AccountResetPasswordRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	accActivation.Used = true
	return a.AccountResetPasswordRepository.Update(ctx, accActivation)
}

func (a accountResetPasswordService) IsActivationValid(accActivation *model.AccountResetPassword) bool {
	t := time.Now()
	if !(accActivation.GenerationDate.Before(t) && accActivation.ExpirationDate.After(t)) || accActivation.Used {
		return false
	}
	return true
}

func (a accountResetPasswordService) GetValidActivationById(ctx context.Context, id string) (*model.AccountResetPassword, error) {
	accActivation, err := a.AccountResetPasswordRepository.GetById(ctx, id)
	if err != nil || !a.IsActivationValid(accActivation) {
		return nil, errors.New("password reset link is not valid")
	}

	return accActivation, err}

