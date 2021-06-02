package intercomm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"relationship-service/conf"
)

type UserClient interface {
	IsPrivate(id string) (bool, error)
}

type userClient struct {
}

var (
	baseUrl = ""
)

func NewUserClient() UserClient {
	baseUrl = fmt.Sprintf("%s%s:%s/api/users", conf.Current.Userservice.Protocol, conf.Current.Userservice.Domain, conf.Current.Userservice.Port)
	return &userClient{}
}

type privateFlag struct {
	isPrivate bool `json:"isPrivate"`
}

func (u userClient) IsPrivate(id string) (bool, error) {
	return true, nil
	resp, err := http.Get(fmt.Sprintf("%s/is-private/%s", baseUrl, id))
	if err != nil || resp.StatusCode != 200 {
		return false, errors.New("could not get user profile private flag")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	result := &privateFlag{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return false, err
	}
	return result.isPrivate, nil
}
