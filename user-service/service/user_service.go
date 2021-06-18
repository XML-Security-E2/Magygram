package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
	"user-service/domain/model"
	"user-service/domain/repository"
	"user-service/domain/service-contracts"
	"user-service/domain/service-contracts/exceptions"
	"user-service/logger"
	"user-service/service/intercomm"
)

type userService struct {
	repository.UserRepository
	service_contracts.AccountActivationService
	service_contracts.ResetPasswordService
	intercomm.AuthClient
	intercomm.RelationshipClient
	intercomm.PostClient
	intercomm.MediaClient
	intercomm.MessageClient
}

var (
	MaxUnsuccessfulLogins = 3
)

func NewAuthService(r repository.UserRepository, a service_contracts.AccountActivationService, ic intercomm.AuthClient, rp service_contracts.ResetPasswordService, rC intercomm.RelationshipClient, pc intercomm.PostClient, mc intercomm.MediaClient, msc intercomm.MessageClient) service_contracts.UserService {
	return &userService{r, a,  rp , ic, rC, pc, mc, msc}
}

func (u *userService) GetUsersForPostNotification(ctx context.Context, userId string) ([]*model.UserInfo, error) {

	followingUsers, err := u.RelationshipClient.GetFollowingUsers(userId)
	if err != nil {
		return []*model.UserInfo{}, err
	}

	var retVal []*model.UserInfo

	for _, followingUserId := range followingUsers.Users {
		followingUser, err := u.UserRepository.GetByID(ctx, followingUserId)
		if err == nil && followingUser.NotificationSettings.NotifyPost {
			retVal = append(retVal, &model.UserInfo{
				Id:       followingUserId,
				Username: followingUser.Username,
				ImageURL: followingUser.ImageUrl,
			})
		}
	}

	return retVal, nil
}

func (u *userService) GetUsersForStoryNotification(ctx context.Context, userId string) ([]*model.UserInfo, error) {

	followingUsers, err := u.RelationshipClient.GetFollowedUsers(userId)
	if err != nil {
		return []*model.UserInfo{}, err
	}

	var retVal []*model.UserInfo

	for _, followingUserId := range followingUsers.Users {
		followingUser, err := u.UserRepository.GetByID(ctx, followingUserId)
		if err == nil && followingUser.NotificationSettings.NotifyStory {
			retVal = append(retVal, &model.UserInfo{
				Id:       followingUserId,
				Username: followingUser.Username,
				ImageURL: followingUser.ImageUrl,
			})
		}
	}

	return retVal, nil}

func (u *userService) CheckIfPostInteractionNotificationEnabled(ctx context.Context, userId string, interactionType string) (bool, error) {
	user, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return false, errors.New("invalid user id")
	}

	if interactionType == "like" {
		return user.NotificationSettings.NotifyLike, nil
	} else if interactionType == "dislike" {
		return user.NotificationSettings.NotifyDislike, nil
	} else if interactionType == "comment" {
		return user.NotificationSettings.NotifyComment, nil
	} else if interactionType == "follow" {
		return user.NotificationSettings.NotifyFollow, nil
	} else if interactionType == "follow-request" {
		return user.NotificationSettings.NotifyFollowRequest, nil
	} else if interactionType == "accepted-follow-request" {
		return user.NotificationSettings.NotifyAcceptFollowRequest, nil
	}
	return false, nil
}

func (u *userService) EditUser(ctx context.Context, bearer string, userId string, userRequest *model.EditUserRequest) (string, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return "", err
	}

	if loggedId != userId {
		logger.LoggingEntry.WithFields(logrus.Fields{"requested_user_id" : userId, "logged_user_id" : loggedId}).Warn("Unauthorized access")
		return "", &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
	}

	user, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return "", errors.New("invalid user id")
	}

	user.Username = userRequest.Username
	user.Name = userRequest.Name
	user.Surname = userRequest.Surname
	user.Number = userRequest.Number
	user.Website = userRequest.Website
	user.Bio = userRequest.Bio
	user.Gender = userRequest.Gender
	if err = validator.New().Struct(user); err!= nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"name": userRequest.Name,
													 "surname" : userRequest.Surname,
													 "number" : userRequest.Number,
													 "gender" : userRequest.Gender,
													 "username" : userRequest.Username,
													 "website" : userRequest.Website,
													 "bio": userRequest.Bio}).Warn("User registration validation failure")
		return "", err
	}

	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"name": userRequest.Name,
													 "surname" : userRequest.Surname,
													 "number" : userRequest.Number,
													 "gender" : userRequest.Gender,
													 "username" : userRequest.Username,
													 "website" : userRequest.Website,
													 "bio": userRequest.Bio}).Error("User database update failure")
		return "", err}

	logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Info("User information updated")

	return user.Id, err
}

