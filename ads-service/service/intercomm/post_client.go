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

type PostClient interface {
	GetPostsFirstMedia(postIds []string) ([]*model.IdMediaWebsiteResponse, error)
}

type postClient struct {}

func NewPostClient() PostClient {
	basePostUrl = fmt.Sprintf("%s%s:%s/api/posts", conf.Current.Postservice.Protocol, conf.Current.Postservice.Domain, conf.Current.Postservice.Port)
	return &postClient{}
}

var (
	basePostUrl = ""
)

type IdsRequests struct {
	Users []string
}


func (p postClient) GetPostsFirstMedia(postIds []string) ([]*model.IdMediaWebsiteResponse, error) {
	requ := IdsRequests{Users: postIds}
	jsonRequest, _ := json.Marshal(requ)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/media/first/preview", basePostUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			return []*model.IdMediaWebsiteResponse{}, nil
		}

		return []*model.IdMediaWebsiteResponse{}, errors.New("posts not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []*model.IdMediaWebsiteResponse{}, err
	}
	var postImage []*model.IdMediaWebsiteResponse
	json.Unmarshal(bodyBytes, &postImage)
	if postImage == nil {
		return []*model.IdMediaWebsiteResponse{}, nil
	}

	return postImage, nil

}
