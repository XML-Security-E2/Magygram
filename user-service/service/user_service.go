package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mime/multipart"
	"regexp"
	"time"
	"user-service/conf"
	"user-service/domain/model"
	"user-service/domain/repository"
	service_contracts "user-service/domain/service-contracts"
	"user-service/domain/service-contracts/exceptions"
	"user-service/logger"
	"user-service/saga"
	"user-service/service/intercomm"

	"github.com/beevik/guid"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

type userService struct {
	repository.UserRepository
	repository.NotificationRulesRepository
	service_contracts.AccountActivationService
	service_contracts.ResetPasswordService
	intercomm.AuthClient
	intercomm.RelationshipClient
	intercomm.PostClient
	intercomm.MediaClient
	intercomm.MessageClient
	intercomm.StoryClient
	saga.Orchestrator
}

var (
	MaxUnsuccessfulLogins = 3
	ImageBytes []byte= nil
	RedisClient *redis.Client = nil
)

func NewAuthService(r repository.UserRepository, nrr repository.NotificationRulesRepository, a service_contracts.AccountActivationService, ic intercomm.AuthClient, rp service_contracts.ResetPasswordService, rC intercomm.RelationshipClient, pc intercomm.PostClient, mc intercomm.MediaClient, msc intercomm.MessageClient, sclient intercomm.StoryClient, orchestrator saga.Orchestrator) service_contracts.UserService {
	return &userService{r, nrr, a, rp, ic, rC, pc, mc, msc,sclient, orchestrator}
}

func (u *userService) GetUsersNotificationsSettings(ctx context.Context, bearer string, userId string) (*model.SettingsRequest, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	notificationRule, err := u.NotificationRulesRepository.GetRuleForUser(ctx, loggedId, userId)
	if err != nil {
		return nil, err
	}

	return &model.SettingsRequest{
		PostNotifications:  notificationRule.PostNotifications,
		StoryNotifications: notificationRule.StoryNotifications,
	}, nil
}

func (u *userService) ChangeUsersNotificationsSettings(ctx context.Context, bearer string, settingsReq *model.SettingsRequest, userId string) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	followedUsers, err := u.RelationshipClient.GetFollowingUsers(userId)
	if err != nil {
		return err
	}

	if !doesUserFollow(followedUsers, loggedId) {
		return errors.New("cannot edit notification settings for not followed user")
	}

	notificationRule, err := u.NotificationRulesRepository.GetRuleForUser(ctx, loggedId, userId)
	if err != nil {
		u.NotificationRulesRepository.Create(ctx, &model.PostStoryNotifications{
			Id:                  guid.New().String(),
			UserId:              loggedId,
			NotificationsFromId: userId,
			PostNotifications:   settingsReq.PostNotifications,
			StoryNotifications:  settingsReq.StoryNotifications,
		})
	} else {
		notificationRule.StoryNotifications = settingsReq.StoryNotifications
		notificationRule.PostNotifications = settingsReq.PostNotifications
		u.NotificationRulesRepository.Update(ctx, notificationRule)
	}
	return nil

}

func (u *userService) GetUsersForPostNotification(ctx context.Context, userId string) ([]*model.UserInfo, error) {

	notifyIds, err := u.NotificationRulesRepository.GetNotifiersForPost(ctx, userId)
	if err != nil || len(notifyIds) == 0 {
		return []*model.UserInfo{}, nil
	}
	var retVal []*model.UserInfo

	for _, notifyId := range notifyIds {
		notifyUser, err := u.UserRepository.GetByID(ctx, notifyId)
		if err == nil {
			retVal = append(retVal, &model.UserInfo{
				Id:       notifyId,
				Username: notifyUser.Username,
				ImageURL: notifyUser.ImageUrl,
			})
		}
	}

	return retVal, nil
}

func (u *userService) GetUsersForStoryNotification(ctx context.Context, userId string) ([]*model.UserInfo, error) {

	notifyIds, err := u.NotificationRulesRepository.GetNotifiersForStory(ctx, userId)
	if err != nil || len(notifyIds) == 0 {
		return []*model.UserInfo{}, nil
	}
	var retVal []*model.UserInfo

	for _, notifyId := range notifyIds {
		notifyUser, err := u.UserRepository.GetByID(ctx, notifyId)
		if err == nil {
			retVal = append(retVal, &model.UserInfo{
				Id:       notifyId,
				Username: notifyUser.Username,
				ImageURL: notifyUser.ImageUrl,
			})
		}
	}

	return retVal, nil
}

