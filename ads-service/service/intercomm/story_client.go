package intercomm

import (
	"ads-service/conf"
	"ads-service/domain/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

type StoryClient interface {
	GetStoryMedia(storyIds []string) ([]*model.IdMediaWebsiteResponse, error)
}

type storyClient struct {}

func NewStoryClient() StoryClient {
	baseStorytUrl = fmt.Sprintf("%s%s:%s/api/story", conf.Current.Storyservice.Protocol, conf.Current.Storyservice.Domain, conf.Current.Storyservice.Port)
	return &storyClient{}
}

var (
	baseStorytUrl = ""
)


func (s storyClient) GetStoryMedia(storyIds []string) ([]*model.IdMediaWebsiteResponse, error) {
	requ := IdsRequests{Users: storyIds}
	jsonRequest, _ := json.Marshal(requ)


	req, err := http.NewRequest("POST", fmt.Sprintf("%s/media/first/preview", baseStorytUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			return []*model.IdMediaWebsiteResponse{}, nil
		}

		return []*model.IdMediaWebsiteResponse{}, errors.New("stories not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []*model.IdMediaWebsiteResponse{}, nil
	}
	var postImage []*model.IdMediaWebsiteResponse
	json.Unmarshal(bodyBytes, &postImage)

	if postImage == nil{
		return []*model.IdMediaWebsiteResponse{}, nil
	}

	return postImage, nil}