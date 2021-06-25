package redisdb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math"
	"message-service/domain/model"
	"message-service/domain/repository"
	"sort"
)

type conversationRepository struct {
	Db *redis.Client
}

func NewConversationRepository(Db *redis.Client) repository.ConversationRepository {
	return &conversationRepository{Db}
}

func (c conversationRepository) CreateConversation(ctx context.Context, conversation *model.Conversation) error {
	jsonString, err := json.Marshal(conversation)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = c.Db.Set(ctx, fmt.Sprintf("%s/%s/%s/%s", model.ConvPrefix, conversation.ParticipantOne.Id, conversation.ParticipantTwo.Id, conversation.Id), jsonString, 0).Err()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (c conversationRepository) GetConversationForUser(ctx context.Context, loggedId string, userId string, limit int64) (*model.Conversation, error) {
	keys, _, err := c.Db.Scan(ctx, 0, fmt.Sprintf("%s/%s/%s/*", model.ConvPrefix, loggedId, userId), limit).Result()

	if err != nil || keys == nil {
		return nil, nil
	}

	if len(keys) > 0 {
		val, err := c.Db.Get(ctx, keys[0]).Bytes()
		if err != nil {
			return nil, nil
		}
		var temp *model.Conversation
		json.Unmarshal(val, &temp)

		return temp, nil
	} else {
		keys, _, err = c.Db.Scan(ctx, 0, fmt.Sprintf("%s/%s/%s/*", model.ConvPrefix, userId, loggedId), limit).Result()
		if err != nil || keys == nil {
			return nil, nil
		}
		if len(keys) > 0 {
			val, err := c.Db.Get(ctx, keys[0]).Bytes()
			if err != nil {
				return nil, nil
			}
			var temp *model.Conversation
			json.Unmarshal(val, &temp)

			return temp, nil
		}
	}

	return nil, nil
}

func (c conversationRepository) GetAllForUser(ctx context.Context, userId string, limit int64) ([]*model.Conversation, error) {
	keys, _, err := c.Db.Scan(ctx, 0, fmt.Sprintf("%s/%s/*", model.ConvPrefix, userId), limit).Result()
	keysOth, _, err := c.Db.Scan(ctx, 0, fmt.Sprintf("%s/*/%s/*", model.ConvPrefix, userId), limit).Result()

	var conversations []*model.Conversation
	for _, key := range keys {
		val, err := c.Db.Get(ctx, key).Bytes()
		if err != nil {
			return nil, err
		}

		var temp *model.Conversation
		json.Unmarshal(val, &temp)
		conversations = append(conversations, temp)
	}

	for _, key := range keysOth {
		val, err := c.Db.Get(ctx, key).Bytes()
		if err != nil {
			return nil, err
		}

		var temp *model.Conversation
		json.Unmarshal(val, &temp)
		conversations = append(conversations, temp)
	}

	sort.Slice(conversations, func(i, j int) bool {
		return conversations[i].LastMessage.Timestamp.After(conversations[j].LastMessage.Timestamp)
	})

	return conversations, err
}

func (c conversationRepository) ViewUsersMessages(ctx context.Context, userId string, conversationId string) error {
	keys, _, err := c.Db.Scan(ctx, 0, fmt.Sprintf("%s/*/*/%s", model.ConvPrefix, conversationId), math.MaxInt64).Result()
	if err != nil || keys == nil {
		return err
	}
	if len(keys) > 0 {
		val, err := c.Db.Get(ctx, keys[0]).Bytes()
		if err != nil {
			return err
		}
		var temp *model.Conversation
		json.Unmarshal(val, &temp)

		if temp.ParticipantOne.Id != userId && temp.ParticipantTwo.Id != userId {
			return errors.New("unauthorized access")
		}

		if userId != temp.LastMessageUserId {
			temp.LastMessage.Viewed = true
		}
		var mess []model.Message
		for _, message := range temp.Messages {
			if userId != message.MessageFromId {
				message.Viewed = true
			}
			mess = append(mess, message)
		}
		temp.Messages = mess

		err = c.Update(ctx, temp)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("invalid conversation id")
}

func (c conversationRepository) ViewUserMediaMessage(ctx context.Context, userId string, conversationId string, messageId string) error {
	keys, _, err := c.Db.Scan(ctx, 0, fmt.Sprintf("%s/*/*/%s", model.ConvPrefix, conversationId), math.MaxInt64).Result()
	if err != nil || keys == nil {
		return err
	}
	if len(keys) > 0 {
		val, err := c.Db.Get(ctx, keys[0]).Bytes()
		if err != nil {
			return err
		}
		var temp *model.Conversation
		json.Unmarshal(val, &temp)

		if temp.ParticipantOne.Id != userId && temp.ParticipantTwo.Id != userId {
			return errors.New("unauthorized access")
		}

		var mess []model.Message
		for _, message := range temp.Messages {
			if message.Id == messageId {
				if userId != message.MessageFromId {
					message.ViewedMedia = true
				}
			}
			mess = append(mess, message)
		}
		temp.Messages = mess

		err = c.Update(ctx, temp)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("invalid conversation id")
}

func (c conversationRepository) GetAllMessagesFromUser(ctx context.Context, loggedId string, userId string, limit int64) ([]model.Message, error) {
	keys, _, err := c.Db.Scan(ctx, 0, fmt.Sprintf("%s/%s/%s/*", model.ConvPrefix, loggedId, userId), limit).Result()

	if err != nil || keys == nil {
		return []model.Message{}, nil
	}

	if len(keys) > 0 {
		val, err := c.Db.Get(ctx, keys[0]).Bytes()
		if err != nil {
			return nil, nil
		}
		var temp *model.Conversation
		json.Unmarshal(val, &temp)
		messages := temp.Messages
		sort.Slice(messages, func(i, j int) bool {
			return messages[i].Timestamp.Before(messages[j].Timestamp)
		})

		return messages, nil
	} else {
		keys, _, err = c.Db.Scan(ctx, 0, fmt.Sprintf("%s/%s/%s/*", model.ConvPrefix, userId , loggedId), limit).Result()
		if err != nil || keys == nil {
			return []model.Message{}, nil
		}
		if len(keys) > 0 {
			val, err := c.Db.Get(ctx, keys[0]).Bytes()
			if err != nil {
				return nil, nil
			}
			var temp *model.Conversation
			json.Unmarshal(val, &temp)
			messages := temp.Messages
			sort.Slice(messages, func(i, j int) bool {
				return messages[i].Timestamp.Before(messages[j].Timestamp)
			})

			return messages, nil
		}
	}

	return []model.Message{}, nil
}

func (c conversationRepository) Update(ctx context.Context, conversation *model.Conversation) error {
	jsonString, err := json.Marshal(conversation)
	if err != nil {
		return err
	}

	err = c.Db.Del(ctx, fmt.Sprintf("%s/%s/%s/%s", model.ConvPrefix, conversation.ParticipantOne.Id, conversation.ParticipantTwo.Id, conversation.Id)).Err()
	if err != nil {
		return err
	}

	err = c.Db.Set(ctx, fmt.Sprintf("%s/%s/%s/%s", model.ConvPrefix, conversation.ParticipantOne.Id, conversation.ParticipantTwo.Id, conversation.Id), jsonString, 0).Err()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (c conversationRepository) CreateMessageRequest(ctx context.Context, request *model.MessageRequest) error {
	jsonString, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = c.Db.Set(ctx, fmt.Sprintf("%s/%s/%s/%s", model.MessageRequestPrefix, request.MessageFrom.Id, request.MessageTo, request.Id), jsonString, 0).Err()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}