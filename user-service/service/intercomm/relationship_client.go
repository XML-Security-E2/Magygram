package intercomm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"user-service/conf"
	"user-service/domain/model"
)

type RelationshipClient interface {
	CreateUser(user *model.User) error
	GetFollowedUsers(userId string) (model.FollowedUsersResponse, error)
	GetFollowingUsers(userId string) (model.FollowedUsersResponse, error)
	FollowRequest(request *model.FollowRequest) (bool,error)
	Unfollow(request *model.FollowRequest) error
	ReturnFollowRequestsForUser(bearer string, objectId string) (bool, error)
}

type relationshipClient struct { }

type userRequest struct {
	Id string `json:"id"`
}

var (
	baseRelationshipUrl = ""
)

func NewRelationshipClient() RelationshipClient {
	baseRelationshipUrl = fmt.Sprintf("%s%s:%s/api/relationship", conf.Current.Relationshipservice.Protocol, conf.Current.Relationshipservice.Domain, conf.Current.Relationshipservice.Port)
	return &relationshipClient{}
}

func (r relationshipClient) CreateUser(user *model.User) error {
	userRequest := &userRequest{Id: user.Id}
	jsonUserRequest, _ := json.Marshal(userRequest)

	resp, err := http.Post(baseRelationshipUrl + "/user",
		"application/json", bytes.NewBuffer(jsonUserRequest))
	if err != nil || resp.StatusCode != 201 {
		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(message)
	}
	return nil
}

func (r relationshipClient) GetFollowedUsers(userId string) (model.FollowedUsersResponse, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/followed-users/%s", baseRelationshipUrl, userId), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		fmt.Println(resp.StatusCode)
		return model.FollowedUsersResponse{}, errors.New("post not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.FollowedUsersResponse{}, err
	}

	var users model.FollowedUsersResponse
	_ = json.Unmarshal(bodyBytes, &users)

	fmt.Println(users)

	return users, nil
}

func (r relationshipClient) GetFollowingUsers(userId string) (model.FollowedUsersResponse, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/following-users/%s", baseRelationshipUrl, userId), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return model.FollowedUsersResponse{}, errors.New("post not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.FollowedUsersResponse{}, err
	}

	var users model.FollowedUsersResponse
	_ = json.Unmarshal(bodyBytes, &users)

	fmt.Println(users)

	return users, nil
}

func (r relationshipClient) FollowRequest(request *model.FollowRequest) (bool,error) {
	jsonRequest, _ := json.Marshal(request)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/follow", baseRelationshipUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 201 {
		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return false, err
		}
		return false, errors.New(message)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var followRequest bool
	json.Unmarshal(bodyBytes, &followRequest)

	return followRequest, nil
}

func (r relationshipClient) Unfollow(request *model.FollowRequest) error {
	jsonRequest, _ := json.Marshal(request)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/unfollow", baseRelationshipUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(message)
	}
	return nil
}

func (r relationshipClient) ReturnFollowRequestsForUser(bearer string, objectId string) (bool, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/follow-requests/%s", baseRelationshipUrl, objectId), nil)
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return false, errors.New("user not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var isSentRequest bool
	json.Unmarshal(bodyBytes, &isSentRequest)
	fmt.Println(isSentRequest)

	return isSentRequest, nil
}
