package model

import (
	"errors"
	"github.com/beevik/guid"
	"time"
)

type Conversation struct {
	Id string `json:"id"`
	ParticipantOne UserInfo `json:"participantOneId"`
	ParticipantTwo UserInfo `json:"participantTwoId"`
	Messages []Message `json:"messages"`
	LastMessage Message `json:"lastMessage"`
	LastMessageUserId string `json:"lastMessageUserId"`
}

type ConversationResponse struct {
	Id string `json:"id"`
	Participant UserInfo `json:"participant"`
	LastMessage Message `json:"lastMessage"`
	LastMessageUserId string `json:"lastMessageUserId"`
}

type Message struct {
	MessageToId string `json:"messageToId"`
	MessageFromId string `json:"messageFromId"`
	MessageType MessageType `json:"messageType"`
	Text string `json:"text"`
	ContentUrl string `json:"contentUrl"`
	Timestamp time.Time `json:"timestamp"`
	Viewed  bool `json:"viewed"`
	ViewedMedia  bool `json:"viewedMedia"`
}

type MessagesResponse struct {
	UserInfo UserInfo `json:"userInfo"`
	Messages []Message `json:"messages"`
}

type MessageSentRequest struct {
	MessageTo string `json:"messageTo"`
	MessageType MessageType `json:"messageType"`
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

func NewConversation(message *Message, participantOne UserInfo, participantTwo UserInfo) *Conversation {
	return &Conversation{
		Id:                  guid.New().String(),
		ParticipantOne:    	 participantOne,
		ParticipantTwo:    	 participantTwo,
		Messages:            []Message{*message},
		LastMessage:         *message,
		LastMessageUserId:   message.MessageFromId,
	}
}

func NewMessage(messageRequest *MessageSentRequest, messageFrom string) (*Message, error) {
	err := validateMessageTypeEnums(messageRequest.MessageType)
	if err != nil {
		return nil, err
	}

	return &Message{
		MessageToId:   messageRequest.MessageTo,
		MessageFromId: messageFrom,
		MessageType:   messageRequest.MessageType,
		Text:          messageRequest.Text,
		ContentUrl:    messageRequest.ContentUrl,
		Timestamp:     time.Now(),
		Viewed:        false,
		ViewedMedia:   false,
	}, nil
}

func validateMessageTypeEnums(messageType MessageType) error {
	switch messageType {
	case "TEXT", "PHOTO", "VIDEO", "POST", "STORY":
		return nil
	}
	return errors.New("invalid message type")
}