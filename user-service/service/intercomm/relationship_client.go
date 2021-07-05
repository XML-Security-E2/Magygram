package intercomm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"user-service/conf"
	"user-service/domain/model"
	"user-service/logger"
	"user-service/tracer"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type RelationshipClient interface {
	CreateUser(user *model.User) error
	GetFollowedUsers(ctx context.Context, userId string) (model.FollowedUsersResponse, error)
	GetFollowingUsers(ctx context.Context, userId string) (model.FollowedUsersResponse, error)
	FollowRequest(ctx context.Context, request *model.FollowRequest) (bool, error)
	Unfollow(ctx context.Context, request *model.FollowRequest) error
	Mute(ctx context.Context, request *model.Mute) error
	Unmute(ctx context.Context, request *model.Mute) error
	ReturnFollowRequestsForUser(ctx context.Context, bearer string, objectId string) (bool, error)
	ReturnFollowRequests(ctx context.Context, bearer string) (model.FollowedUsersResponse, error)
	AcceptFollowRequest(ctx context.Context, bearer string, userId string) error
	IsMuted(ctx context.Context, mute model.Mute) (bool, error)
	GetRecommendedUsers(userId string) (model.RecommendedUsersResponse, error)
}

type relationshipClient struct{}

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

	req, err := http.NewRequest("POST", baseRelationshipUrl+"/user", bytes.NewBuffer(jsonUserRequest))
	req.Header.Add("Content-Type", "application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	//resp, err := http.Post(baseRelationshipUrl + "/user","application/json", bytes.NewBuffer(jsonUserRequest))

	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"name": user.Name,
				"surname":  user.Surname,
				"email":    user.Email,
				"username": user.Username}).Error("Relationship-service not available")
			return err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"name": user.Name, "surname": user.Surname, "email": user.Email, "username": user.Username}).Error("Relationship-service user registration")

		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(message)
	}
	return nil
}

func (r relationshipClient) GetFollowedUsers(ctx context.Context, userId string) (model.FollowedUsersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "RelationshipClientGetFollowedUsers")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/followed-users/%s", baseRelationshipUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Relationship-service not available")
			return model.FollowedUsersResponse{}, errors.New("post not found")
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Relationship-service get followed users")
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

func (r relationshipClient) GetFollowingUsers(ctx context.Context, userId string) (model.FollowedUsersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "RelationshipClientGetFollowingUsers")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/following-users/%s", baseRelationshipUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Relationship-service not available")
			return model.FollowedUsersResponse{}, errors.New("post not found")
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Relationship-service get following users")
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

func (r relationshipClient) FollowRequest(ctx context.Context, request *model.FollowRequest) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "RelationshipClientFollowRequest")
	defer span.Finish()

	jsonRequest, _ := json.Marshal(request)

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/follow", baseRelationshipUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": request.SubjectId, "object_id": request.ObjectId}).Error("Relationship-service not available")
			return false, errors.New("user not found")
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": request.SubjectId, "object_id": request.ObjectId}).Error("Relationship-service follow user")
		return false, errors.New("user not found")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var followRequest bool
	json.Unmarshal(bodyBytes, &followRequest)

	return followRequest, nil
}

func (r relationshipClient) Mute(ctx context.Context, request *model.Mute) error {
	span := tracer.StartSpanFromContext(ctx, "RelationshipClientMute")
	defer span.Finish()

	jsonRequest, _ := json.Marshal(request)

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/mute", baseRelationshipUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	_, err = client.Do(req)

	if err != nil {
		return errors.New("user not found")
	}

	return nil
}

func (r relationshipClient) Unmute(ctx context.Context, request *model.Mute) error {
	span := tracer.StartSpanFromContext(ctx, "RelationshipClientUnmute")
	defer span.Finish()

	jsonRequest, _ := json.Marshal(request)

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/unmute", baseRelationshipUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	_, err = client.Do(req)

	if err != nil {
		return errors.New("user not found")
	}

	return nil
}

func (r relationshipClient) Unfollow(ctx context.Context, request *model.FollowRequest) error {
	span := tracer.StartSpanFromContext(ctx, "RelationshipClientUnfollow")
	defer span.Finish()

	jsonRequest, _ := json.Marshal(request)

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/unfollow", baseRelationshipUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	req = req.WithContext(ctx)

	client := &http.Client{}
	tracer.Inject(span, req)
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": request.SubjectId, "object_id": request.ObjectId}).Error("Relationship-service not available")
			return err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": request.SubjectId, "object_id": request.ObjectId}).Error("Relationship-service unfollow user")
		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(message)
	}
	return nil
}

func (r relationshipClient) ReturnFollowRequestsForUser(ctx context.Context, bearer string, objectId string) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "RelationshipClientReturnFollowRequestsForUser")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/follow-requests/%s", baseRelationshipUrl, objectId), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id": objectId}).Error("Relationship-service not available")
			return false, errors.New("user not found")
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": objectId}).Error("Relationship-service get follow requests")
		return false, errors.New("user not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LoggingEntry.Error("Parsing relationship-service get follow requests response")
		return false, err
	}
	var isSentRequest bool
	json.Unmarshal(bodyBytes, &isSentRequest)
	fmt.Println(isSentRequest)

	return isSentRequest, nil
}

func (r relationshipClient) ReturnFollowRequests(ctx context.Context, bearer string) (model.FollowedUsersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "RelationshipClientReturnFollowRequests")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/follow-requests", baseRelationshipUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.Error("Relationship-service not available")
			return model.FollowedUsersResponse{}, errors.New("user not found")
		}

		logger.LoggingEntry.Error("Relationship-service get follow requests")
		return model.FollowedUsersResponse{}, errors.New("user not found")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LoggingEntry.Error("Parsing relationship-service get follow requests response")
		return model.FollowedUsersResponse{}, err
	}

	var users model.FollowedUsersResponse
	_ = json.Unmarshal(bodyBytes, &users)

	fmt.Println(users)

	return users, nil
}

func (r relationshipClient) AcceptFollowRequest(ctx context.Context, bearer string, userId string) error {
	span := tracer.StartSpanFromContext(ctx, "RelationshipClientAcceptFollowRequest")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"POST", fmt.Sprintf("%s/accept-follow-request/%s", baseRelationshipUrl, userId), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Relationship-service not available")
			return errors.New("user not found")
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Relationship-service accept follow request")
		return errors.New("user not found")
	}
	return nil
}

func (r relationshipClient) IsMuted(ctx context.Context, mute model.Mute) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "RelationshipClientIsMuted")
	defer span.Finish()

	jsonMute, _ := json.Marshal(mute)

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/is-muted", baseRelationshipUrl), bytes.NewBuffer(jsonMute))
	req.Header.Set("Content-Type", "application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp == nil {
		return false, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	var isMuted bool
	_ = json.Unmarshal(bodyBytes, &isMuted)

	return isMuted, err
}

func (r relationshipClient) GetRecommendedUsers(userId string) (model.RecommendedUsersResponse, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/recommended-users/%s", baseRelationshipUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Relationship-service not available")
			return model.RecommendedUsersResponse{}, errors.New("post not found")
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Relationship-service get followed users")
		fmt.Println(resp.StatusCode)
		return model.RecommendedUsersResponse{}, errors.New("post not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.RecommendedUsersResponse{}, err
	}

	var users model.RecommendedUsersResponse
	_ = json.Unmarshal(bodyBytes, &users)

	return users, nil
}
