package service

import (
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"relationship-service/domain/model"
	"relationship-service/infrastructure/persistence/neo4jdb"
	"relationship-service/logger"
	"relationship-service/service/intercomm"
)

type FollowService interface {
	FollowRequest(followRequest *model.FollowRequest) (bool, error)
	Unfollow(followRequest *model.FollowRequest) error
	IsUserFollowed(followRequest *model.FollowRequest) (interface{}, error)
	IsMuted(mute *model.Mute) (interface{}, error)
	AcceptFollowRequest(bearer string, userId string) error
	CreateUser(user *model.User) error
	ReturnFollowedUsers(user *model.User) (interface{}, error)
	ReturnFollowingUsers(user *model.User) (interface{}, error)
	ReturnFollowRequests(bearer string) (interface{}, error)
	ReturnFollowRequestsForUser(bearer string, objectId string) (interface{}, error)
	Mute(mute *model.Mute) error
}

type followService struct {
	neo4jdb.FollowRepository
	intercomm.UserClient
	intercomm.AuthClient
}

func NewFollowService(r neo4jdb.FollowRepository, userClient intercomm.UserClient, ac intercomm.AuthClient) FollowService {
	return &followService{r, userClient, ac}
}

func (f *followService) Mute(mute *model.Mute) error {
	if err := validator.New().Struct(mute); err != nil {
		return err
	}
	if err := f.FollowRepository.Mute(mute); err != nil {
		return err
	}
	return nil
}

func (f *followService) Unmute(mute *model.Mute) error {
	if err := validator.New().Struct(mute); err != nil {
		return err
	}
	if err := f.FollowRepository.Mute(mute); err != nil {
		return err
	}
	return nil
}

func (f *followService) FollowRequest(followRequest *model.FollowRequest) (bool, error) {
	if err := validator.New().Struct(followRequest); err != nil {
		return false, err
	}
	isPrivate, err := f.UserClient.IsPrivate(followRequest.ObjectId)
	if err != nil {
		return false, nil
	}
	if isPrivate {
		if err:= f.FollowRepository.CreateFollowRequest(followRequest); err != nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
														 "object_id" : followRequest.ObjectId}).Error("Follow request create, database failure")
			return false, err
		}
		logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
													 "object_id" : followRequest.ObjectId}).Info("Follow request created")
		return true, nil
	} else {
		if err:= f.FollowRepository.CreateFollow(followRequest); err != nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
														 "object_id" : followRequest.ObjectId}).Error("Follow user, database failure")
			return false, err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
												     "object_id" : followRequest.ObjectId}).Info("User followed")
	}
	return false, nil
}

func (f *followService) Unfollow(followRequest *model.FollowRequest) error {
	err := f.FollowRepository.Unfollow(followRequest)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
													 "object_id" : followRequest.ObjectId}).Error("Unfollow user, database failure")

		return err
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
												 "object_id" : followRequest.ObjectId}).Info("User unfollowed")

	return err
}

func (f *followService) IsUserFollowed(followRequest *model.FollowRequest) (interface{}, error) {
	if err := validator.New().Struct(followRequest); err != nil {
		return false, err
	}

	exists, err := f.FollowRepository.IsUserFollowed(followRequest)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
													 "object_id" : followRequest.ObjectId}).Error("Check if follows, database failure")
		return false, err
	}

	return exists, nil
}

func (f *followService) IsMuted(mute *model.Mute) (interface{}, error) {
	if err := validator.New().Struct(mute); err != nil {
		return false, err
	}

	exists, err := f.FollowRepository.IsMuted(mute)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (f *followService) CreateUser(user *model.User) error {
	if err := validator.New().Struct(user); err != nil {
		return err
	}
	if err := f.FollowRepository.CreateUser(user); err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": user.Id}).Error("Create user node, database failure")
		return err
	}
	logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : user.Id}).Info("User node created")

	return nil
}

func (f *followService) AcceptFollowRequest(bearer string, userId string) error {
	loggedId, err := f.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return  err
	}

	err = f.FollowRepository.AcceptFollowRequest(&model.FollowRequest{
		SubjectId: userId,
		ObjectId:  loggedId,
	})
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": userId,
													 "object_id" : loggedId}).Error("Accept follow request, database failure")
		return  err
	}
	logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": userId,
													"object_id" : loggedId}).Info("Follow request accepted")
	return err
}

func (f *followService) ReturnFollowedUsers(user *model.User) (interface{}, error) {
	retVal, err := f.FollowRepository.ReturnFollowedUsers(user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": user.Id}).Error("Get followed users, database fetch failure")
		return retVal, err
	}
	return retVal, err
}

func (f *followService) ReturnFollowingUsers(user *model.User) (interface{}, error) {
	retVal, err := f.FollowRepository.ReturnFollowingUsers(user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": user.Id}).Error("Get following users, database fetch failure")
		return retVal, err
	}
	return retVal, err
}

func (f *followService) ReturnFollowRequests(bearer string) (interface{}, error) {
	loggedId, err := f.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return false, err
	}

	retVal, err := f.FollowRepository.ReturnFollowRequests(&model.User{Id: loggedId})
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"logged_user_id": loggedId}).Error("Get follow requests for user, database fetch failure")
		return retVal, err
	}
	return retVal, err
}

func (f *followService) ReturnFollowRequestsForUser(bearer string, objectId string) (interface{}, error) {
	loggedId, err := f.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return false, err
	}

	retVal, err := f.FollowRepository.ReturnFollowRequestsForUser(&model.User{Id: objectId}, loggedId)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"logged_user_id": loggedId, "object_id" : objectId}).Error("Get follow requests for user, database fetch failure")
		return retVal, err
	}
	return retVal, err
}