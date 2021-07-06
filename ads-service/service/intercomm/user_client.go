package intercomm

import (
	"ads-service/conf"
	"ads-service/domain/model"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

type UserClient interface {
	GetLoggedUserTargetGroup(bearer string) (*model.UserTargetGroup,error)
	GetLoggedUserInfo(bearer string) (*model.UserInfo,error)
}

type userClient struct {}

func NewUserClient() UserClient {
	baseUsersUrl = fmt.Sprintf("%s%s:%s/api/users", conf.Current.Userservice.Protocol, conf.Current.Userservice.Domain, conf.Current.Userservice.Port)
	return &userClient{}
}

var (
	baseUsersUrl = ""
)

func (u userClient) GetLoggedUserInfo(bearer string) (*model.UserInfo, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/logged", baseUsersUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			return &model.UserInfo{}, err
		}

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

func (u userClient) GetLoggedUserTargetGroup(bearer string) (*model.UserTargetGroup, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/logged/target-group", baseUsersUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			return &model.UserTargetGroup{}, err
		}

		return &model.UserTargetGroup{}, errors.New("unauthorized")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &model.UserTargetGroup{}, err
	}
	var userInfo model.UserTargetGroup
	_ = json.Unmarshal(bodyBytes, &userInfo)

	return &userInfo, nil
}