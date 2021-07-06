package intercomm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"message-service/conf"
	"message-service/domain/model"
	"message-service/tracer"
	"net/http"
)

type UserClient interface {
	GetUsersForPostNotification(ctx context.Context, userId string) ([]*model.UserInfo, error)
	GetUsersForStoryNotification(ctx context.Context, userId string) ([]*model.UserInfo, error)
	GetLoggedUserInfo(ctx context.Context, bearer string) (*model.UserInfo, error)
	GetUsersInfo(ctx context.Context, userId string) (*model.UserInfo, error)
	CheckIfPostInteractionNotificationEnabled(ctx context.Context, userId string, userFromId string, interactionType string) (bool, error)
	IsUserPrivate(ctx context.Context, userId string) (bool, error)
}

type userClient struct {}

func NewUserClient() UserClient {
	baseUsersUrl = fmt.Sprintf("%s%s:%s/api/users", conf.Current.Userservice.Protocol, conf.Current.Userservice.Domain, conf.Current.Userservice.Port)
	return &userClient{}
}

var (
	baseUsersUrl = ""
)

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

func (u userClient) GetUsersInfo(ctx context.Context, userId string) (*model.UserInfo, error) {
	span := tracer.StartSpanFromContext(ctx, "UserClientGetUsersInfo")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/info/%s", baseUsersUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return &model.UserInfo{}, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &model.UserInfo{}, err
	}
	var userInfo model.UserInfo
	_ = json.Unmarshal(bodyBytes, &userInfo)

	return &userInfo, nil
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
		return &model.UserInfo{}, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &model.UserInfo{}, err
	}
	var userInfo model.UserInfo
	_ = json.Unmarshal(bodyBytes, &userInfo)

	return &userInfo, nil
}

func (u userClient) GetUsersForPostNotification(ctx context.Context, userId string) ([]*model.UserInfo, error) {
	span := tracer.StartSpanFromContext(ctx, "UserClientGetUsersForPostNotification")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/%s/notify/post", baseUsersUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			return []*model.UserInfo{}, err
		}

		return []*model.UserInfo{}, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []*model.UserInfo{}, err
	}
	var userInfo []*model.UserInfo
	_ = json.Unmarshal(bodyBytes, &userInfo)

	return userInfo, nil
}

func (u userClient) GetUsersForStoryNotification(ctx context.Context, userId string) ([]*model.UserInfo, error) {
	span := tracer.StartSpanFromContext(ctx, "UserClientGetUsersForStoryNotification")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/%s/notify/story", baseUsersUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			return []*model.UserInfo{}, err
		}

		return []*model.UserInfo{}, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []*model.UserInfo{}, err
	}
	var userInfo []*model.UserInfo
	_ = json.Unmarshal(bodyBytes, &userInfo)

	return userInfo, nil
}

func (u userClient) CheckIfPostInteractionNotificationEnabled(ctx context.Context, userId string, userFromId string, interactionType string) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "UserClientCheckIfPostInteractionNotificationEnabled")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/%s/%s/notify/%s", baseUsersUrl, userId, userFromId, interactionType), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			return false, err
		}

		return false, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var userInfo bool
	_ = json.Unmarshal(bodyBytes, &userInfo)

	return userInfo, nil
}