func (u *userService) CheckIfPostInteractionNotificationEnabled(ctx context.Context, userId string, userFromId string, interactionType string) (bool, error) {
	user, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return false, errors.New("invalid user id")
	}

	if interactionType == "like" {
		if user.NotificationSettings.NotifyLike == model.FromEveryOne {
			return true, nil
		} else if user.NotificationSettings.NotifyLike == model.FromPeopleIFollow {
			followedUsers, err := u.RelationshipClient.GetFollowedUsers(user.Id)
			if err == nil {
				for _, followedUser := range followedUsers.Users {
					if followedUser == userFromId {
						return true, nil
					}
				}
			}
			return false, nil
		} else {
			return false, nil
		}
	} else if interactionType == "dislike" {
		if user.NotificationSettings.NotifyDislike == model.FromEveryOne {
			return true, nil
		} else if user.NotificationSettings.NotifyDislike == model.FromPeopleIFollow {
			followedUsers, err := u.RelationshipClient.GetFollowedUsers(user.Id)
			if err == nil {
				for _, followedUser := range followedUsers.Users {
					if followedUser == userFromId {
						return true, nil
					}
				}
			}
			return false, nil
		} else {
			return false, nil
		}
	} else if interactionType == "comment" {
		if user.NotificationSettings.NotifyComment == model.FromEveryOne {
			return true, nil
		} else if user.NotificationSettings.NotifyComment == model.FromPeopleIFollow {
			followedUsers, err := u.RelationshipClient.GetFollowedUsers(user.Id)
			if err == nil {
				for _, followedUser := range followedUsers.Users {
					if followedUser == userFromId {
						return true, nil
					}
				}
			}
			return false, nil
		} else {
			return false, nil
		}
	} else if interactionType == "follow" {
		return user.NotificationSettings.NotifyFollow, nil
	} else if interactionType == "follow-request" {
		return user.NotificationSettings.NotifyFollowRequest, nil
	} else if interactionType == "accepted-follow-request" {
		return user.NotificationSettings.NotifyAcceptFollowRequest, nil
	}
	return false, nil
}

func (u userService) DeleteUser(ctx context.Context, requestId string) error {
	request, err := u.UserRepository.GetByID(ctx, requestId)
	if err != nil {
		return errors.New("Request not found")
	}

	request.IsDeleted = true

	u.UserRepository.DeleteUser(ctx, request)

	return nil
}

func (u *userService) EditUser(ctx context.Context, bearer string, userId string, userRequest *model.EditUserRequest) (string, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return "", err
	}

	if loggedId != userId {
		logger.LoggingEntry.WithFields(logrus.Fields{"requested_user_id": userId, "logged_user_id": loggedId}).Warn("Unauthorized access")
		return "", &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
	}

	user, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return "", errors.New("invalid user id")
	}

	temp := user.Username

	user.Username = userRequest.Username
	user.Name = userRequest.Name
	user.Surname = userRequest.Surname
	user.Number = userRequest.Number
	user.Website = userRequest.Website
	user.Bio = userRequest.Bio
	user.Gender = userRequest.Gender
	if err = validator.New().Struct(user); err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"name": userRequest.Name,
			"surname":  userRequest.Surname,
			"number":   userRequest.Number,
			"gender":   userRequest.Gender,
			"username": userRequest.Username,
			"website":  userRequest.Website,
			"bio":      userRequest.Bio}).Warn("User registration validation failure")
		return "", err
	}

	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"name": userRequest.Name,
			"surname":  userRequest.Surname,
			"number":   userRequest.Number,
			"gender":   userRequest.Gender,
			"username": userRequest.Username,
			"website":  userRequest.Website,
			"bio":      userRequest.Bio}).Error("User database update failure")
		return "", err
	}
	fmt.Println(temp)
	if temp != userRequest.Username {
		err = u.editSharedUserInfo(bearer, user, userRequest.Username, user.ImageUrl)
		if err != nil {
			return "", err
		}
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Info("User information updated")

	return user.Id, err
}