func (u *userService) EditUserImage(ctx context.Context, bearer string, userId string, userImage []*multipart.FileHeader) (string, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return "", err
	}

	if loggedId != userId {
		logger.LoggingEntry.WithFields(logrus.Fields{"requested_user_id" : userId, "logged_user_id" : loggedId}).Warn("Unauthorized access")
		return "", &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
	}

	user, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return "", errors.New("invalid user id")
	}

	media, err := u.MediaClient.SaveMedia(userImage)
	if err != nil { return "", err}

	if len(media) == 0 {
		return "", errors.New("error while saving image")
	}
	user.ImageUrl = media[0].Url

	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("User profile picture update failure")
		return "", err}

	logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Info("User profile picture updated")

	return media[0].Url ,err
}

func (u *userService) EditUsersNotifications(ctx context.Context, bearer string, notificationReq *model.NotificationSettings) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		return errors.New("invalid user id")
	}

	user.NotificationSettings = *notificationReq
	_, err = u.UserRepository.Update(ctx, user)

	if err != nil {
		return err
	}
	return nil
}

func (u *userService) EditUsersPrivacySettings(ctx context.Context, bearer string, privacySettingsReq *model.PrivacySettings) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		return errors.New("invalid user id")
	}

	user.PrivacySettings = *privacySettingsReq
	_, err = u.UserRepository.Update(ctx, user)

	if err != nil {
		return err


	}
	return nil
}

func (u *userService) RegisterUser(ctx context.Context, userRequest *model.UserRequest) (*http.Response, error) {
	user, _ := model.NewUser(userRequest)
	if err := validator.New().Struct(user); err!= nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"name": userRequest.Name, "surname" : userRequest.Surname, "email" : userRequest.Email, "username" : userRequest.Username}).Warn("User registration validation failure")
		return nil, err
	}

	result, err := u.UserRepository.Create(ctx, user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"name": userRequest.Name, "surname" : userRequest.Surname, "email" : userRequest.Email, "username" : userRequest.Username}).Error("User database create failure")
		return nil, err
	}

	resp, err := u.AuthClient.RegisterUser(user, userRequest.Password, userRequest.RepeatedPassword)
	if err != nil {
		return nil, err
	}

	err = u.RelationshipClient.CreateUser(user)
	if err != nil {
		return nil, err
	}

	accActivationId, _ :=u.AccountActivationService.Create(ctx, user.Id)


	go SendActivationMail(userRequest.Email, userRequest.Name, accActivationId)

	if userId, ok := result.InsertedID.(string); ok {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : userId}).Info("User registered")
		return resp, nil
	}
	return resp, err
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

	logger.LoggingEntry.WithFields(logrus.Fields{"activation_id" : activationId, "user_id" : accActivation.UserId}).Info("User activated")

	return true, err
}

func (u *userService) ResendActivationLink(ctx context.Context, activateLinkRequest *model.ActivateLinkRequest) (bool, error) {
	user, err := u.UserRepository.GetByEmail(ctx, activateLinkRequest.Email)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email" : activateLinkRequest.Email}).Warn("Invalid email address")
		return false, err
	}

	accActivationId, _ := u.AccountActivationService.Create(ctx, user.Id)
	go SendActivationMail(user.Email, user.Name, accActivationId)

	logger.LoggingEntry.WithFields(logrus.Fields{"activation_id" : accActivationId, "user_id" : user.Id}).Info("Account activation link created")
	return true, nil
}

