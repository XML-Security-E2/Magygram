package intercomm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"story-service/conf"
	"story-service/domain/model"
)

type UserClient interface {
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

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
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