package service

import (
	"context"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
	"magyAgent/domain/service-contracts"
)

type authService struct {
	repository.UserRepository
	service_contracts.AccountActivationService
}

func NewAuthService(r repository.UserRepository, a service_contracts.AccountActivationService) service_contracts.AuthService {
	return &authService{r, a}
}

func (u *authService) RegisterUser(ctx context.Context, userRequest *model.UserRequest) (*model.User, error) {
	user := model.NewUser(userRequest)
	accActivation, _ :=u.AccountActivationService.Create(ctx, user.Id)
	SendMessage(userRequest.Email, userRequest.Name, accActivation.Id)
	return u.UserRepository.Create(ctx, user)
}

func (u *authService) ActivateUser(ctx context.Context, activationId string) (bool, error) {
	accActivation, err := u.AccountActivationService.GetValidActivationById(ctx, activationId)
	if accActivation == nil || err != nil {
		return false, err
	}

	user, err := u.UserRepository.GetByID(ctx, accActivation.UserId)
	if err != nil {
		return false, err
	}
	user.Active = true
	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		return false, err
	}
	return true, err
}
