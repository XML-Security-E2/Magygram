package intercomm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"net/http"
	"post-service/conf"
	"post-service/domain/model"
	"post-service/tracer"
)

type AdsClient interface {
	CreatePostCampaign(ctx context.Context, bearer string, campaignReq *model.CampaignRequest) error
	GetAllActiveAgentsPostCampaigns(ctx context.Context, bearer string) ([]string,error)
	DeleteCampaign(ctx context.Context, bearer string, postId string) error
	GetPostCampaignSuggestion(ctx context.Context, bearer string,count string) ([]string,error)
	UpdateCampaignVisitor(ctx context.Context, bearer string,postId string) error
}

type adsClient struct {}


func NewAdsClient() AdsClient {
	baseAdsUrl = fmt.Sprintf("%s%s:%s/api/ads", conf.Current.Adsservice.Protocol, conf.Current.Adsservice.Domain, conf.Current.Adsservice.Port)
	return &adsClient{}
}

var (
	baseAdsUrl = ""
)

type errMessage struct {
	Message string `json:"message"`
}

func (a adsClient) CreatePostCampaign(ctx context.Context, bearer string, campaignReq *model.CampaignRequest) error {
	span := tracer.StartSpanFromContext(ctx, "AdsClientCreatePostCampaign")
	defer span.Finish()

	jsonStr, err:= json.Marshal(campaignReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx,"POST", fmt.Sprintf("%s/campaign", baseAdsUrl), bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			return errors.New("internal server error")
		}
		message, err := getErrorMessageFromRequestBody(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(message)
	}

	return nil
}

func (a adsClient) GetAllActiveAgentsPostCampaigns(ctx context.Context, bearer string) ([]string,error) {
	span := tracer.StartSpanFromContext(ctx, "AdsClientGetAllActiveAgentsPostCampaigns")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/campaign/post", baseAdsUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var campaigns []string
	_ = json.Unmarshal(bodyBytes, &campaigns)

	return campaigns, nil
}

func (a adsClient) DeleteCampaign(ctx context.Context, bearer string, postId string) error {
	span := tracer.StartSpanFromContext(ctx, "AdsClientDeleteCampaign")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"DELETE", fmt.Sprintf("%s/campaign/post/%s", baseAdsUrl, postId), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return err
	}

	fmt.Println(resp.StatusCode)
	return nil
}

func getErrorMessageFromRequestBody(body io.ReadCloser) (string ,error){
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}
	result := &errMessage{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return "", err
	}
	return result.Message, nil
}

func (a adsClient) GetPostCampaignSuggestion(ctx context.Context, bearer string,count string) ([]string,error) {
	span := tracer.StartSpanFromContext(ctx, "AdsClientGetPostCampaignSuggestion")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"GET", fmt.Sprintf("%s/campaign/post/suggestion/" + count, baseAdsUrl), nil)

	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var campaigns []string
	_ = json.Unmarshal(bodyBytes, &campaigns)

	return campaigns, nil
}

func (a adsClient) UpdateCampaignVisitor(ctx context.Context, bearer string, postId string) error {
	span := tracer.StartSpanFromContext(ctx, "AdsClientUpdateCampaignVisitor")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx,"PUT", fmt.Sprintf("%s/campaign/post/visited/%s", baseAdsUrl, postId), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return err
	}

	fmt.Println(resp.StatusCode)
	return nil
}