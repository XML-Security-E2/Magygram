package intercomm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"request-service/conf"
	"request-service/domain/model"
)

type UserClient interface {
	VerifyAccount(verifyAccountDTO model.VerifyAccountDTO) error
}

type userClient struct {}

func NewUserClient() UserClient {
	baseUsersUrl = fmt.Sprintf("%s%s:%s/api/users", conf.Current.Userservice.Protocol, conf.Current.Userservice.Domain, conf.Current.Userservice.Port)
	return &userClient{}
}

var (
	baseUsersUrl = ""
)

func (u userClient) VerifyAccount(verifyAccountDTO model.VerifyAccountDTO) error {
	json, err := json.Marshal(verifyAccountDTO)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/verify", baseUsersUrl), bytes.NewBuffer(json))

	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		log.Println("test6")
		log.Println(resp.StatusCode)
		return err
	}
	log.Println("7")

	log.Println(resp.StatusCode)

	return nil
}
