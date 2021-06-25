package intercomm

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"message-service/conf"
	"message-service/domain/model"
	"net/http"
)

type RelationshipClient interface {
	GetFollowedUsers(userId string) (model.FollowedUsersResponse, error)
	GetFollowingUsers(userId string) (model.FollowedUsersResponse, error)
}

type relationshipClient struct { }

var (
	baseRelationshipUrl = ""
)

func NewRelationshipClient() RelationshipClient {
	baseRelationshipUrl = fmt.Sprintf("%s%s:%s/api/relationship", conf.Current.Relationshipservice.Protocol, conf.Current.Relationshipservice.Domain, conf.Current.Relationshipservice.Port)
	return &relationshipClient{}
}


func (r relationshipClient) GetFollowedUsers(userId string) (model.FollowedUsersResponse, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/followed-users/%s", baseRelationshipUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		return model.FollowedUsersResponse{}, errors.New("not found")
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
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return model.FollowedUsersResponse{}, errors.New("not found")
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
