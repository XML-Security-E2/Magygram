package service

import (
	"context"
	"errors"
	"github.com/go-playground/validator"
	"user-service/domain/model"
	"user-service/domain/repository"
	"user-service/domain/service-contracts"
	"user-service/domain/service-contracts/exceptions"
	"user-service/service/intercomm"
)

type userService struct {
	repository.UserRepository
	service_contracts.AccountActivationService
	service_contracts.ResetPasswordService
	intercomm.AuthClient
	intercomm.RelationshipClient
	intercomm.PostClient
}

var (
	MaxUnsuccessfulLogins = 3
)

func NewAuthService(r repository.UserRepository, a service_contracts.AccountActivationService, ic intercomm.AuthClient, rp service_contracts.ResetPasswordService, rC intercomm.RelationshipClient, pc intercomm.PostClient) service_contracts.UserService {
	return &userService{r, a,  rp , ic, rC, pc}
}

func (u *userService) EditUser(ctx context.Context, userRequest *model.EditUserRequest) (string, error) {
	user, _ := model.NewEditUser(userRequest)
	if err := validator.New().Struct(user); err!= nil {
		return "", err
	}


	result, err := u.UserRepository.UpdateUserDetails(ctx, user)

	if err != nil { return "", err}

	if userId, ok := result.UpsertedID.(string); ok {
		return userId, nil
	}
	return "", err
}

func (u *userService) RegisterUser(ctx context.Context, userRequest *model.UserRequest) (string, error) {
	user, _ := model.NewUser(userRequest)
	if err := validator.New().Struct(user); err!= nil {
		return "", err
	}

	err := u.AuthClient.RegisterUser(user, userRequest.Password, userRequest.RepeatedPassword)
	if err != nil { return "", err}

	err = u.RelationshipClient.CreateUser(user)
	if err != nil { return "", err}

	accActivationId, _ :=u.AccountActivationService.Create(ctx, user.Id)

	result, err := u.UserRepository.Create(ctx, user)

	if err != nil { return "", err}
	go SendActivationMail(userRequest.Email, userRequest.Name, accActivationId)

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

func (u *userService) SearchForUsersByUsername(ctx context.Context, username string) ([]model.User, error) {
	users, err := u.UserRepository.SearchForUsersByUsername(ctx, username)

	if err != nil {
		return nil, errors.New("Couldn't find any users")
	}

	return users, err
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

func (u *userService) GetUserProfileById(ctx context.Context, userId string) (*model.UserProfileResponse, error) {

	user, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	followingUsers, err := u.RelationshipClient.GetFollowedUsers(userId)
	if err != nil {
		return nil, err
	}

	followedUsers, err := u.RelationshipClient.GetFollowingUsers(userId)

	if err != nil {
		return nil, err
	}

	postsCount, err := u.PostClient.GetUsersPostsCount(userId)
	if err != nil {
		return nil, err
	}

	retVal := &model.UserProfileResponse{
		Username:        user.Username,
		Name:            user.Name,
		Surname:         user.Surname,
		Website:         user.Website,
		Bio:             user.Bio,
		Number:          user.Number,
		Gender:          user.Gender,
		ImageUrl:        user.ImageUrl,
		PostNumber:      postsCount,
		FollowersNumber: len(followedUsers.Users),
		FollowingNumber: len(followingUsers.Users),
	}
	return retVal, nil
}

func (p userService) checkIfUserInfoIsAccessible(bearer string, owner *model.User, loggedUserId string, followedUsers model.FollowedUsersResponse) bool {

	if owner.IsPrivate {
		if bearer == "" {
			return false
		}

		if loggedUserId != owner.Id {
			for _, usrId := range followedUsers.Users {
				if owner.Id == usrId {
					return true
				}
			}
			return false
		}
	}

	return true
}

func (u *userService) GetFollowedUsers(ctx context.Context, bearer string, userId string) ([]*model.UserFollowingResponse, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	owner, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	followingUsers, err := u.RelationshipClient.GetFollowedUsers(loggedId)
	if err != nil {
		return nil, err
	}

	followedUsers, err := u.RelationshipClient.GetFollowingUsers(userId)
	if err != nil {
		return nil, err
	}

	if !u.checkIfUserInfoIsAccessible(bearer, owner, loggedId, followingUsers) {
		return nil, &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
	}

	var userInfos []*model.UserFollowingResponse

	for _, followedId := range followedUsers.Users {
		folUsr, err := u.UserRepository.GetByID(ctx, followedId)
		if err != nil {
			return nil, errors.New("invalid user id")
		}

		following := false
		if bearer != "" {
			following = doesUserFollow(followingUsers, followedId)
		}
		userInfos = append(userInfos, &model.UserFollowingResponse{
			Following: following,
			UserInfo:  &model.UserInfo{
				Id:       followedId,
				Username: folUsr.Username,
				ImageURL: folUsr.ImageUrl,
			},
		})
	}

	return userInfos, nil
}

func doesUserFollow(followingUsers model.FollowedUsersResponse, followedId string) bool {
	for _, folId := range followingUsers.Users {
		if folId == followedId {
			return true
		}
	}
	return false
}

func (u *userService) GetFollowingUsers(ctx context.Context, bearer string, userId string) ([]*model.UserFollowingResponse, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	owner, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	followingUsers, err := u.RelationshipClient.GetFollowedUsers(loggedId)
	if err != nil {
		return nil, err
	}

	followingUsersRet, err := u.RelationshipClient.GetFollowedUsers(userId)
	if err != nil {
		return nil, err
	}

	if !u.checkIfUserInfoIsAccessible(bearer, owner, loggedId, followingUsers) {
		return nil, &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
	}

	var userInfos []*model.UserFollowingResponse

	for _, followedId := range followingUsersRet.Users {
		folUsr, err := u.UserRepository.GetByID(ctx, followedId)
		if err != nil {
			return nil, errors.New("invalid user id")
		}

		following := false
		if bearer != "" {
			following = doesUserFollow(followingUsers, followedId)
		}
		userInfos = append(userInfos, &model.UserFollowingResponse{
			Following: following,
			UserInfo:  &model.UserInfo{
				Id:       followedId,
				Username: folUsr.Username,
				ImageURL: folUsr.ImageUrl,
			},
		})
	}

	return userInfos, nil
}

func (u *userService) FollowUser(ctx context.Context, bearer string, userId string) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	err = u.RelationshipClient.FollowRequest(&model.FollowRequest{
		SubjectId: loggedId,
		ObjectId:  userId,
	})
	return err
}

func (u *userService) UnfollowUser(ctx context.Context, bearer string, userId string) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	err = u.RelationshipClient.Unfollow(&model.FollowRequest{
		SubjectId: loggedId,
		ObjectId:  userId,
	})
	return err
}
