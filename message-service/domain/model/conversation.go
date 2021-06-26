package model

import (
	"errors"
	"github.com/beevik/guid"
	"mime/multipart"
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

type MessageSendResponse struct {
	IsMessageRequest bool `json:"isMessageRequest"`
	Conversation *ConversationResponse `json:"conversation"`
	ConversationRequest *ConversationRequest `json:"conversationRequest"`
}

type ConversationResponse struct {
	Id string `json:"id"`
	Participant UserInfo `json:"participant"`
	LastMessage Message `json:"lastMessage"`
	LastMessageUserId string `json:"lastMessageUserId"`
}

type Message struct {
	Id string `json:"id"`
	MessageToId string `json:"messageToId"`
	MessageFromId string `json:"messageFromId"`
	MessageType MessageType `json:"messageType"`
	Text string `json:"text"`
	ContentId string `json:"contentId"`
	Media *Media `json:"media"`
	Timestamp time.Time `json:"timestamp"`
	Viewed  bool `json:"viewed"`
	ViewedMedia  bool `json:"viewedMedia"`
}

type Media struct {
	Url       string `json:"url"`
	MediaType string `json:"mediaType"`
}

type FollowedUsersResponse struct {
	Users []string
}

type MessagesResponse struct {
	UserInfo UserInfo `json:"userInfo"`
	Messages []Message `json:"messages"`
}

type MessageSentRequest struct {
	MessageTo string `json:"messageTo"`
	MessageType MessageType `json:"messageType"`
	Media *multipart.FileHeader `json:"media"`
	Text string `json:"text"`
	ContentId string `json:"contentId"`
}

type ConversationRequest struct {
	Id string `json:"id"`
	RequestFrom UserInfo `json:"participantOne"`
	RequestTo UserInfo `json:"participantTwo"`
	Messages []Message `json:"messages"`
	RequestStatus  MessageRequestStatus `json:"requestStatus"`
	LastMessage Message `json:"lastMessage"`
	LastMessageUserId string `json:"lastMessageUserId"`
}

type MessageRequest struct {
	Id string `json:"id"`
	MessageTo UserInfo `json:"messageTo"`
	MessageFrom UserInfo `json:"messageFrom"`
	MessageType MessageType `json:"messageType"`
	Text string `json:"text"`
	Media *Media `json:"media"`
	ContentId string `json:"contentId"`
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
	DENIED
)

type MessageType string

const(
	TEXT = iota
	MEDIA
	POST
	STORY
)

func NewMessageRequest(message *Message, messageFrom UserInfo, messageTo UserInfo) *MessageRequest {
	return &MessageRequest{
		Id:                   guid.New().String(),
		MessageTo:            messageFrom,
		MessageFrom:          messageTo,
		MessageType:          message.MessageType,
		Text:                 message.Text,
		Media:                message.Media,
		ContentId:            message.ContentId,
		Timestamp:            time.Now(),
		MessageRequestStatus: "PENDING",
	}
}

func NewConversationRequest(message *Message, requestFrom UserInfo, requestTo UserInfo) *ConversationRequest {
	return &ConversationRequest{
		Id:                  guid.New().String(),
		RequestFrom:    	 requestFrom,
		RequestTo:    	     requestTo,
		Messages:            []Message{*message},
		LastMessage:         *message,
		LastMessageUserId:   message.MessageFromId,
		RequestStatus:       "PENDING",
	}
}

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

func NewMessage(messageRequest *MessageSentRequest, messageFrom string, media *Media) (*Message, error) {
	err := validateMessageTypeEnums(messageRequest.MessageType)
	if err != nil {
		return nil, err
	}

	return &Message{
		Id: 		   guid.New().String(),
		MessageToId:   messageRequest.MessageTo,
		MessageFromId: messageFrom,
		MessageType:   messageRequest.MessageType,
		Text:          messageRequest.Text,
		ContentId:     messageRequest.ContentId,
		Timestamp:     time.Now(),
		Media: 		   media,
		Viewed:        false,
		ViewedMedia:   false,
	}, nil
}

func validateMessageTypeEnums(messageType MessageType) error {
	switch messageType {
	case "TEXT", "MEDIA", "POST", "STORY":
		return nil
	}
	return errors.New("invalid message type")
}