func (u *userService) EditUserImage(ctx context.Context, bearer string, userId string, userImage []*multipart.FileHeader) (string, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return "", err
	}

	if loggedId != userId {
		logger.LoggingEntry.WithFields(logrus.Fields{"requested_user_id": userId, "logged_user_id": loggedId}).Warn("Unauthorized access")
		return "", &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
	}

	user, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return "", errors.New("invalid user id")
	}

	media, err := u.MediaClient.SaveMedia(userImage)
	if err != nil {
		return "", err
	}

	if len(media) == 0 {
		return "", errors.New("error while saving image")
	}
	user.ImageUrl = media[0].Url

	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("User profile picture update failure")
		return "", err
	}

	err = u.editSharedUserInfo(bearer, user, user.Username, user.ImageUrl)

	if err != nil {
		return "", err
	}


	logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Info("User profile picture updated")

	return media[0].Url, err
}

func (u *userService) editSharedUserInfo(bearer string, user *model.User, username string, imageUrl string) error {
	err := u.PostClient.EditPostOwnerInfo(bearer, model.UserInfo{
		Id:       user.Id,
		Username: username,
		ImageURL: imageUrl,
	})

	if err != nil {
		return err
	}

	err = u.StoryClient.EditStoryOwnerInfo(bearer, model.UserInfo{
		Id:       user.Id,
		Username: username,
		ImageURL: imageUrl,
	})

	if err != nil {
		return err
	}

	if len(user.LikedPosts) > 0 {
		err = u.PostClient.EditLikedByInfo(bearer, model.UserInfoEdit{
			Id:       user.Id,
			Username: username,
			ImageURL: imageUrl,
			PostIds:  user.LikedPosts,
		})

		if err != nil {
			return err
		}
	}

	if len(user.DislikedPosts) > 0 {
		err = u.PostClient.EditDislikedByInfo(bearer, model.UserInfoEdit{
			Id:       user.Id,
			Username: username,
			ImageURL: imageUrl,
			PostIds:  user.DislikedPosts,
		})

		if err != nil {
			return err
		}
	}

	if len(user.CommentedPosts) > 0 {
		err = u.PostClient.EditCommentedByInfo(bearer, model.UserInfoEdit{
			Id:       user.Id,
			Username: username,
			ImageURL: imageUrl,
			PostIds:  user.CommentedPosts,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (u *userService) EditUsersNotifications(ctx context.Context, bearer string, notificationReq *model.NotificationSettingsUpdateReq) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		return errors.New("invalid user id")
	}

	notifications := &model.NotificationSettings{
		NotifyLike:                "",
		NotifyDislike:             "",
		NotifyFollow:              notificationReq.NotifyFollow,
		NotifyFollowRequest:       notificationReq.NotifyFollowRequest,
		NotifyAcceptFollowRequest: notificationReq.NotifyAcceptFollowRequest,
		NotifyComment:             "",
	}

	if notificationReq.NotifyLike == 0 {
		notifications.NotifyLike = model.MutE
	} else if notificationReq.NotifyLike == 1 {
		notifications.NotifyLike = model.FromPeopleIFollow
	} else {
		notifications.NotifyLike = model.FromEveryOne
	}

	if notificationReq.NotifyDislike == 0 {
		notifications.NotifyDislike = model.MutE
	} else if notificationReq.NotifyDislike == 1 {
		notifications.NotifyDislike = model.FromPeopleIFollow
	} else {
		notifications.NotifyDislike = model.FromEveryOne
	}

	if notificationReq.NotifyComment == 0 {
		notifications.NotifyComment = model.MutE
	} else if notificationReq.NotifyComment == 1 {
		notifications.NotifyComment = model.FromPeopleIFollow
	} else {
		notifications.NotifyComment = model.FromEveryOne
	}

	user.NotificationSettings = *notifications
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
	user.IsPrivate = (*privacySettingsReq).IsPrivate
	_, err = u.UserRepository.Update(ctx, user)

	if err != nil {
		return err

	}
	return nil
}

func (u *userService) RegisterUser(ctx context.Context, userRequest *model.UserRequest) ([]byte, error) {
	user, _ := model.NewUser(userRequest)
	if err := validator.New().Struct(user); err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"name": userRequest.Name, "surname": userRequest.Surname, "email": userRequest.Email, "username": userRequest.Username}).Warn("User registration validation failure")
		return nil, err
	}

	result, err := u.UserRepository.Create(ctx, user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"name": userRequest.Name, "surname": userRequest.Surname, "email": userRequest.Email, "username": userRequest.Username}).Error("User database create failure")
		return nil, err
	}

	sagaUserRequest := saga.UserRequest{
		Id: user.Id,
		Email: user.Email,
		Password: userRequest.Password,
		RepeatedPassword: userRequest.RepeatedPassword,
	}

	registrationMessage := saga.RegisterUserMessage{Service: saga.ServiceAuth, SenderService: saga.ServiceUser, Action: saga.ActionStart, User: sagaUserRequest}
	u.Orchestrator.Next(saga.AuthChannel, saga.ServiceAuth, registrationMessage)

	finished := make(chan bool)

	go u.RedisConnection(finished)

	select {
	case redisResult := <-finished:
		if !redisResult {
			return nil, errors.New("Internal server error")
		}

		accActivationId, _ := u.AccountActivationService.Create(ctx, user.Id)

		go SendActivationMail(userRequest.Email, userRequest.Name, accActivationId)

		if userId, ok := result.InsertedID.(string); ok {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Info("User registered")
			return ImageBytes, nil
		}
		return ImageBytes, err
	case <-time.After(10 * time.Second):
		fmt.Println("ROLLBACK ODRADI KAD JEDAN MS NE RADI")
		sendToReplyChannel(RedisClient, &registrationMessage, saga.ActionError, saga.ServiceAuth, saga.ServiceUser)
		sendToReplyChannel(RedisClient, &registrationMessage, saga.ActionError, saga.ServiceUser, saga.ServiceUser)

		return nil, errors.New("Internal server error")
	}

	return nil, errors.New("Internal server error")
}

