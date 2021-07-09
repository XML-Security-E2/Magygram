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

type PostClient interface {
	GetPostsFirstImage(ctx context.Context, postId string) (*model.Media, error)
	GetUsersPostsCount(ctx context.Context, userId string) (int, error)
	EditPostOwnerInfo(ctx context.Context, bearer string, userInfo model.UserInfo) error
	EditLikedByInfo(ctx context.Context, bearer string,userInfo model.UserInfoEdit) error
	EditDislikedByInfo(ctx context.Context, bearer string, userInfo model.UserInfoEdit) error
	EditCommentedByInfo(ctx context.Context, bearer string, userInfo model.UserInfoEdit) error
}

type postClient struct {}

func NewPostClient() PostClient {
	basePostUrl = fmt.Sprintf("%s%s:%s/api/posts", conf.Current.Postservice.Protocol, conf.Current.Postservice.Domain, conf.Current.Postservice.Port)
	return &postClient{}
}

var (
	basePostUrl = ""
)


func (a postClient) EditPostOwnerInfo(ctx context.Context, bearer string, userInfo model.UserInfo) error {
	span := tracer.StartSpanFromContext(ctx, "PostClientEditPostOwnerInfo")
	defer span.Finish()

	jsonRequest, _ := json.Marshal(userInfo)

	req, err := http.NewRequestWithContext(ctx,"PUT", fmt.Sprintf("%s/user-info", basePostUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return err
	}

	return nil
}

func (a postClient) EditLikedByInfo(ctx context.Context, bearer string, userInfo model.UserInfoEdit) error {
	span := tracer.StartSpanFromContext(ctx, "PostClientEditLikedByInfo")
	defer span.Finish()

	jsonRequest, _ := json.Marshal(userInfo)

	req, err := http.NewRequestWithContext(ctx,"PUT", fmt.Sprintf("%s/liked-by/user-info", basePostUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return err
	}

	return nil
}

func (a postClient) EditDislikedByInfo(ctx context.Context, bearer string, userInfo model.UserInfoEdit) error {
	span := tracer.StartSpanFromContext(ctx, "PostClientEditDislikedByInfo")
	defer span.Finish()

	jsonRequest, _ := json.Marshal(userInfo)

	req, err := http.NewRequestWithContext(ctx,"PUT", fmt.Sprintf("%s/disliked-by/user-info", basePostUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return err
	}

	return nil
}


func (a postClient) EditCommentedByInfo(ctx context.Context, bearer string, userInfo model.UserInfoEdit) error {
	span := tracer.StartSpanFromContext(ctx, "PostClientEditCommentedByInfo")
	defer span.Finish()

	jsonRequest, _ := json.Marshal(userInfo)

	req, err := http.NewRequestWithContext(ctx,"PUT", fmt.Sprintf("%s/commented/user-info", basePostUrl), bytes.NewBuffer(jsonRequest))
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

func (a postClient) GetPostsFirstImage(ctx context.Context, postId string) (*model.Media, error) {
	span := tracer.StartSpanFromContext(ctx, "PostClientGetPostsFirstImage")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/%s/image", basePostUrl, postId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

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

func (a postClient) GetUsersPostsCount(ctx context.Context, userId string) (int, error) {
	span := tracer.StartSpanFromContext(ctx, "PostClientGetUsersPostsCount")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/%s/count", basePostUrl, userId), nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

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