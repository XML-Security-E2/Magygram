package intercomm

import (
	"ads-service/conf"
	"ads-service/domain/model"
	"ads-service/tracer"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type StoryClient interface {
	GetStoryMedia(ctx context.Context, storyIds []string) ([]*model.IdMediaWebsiteResponse, error)
	CreateStoryCampagin(ctx context.Context, bearer string, request *multipart.FileHeader) (string, error)
}

type storyClient struct {}

func NewStoryClient() StoryClient {
	baseStorytUrl = fmt.Sprintf("%s%s:%s/api/story", conf.Current.Storyservice.Protocol, conf.Current.Storyservice.Domain, conf.Current.Storyservice.Port)
	return &storyClient{}
}

var (
	baseStorytUrl = ""
)


func (s storyClient) GetStoryMedia(ctx context.Context, storyIds []string) ([]*model.IdMediaWebsiteResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "StoryClientGetStoryMedia")
	defer span.Finish()

	requ := IdsRequests{Users: storyIds}
	jsonRequest, _ := json.Marshal(requ)

	req, err := http.NewRequestWithContext(ctx,"POST", fmt.Sprintf("%s/media/first/preview", baseStorytUrl), bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

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

	return postImage, nil
}



func (s storyClient) CreateStoryCampagin(ctx context.Context, bearer string, media *multipart.FileHeader) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "StoryClientCreateStoryCampagin")
	defer span.Finish()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	var fw io.Writer

	src, err := media.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	fw, err = writer.CreateFormFile("images", media.Filename)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(fw, src)
	if err != nil {
		return "", err
	}
	writer.Close()
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx,"POST", fmt.Sprintf("%s/campaign/agent", baseStorytUrl), bytes.NewReader(body.Bytes()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			return "", err
		}

		return "", errors.New("error while creating post campaign")
	}


	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var contentId string
	json.Unmarshal(bodyBytes, &contentId)

	return contentId, nil
}
