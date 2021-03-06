package service

import (
	"context"
	"errors"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"log"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
	"magyAgent/domain/service-contracts"
)

type authService struct {
	repository.UserRepository
	service_contracts.AccountActivationService
	repository.LoginEventRepository
	service_contracts.AccountResetPasswordService
}


var (
	MAX_UNSUCCESSFUL_LOGINS = 3
)

func NewAuthService(r repository.UserRepository, a service_contracts.AccountActivationService, l repository.LoginEventRepository, rp service_contracts.AccountResetPasswordService) service_contracts.AuthService {
	return &authService{r, a, l, rp }
}

func (u *authService) RegisterUser(ctx context.Context, userRequest *model.UserRequest) (*model.User, error) {
	user, err := model.NewUser(userRequest)
	if err != nil { return nil, err}
	if error := validator.New().Struct(user); error!= nil {
		return nil, error
	}
	accActivation, _ :=u.AccountActivationService.Create(ctx, user.Id)
	user, err = u.UserRepository.Create(ctx, user)
	if err != nil { return nil, err}
	go SendActivationMail(userRequest.Email, userRequest.Name, accActivation.Id)
	return user, err
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

	_, err = u.UseAccountActivation(ctx, activationId)
	if err != nil {
		return false, err
	}

	u.HandleLoginEventAndAccountActivation(ctx, user.Email, true, model.ActivatedAccount)

	return true, err
}

func (u *authService) DeactivateUser(ctx context.Context, userEmail string) (bool, error){

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

func (u *authService) ResendActivationLink(ctx context.Context, activateLinkRequest *model.ActivateLinkRequest) (bool, error) {
	user, err := u.UserRepository.GetByEmail(ctx, activateLinkRequest.Email)
	if err != nil {
		return false, err
	}

	accActivation, _ := u.AccountActivationService.Create(ctx, user.Id)
	go SendActivationMail(user.Email, user.Name, accActivation.Id)

	return true, nil
}

func (u *authService) AuthenticateUser(ctx context.Context, loginRequest *model.LoginRequest) (*model.User, error) {
	user, err := u.UserRepository.GetByEmailEagerly(ctx, loginRequest.Email)
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

func (u *authService) HandleLoginEventAndAccountActivation(ctx context.Context, userEmail string, successful bool, eventType string) {
	if successful {
		u.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, 0))
		return
	}
	loginEvent, err := u.LoginEventRepository.GetLastByUserEmail(ctx, userEmail)

	if err != nil || loginEvent == nil {
		u.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, 1))
		return
	}

	newLoginEvent, _ := u.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, loginEvent.RepetitionNumber + 1))
	if newLoginEvent.RepetitionNumber > MAX_UNSUCCESSFUL_LOGINS {
		u.DeactivateUser(ctx, userEmail)
	}
}

func (u *authService) HasUserPermission(desiredPermission string, userId string) (bool, error) {
	roles, err := u.UserRepository.GetAllRolesByUserId(userId)
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

func (u *authService) ResetPassword(ctx context.Context, userEmail string) (bool, error) {
	user, err := u.GetByEmail(ctx,userEmail)
	//pokrivena invalid email
	if err != nil {
		return false, errors.New("invalid email address")
	}

	accResetPassword, _ :=u.AccountResetPasswordService.Create(ctx, user.Id)

	go SendResetPasswordMail(user.Email, user.Name, accResetPassword.Id)

	return true, nil
}

func (u *authService) ResetPasswordActivation(ctx context.Context, activationId string) (bool, error) {
	accReset, err := u.AccountResetPasswordService.GetValidActivationById(ctx, activationId)
	if accReset == nil || err != nil {
		return false, err
	}

	return true, err
}

func (u *authService) ChangeNewPassword(ctx context.Context, changePasswordRequest *model.ChangeNewPasswordRequest) (bool, error) {
	hashAndSalt, err := model.HashAndSaltPasswordIfStrongAndMatching(changePasswordRequest.Password, changePasswordRequest.PasswordRepeat)
	if err != nil {
		return false, err
	}

	accReset, err := u.AccountResetPasswordService.GetValidActivationById(ctx, changePasswordRequest.ResetPasswordId)
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



func (u *authService) GetUserEmailIfUserExist(ctx context.Context, userId string) (*model.User, error) {
	user, err := u.UserRepository.GetByID(ctx, userId)

	if err != nil {
		return nil, errors.New("invalid user id")
	}

	return user, err
}

func (u *authService) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	user, err := u.UserRepository.GetByID(ctx, userId)

	if err != nil {
		return nil, errors.New("invalid user id")
	}

	return user, err
}