func (u *userService) RegisterAgentByAdmin(ctx context.Context, agentRequest *model.AgentRequest) (string, error) {
	agentRegistrationDTO := model.AgentRegistrationDTO{
		Name: agentRequest.Name,
		Surname: agentRequest.Surname,
		Email: agentRequest.Email,
		Website: agentRequest.WebSite,
		Username: agentRequest.Username,
		Password: "",
	}

	user, _ := model.NewAgent(&agentRegistrationDTO)
	if err := validator.New().Struct(user); err != nil {
		return "", err
	}

	result, err := u.UserRepository.Create(ctx, user)
	if err != nil {
		return "", err
	}

	hashAndSalt, err := HashAndSaltPasswordIfStrongAndMatching(agentRequest.Password, agentRequest.RepeatedPassword)
	if err != nil {
		return "", err
	}

	err = u.AuthClient.RegisterAgent(user, hashAndSalt)
	if err != nil {
		return "", err
	}

	err = u.RelationshipClient.CreateUser(user)
	if err != nil {
		return "", err
	}

	accActivationId, _ := u.AccountActivationService.Create(ctx, user.Id)

	go SendActivationMail(agentRequest.Email, agentRequest.Name, accActivationId)

	if userId, ok := result.InsertedID.(string); ok {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Info("User registered")
		return userId, nil
	}
	return user.Id, err
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

	logger.LoggingEntry.WithFields(logrus.Fields{"activation_id": activationId, "user_id": accActivation.UserId}).Info("User activated")

	return true, err
}

func (u *userService) ResendActivationLink(ctx context.Context, activateLinkRequest *model.ActivateLinkRequest) (bool, error) {
	user, err := u.UserRepository.GetByEmail(ctx, activateLinkRequest.Email)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email": activateLinkRequest.Email}).Warn("Invalid email address")
		return false, err
	}

	accActivationId, _ := u.AccountActivationService.Create(ctx, user.Id)
	go SendActivationMail(user.Email, user.Name, accActivationId)

	logger.LoggingEntry.WithFields(logrus.Fields{"activation_id": accActivationId, "user_id": user.Id}).Info("Account activation link created")
	return true, nil
}

