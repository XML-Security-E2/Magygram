package intercomm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"post-service/conf"
	"post-service/domain/model"
)

type UserClient interface {
	GetLoggedUserInfo(bearer string) (*model.UserInfo,error)
	MapPostsToFavourites(bearer string, postIds []string) ([]*model.PostIdFavouritesFlag,error)
	IsUserPrivate(userId string) (bool, error)
}

type userClient struct {}

func NewUserClient() UserClient {
	baseUsersUrl = fmt.Sprintf("%s%s:%s/api/users", conf.Current.Userservice.Protocol, conf.Current.Userservice.Domain, conf.Current.Userservice.Port)
	return &userClient{}
}

var (
	baseUsersUrl = ""
)

func (u userClient) GetLoggedUserInfo(bearer string) (*model.UserInfo, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/logged", baseUsersUrl), nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return &model.UserInfo{}, errors.New("unauthorized")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &model.UserInfo{}, err
	}
	var userInfo model.UserInfo
	_ = json.Unmarshal(bodyBytes, &userInfo)

	return &userInfo, nil
}

func (u userClient) MapPostsToFavourites(bearer string, postIds []string) ([]*model.PostIdFavouritesFlag, error) {

	jsonStr, err:= json.Marshal(postIds)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/collections/check-favourites", baseUsersUrl), bytes.NewReader(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("unauthorized")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var postIdFav []*model.PostIdFavouritesFlag
	_ = json.Unmarshal(bodyBytes, &postIdFav)

	return postIdFav, nil
}


func (u userClient) IsUserPrivate(userId string) (bool, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/is-private", baseUsersUrl, userId), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return false, errors.New("user not found")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var isPrivate bool
	json.Unmarshal(bodyBytes, &isPrivate)

	return isPrivate, nil
}
