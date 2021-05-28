package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"user-service/domain/model"
	"user-service/domain/repository"
	"user-service/domain/service-contracts"
)

type userService struct {
	repository.UserRepository
	service_contracts.AccountActivationService
	repository.LoginEventRepository
	service_contracts.ResetPasswordService
}

var (
	MaxUnsuccessfulLogins = 3
)

func NewAuthService(r repository.UserRepository, a service_contracts.AccountActivationService, l repository.LoginEventRepository, rp service_contracts.ResetPasswordService) service_contracts.UserService {
	return &userService{r, a, l, rp }
}

func (u *userService) RegisterUser(ctx context.Context, userRequest *model.UserRequest) (string, error) {
	user, err := model.NewUser(userRequest)
	if err != nil { return "", err}
	if err := validator.New().Struct(user); err!= nil {
		return "", err
	}
	accActivationId, _ :=u.AccountActivationService.Create(ctx, user.Id)
	fmt.Println(accActivationId)
	result, err := u.UserRepository.Create(ctx, user)
	if err != nil { return "", err}
	go SendActivationMail(userRequest.Email, userRequest.Name, accActivationId)
	if userId, ok := result.InsertedID.(primitive.ObjectID); ok {
		return userId.String(), nil
	}
	return "", err
}

func (u *userService) ActivateUser(ctx context.Context, activationId string) (bool, error) {

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

	_, err = u.UseAccountActivation(ctx, activationId)
	if err != nil {
		return false, err
	}

	u.HandleLoginEventAndAccountActivation(ctx, user.Email, true, model.ActivatedAccount)

	return true, err
}

func (u *userService) DeactivateUser(ctx context.Context, userEmail string) (bool, error){

	user, err := u.UserRepository.GetByEmail(ctx, userEmail)
	if err != nil {
		return false, err
	}
	user.Active = false
	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		return false, err
	}
	return true, err
}

func (u *userService) ResendActivationLink(ctx context.Context, activateLinkRequest *model.ActivateLinkRequest) (bool, error) {
	user, err := u.UserRepository.GetByEmail(ctx, activateLinkRequest.Email)
	if err != nil {
		return false, err
	}

	accActivationId, _ := u.AccountActivationService.Create(ctx, user.Id)
	go SendActivationMail(user.Email, user.Name, accActivationId)

	return true, nil
}

func (u *userService) AuthenticateUser(ctx context.Context, loginRequest *model.LoginRequest) (*model.User, error) {
	user, err := u.UserRepository.GetByEmail(ctx, loginRequest.Email)
	if err != nil {
		return nil, errors.New("invalid email address")
	}
	if !user.Active {
		return user, errors.New("user account is not activated")
	}
	if !equalPasswords(user.Password, loginRequest.Password) {
		u.HandleLoginEventAndAccountActivation(ctx, user.Email, false, model.UnsuccessfulLogin)
		return nil, errors.New("invalid password")
	}
	u.HandleLoginEventAndAccountActivation(ctx, user.Email, true, model.SuccessfulLogin)
	return user, err
}

func (u *userService) HandleLoginEventAndAccountActivation(ctx context.Context, userEmail string, successful bool, eventType string) {
	if successful {
		u.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, 0))
		return
	}
	loginEvent, err := u.LoginEventRepository.GetLastByUserEmail(ctx, userEmail)

	if err != nil || loginEvent == nil {
		u.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, 1))
		return
	}

	_, _ = u.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, loginEvent.RepetitionNumber+1))
	if loginEvent.RepetitionNumber + 1 > MaxUnsuccessfulLogins {
		u.DeactivateUser(ctx, userEmail)
	}
}

func (u *userService) HasUserPermission(ctx context.Context, desiredPermission string, userId string) (bool, error) {
	roles, err := u.UserRepository.GetAllRolesByUserId(ctx, userId)
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

func (u *userService) ResetPassword(ctx context.Context, userEmail string) (bool, error) {
	user, err := u.GetByEmail(ctx,userEmail)
	//pokrivena invalid email
	if err != nil {
		return false, errors.New("invalid email address")
	}

	accResetPasswordId, _ := u.ResetPasswordService.Create(ctx, user.Id)
	go SendResetPasswordMail(user.Email, user.Name, accResetPasswordId)

	return true, nil
}

func (u *userService) ResetPasswordActivation(ctx context.Context, activationId string) (bool, error) {
	accReset, err := u.ResetPasswordService.GetValidActivationById(ctx, activationId)
	if accReset == nil || err != nil {
		return false, err
	}

	return true, err
}

func (u *userService) ChangeNewPassword(ctx context.Context, changePasswordRequest *model.ChangeNewPasswordRequest) (bool, error) {
	hashAndSalt, err := model.HashAndSaltPasswordIfStrongAndMatching(changePasswordRequest.Password, changePasswordRequest.PasswordRepeat)
	if err != nil {
		return false, err
	}

	accReset, err := u.ResetPasswordService.GetValidActivationById(ctx, changePasswordRequest.ResetPasswordId)
	if accReset == nil || err != nil {
		return false, err
	}

	user, err := u.UserRepository.GetByID(ctx, accReset.UserId)
	if err != nil {
		return false, err
	}
	user.Password = hashAndSalt
	user.Active = true
	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		return false, err
	}

	_, err = u.UseAccountReset(ctx, changePasswordRequest.ResetPasswordId)
	if err != nil {
		return false, err
	}

	return true, err
}



func (u *userService) GetUserEmailIfUserExist(ctx context.Context, userId string) (*model.User, error) {
	user, err := u.UserRepository.GetByID(ctx, userId)

	if err != nil {
		return nil, errors.New("invalid user id")
	}

	return user, err
}

func (u *userService) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	user, err := u.UserRepository.GetByID(ctx, userId)

	if err != nil {
		return nil, errors.New("invalid user id")
	}

	return user, err
}
