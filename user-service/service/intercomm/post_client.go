package intercomm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"user-service/conf"
	"user-service/domain/model"
)

type PostClient interface {
	GetPostsFirstImage(postId string) (*model.Media, error)
}

type postClient struct {}

func NewPostClient() PostClient {
	basePostUrl = fmt.Sprintf("%s%s:%s/api/posts", conf.Current.Postservice.Protocol, conf.Current.Postservice.Domain, conf.Current.Postservice.Port)
	return &postClient{}
}

var (
	basePostUrl = ""
)

func (a postClient) GetPostsFirstImage(postId string) (*model.Media, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/image", basePostUrl, postId), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return nil, errors.New("post not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var postImage model.Media
	json.Unmarshal(bodyBytes, &postImage)

	return &postImage, nil
}