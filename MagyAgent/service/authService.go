package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
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
	go SendActivationMail(userRequest.Email, userRequest.Name, accActivation.Id)
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

func (u *authService) AuthenticateUser(ctx context.Context, loginRequest *model.LoginRequest) (*model.User, error) {
	user, err := u.UserRepository.GetByEmail(ctx, loginRequest.Email)
	if err != nil {
		return nil, errors.New("invalid email address")
	}
	if !user.Active {
		return nil, errors.New("user account is not activated")
	}
	if !equalPasswords(user.Password, loginRequest.Password) {
		return nil, errors.New("invalid password")
	}
	return user, err
}

func equalPasswords(hashedPwd string, passwordRequest string) bool {

	byteHash := []byte(hashedPwd)
	plainPwd := []byte(passwordRequest)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}