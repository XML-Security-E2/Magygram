package intercomm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"user-service/conf"
	"user-service/domain/model"
	"user-service/logger"
)

type AuthClient interface {
	RegisterUser(user *model.User, password string, passwordRepeat string) error
	ActivateUser(userId string) error
	GetLoggedUserId(bearer string) (string,error)
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
	baseUrl = fmt.Sprintf("%s%s:%s/api/auth", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port)
	return &authClient{}
}

var (
	baseUrl = ""
)

func (a authClient) GetLoggedUserId(bearer string) (string,error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/logged-user", baseUrl), nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.Error("Auth-service not available")
			return "", err
		}

		logger.LoggingEntry.Error("Auth-service get logged user")
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

	fmt.Println(fmt.Sprintf("%s%s:%s/api/users", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port))
	resp, err := http.Post(fmt.Sprintf("%s%s:%s/api/users", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port),
		"application/json", bytes.NewBuffer(jsonUserRequest))
	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"name": user.Name,
														 "surname" : user.Surname,
														 "email" : user.Email,
														 "username" : user.Username}).Error("Auth-service not available")
			return err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"name": user.Name, "surname" : user.Surname, "email" : user.Email, "username" : user.Username}).Error("Auth-service user registration")
		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(message)
	}
	return nil
}

func (a authClient) ActivateUser(userId string) error {

	resp, err := http.Get(fmt.Sprintf("%s%s:%s/api/users/activate/%s", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port, userId))
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : userId}).Error("Auth-service not available")
			return err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : userId}).Error("Auth-service user activation")
		return errors.New("failed updating user")
	}
	return nil
}


func (a authClient) ChangePassword(userId string, password string, passwordRepeat string) error {
	passwordRequest := &passwordChangeRequest{UserId: userId, Password: password, PasswordRepeat: passwordRepeat}
	jsonPasswordRequest, _ := json.Marshal(passwordRequest)

	resp, err := http.Post(fmt.Sprintf("%s%s:%s/api/users/reset-password", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port), "application/json", bytes.NewBuffer(jsonPasswordRequest))
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Auth-service not available")
			return err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Auth-service reset password")
		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(message)
	}
	return nil
}