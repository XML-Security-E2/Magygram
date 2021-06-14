package service

import (
	"context"
	"errors"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
	service_contracts "magyAgent/domain/service-contracts"
	"time"
)

type accountActivationService struct {
	repository.AccountActivationRepository
}

func NewAccountActivationService(r repository.AccountActivationRepository) service_contracts.AccountActivationService {
	return &accountActivationService{r}
}

func (a *accountActivationService) Create(ctx context.Context, userId string) (*model.AccountActivation, error) {
	return a.AccountActivationRepository.Create(ctx, model.NewAccountActivation(userId))
}
func (a *accountActivationService) UseAccountActivation(ctx context.Context, id string) (*model.AccountActivation, error) {
	accActivation, err := a.AccountActivationRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	accActivation.Used = true
	return a.AccountActivationRepository.Update(ctx, accActivation)
}

func (a *accountActivationService) IsActivationValid(accActivation *model.AccountActivation) bool {
	t := time.Now()
	if !(accActivation.GenerationDate.Before(t) && accActivation.ExpirationDate.After(t)) || accActivation.Used {
		return false
	}
	return true
}

func (a *accountActivationService) GetValidActivationById(ctx context.Context, id string) (*model.AccountActivation, error) {
	accActivation, err := a.AccountActivationRepository.GetById(ctx, id)
	if err != nil || !a.IsActivationValid(accActivation) {
		return nil, errors.New("account activation link is not valid")
	}

	return accActivation, err
}
