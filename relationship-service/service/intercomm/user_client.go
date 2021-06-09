package intercomm

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"relationship-service/conf"
	"relationship-service/logger"
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
	resp, err := http.Get(fmt.Sprintf("%s/%s/is-private", baseUrl, id))
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : id}).Error("User-service not available")
			return false, err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : id}).Error("User-service check user privacy")
		return false, errors.New("could not get user profile private flag")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var isPrivate bool
	json.Unmarshal(bodyBytes, &isPrivate)

	return isPrivate, nil
}
