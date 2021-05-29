package intercomm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"user-service/conf"
	"user-service/domain/model"
)

type AuthClient interface {
	RegisterUser(user *model.User) error
}

type userAuthRequest struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
}

type authClient struct {}

func NewAuthClient() AuthClient {
	baseUrl = fmt.Sprintf("%s%s:%s/users", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port)
	return &authClient{}
}

var (
	baseUrl = ""
)

func (a authClient) RegisterUser(user *model.User) error {
	userRequest := &userAuthRequest{Id: user.Id, Email: user.Email, Password: user.Password, RepeatedPassword: user.Password}
	jsonUserRequest, _ := json.Marshal(userRequest)
	fmt.Println(baseUrl)
	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(jsonUserRequest))
	if err != nil || resp.StatusCode != 201 {
		fmt.Println(resp.StatusCode)
		return errors.New("failed creating user")
	}
	return nil
}