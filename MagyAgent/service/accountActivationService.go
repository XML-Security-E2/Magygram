package service

import (
	"context"
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

func (a *accountActivationService) IsActivationValid(accActivation *model.AccountActivation) bool {
	t := time.Now()
	if !(accActivation.GenerationDate.Before(t) && accActivation.ExpirationDate.After(t)) {
		return false
	}
	return true
}

func (a *accountActivationService) GetValidActivationById(ctx context.Context, id string) (*model.AccountActivation, error) {
	accActivation, err := a.AccountActivationRepository.GetById(ctx, id)
	if err != nil || !a.IsActivationValid(accActivation) {
		return nil, err
	}

	return accActivation, err
}
