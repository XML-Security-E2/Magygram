package intercomm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"relationship-service/conf"
	"relationship-service/logger"
	"relationship-service/tracer"
)

type UserClient interface {
	IsPrivate(ctx context.Context, id string) (bool, error)
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

func (u userClient) IsPrivate(ctx context.Context, id string) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "UserClientIsPrivate")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/%s/is-private", baseUrl, id),nil)

	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)

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
