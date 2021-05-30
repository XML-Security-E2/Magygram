package intercomm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"user-service/conf"
	"user-service/domain/model"
)

type AuthClient interface {
	RegisterUser(user *model.User, password string, passwordRepeat string) error
	ActivateUser(userId string) error
	ChangePassword(userId string, password string, passwordRepeat string) error
}

type userAuthRequest struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
}

type passwordChangeRequest struct {
	UserId string `json:"userId"`
	Password string `json:"password"`
	PasswordRepeat string `json:"passwordRepeat"`
}

type errMessage struct {
	Message string `json:"message"`
}

type authClient struct {}


func NewAuthClient() AuthClient {
	baseUrl = fmt.Sprintf("%s%s:%s/api/users", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port)
	return &authClient{}
}

var (
	baseUrl = ""
)

func getErrorMessageFromRequestBody(body io.ReadCloser) (string ,error){
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}
	result := &errMessage{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return "", err
	}
	return result.Message, nil
}

func (a authClient) RegisterUser(user *model.User, password string, passwordRepeat string) error {
	userRequest := &userAuthRequest{Id: user.Id, Email: user.Email, Password: password, RepeatedPassword: passwordRepeat}
	jsonUserRequest, _ := json.Marshal(userRequest)

	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(jsonUserRequest))
	if err != nil || resp.StatusCode != 201 {
		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(message)
	}
	return nil
}

func (a authClient) ActivateUser(userId string) error {

	resp, err := http.Get(fmt.Sprintf("%s/activate/%s", baseUrl, userId))
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return errors.New("failed updating user")
	}
	return nil
}


func (a authClient) ChangePassword(userId string, password string, passwordRepeat string) error {
	passwordRequest := &passwordChangeRequest{UserId: userId, Password: password, PasswordRepeat: passwordRepeat}
	jsonPasswordRequest, _ := json.Marshal(passwordRequest)

	resp, err := http.Post(fmt.Sprintf("%s/reset-password", baseUrl), "application/json", bytes.NewBuffer(jsonPasswordRequest))
	if err != nil || resp.StatusCode != 200 {
		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(message)
	}
	return nil
}