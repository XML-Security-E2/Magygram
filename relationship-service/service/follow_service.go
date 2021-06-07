package service

import (
	"github.com/go-playground/validator"
	"relationship-service/domain/model"
	"relationship-service/infrastructure/persistence/neo4jdb"
	"relationship-service/service/intercomm"
)

type FollowService interface {
	FollowRequest(followRequest *model.FollowRequest) (bool, error)
	Unfollow(followRequest *model.FollowRequest) error
	IsUserFollowed(followRequest *model.FollowRequest) (interface{}, error)
	AcceptFollowRequest(followRequest *model.FollowRequest) error
	CreateUser(user *model.User) error
	ReturnFollowedUsers(user *model.User) (interface{}, error)
	ReturnFollowingUsers(user *model.User) (interface{}, error)
	ReturnFollowRequests(user *model.User) (interface{}, error)
}

type followService struct {
	neo4jdb.FollowRepository
	intercomm.UserClient
}

func NewFollowService(r neo4jdb.FollowRepository, userClient intercomm.UserClient) FollowService {
	return &followService{r, userClient}
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
			return false, err
		}
	} else {
		if err:= f.FollowRepository.CreateFollow(followRequest); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (f *followService) Unfollow(followRequest *model.FollowRequest) error {
	return f.FollowRepository.Unfollow(followRequest)
}

func (f *followService) IsUserFollowed(followRequest *model.FollowRequest) (interface{}, error) {
	if err := validator.New().Struct(followRequest); err != nil {
		return false, err
	}

	exists, err := f.FollowRepository.IsUserFollowed(followRequest);
	if err != nil {
			return false, err
	}

	return exists, nil
}

func (f *followService) CreateUser(user *model.User) error {
	if err := validator.New().Struct(user); err != nil {
		return err
	}
	if err:= f.FollowRepository.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (f *followService) AcceptFollowRequest(followRequest *model.FollowRequest) error {
	return f.FollowRepository.AcceptFollowRequest(followRequest)
}

func (f *followService) ReturnFollowedUsers(user *model.User) (interface{}, error) {
	return f.FollowRepository.ReturnFollowedUsers(user)
}

func (f *followService) ReturnFollowingUsers(user *model.User) (interface{}, error) {
	return f.FollowRepository.ReturnFollowingUsers(user)
}

func (f *followService) ReturnFollowRequests(user *model.User) (interface{}, error) {
	return f.FollowRepository.ReturnFollowRequests(user)
}