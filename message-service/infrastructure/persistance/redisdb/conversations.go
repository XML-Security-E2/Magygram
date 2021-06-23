package redisdb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
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
	err = c.Db.Set(ctx, fmt.Sprintf("%s/%s/%s/%s", model.ConvPrefix, conversation.ParticipantOneId, conversation.ParticipantTwoId, conversation.Id), jsonString, 0).Err()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
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

func (c conversationRepository) GetAllMessagesFromUser(ctx context.Context, loggedId string, userId string, limit int64) ([]model.Message, error) {
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
		messages := temp.Messages
		sort.Slice(messages, func(i, j int) bool {
			return messages[i].Timestamp.After(messages[j].Timestamp)
		})

		return messages, nil
	} else {
		keys, _, err = c.Db.Scan(ctx, 0, fmt.Sprintf("%s/%s/%s/*", model.ConvPrefix, userId , loggedId), limit).Result()
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
			messages := temp.Messages
			sort.Slice(messages, func(i, j int) bool {
				return messages[i].Timestamp.After(messages[j].Timestamp)
			})

			return messages, nil
		}
	}

	return nil, nil
}

func (c conversationRepository) Update(ctx context.Context, conversation *model.Conversation) error {
	jsonString, err := json.Marshal(conversation)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = c.Db.Set(ctx, fmt.Sprintf("%s/%s/%s/%s", model.ConvPrefix, conversation.ParticipantOneId, conversation.ParticipantTwoId, conversation.Id), jsonString, 0).Err()
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