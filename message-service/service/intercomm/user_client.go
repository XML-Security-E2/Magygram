package intercomm

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"message-service/conf"
	"message-service/domain/model"
	"net/http"
)

type UserClient interface {
	GetUsersForPostNotification(userId string) ([]*model.UserInfo, error)
	GetUsersForStoryNotification(userId string) ([]*model.UserInfo, error)
	GetLoggedUserInfo(bearer string) (*model.UserInfo, error)
	GetUsersInfo(userId string) (*model.UserInfo, error)
	CheckIfPostInteractionNotificationEnabled(userId string, userFromId string, interactionType string) (bool, error)
}

type userClient struct {}

func NewUserClient() UserClient {
	baseUsersUrl = fmt.Sprintf("%s%s:%s/api/users", conf.Current.Userservice.Protocol, conf.Current.Userservice.Domain, conf.Current.Userservice.Port)
	return &userClient{}
}

var (
	baseUsersUrl = ""
)

func (u userClient) GetUsersInfo(userId string) (*model.UserInfo, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/info/%s", baseUsersUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

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

func (u userClient) GetLoggedUserInfo(bearer string) (*model.UserInfo, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/logged", baseUsersUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

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

func (u userClient) GetUsersForPostNotification(userId string) ([]*model.UserInfo, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/notify/post", baseUsersUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

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

func (u userClient) GetUsersForStoryNotification(userId string) ([]*model.UserInfo, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/notify/story", baseUsersUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

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

func (u userClient) CheckIfPostInteractionNotificationEnabled(userId string, userFromId string, interactionType string) (bool, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s/notify/%s", baseUsersUrl, userId, userFromId, interactionType), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

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