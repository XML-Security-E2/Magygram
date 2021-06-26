package service

import (
	"context"
	"errors"
	"math"
	"message-service/domain/model"
	"message-service/domain/repository"
	"message-service/domain/service-contracts"
	"message-service/domain/service-contracts/exceptions/denied"
	"message-service/service/intercomm"
	"mime/multipart"
)

var (
	limitConv int64 = math.MaxInt64
)

type conversationService struct {
	repository.ConversationRepository
	intercomm.AuthClient
	intercomm.UserClient
	intercomm.MediaClient
	intercomm.RelationshipClient
}

func NewConversationService(r repository.ConversationRepository, ac intercomm.AuthClient, uc intercomm.UserClient, mc intercomm.MediaClient, rc intercomm.RelationshipClient) service_contracts.ConversationService {
	return &conversationService{r, ac, uc, mc, rc}
}

func (c conversationService) AcceptConversationRequest(ctx context.Context, bearer string, requestId string) error {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	conversationReq, err := c.ConversationRepository.GetConversationRequestById(ctx, requestId)
	if err == nil && conversationReq == nil{
		return errors.New("request not found")
	}

	if conversationReq.RequestTo.Id != loggedId {
		return errors.New("unauthorized")
	}

	conv := &model.Conversation{
		Id:                conversationReq.Id,
		ParticipantOne:    conversationReq.RequestFrom,
		ParticipantTwo:    conversationReq.RequestTo,
		Messages:          conversationReq.Messages,
		LastMessage:       conversationReq.LastMessage,
		LastMessageUserId: conversationReq.LastMessageUserId,
	}
	err = c.ConversationRepository.CreateConversation(ctx, conv)
	if err != nil {
		return err
	}

	err = c.ConversationRepository.DeleteConversationRequest(ctx, conversationReq.RequestFrom.Id, conversationReq.RequestTo.Id, requestId)
	if err != nil {
		return err
	}

	return nil
}

func (c conversationService) DenyConversationRequest(ctx context.Context, bearer string, requestId string) error {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	conversationReq, err := c.ConversationRepository.GetConversationRequestById(ctx, requestId)
	if err == nil && conversationReq == nil{
		return errors.New("request not found")
	}

	if conversationReq.RequestTo.Id != loggedId {
		return errors.New("unauthorized")
	}

	conversationReq.RequestStatus = "DENIED"
	err = c.ConversationRepository.UpdateConversationRequest(ctx, conversationReq)
	if err != nil {
		return err
	}

	return nil
}

func (c conversationService) DeleteConversationRequest(ctx context.Context, bearer string, requestId string) error {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	conversationReq, err := c.ConversationRepository.GetConversationRequestById(ctx, requestId)
	if err == nil && conversationReq == nil{
		return errors.New("request not found")
	}

	if conversationReq.RequestTo.Id != loggedId {
		return errors.New("unauthorized")
	}

	err = c.ConversationRepository.DeleteConversationRequest(ctx, conversationReq.RequestFrom.Id, conversationReq.RequestTo.Id, requestId)
	if err != nil {
		return err
	}

	return nil
}


