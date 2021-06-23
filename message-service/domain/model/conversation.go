package model

import (
	"errors"
	"github.com/beevik/guid"
	"time"
)

type Conversation struct {
	Id string `json:"id"`
	ParticipantOneId string `json:"participantOneId"`
	ParticipantTwoId string `json:"participantTwoId"`
	Messages []Message `json:"messages"`
	LastMessage Message `json:"lastMessage"`
	LastMessageUserInfo UserInfo `json:"lastMessageUserInfo"`
}

type Message struct {
	MessageTo UserInfo `json:"messageTo"`
	MessageFrom UserInfo `json:"messageFrom"`
	MessageType string `json:"messageType"`
	Text string `json:"text"`
	ContentUrl string `json:"contentUrl"`
	Timestamp time.Time `json:"timestamp"`
	Viewed  bool `json:"viewed"`
	ViewedMedia  bool `json:"viewedMedia"`
}

type MessageSentRequest struct {
	MessageTo string `json:"messageTo"`
	MessageType string `json:"messageType"`
	Text string `json:"text"`
	ContentUrl string `json:"contentUrl"`
}

type MessageRequest struct {
	Id string `json:"id"`
	MessageTo string `json:"messageTo"`
	MessageFrom UserInfo `json:"messageFrom"`
	MessageType string `json:"messageType"`
	Text string `json:"text"`
	ContentUrl string `json:"contentUrl"`
	Timestamp time.Time `json:"timestamp"`
	MessageRequestStatus  MessageRequestStatus `json:"messageRequestStatus"`
}

type MessageRequestStatus string

var (
	ConvPrefix = "conversation"
	MessageRequestPrefix = "message-request"
)

const(
	PENDING = iota
	ACCEPTED
	DENIED
	DELETED
)

type MessageType string

const(
	TEXT = iota
	PHOTO
	VIDEO
	POST
	STORY
)

func NewConversation(message *Message) *Conversation {
	return &Conversation{
		Id:                  guid.New().String(),
		ParticipantOneId:    message.MessageFrom.Id,
		ParticipantTwoId:    message.MessageTo.Id,
		Messages:            []Message{*message},
		LastMessage:         *message,
		LastMessageUserInfo: message.MessageFrom,
	}
}

func NewMessage(messageRequest *MessageSentRequest, messageType MessageType, messageFrom UserInfo, messageTo UserInfo) (*Message, error) {
	err := validateMessageTypeEnums(messageType)
	if err != nil {
		return nil, err
	}

	return &Message{
		MessageTo:   messageTo,
		MessageFrom: messageFrom,
		MessageType: messageRequest.MessageType,
		Text:        messageRequest.Text,
		ContentUrl:  messageRequest.ContentUrl,
		Timestamp:   time.Now(),
		Viewed:      false,
		ViewedMedia: false,
	}, nil
}

func validateMessageTypeEnums(messageType MessageType) error {
	switch messageType {
	case "TEXT", "PHOTO", "VIDEO", "POST", "STORY":
		return nil
	}
	return errors.New("invalid message type")
}