func (u *userService) ResetPassword(ctx context.Context, userEmail string) (bool, error) {
	user, err := u.GetByEmail(ctx, userEmail)
	//pokrivena invalid email
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email": userEmail}).Warn("Invalid email address")
		return false, errors.New("invalid email address")
	}

	accResetPasswordId, _ := u.ResetPasswordService.Create(ctx, user.Id)
	go SendResetPasswordMail(user.Email, user.Name, accResetPasswordId)

	logger.LoggingEntry.WithFields(logrus.Fields{"reset_password_id": accResetPasswordId, "user_id": user.Id}).Info("Reset password link created")
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

	logger.LoggingEntry.WithFields(logrus.Fields{"reset_password_id": changePasswordRequest.ResetPasswordId, "user_id": user.Id}).Info("Users password changed")
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
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Warn("Invalid user id")
		return nil, errors.New("invalid user id")
	}

	return user, err
}

func (u *userService) SearchForUsersByUsername(ctx context.Context, username string, bearer string) ([]model.User, error) {
	userId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	users, err := u.UserRepository.SearchForUsersByUsername(ctx, username, userId)

	if err != nil {
		return nil, errors.New("Couldn't find any users")
	}

	return users, err
}

func (u *userService) SearchForInfluencerByUsername(ctx context.Context, username string, bearer string) ([]model.User, error) {
	userId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	users, err := u.UserRepository.SearchForInfluencerByUsername(ctx, username, userId)

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

func (u *userService) GetUsersInfo(ctx context.Context, userId string) (*model.UserInfo, error) {
	user, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	return &model.UserInfo{
		Id:       userId,
		Username: user.Username,
		ImageURL: user.ImageUrl,
	}, nil}


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

func (u *userService) GetUserProfileById(ctx context.Context, bearer string, userId string) (*model.UserProfileResponse, error) {

	log.Println(userId)

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
	if loggedId != "" {
		for _, blockedUser := range user.BlockedUsers {
			if blockedUser == loggedId {
				return nil, errors.New("invalid user id")
			}
		}
	}
	sentReq := false
	muted := false
	blocked := false
	if userId != loggedId {
		sentReq, _ = u.RelationshipClient.ReturnFollowRequestsForUser(bearer, userId)
		muted, _ = u.RelationshipClient.IsMuted(model.Mute{SubjectId: loggedId, ObjectId: userId})
		blocked, _ = u.UserRepository.IsBlocked(ctx, loggedId, userId)
	}

	notificationSettings := model.NotificationSettingsUpdateReq{
		NotifyLike:                0,
		NotifyDislike:             0,
		NotifyFollow:              user.NotificationSettings.NotifyFollow,
		NotifyFollowRequest:       user.NotificationSettings.NotifyFollowRequest,
		NotifyAcceptFollowRequest: user.NotificationSettings.NotifyAcceptFollowRequest,
		NotifyComment:             0,
	}

	if user.NotificationSettings.NotifyLike == model.MutE {
		notificationSettings.NotifyLike = 0
	} else if user.NotificationSettings.NotifyLike == model.FromPeopleIFollow {
		notificationSettings.NotifyLike = 1
	} else {
		notificationSettings.NotifyLike = 2
	}

	if user.NotificationSettings.NotifyDislike == model.MutE {
		notificationSettings.NotifyDislike = 0
	} else if user.NotificationSettings.NotifyDislike == model.FromPeopleIFollow {
		notificationSettings.NotifyDislike = 1
	} else {
		notificationSettings.NotifyDislike = 2
	}

	if user.NotificationSettings.NotifyComment == model.MutE {
		notificationSettings.NotifyComment = 0
	} else if user.NotificationSettings.NotifyComment == model.FromPeopleIFollow {
		notificationSettings.NotifyComment = 1
	} else {
		notificationSettings.NotifyComment = 2
	}

	fmt.Println(sentReq)
	retVal := &model.UserProfileResponse{
		Username:             user.Username,
		Name:                 user.Name,
		Surname:              user.Surname,
		Website:              user.Website,
		Bio:                  user.Bio,
		Number:               user.Number,
		Gender:               user.Gender,
		Category:             user.Category,
		ImageUrl:             user.ImageUrl,
		PostNumber:           postsCount,
		Following:            following,
		Muted:                muted,
		Blocked:              blocked,
		Email:                user.Email,
		FollowersNumber:      len(followedUsers.Users),
		FollowingNumber:      len(followingUsers.Users),
		SentFollowRequest:    sentReq,
		PrivacySettings:      user.PrivacySettings,
		NotificationSettings: notificationSettings,
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
			UserInfo: &model.UserInfo{
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
			UserInfo: &model.UserInfo{
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
	for index, blockedUserId := range user.BlockedUsers {
		if blockedUserId == userId {
			return true, index
		}
	}
	return false, 0
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
			UserInfo: &model.UserInfo{
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

	var result, index = didUserLikedPost(user, postId)

	if result {
		user.LikedPosts = append(user.LikedPosts[:index], user.LikedPosts[index+1:]...)
	} else {
		user.LikedPosts = append(user.LikedPosts, postId)
	}

	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		errors.New("user not modified")
	}

	return nil
}

func (u *userService) AddComment(ctx context.Context, bearer string, postId string) error {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		errors.New("invalid user id")
	}

	commented := didUserCommentedPost(user, postId)

	if !commented {
		user.CommentedPosts = append(user.CommentedPosts, postId)
		_, err = u.UserRepository.Update(ctx, user)
		if err != nil {
			errors.New("user not modified")
		}
	}

	return nil
}

func didUserCommentedPost(user *model.User, postId string) bool {
	for _, commentedPostId := range user.CommentedPosts {
		if commentedPostId == postId {
			return true
		}
	}
	return false
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

	var result, index = didUserDislikedPost(user, postId)

	if result {
		user.DislikedPosts = append(user.DislikedPosts[:index], user.DislikedPosts[index+1:]...)
	} else {
		user.DislikedPosts = append(user.DislikedPosts, postId)
	}

	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		errors.New("user not modified")
	}

	return nil
}

func didUserDislikedPost(user *model.User, postId string) (bool, int) {
	for index, dislikedPostId := range user.DislikedPosts {
		if dislikedPostId == postId {
			return true, index
		}
	}
	return false, 0
}

func didUserLikedPost(user *model.User, postId string) (bool, int) {
	for index, likedPostId := range user.LikedPosts {
		if likedPostId == postId {
			return true, index
		}
	}
	return false, 0
}

func (u *userService) GetUserLikedPost(ctx context.Context, bearer string) ([]string, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return []string{}, err
	}

	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		return []string{}, errors.New("invalid user id")
	}

	return user.LikedPosts, nil
}

