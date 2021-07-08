
package intercomm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"story-service/conf"
	"story-service/domain/model"
	"story-service/logger"
	"story-service/tracer"
)

type RelationshipClient interface {
	GetFollowedUsers(ctx context.Context, userId string) (model.FollowedUsersResponse, error)
	GetUnmutedFollowedUsers(ctx context.Context, userId string) (model.FollowedUsersResponse, error)
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

func (r relationshipClient) GetFollowedUsers(ctx context.Context, userId string) (model.FollowedUsersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "RelationshipClientGetFollowedUsers")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/followed-users/%s", baseRelationshipUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Relationship-service not available")
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Relationship-service get followed users")
		return model.FollowedUsersResponse{}, errors.New("post not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.FollowedUsersResponse{}, err
	}

	var users model.FollowedUsersResponse
	_ = json.Unmarshal(bodyBytes, &users)

	return users, nil
}

func (r relationshipClient) GetUnmutedFollowedUsers(ctx context.Context, userId string) (model.FollowedUsersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "RelationshipClientGetUnmutedFollowedUsers")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/unmuted-followed-users/%s", baseRelationshipUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Relationship-service not available")
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Relationship-service get followed users")
		return model.FollowedUsersResponse{}, errors.New("post not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.FollowedUsersResponse{}, err
	}

	var users model.FollowedUsersResponse
	_ = json.Unmarshal(bodyBytes, &users)

	return users, nil
}
