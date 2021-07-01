package saga

import (
	"encoding/json"
)

const (
	UserChannel    string = "UserChannel"
	AuthChannel    string = "AuthChannel"
	RelationshipChannel string = "RelationshipChannel"
	ReplyChannel    string = "ReplyChannel"
	ServiceUser    string = "User"
	ServiceAuth    string = "Auth"
	ServiceRelationship string = "Relationship"
	ActionStart     string = "Start"
	ActionDone      string = "DoneMsg"
	ActionError     string = "ErrorMsg"
	ActionRollback  string = "RollbackMsg"
)

type RegisterUserMessage struct {
	Service       string         `json:"service"`
	SenderService string         `json:"sender_service"`
	Action        string         `json:"action"`
	User       UserRequest            `json:"user_request"`
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
