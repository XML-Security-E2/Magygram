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

type StoryClient interface {
	GetStoryHighlightIfValid(bearer string, request *model.HighlightRequest) (*model.HighlightImageWithMedia, error)
}

type storyClient struct {}

func NewStoryClient() StoryClient {
	baseStorytUrl = fmt.Sprintf("%s%s:%s/api/story", conf.Current.Storyservice.Protocol, conf.Current.Storyservice.Domain, conf.Current.Storyservice.Port)
	return &storyClient{}
}

var (
	baseStorytUrl = ""
)

func (s storyClient) GetStoryHighlightIfValid(bearer string, request *model.HighlightRequest) (*model.HighlightImageWithMedia, error) {
	jsonRequest, _ := json.Marshal(request)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/highlights", baseStorytUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
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