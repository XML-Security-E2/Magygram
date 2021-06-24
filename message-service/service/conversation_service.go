package service

import (
	"context"
	"fmt"
	"math"
	"message-service/domain/model"
	"message-service/domain/repository"
	"message-service/domain/service-contracts"
	"message-service/service/intercomm"
)

var (
	limitConv int64 = math.MaxInt64
)

type conversationService struct {
	repository.ConversationRepository
	intercomm.AuthClient
	intercomm.UserClient
}

func NewConversationService(r repository.ConversationRepository, ac intercomm.AuthClient, uc intercomm.UserClient) service_contracts.ConversationService {
	return &conversationService{r, ac, uc}
}

//PROVERITI ZA PRIVATNOST ITD
func (c conversationService) SendMessage(ctx context.Context, bearer string, messageRequest *model.MessageSentRequest) (*model.ConversationResponse, error) {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	message, err := model.NewMessage(messageRequest, loggedId)
	if err != nil {
		return nil, err
	}

	conversation, err := c.ConversationRepository.GetConversationForUser(ctx, loggedId, messageRequest.MessageTo, limitConv)
	if err == nil && conversation == nil {
		fmt.Println("USAOOO")
		participantOne, err := c.UserClient.GetLoggedUserInfo(bearer)
		if err != nil {
			return nil, err
		}
		participantTwo, err := c.UserClient.GetUsersInfo(messageRequest.MessageTo)
		if err != nil {
			return nil, err
		}

		conv := model.NewConversation(message, *participantOne, *participantTwo)
		err = c.ConversationRepository.CreateConversation(ctx, conv)
		if err != nil {
			return nil, err
		}

		return &model.ConversationResponse{
			Id:                conv.Id,
			Participant:       conv.ParticipantTwo,
			LastMessage:       *message,
			LastMessageUserId: loggedId,
		}, nil
	}

	fmt.Println(len(conversation.Messages))

	conversation.Messages = append(conversation.Messages, *message)
	conversation.LastMessage = *message
	conversation.LastMessageUserId = loggedId

	fmt.Println(len(conversation.Messages))

	err = c.ConversationRepository.Update(ctx, conversation)
	if err != nil {
		return nil, err
	}

	return &model.ConversationResponse{
		Id:                conversation.Id,
		Participant:       conversation.ParticipantTwo,
		LastMessage:       *message,
		LastMessageUserId: loggedId,
	}, nil
}

func (c conversationService) GetAllConversationsForUser(ctx context.Context, bearer string) ([]*model.ConversationResponse, error) {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}
	var convResp []*model.ConversationResponse
	conversations, err := c.ConversationRepository.GetAllForUser(ctx, loggedId, limitConv)

	for _, conversation := range conversations {
		if loggedId == conversation.ParticipantOne.Id {
			convResp = append(convResp, &model.ConversationResponse{
				Id:                conversation.Id,
				Participant:       conversation.ParticipantTwo,
				LastMessage:       conversation.LastMessage,
				LastMessageUserId: conversation.LastMessage.MessageFromId,
			})
		} else {
			convResp = append(convResp, &model.ConversationResponse{
				Id:                conversation.Id,
				Participant:       conversation.ParticipantOne,
				LastMessage:       conversation.LastMessage,
				LastMessageUserId: conversation.LastMessage.MessageFromId,
			})
		}
	}

	return convResp, nil
}

func (c conversationService) GetAllMessagesFromUser(ctx context.Context, bearer string, userId string) (*model.MessagesResponse, error) {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	messages, err := c.ConversationRepository.GetAllMessagesFromUser(ctx, loggedId, userId, limitConv)
	if err != nil {
		return nil, err
	}
	participant, err := c.UserClient.GetUsersInfo(userId)
	if err != nil {
		return nil, err
	}


	return &model.MessagesResponse{
		UserInfo: *participant,
		Messages: messages,
	}, nil
}

func (c conversationService) ViewUsersMessages(ctx context.Context, bearer string, conversationId string) error {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}


	return c.ConversationRepository.ViewUsersMessages(ctx, loggedId, conversationId)
}
