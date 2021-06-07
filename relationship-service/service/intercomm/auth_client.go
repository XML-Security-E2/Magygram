package intercomm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"relationship-service/conf"
)

type AuthClient interface {
	GetLoggedUserId(bearer string) (string,error)
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

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
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
