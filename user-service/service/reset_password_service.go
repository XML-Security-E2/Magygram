package service

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"user-service/domain/model"
	"user-service/domain/repository"
	service_contracts "user-service/domain/service-contracts"
	"user-service/logger"
)

type resetPasswordService struct {
	repository.ResetPasswordRepository
}

func NewResetPasswordService(r repository.ResetPasswordRepository) service_contracts.ResetPasswordService {
	return &resetPasswordService{r}
}

func (a resetPasswordService) Create(ctx context.Context, userId string) (string, error) {
	result, err :=a.ResetPasswordRepository.Create(ctx, model.NewResetPassword(userId))
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : userId}).Warn("Account password reset unsuccessful creating")
		return "", err}

	return result.InsertedID.(string), err
}

func (a *resetPasswordService) UseAccountReset(ctx context.Context, id string) (string, error) {
	accActivation, err := a.ResetPasswordRepository.GetById(ctx, id)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"reset_password_id" : id}).Warn("Invalid account password reset")
		return "", err
	}
	accActivation.Used = true
	result, _ := a.ResetPasswordRepository.Update(ctx, accActivation)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"reset_password_id" : id}).Warn("Account password reset unsuccessful update")
		return "", err
	}
	if resetPasswordId, ok := result.UpsertedID.(primitive.ObjectID); ok {
		return resetPasswordId.String(), nil
	}
	return "", err
}

func (a resetPasswordService) IsActivationValid(accActivation *model.ResetPassword) bool {
	t := time.Now()
	if !(accActivation.GenerationDate.Before(t) && accActivation.ExpirationDate.After(t)) || accActivation.Used {
		return false
	}
	return true
}

func (a resetPasswordService) GetValidActivationById(ctx context.Context, id string) (*model.ResetPassword, error) {
	accActivation, err := a.ResetPasswordRepository.GetById(ctx, id)
	if err != nil || !a.IsActivationValid(accActivation) {
		logger.LoggingEntry.WithFields(logrus.Fields{"reset_password_id" : id}).Warn("Invalid reset password link")
		return nil, errors.New("password reset link is not valid")
	}

	return accActivation, err}
