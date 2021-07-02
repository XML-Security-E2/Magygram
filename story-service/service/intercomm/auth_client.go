package intercomm

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"story-service/conf"
	"story-service/logger"
)

type AuthClient interface {
	GetLoggedUserId(bearer string) (string,error)
	HasRole(bearer string, role string) (bool,error)
}

type authClient struct {}

func NewAuthClient() AuthClient {
	baseAuthUrl = fmt.Sprintf("%s%s:%s/api/auth", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port)
	return &authClient{}
}
var (
	baseAuthUrl = ""
)

func (a authClient) GetLoggedUserId(bearer string) (string,error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/logged-user", baseAuthUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.Error("Auth-service not available")
			return "", err
		}

		logger.LoggingEntry.Error("Auth-service get logged user")
		return "", errors.New("unauthorized")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var userId string
	json.Unmarshal(bodyBytes, &userId)

	return userId, nil
}

func (a authClient) HasRole(bearer string,role string) (bool,error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/has-role", baseAuthUrl), nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("X-permissions", "[" + "\"" + role+ "\"" +"]")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			return false, err
		}

		return false, nil
	}

	return true, nil
}