func (u *userService) ResetPassword(ctx context.Context, userEmail string) (bool, error) {
	user, err := u.GetByEmail(ctx,userEmail)
	//pokrivena invalid email
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email" : userEmail}).Warn("Invalid email address")
		return false, errors.New("invalid email address")
	}

	accResetPasswordId, _ := u.ResetPasswordService.Create(ctx, user.Id)
	go SendResetPasswordMail(user.Email, user.Name, accResetPasswordId)

	logger.LoggingEntry.WithFields(logrus.Fields{"reset_password_id" : accResetPasswordId, "user_id" : user.Id}).Info("Reset password link created")
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

	err = u.AuthClient.ChangePassword(user.Id, changePasswordRequest.Password, changePasswordRequest.PasswordRepeat)
	if err != nil {
		return false, err
	}

	_, err = u.UseAccountReset(ctx, changePasswordRequest.ResetPasswordId)
	if err != nil {
		return false, err
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"reset_password_id" : changePasswordRequest.ResetPasswordId, "user_id" : user.Id}).Info("Users password changed")
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
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : userId}).Warn("Invalid user id")
		return nil, errors.New("invalid user id")
	}

	return user, err
}

func (u *userService) SearchForUsersByUsername(ctx context.Context, username string, bearer string) ([]model.User, error) {
	userId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	users, err := u.UserRepository.SearchForUsersByUsername(ctx, username,userId)

	if err != nil {
		return nil, errors.New("Couldn't find any users")
	}

	return users, err
}

func (u *userService) SearchForUsersByUsernameByGuest(ctx context.Context, username string) ([]model.User, error) {
	users, err := u.UserRepository.SearchForUsersByUsernameByGuest(ctx, username)

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
		ImageURL: user.ImageUrl,
	}, nil
}

func (u *userService) GetUserProfileById(ctx context.Context,bearer string, userId string) (*model.UserProfileResponse, error) {

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

	loggedId, _ := u.AuthClient.GetLoggedUserId(bearer)
	following := false
	if loggedId != "" {
		following = doesUserFollow(followedUsers, loggedId)
	}
	sentReq := false
	if userId != loggedId {
		sentReq, _ = u.RelationshipClient.ReturnFollowRequestsForUser(bearer, userId)
	}

	fmt.Println(sentReq)
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
		Following: 		 following,
		Email:			 user.Email,
		FollowersNumber: len(followedUsers.Users),
		FollowingNumber: len(followingUsers.Users),
		SentFollowRequest: sentReq,
		NotificationSettings: user.NotificationSettings,
		PrivacySettings: user.PrivacySettings,
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

func (u *userService) FollowUser(ctx context.Context, bearer string, userId string) (bool, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return false, err
	}

	user, _ := u.UserRepository.GetByID(ctx, loggedId)

	followRequest, err := u.RelationshipClient.FollowRequest(&model.FollowRequest{
		SubjectId: loggedId,
		ObjectId:  userId,
	})
	if err == nil {
		if followRequest {
			err = u.MessageClient.CreateNotification(&intercomm.NotificationRequest{
				Username:  user.Username,
				UserId:    userId,
				NotifyUrl: "TODO",
				ImageUrl:  user.ImageUrl,
				Type:      intercomm.FollowRequest,
			})
		} else {
			err = u.MessageClient.CreateNotification(&intercomm.NotificationRequest{
				Username:  user.Username,
				UserId:    userId,
				NotifyUrl: "TODO",
				ImageUrl:  user.ImageUrl,
				Type:      intercomm.Followed,
			})
		}
	}

	return followRequest, err
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

func (u *userService) MuteUser(ctx context.Context, bearer string, userId string) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	err = u.RelationshipClient.Mute(&model.Mute{
		SubjectId: loggedId,
		ObjectId:  userId,
	})

	return err
}

func (u *userService) UnmuteUser(ctx context.Context, bearer string, userId string) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	err = u.RelationshipClient.Unmute(&model.Mute{
		SubjectId: loggedId,
		ObjectId:  userId,
	})

	return err
}

func (u *userService) BlockUser(ctx context.Context, bearer string, userId string) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}
	user, err := u.UserRepository.GetByID(ctx, loggedId)
	user.BlockedUsers = append(user.BlockedUsers, userId)
	if _, err = u.UserRepository.Update(ctx, user); err != nil {
		return err
	}
	if err = u.RelationshipClient.Unfollow(&model.FollowRequest{SubjectId: loggedId, ObjectId: userId}); err != nil {
		return err
	}
	if err = u.RelationshipClient.Unfollow(&model.FollowRequest{SubjectId: userId, ObjectId: loggedId}); err != nil {
		return err
	}
	return err
}

