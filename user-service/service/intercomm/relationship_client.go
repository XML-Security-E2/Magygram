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

type RelationshipClient interface {
	CreateUser(user *model.User) error
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

type userRequest struct {
	Id string `json:"id"`
}

