package intercomm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"request-service/conf"
	"request-service/tracer"
)

type AuthClient interface {
	GetLoggedUserId(ctx context.Context, bearer string) (string,error)
}


type authClient struct {}

func NewAuthClient() AuthClient {
	baseUrl = fmt.Sprintf("%s%s:%s/api/auth", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port)
	return &authClient{}
}

var (
	baseUrl = ""
)

func (a authClient) GetLoggedUserId(ctx context.Context, bearer string) (string,error) {
	span := tracer.StartSpanFromContext(ctx, "AuthClientGetLoggedUserId")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/logged-user", baseUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			return "", err
		}

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