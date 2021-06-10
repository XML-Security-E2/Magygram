package service

import (
	"auth-service/domain/model"
	"auth-service/domain/repository"
	"auth-service/domain/service-contracts"
	"auth-service/logger"
	"context"
	"errors"
	"github.com/pquerna/otp/totp"
	"github.com/sirupsen/logrus"
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
		logger.LoggingEntry.WithFields(logrus.Fields{"email": loginRequest.Email}).Warn("Invalid email address")
		return nil, errors.New("invalid email address")
	}
	if !user.Active {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : user.Id}).Warn("User account not activated")
		return user, errors.New("user account is not activated")
	}
	if !equalPasswords(user.Password, loginRequest.Password) {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : user.Id}).Warn("Wrong password")
		a.HandleLoginEventAndAccountActivation(ctx, user.Email, false, model.UnsuccessfulLogin)
		return nil, errors.New("invalid password")
	}
	a.HandleLoginEventAndAccountActivation(ctx, user.Email, true, model.SuccessfulLogin)
	return user, err
}

func (a authService) AuthenticateTwoFactoryUser(ctx context.Context, loginRequest *model.LoginTwoFactoryRequest) (*model.User, error) {
	user, err := a.UserService.GetByEmail(ctx, loginRequest.Email)

	valid := totp.Validate(loginRequest.Token, user.TotpToken)

	a.HandleLoginEventAndAccountActivation(ctx, user.Email, true, model.SuccessfulLogin)

	if valid{
		return user, err
	}

	return nil, err
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
		_, err := a.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, 0))
		if err != nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_email" : userEmail}).Warn("Create success login event, database failure")
		}
		logger.LoggingEntry.WithFields(logrus.Fields{"user_email" : userEmail}).Info("User logged in successfully")
		return
	}
	loginEvent, err := a.LoginEventRepository.GetLastByUserEmail(ctx, userEmail)

	if err != nil || loginEvent == nil {
		_, err = a.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, 1))
		if err != nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_email" : userEmail}).Warn("Create fail login event, database failure")
		}
		return
	}

	_, err = a.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, loginEvent.RepetitionNumber+1))
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_email" : userEmail}).Warn("Create fail login event, database failure")
	}
	if loginEvent.RepetitionNumber + 1 > MaxUnsuccessfulLogins {
		_, _ = a.DeactivateUser(ctx, userEmail)
	}
}


