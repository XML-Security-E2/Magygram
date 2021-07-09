package intercomm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"user-service/conf"
	"user-service/domain/model"
	"user-service/logger"
	"user-service/tracer"
)

type StoryClient interface {
	GetStoryHighlightIfValid(ctx context.Context, bearer string, request *model.HighlightRequest) (*model.HighlightImageWithMedia, error)
	EditStoryOwnerInfo(ctx context.Context, bearer string, userInfo model.UserInfo) error
}

type storyClient struct {}

func NewStoryClient() StoryClient {
	baseStorytUrl = fmt.Sprintf("%s%s:%s/api/story", conf.Current.Storyservice.Protocol, conf.Current.Storyservice.Domain, conf.Current.Storyservice.Port)
	return &storyClient{}
}

var (
	baseStorytUrl = ""
)


func (s storyClient) EditStoryOwnerInfo(ctx context.Context, bearer string, userInfo model.UserInfo) error {
	span := tracer.StartSpanFromContext(ctx, "StoryClientEditStoryOwnerInfo")
	defer span.Finish()

	jsonRequest, _ := json.Marshal(userInfo)

	req, err := http.NewRequestWithContext(ctx,"PUT", fmt.Sprintf("%s/user-info", baseStorytUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	tracer.Inject(span, req)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return err
	}

	return nil
}


func (s storyClient) GetStoryHighlightIfValid(ctx context.Context, bearer string, request *model.HighlightRequest) (*model.HighlightImageWithMedia, error) {
	span := tracer.StartSpanFromContext(ctx, "StoryClientGetStoryHighlightIfValid")
	defer span.Finish()

	jsonRequest, _ := json.Marshal(request)

	req, err := http.NewRequestWithContext(ctx,"POST", fmt.Sprintf("%s/highlights", baseStorytUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	tracer.Inject(span, req)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"highlights_name": request.Name, "story_ids" : request.StoryIds}).Error("Story-service not available")
			return nil, err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"highlights_name": request.Name, "story_ids" : request.StoryIds}).Error("Story-service get stories for highlights")
		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(message)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var highlight model.HighlightImageWithMedia
	json.Unmarshal(bodyBytes, &highlight)
	return &highlight, nil
}
