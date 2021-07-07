package intercomm

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"magyAgent/conf"
	"net/http"
)

type MediaClient interface {
	GetMedia(mediaName string) ([]byte, error)
}

type mediaClient struct {}

func NewMediaClient() MediaClient {
	return &mediaClient{}
}


func (m mediaClient) GetMedia(mediaName string) ([]byte, error) {
	fmt.Println(mediaName)
	req, err := http.NewRequest("GET",  mediaName, nil)
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			fmt.Println(err.Error())

			return nil, err
		}
		fmt.Println(resp.StatusCode)

		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}