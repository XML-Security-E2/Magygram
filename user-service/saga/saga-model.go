package saga

import (
	"encoding/json"
)

type RegisterUserMessage struct {
	Service       string         `json:"service"`
	SenderService string         `json:"sender_service"`
	UserType 	  string          `json:"user_type"`
	Action        string         `json:"action"`
	User          UserRequest       `json:"user_request"`
	ImageByte []byte `json:"image"`
}

type UserRequest struct {
	Id               string `json:"id"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
}

func (m RegisterUserMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
