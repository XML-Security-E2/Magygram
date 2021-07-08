package intercomm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"post-service/conf"
	"post-service/domain/model"
	"post-service/logger"
	"post-service/tracer"
)

type UserClient interface {
	GetLoggedUserInfo(ctx context.Context, bearer string) (*model.UserInfo,error)
	GetLoggedAgentInfo(ctx context.Context, bearer string) (*model.AgentInfo,error)
	MapPostsToFavourites(bearer string, postIds []string) ([]*model.PostIdFavouritesFlag,error)
	IsUserPrivate(ctx context.Context, userId string) (bool, error)
	UpdateLikedPosts(bearer string, postId string) error
	AddComment(bearer string, postId string) error
	UpdateDislikedPosts(bearer string, postId string) error
	GetLikedPosts(ctx context.Context, bearer string) ([]string, error)
	GetDislikedPosts(ctx context.Context, bearer string) ([]string, error)
}

type userClient struct {}

func NewUserClient() UserClient {
	baseUsersUrl = fmt.Sprintf("%s%s:%s/api/users", conf.Current.Userservice.Protocol, conf.Current.Userservice.Domain, conf.Current.Userservice.Port)
	return &userClient{}
}

var (
	baseUsersUrl = ""
)


func (u userClient) AddComment(bearer string, postId string) error {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/post/comment/%s", baseUsersUrl, postId), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return err
	}

	return nil
}

func (u userClient) GetLoggedUserInfo(ctx context.Context, bearer string) (*model.UserInfo, error) {
	span := tracer.StartSpanFromContext(ctx, "UserClientGetLoggedUserInfo")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/logged", baseUsersUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.Error("User-service not available")
			return &model.UserInfo{}, err
		}

		logger.LoggingEntry.Error("User-service get logged user info")

		return &model.UserInfo{}, errors.New("unauthorized")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &model.UserInfo{}, err
	}
	var userInfo model.UserInfo
	_ = json.Unmarshal(bodyBytes, &userInfo)

	return &userInfo, nil
}

func (u userClient) GetLoggedAgentInfo(ctx context.Context, bearer string) (*model.AgentInfo, error) {
	span := tracer.StartSpanFromContext(ctx, "UserClientGetLoggedUserInfo")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/logged/agent", baseUsersUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			return &model.AgentInfo{}, err
		}
		return &model.AgentInfo{}, errors.New("unauthorized")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &model.AgentInfo{}, err
	}
	var userInfo model.AgentInfo
	_ = json.Unmarshal(bodyBytes, &userInfo)

	return &userInfo, nil}

func (u userClient) UpdateLikedPosts(bearer string, postId string) error {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/post/like/" + postId, baseUsersUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return err
	}

	return nil
}

func (u userClient) UpdateDislikedPosts(bearer string, postId string) error {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/post/dislike/" + postId, baseUsersUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return err
	}

	return nil
}

func (u userClient) MapPostsToFavourites(bearer string, postIds []string) ([]*model.PostIdFavouritesFlag, error) {

	jsonStr, err:= json.Marshal(postIds)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/collections/check-favourites", baseUsersUrl), bytes.NewReader(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"post_ids" : postIds}).Error("User-service not available")
			return nil, err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"post_ids" : postIds}).Error("User-service map posts to favourites")
		return nil, errors.New("unauthorized")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var postIdFav []*model.PostIdFavouritesFlag
	_ = json.Unmarshal(bodyBytes, &postIdFav)

	return postIdFav, nil
}


func (u userClient) IsUserPrivate(ctx context.Context, userId string) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "UserClientIsUserPrivate")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/%s/is-private", baseUsersUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : userId}).Error("User-service not available")
			return false, err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : userId}).Error("User-service check user privacy")
		return false, errors.New("user not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var isPrivate bool
	json.Unmarshal(bodyBytes, &isPrivate)

	return isPrivate, nil
}

func (u userClient) GetLikedPosts(ctx context.Context, bearer string) ([]string, error) {
	span := tracer.StartSpanFromContext(ctx, "UserClientGetLikedPosts")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/post/liked", baseUsersUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return []string{} ,err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	var users []string
	_ = json.Unmarshal(bodyBytes, &users)

	fmt.Println(users)

	return users, nil
}

func (u userClient) GetDislikedPosts(ctx context.Context, bearer string) ([]string, error) {
	span := tracer.StartSpanFromContext(ctx, "UserClientGetLikedPosts")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/post/disliked", baseUsersUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return []string{} ,err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	var users []string
	_ = json.Unmarshal(bodyBytes, &users)

	fmt.Println(users)

	return users, nil}