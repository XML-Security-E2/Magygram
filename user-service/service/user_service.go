package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"user-service/domain/model"
	"user-service/domain/repository"
	"user-service/domain/service-contracts"
	"user-service/service/intercomm"
)

type userService struct {
	repository.UserRepository
	service_contracts.AccountActivationService
	service_contracts.ResetPasswordService
	intercomm.AuthClient
}

var (
	MaxUnsuccessfulLogins = 3
)

func NewAuthService(r repository.UserRepository, a service_contracts.AccountActivationService, ic intercomm.AuthClient, rp service_contracts.ResetPasswordService) service_contracts.UserService {
	return &userService{r, a,  rp , ic}
}

func (u *userService) RegisterUser(ctx context.Context, userRequest *model.UserRequest) (string, error) {
	user := model.NewUser(userRequest)
	if err := validator.New().Struct(user); err!= nil {
		return "", err
	}

	err := u.AuthClient.RegisterUser(user, userRequest.Password, userRequest.RepeatedPassword)
	if err != nil { return "", err}

	accActivationId, _ :=u.AccountActivationService.Create(ctx, user.Id)

	result, err := u.UserRepository.Create(ctx, user)

	if err != nil { return "", err}
	go SendActivationMail(userRequest.Email, userRequest.Name, accActivationId)
	fmt.Println(result.InsertedID.(string))
	if userId, ok := result.InsertedID.(string); ok {
		return userId, nil
	}
	return "", err
}

func (u *userService) ActivateUser(ctx context.Context, activationId string) (bool, error) {

	accActivation, err := u.AccountActivationService.GetValidActivationById(ctx, activationId)
	if accActivation == nil || err != nil {
		return false, err
	}

	err = u.AuthClient.ActivateUser(accActivation.UserId)
	if err != nil {
		return false, err
	}

	_, err = u.UseAccountActivation(ctx, activationId)
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

	accReset, err := u.ResetPasswordService.GetValidActivationById(ctx, changePasswordRequest.ResetPasswordId)
	if accReset == nil || err != nil {
		return false, err
	}

	user, err := u.UserRepository.GetByID(ctx, accReset.UserId)
	if err != nil {
		return false, err
	}

	err = 	u.AuthClient.ChangePassword(user.Id, changePasswordRequest.Password, changePasswordRequest.PasswordRepeat)
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

func (u *userService) GetLoggedUserInfo(ctx context.Context, bearer string) (*model.UserInfo, error) {

	userId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	user, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	return &model.UserInfo{
		Id:       userId,
		Username: user.Username,
		ImageURL: "",
	}, nil
}
