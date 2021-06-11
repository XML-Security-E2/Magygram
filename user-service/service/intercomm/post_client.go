package intercomm

import (
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
)

type PostClient interface {
	GetPostsFirstImage(postId string) (*model.Media, error)
	GetUsersPostsCount(userId string) (int, error)
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
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"post_id": postId}).Error("Post-service not available")
			return nil, err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"post_id": postId}).Error("Post-service get posts first image")
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

func (a postClient) GetUsersPostsCount(userId string) (int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/count", basePostUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Post-service not available")
			return 0, err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("Post-service get post count")
		fmt.Println(resp.StatusCode)
		return 0, errors.New("posts not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var postsCount int
	json.Unmarshal(bodyBytes, &postsCount)

	return postsCount, nil
}