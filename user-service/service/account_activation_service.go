package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"user-service/domain/model"
	"user-service/domain/repository"
	"user-service/domain/service-contracts"
)

type accountActivationService struct {
	repository.AccountActivationRepository
}

func NewAccountActivationService(r repository.AccountActivationRepository) service_contracts.AccountActivationService {
	return &accountActivationService{r}
}

func (a *accountActivationService) Create(ctx context.Context, userId string) (string, error) {
	result, err :=a.AccountActivationRepository.Create(ctx, model.NewAccountActivation(userId))
	if err != nil { return "", err}

	return result.InsertedID.(string), err
}
func (a *accountActivationService) UseAccountActivation(ctx context.Context, id string) (string, error) {
	accActivation, err := a.AccountActivationRepository.GetById(ctx, id)
	if err != nil {
		return "", err
	}
	accActivation.Used = true
	result, _ := a.AccountActivationRepository.Update(ctx, accActivation)
	if err != nil { return "", err}
	if activationId, ok := result.UpsertedID.(primitive.ObjectID); ok {
		return activationId.String(), nil
	}
	return "", err
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