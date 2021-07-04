package intercomm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"net/http"
	"story-service/conf"
	"story-service/domain/model"
)

type AdsClient interface {
	CreatePostCampaign(bearer string, campaignReq *model.CampaignRequest) error
	GetAllActiveAgentsStoryCampaigns(bearer string) ([]string, error)
	DeleteCampaign(bearer string, postId string) error
	GetStoryCampaignSuggestion(bearer string,count string) ([]string,error)
	UpdateCampaignVisitor(bearer string,storyId string) error

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

func (a adsClient) CreatePostCampaign(bearer string, campaignReq *model.CampaignRequest) error {
	jsonStr, err:= json.Marshal(campaignReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/campaign", baseAdsUrl), bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

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

func (a adsClient) DeleteCampaign(bearer string, postId string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/campaign/story/%s", baseAdsUrl, postId), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return err
	}

	return nil
}

func (a adsClient) GetAllActiveAgentsStoryCampaigns(bearer string) ([]string, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/campaign/story", baseAdsUrl), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

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

func (a adsClient) GetStoryCampaignSuggestion(bearer string,count string) ([]string,error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/campaign/story/suggestion/" + count, baseAdsUrl), nil)

	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

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

func (a adsClient) UpdateCampaignVisitor(bearer string, storyId string) error {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/campaign/story/visited/%s", baseAdsUrl, storyId), nil)
	req.Header.Add("Authorization", bearer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return err
	}

	fmt.Println(resp.StatusCode)
	return nil
}