func (u *userService) UnblockUser(ctx context.Context, bearer string, userId string) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}
	user, err := u.UserRepository.GetByID(ctx, loggedId)
	result, index := isUserBlocked(user, userId)
	if result {
		user.BlockedUsers = append(user.BlockedUsers[:index], user.BlockedUsers[index+1:]...)
		u.UserRepository.Update(ctx, user)
	}

	return err
}

func isUserBlocked(user *model.User, userId string) (bool, int) {
	for index, blockedUserId := range user.BlockedUsers{
		if blockedUserId == userId {
			return true, index
		}
	}
	return false ,0
}

func (u *userService) GetFollowRequests(ctx context.Context, bearer string) ([]*model.UserFollowingResponse, error) {
	requestsFrom, err := u.RelationshipClient.ReturnFollowRequests(bearer)
	if err != nil {
		return nil, err
	}

	var userInfos []*model.UserFollowingResponse

	fmt.Println(len(requestsFrom.Users))
	for _, followedId := range requestsFrom.Users {
		folUsr, err := u.UserRepository.GetByID(ctx, followedId)
		if err != nil {
			return nil, errors.New("invalid user id")
		}

		userInfos = append(userInfos, &model.UserFollowingResponse{
			Following: false,
			UserInfo:  &model.UserInfo{
				Id:       followedId,
				Username: folUsr.Username,
				ImageURL: folUsr.ImageUrl,
			},
		})
	}

	return userInfos, nil
}

func (u *userService) AcceptFollowRequest(ctx context.Context, bearer string, userId string) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	user, _ := u.UserRepository.GetByID(ctx, loggedId)

	err = u.RelationshipClient.AcceptFollowRequest(bearer, userId)
	if err != nil {
		return err
	}

	err = u.MessageClient.CreateNotification(&intercomm.NotificationRequest{
		Username:  user.Username,
		UserId:    userId,
		NotifyUrl: "TODO",
		ImageUrl:  user.ImageUrl,
		Type:      intercomm.AcceptedFollowRequest,
	})

	return nil
}

func (u *userService) UpdateLikedPost(ctx context.Context, bearer string, postId string) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		errors.New("invalid user id")
	}

	var result, index = didUserLikedPost(user,postId)

	if result{
		user.LikedPosts = append(user.LikedPosts[:index], user.LikedPosts[index+1:]...)
	}else{
		user.LikedPosts = append(user.LikedPosts, postId)
	}

	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		errors.New("user not modified")
	}

	return nil
}

func (u *userService) UpdateDislikedPost(ctx context.Context, bearer string, postId string) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		errors.New("invalid user id")
	}

	var result, index = didUserDislikedPost(user,postId)

	if result{
		user.DislikedPosts = append(user.DislikedPosts[:index], user.DislikedPosts[index+1:]...)
	}else{
		user.DislikedPosts = append(user.DislikedPosts, postId)
	}

	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		errors.New("user not modified")
	}

	return nil}

func didUserDislikedPost(user *model.User, postId string) (bool, int) {
	for index,dislikedPostId := range user.DislikedPosts{
		if dislikedPostId==postId{
			return true, index
		}
	}
	return false ,0
}

func didUserLikedPost(user *model.User, postId string) (bool, int) {
	for index,likedPostId := range user.LikedPosts{
		if likedPostId==postId{
			return true, index
		}
	}
	return false ,0
}

func (u *userService) GetUserLikedPost(ctx context.Context, bearer string) ([]string, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return []string{},err
	}

	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		return []string{},errors.New("invalid user id")
	}

	return user.LikedPosts,nil
}

func (u *userService) GetUserDislikedPost(ctx context.Context, bearer string) ([]string, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return []string{},err
	}

	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		return []string{},errors.New("invalid user id")
	}

	return user.DislikedPosts,nil
}

func (u *userService) VerifyUser(ctx context.Context, dto *model.VerifyAccountDTO) error {
	user, err := u.UserRepository.GetByID(ctx, dto.UserId)
	if err != nil {
		return errors.New("invalid user id")
	}

	user.IsVerified=true
	user.Category= model.Category(dto.Category)

	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		errors.New("user not modified")
	}

	return nil
}

func (u *userService) CheckIfUserVerified(ctx context.Context, bearer string) (bool, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return false,err
	}

	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		return false, errors.New("invalid user id")
	}

	return user.IsVerified,nil
}