func (u *userService) GetUserDislikedPost(ctx context.Context, bearer string) ([]string, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return []string{}, err
	}

	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		return []string{}, errors.New("invalid user id")
	}

	return user.DislikedPosts, nil
}

func (u *userService) VerifyUser(ctx context.Context, dto *model.VerifyAccountDTO) error {
	user, err := u.UserRepository.GetByID(ctx, dto.UserId)
	if err != nil {
		return errors.New("invalid user id")
	}

	user.IsVerified = true
	user.Category = model.Category(dto.Category)

	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		errors.New("user not modified")
	}

	return nil
}

func (u *userService) CheckIfUserVerified(ctx context.Context, bearer string) (bool, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return false, err
	}

	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		return false, errors.New("invalid user id")
	}

	return user.IsVerified, nil
}

func (u *userService) CheckIfUserVerifiedById(ctx context.Context, userId string) (bool, error) {
	log.Println(userId)
	user, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return false, errors.New("invalid user id")
	}

	return user.IsVerified, nil
}


func (u *userService) GetFollowRecommendation(ctx context.Context, bearer string) (*model.FollowRecommendationResponse, error) {
	loggedId, err := u.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	log.Println("test")
	user, err := u.UserRepository.GetByID(ctx, loggedId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}
	log.Println("test1")


	var recommendedUsers model.RecommendedUsersResponse
	recommendedUsers, err = u.RelationshipClient.GetRecommendedUsers(loggedId)
	if err != nil {
		return nil, err
	}

	retVal := model.FollowRecommendationResponse{
		Name: user.Name,
		Surname: user.Surname,
		ImageURL: user.ImageUrl,
		Username: user.Username,
		RecommendedUsers: []*model.RecommendUserInfo{},
	}

	for _, userId := range recommendedUsers.Users {
		user, err := u.UserRepository.GetByID(ctx, userId)
		if err != nil {
			return nil, errors.New("invalid user id")
		}

		var newUserInfo = &model.RecommendUserInfo{
			Id:       userId,
			Username: user.Username,
			ImageURL: user.ImageUrl,
			SendedRequest: false,
			Followed:false,
		}
		retVal.RecommendedUsers=append(retVal.RecommendedUsers, newUserInfo)
	}

	return &retVal,nil
}

