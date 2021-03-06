package intercomm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"net/http"
	"user-service/conf"
	"user-service/domain/model"
	"user-service/logger"
	"user-service/tracer"
)

type AuthClient interface {
	RegisterUser(ctx context.Context, user *model.User, password string, passwordRepeat string) (*http.Response, error )
	ActivateUser(ctx context.Context, userId string) error
	GetLoggedUserId(ctx context.Context, bearer string) (string,error)
	ChangePassword(ctx context.Context, userId string, password string, passwordRepeat string) error
	HasRole(ctx context.Context, bearer string, role string) (bool,error)
	RegisterAgent(ctx context.Context, user *model.User, password string) error
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

func (a authClient) RegisterUser(ctx context.Context, user *model.User, password string, passwordRepeat string) (*http.Response, error ){
	span := tracer.StartSpanFromContext(ctx, "AuthClientRegisterUser")
	defer span.Finish()

	userRequest := &userAuthRequest{Id: user.Id, Email: user.Email, Password: password, RepeatedPassword: passwordRepeat}
	jsonUserRequest, _ := json.Marshal(userRequest)

	req, err := http.NewRequestWithContext(ctx,"POST", fmt.Sprintf("%s%s:%s/api/users", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port),
										bytes.NewBuffer(jsonUserRequest))
	req.Header.Add("Content-Type","application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)

	//resp, err := http.Post(fmt.Sprintf("%s%s:%s/api/users", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port),
	//	"application/json", bytes.NewBuffer(jsonUserRequest))
	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"name": user.Name,
														 "surname" : user.Surname,
														 "email" : user.Email,
														 "username" : user.Username}).Error("Auth-service not available")
			return nil,err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"name": user.Name, "surname" : user.Surname, "email" : user.Email, "username" : user.Username}).Error("Auth-service user registration")
		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(message)
	}

	return resp,nil
}

func (a authClient) RegisterAgent(ctx context.Context, user *model.User, password string) error {
	span := tracer.StartSpanFromContext(ctx, "AuthClientRegisterAgent")
	defer span.Finish()

	userRequest := &userAuthRequest{Id: user.Id, Email: user.Email, Password: password, RepeatedPassword: password}
	jsonUserRequest, _ := json.Marshal(userRequest)

	req, err := http.NewRequestWithContext(ctx,"POST", fmt.Sprintf("%s%s:%s/api/users/agent", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port),
		bytes.NewBuffer(jsonUserRequest))
	req.Header.Add("Content-Type","application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			return err
		}

		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(message)
	}

	return nil
}

func (a authClient) ActivateUser(ctx context.Context, userId string) error {
	span := tracer.StartSpanFromContext(ctx, "AuthClientActivateUser")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s%s:%s/api/users/activate/%s", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port, userId),
								  nil)
	//resp, err := http.Get(fmt.Sprintf("%s%s:%s/api/users/activate/%s", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port, userId))

	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)

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


func (a authClient) ChangePassword(ctx context.Context, userId string, password string, passwordRepeat string) error {
	span := tracer.StartSpanFromContext(ctx, "ChangePasswordChangePassword")
	defer span.Finish()

	passwordRequest := &passwordChangeRequest{UserId: userId, Password: password, PasswordRepeat: passwordRepeat}
	jsonPasswordRequest, _ := json.Marshal(passwordRequest)

	//resp, err := http.Post(fmt.Sprintf("%s%s:%s/api/users/reset-password", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port), "application/json", bytes.NewBuffer(jsonPasswordRequest))

	req, err := http.NewRequestWithContext(ctx,"POST", fmt.Sprintf("%s%s:%s/api/users/reset-password", conf.Current.Authservice.Protocol, conf.Current.Authservice.Domain, conf.Current.Authservice.Port),
										bytes.NewBuffer(jsonPasswordRequest))
	req.Header.Add("Content-Type","application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)

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

func (a authClient) HasRole(ctx context.Context, bearer string,role string) (bool,error) {
	span := tracer.StartSpanFromContext(ctx, "ChangePasswordHasRole")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/has-role", baseUrl), nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("X-permissions", "[" + "\"" + role+ "\"" +"]")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

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