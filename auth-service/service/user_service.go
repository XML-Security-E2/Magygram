package service

import (
	"auth-service/domain/model"
	"auth-service/domain/repository"
	"auth-service/domain/service-contracts"
	"context"
	"errors"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



type userService struct {
	repository.UserRepository
	repository.LoginEventRepository
}

func NewUserService(r repository.UserRepository,a repository.LoginEventRepository) service_contracts.UserService {
	return &userService{r,a}
}

func (u userService) RegisterUser(ctx context.Context, userRequest *model.UserRequest) (string, error) {
	user, err := model.NewUser(userRequest)
	if err != nil { return "", err}
	if err := validator.New().Struct(user); err!= nil {
		return "", err
	}
	result, err := u.UserRepository.Create(ctx, user)
	if err != nil { return "", err}
	if userId, ok := result.InsertedID.(primitive.ObjectID); ok {
		return userId.String(), nil
	}
	return "", err
}

func (u userService) ActivateUser(ctx context.Context, userId string) (bool, error) {

	user, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return false, err
	}
	user.Active = true
	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		return false, err
	}

	u.HandleLoginEventAndAccountActivation(ctx, user.Email, true, model.ActivatedAccount)

	return true, err
}

func (u userService) HandleLoginEventAndAccountActivation(ctx context.Context, userEmail string, successful bool, eventType string) {
	if successful {
		_, _ = u.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, 0))
		return
	}
	loginEvent, err := u.LoginEventRepository.GetLastByUserEmail(ctx, userEmail)

	if err != nil || loginEvent == nil {
		_, _ = u.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, 1))
		return
	}

	_, _ = u.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, loginEvent.RepetitionNumber+1))
	if loginEvent.RepetitionNumber + 1 > MaxUnsuccessfulLogins {
		_, _ = u.DeactivateUser(ctx, userEmail)
	}
}

func (u userService) DeactivateUser(ctx context.Context, userEmail string) (bool, error) {
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

func (u userService) ResetPassword(ctx context.Context, changePasswordRequest *model.PasswordChangeRequest) (bool, error) {
	hashAndSalt, err := model.HashAndSaltPasswordIfStrongAndMatching(changePasswordRequest.Password, changePasswordRequest.PasswordRepeat)
	if err != nil {
		return false, err
	}

	user, err := u.UserRepository.GetByID(ctx, changePasswordRequest.UserId)
	if err != nil {
		return false, err
	}
	user.Password = hashAndSalt
	user.Active = true
	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		return false, err
	}

	return true, err
}

func (u userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := u.UserRepository.GetByEmail(ctx, email)

	if err != nil {
		return nil, errors.New("invalid user id")
	}

	return user, err
}

func (u userService) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	user, err := u.UserRepository.GetByID(ctx, userId)

	if err != nil {
		return nil, errors.New("invalid user id")
	}

	return user, err
}

func (u userService) GetAllRolesByUserId(ctx context.Context, userId string) ([]model.Role, error) {
	return u.UserRepository.GetAllRolesByUserId(ctx, userId)
}