func (u *userService) RegisterAgent(ctx context.Context, agentRegistrationDTO *model.AgentRegistrationDTO) (string, error) {
	user, _ := model.NewAgent(agentRegistrationDTO)
	if err := validator.New().Struct(user); err != nil {
		return "", err
	}
	result, err := u.UserRepository.Create(ctx, user)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = u.AuthClient.RegisterAgent(user, agentRegistrationDTO.Password)
	if err != nil {
		return "", err
	}

	err = u.RelationshipClient.CreateUser(user)
	if err != nil {
		return "", err
	}

	accActivationId, _ := u.AccountActivationService.Create(ctx, user.Id)

	go SendActivationMail(agentRegistrationDTO.Email, agentRegistrationDTO.Name, accActivationId)

	if userId, ok := result.InsertedID.(string); ok {
		return userId, nil
	}

	return user.Id, err
}

func HashAndSaltPasswordIfStrongAndMatching(password string, repeatedPassword string) (string, error) {
	isMatching := password == repeatedPassword
	if !isMatching {
		return "", errors.New("passwords are not matching")
	}
	isWeak, _ := regexp.MatchString("^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[^!@#$%^&*(),.?\":{}|<>~'_+=]*)$", password)

	if isWeak {
		return "", errors.New("password must contain minimum eight characters, at least one capital letter, one number and one special character")
	}
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash), err
}

func (u *userService) RedisConnection(finished chan bool) {
	// create client and ping redis
	var err error


	client := redis.NewClient(&redis.Options{Addr: conf.Current.RedisDatabase.Host+":"+ conf.Current.RedisDatabase.Port, Password: "", DB: 0})
	if _, err = client.Ping().Result(); err != nil {
		log.Fatalf("error creating redis client %s", err)
	}

	RedisClient=client


	// subscribe to the required channels
	pubsub := client.Subscribe(saga.UserChannel, saga.ReplyChannel)
	if _, err = pubsub.Receive(); err != nil {
		log.Fatalf("error subscribing %s", err)
	}
	defer func() { _ = pubsub.Close() }()
	ch := pubsub.Channel()

	log.Println("starting the order service")
	for {
		select {
		case msg := <-ch:
			m := saga.RegisterUserMessage{}
			err := json.Unmarshal([]byte(msg.Payload), &m)
			if err != nil {
				log.Println(err)
				continue
			}

			switch msg.Channel {
			case saga.UserChannel:
				// Happy Flow
				if m.Action == saga.ActionStart {
					if m.SenderService == saga.ServiceRelationship{
						ImageBytes = m.ImageByte
						finished <- true
					}
				}

				// Rollback flow
				if m.Action == saga.ActionRollback {
					log.Println("TEST")
					u.UserRepository.PhysicalDelete(context.TODO(), m.User.Id)

					finished <- false
				}
			}
		}
	}
}


func sendToReplyChannel(client *redis.Client, m *saga.RegisterUserMessage, action string, service string, senderService string) {
	var err error
	m.Action = action
	m.Service = service
	m.SenderService = senderService
	if err = client.Publish(saga.ReplyChannel, m).Err(); err != nil {
		log.Printf("error publishing done-message to %s channel", saga.ReplyChannel)
	}
	log.Printf("done message published to channel :%s", saga.ReplyChannel)
}