//PROVERITI ZA PRIVATNOST ITD
func (c conversationService) SendMessage(ctx context.Context, bearer string, messageRequest *model.MessageSentRequest) (*model.MessageSendResponse, error) {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	var message *model.Message
	if messageRequest.MessageType == "MEDIA" {
		media, err := c.MediaClient.SaveMedia([]*multipart.FileHeader{messageRequest.Media})
		if err != nil {
			return nil, err
		}

		if len(media) == 0 {
			return nil, errors.New("error while saving file")
		}

		message, err = model.NewMessage(messageRequest, loggedId, &media[0])
	} else {
		message, err = model.NewMessage(messageRequest, loggedId, nil)
	}
	if err != nil {
		return nil, err
	}

	conversation, err := c.ConversationRepository.GetConversationForUser(ctx, loggedId, messageRequest.MessageTo, limitConv)
	if err == nil && conversation == nil {
		isPrivate, err := c.UserClient.IsUserPrivate(messageRequest.MessageTo)
		if err != nil {
			return nil, err
		}

		if isPrivate {
			followers, err := c.RelationshipClient.GetFollowedUsers(loggedId)
			if err != nil {
				return nil, err
			}

			if !isInFollowers(followers, messageRequest.MessageTo){
				request, err := c.sendMessageRequest(ctx, loggedId, messageRequest.MessageTo, message)
				if err != nil {
					return nil, err
				}
				return &model.MessageSendResponse{
					ConversationRequest: request,
					Conversation:   nil,
					IsMessageRequest: true,
				}, nil
			}
		}

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

		return &model.MessageSendResponse{
			ConversationRequest: nil,
			Conversation:   &model.ConversationResponse{
				Id:                conv.Id,
				Participant:       conv.ParticipantTwo,
				LastMessage:       *message,
				LastMessageUserId: loggedId,
			},
			IsMessageRequest: false,
		}, nil
	}

	conversation.Messages = append(conversation.Messages, *message)
	conversation.LastMessage = *message
	conversation.LastMessageUserId = loggedId

	err = c.ConversationRepository.Update(ctx, conversation)
	if err != nil {
		return nil, err
	}

	return &model.MessageSendResponse{
		ConversationRequest: nil,
		Conversation:  &model.ConversationResponse{
			Id:                conversation.Id,
			Participant:       conversation.ParticipantTwo,
			LastMessage:       *message,
			LastMessageUserId: loggedId,
		},
		IsMessageRequest: false,

	} , nil
}

func (c conversationService) sendMessageRequest(ctx context.Context, loggedId string, userId string, message *model.Message) (*model.ConversationRequest,error) {
	messageFrom, err := c.UserClient.GetUsersInfo(loggedId)
	if err != nil {
		return nil, err
	}
	messageTo, err := c.UserClient.GetUsersInfo(userId)
	if err != nil {
		return  nil, err
	}

	conversationReq, err := c.ConversationRepository.GetConversationRequestFromUser(ctx, loggedId, userId, limitConv)
	if err == nil && conversationReq == nil {
		conversationRequest := model.NewConversationRequest(message, *messageFrom, *messageTo)
		err = c.ConversationRepository.CreateConversationRequest(ctx, conversationRequest)
		return conversationRequest, nil
	}

	if conversationReq.RequestStatus == "DENIED" {
		return nil, &denied.MessageRequestDeniedError{Msg: "message request denied"}
	}

	conversationReq.Messages = append(conversationReq.Messages, *message)
	conversationReq.LastMessage = *message
	conversationReq.LastMessageUserId = loggedId

	err = c.ConversationRepository.UpdateConversationRequest(ctx, conversationReq)
	if err != nil {
		return nil, err
	}

	return conversationReq, nil
}

func isInFollowers(followingUsers model.FollowedUsersResponse, userId string) bool {
	for _, folId := range followingUsers.Users {
		if folId == userId {
			return true
		}
	}
	return false
}

func (c conversationService) GetAllMessageRequestsForUser(ctx context.Context, bearer string) ([]*model.ConversationResponse, error) {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}
	var convResp []*model.ConversationResponse
	conversations, err := c.ConversationRepository.GetAllConversationRequestsForUser(ctx, loggedId, limitConv)

	for _, conversation := range conversations {
		if conversation.RequestStatus == "PENDING" {
			convResp = append(convResp, &model.ConversationResponse{
				Id:                conversation.Id,
				Participant:       conversation.RequestFrom,
				LastMessage:       conversation.LastMessage,
				LastMessageUserId: conversation.LastMessage.MessageFromId,
			})
		}
	}

	return convResp, nil
}

func (c conversationService) GetAllMessagesFromUserFromRequest(ctx context.Context, bearer string, userId string) (*model.MessagesResponse, error) {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return nil, err
	}

	convReq, err := c.ConversationRepository.GetConversationRequestFromUser(ctx, userId, loggedId, limitConv)
	if err != nil {
		return nil, err
	}
	participant, err := c.UserClient.GetUsersInfo(userId)
	if err != nil {
		return nil, err
	}


	return &model.MessagesResponse{
		UserInfo: *participant,
		Messages: convReq.Messages,
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

func (c conversationService) ViewUserMediaMessages(ctx context.Context, bearer string, conversationId string, messageId string) error {
	loggedId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	return c.ConversationRepository.ViewUserMediaMessage(ctx, loggedId, conversationId, messageId)
}
