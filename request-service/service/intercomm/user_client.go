package intercomm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"request-service/conf"
	"request-service/domain/model"
)

type UserClient interface {
	VerifyAccount(verifyAccountDTO model.VerifyAccountDTO) error
	RegisterAgent(agentRegistrationDTO model.AgentRegistrationDTO) error
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
		return err
	}

	return nil
}

func (u userClient) RegisterAgent(agentRegistrationDTO model.AgentRegistrationDTO) error {
	json, err := json.Marshal(agentRegistrationDTO)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/agent", baseUsersUrl), bytes.NewBuffer(json))

	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 201 {
		if resp.StatusCode != 201{
			return errors.New("Nije moguce odobriti datog korisnika")
		}

		return err
	}

	return nil
}
