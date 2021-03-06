package service

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"user-service/domain/model"
	"user-service/domain/repository"
	"user-service/domain/service-contracts"
	"user-service/logger"
	"user-service/tracer"
)

type accountActivationService struct {
	repository.AccountActivationRepository
}

func NewAccountActivationService(r repository.AccountActivationRepository) service_contracts.AccountActivationService {
	return &accountActivationService{r}
}

func (a *accountActivationService) Create(ctx context.Context, userId string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "AccountActivationServiceCreate")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	result, err :=a.AccountActivationRepository.Create(ctx, model.NewAccountActivation(userId))
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : userId}).Warn("Account activation unsuccessful creating")
		return "", err
	}

	return result.InsertedID.(string), err
}
func (a *accountActivationService) UseAccountActivation(ctx context.Context, id string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "AccountActivationServiceUseAccountActivation")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	accActivation, err := a.AccountActivationRepository.GetById(ctx, id)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"activation_id" : id}).Warn("Invalid account activation")
		return "", err
	}
	accActivation.Used = true
	result, _ := a.AccountActivationRepository.Update(ctx, accActivation)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"activation_id" : id}).Warn("Account activation unsuccessful update")
		return "", err
	}
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
	span := tracer.StartSpanFromContext(ctx, "AccountActivationServiceGetValidActivationById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	accActivation, err := a.AccountActivationRepository.GetById(ctx, id)
	if err != nil || !a.IsActivationValid(accActivation) {
		logger.LoggingEntry.WithFields(logrus.Fields{"activation_id" : id}).Warn("Invalid account activation")
		return nil, errors.New("account activation link is not valid")
	}

	return accActivation, err
}