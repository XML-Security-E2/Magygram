
package intercomm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"post-service/conf"
	"post-service/domain/model"
)

type RelationshipClient interface {
	GetFollowedUsers(userId string) (model.FollowedUsersResponse, error)
}

type relationshipClient struct {

}

var (
	baseRelationshipUrl = ""
)

func NewRelationshipClient() RelationshipClient {
	baseRelationshipUrl = fmt.Sprintf("%s%s:%s/api/relationship", conf.Current.Relationshipservice.Protocol, conf.Current.Relationshipservice.Domain, conf.Current.Relationshipservice.Port)
	return &relationshipClient{}
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
