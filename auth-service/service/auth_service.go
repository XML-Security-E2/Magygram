package service

import (
	"auth-service/domain/model"
	"auth-service/domain/repository"
	service_contracts "auth-service/domain/service-contracts"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var (
	MaxUnsuccessfulLogins = 3
)

type authService struct {
	repository.LoginEventRepository
	service_contracts.UserService
}

func NewAuthService(r repository.LoginEventRepository,u service_contracts.UserService) service_contracts.AuthService {
	return &authService{r,u }
}

func (a authService) AuthenticateUser(ctx context.Context, loginRequest *model.LoginRequest) (*model.User, error) {
	user, err := a.UserService.GetByEmail(ctx, loginRequest.Email)
	if err != nil {
		return nil, errors.New("invalid email address")
	}
	if !user.Active {
		return user, errors.New("user account is not activated")
	}
	if !equalPasswords(user.Password, loginRequest.Password) {
		a.HandleLoginEventAndAccountActivation(ctx, user.Email, false, model.UnsuccessfulLogin)
		return nil, errors.New("invalid password")
	}
	a.HandleLoginEventAndAccountActivation(ctx, user.Email, true, model.SuccessfulLogin)
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

func (a authService) HasUserPermission(ctx context.Context, desiredPermission string, userId string) (bool, error) {
	roles, err := a.UserService.GetAllRolesByUserId(ctx, userId)
	if err != nil {
		return false, errors.New("invalid email address")
	}
	for _, role := range roles {
		for _, permission := range role.Permissions {
			if permission.Name == desiredPermission {
				return true, err
			}
		}
	}
	return false, err
}

func (a authService) HandleLoginEventAndAccountActivation(ctx context.Context, userEmail string, successful bool, eventType string) {
	if successful {
		_, _ = a.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, 0))
		return
	}
	loginEvent, err := a.LoginEventRepository.GetLastByUserEmail(ctx, userEmail)

	if err != nil || loginEvent == nil {
		_, _ = a.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, 1))
		return
	}

	_, _ = a.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, loginEvent.RepetitionNumber+1))
	if loginEvent.RepetitionNumber + 1 > MaxUnsuccessfulLogins {
		_, _ = a.DeactivateUser(ctx, userEmail)
	}
}


