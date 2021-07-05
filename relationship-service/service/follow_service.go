package service

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"log"
	"relationship-service/conf"
	"relationship-service/domain/model"
	"relationship-service/infrastructure/persistence/neo4jdb"
	"relationship-service/logger"
	"relationship-service/saga"
	"relationship-service/service/intercomm"
	"relationship-service/tracer"
)

type FollowService interface {
	FollowRequest(ctx context.Context, followRequest *model.FollowRequest) (bool, error)
	Unfollow(ctx context.Context, followRequest *model.FollowRequest) error
	IsUserFollowed(ctx context.Context, followRequest *model.FollowRequest) (interface{}, error)
	IsMuted(ctx context.Context, mute *model.Mute) (interface{}, error)
	AcceptFollowRequest(ctx context.Context, bearer string, userId string) error
	CreateUser(ctx context.Context, user *model.User) error
	ReturnFollowedUsers(ctx context.Context, user *model.User) (interface{}, error)
	ReturnUnmutedFollowedUsers(ctx context.Context, user *model.User) (interface{}, error)
	ReturnFollowingUsers(ctx context.Context, user *model.User) (interface{}, error)
	ReturnFollowRequests(ctx context.Context, bearer string) (interface{}, error)
	ReturnFollowRequestsForUser(ctx context.Context, bearer string, objectId string) (interface{}, error)
	Mute(ctx context.Context, mute *model.Mute) error
	Unmute(ctx context.Context, mute *model.Mute) error
	ReturnRecommendedUsers(ctx context.Context, user *model.User) (interface{}, error)
	RedisConnection()
}

type followService struct {
	neo4jdb.FollowRepository
	intercomm.UserClient
	intercomm.AuthClient
}

func NewFollowService(r neo4jdb.FollowRepository, userClient intercomm.UserClient, ac intercomm.AuthClient) FollowService {
	return &followService{r, userClient, ac}
}

func (f *followService) Mute(ctx context.Context, mute *model.Mute) error {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceMute")
	defer span.Finish()

	if err := validator.New().Struct(mute); err != nil {
		return err
	}
	if err := f.FollowRepository.Mute(mute); err != nil {
		return err
	}
	return nil
}

func (f *followService) Unmute(ctx context.Context, mute *model.Mute) error {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceUnmute")
	defer span.Finish()

	if err := validator.New().Struct(mute); err != nil {
		return err
	}
	if err := f.FollowRepository.Unmute(mute); err != nil {
		return err
	}
	return nil
}

func (f *followService) FollowRequest(ctx context.Context, followRequest *model.FollowRequest) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceFollowRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	if err := validator.New().Struct(followRequest); err != nil {
		return false, err
	}
	isPrivate, err := f.UserClient.IsPrivate(ctx, followRequest.ObjectId)
	if err != nil {
		return false, nil
	}
	if isPrivate {
		if err:= f.FollowRepository.CreateFollowRequest(followRequest); err != nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
														 "object_id" : followRequest.ObjectId}).Error("Follow request create, database failure")
			return false, err
		}
		logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
													 "object_id" : followRequest.ObjectId}).Info("Follow request created")
		return true, nil
	} else {
		if err:= f.FollowRepository.CreateFollow(followRequest); err != nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
														 "object_id" : followRequest.ObjectId}).Error("Follow user, database failure")
			return false, err
		}

		logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
												     "object_id" : followRequest.ObjectId}).Info("User followed")
	}
	return false, nil
}

func (f *followService) Unfollow(ctx context.Context, followRequest *model.FollowRequest) error {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceUnfollow")
	defer span.Finish()

	err := f.FollowRepository.Unfollow(followRequest)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
													 "object_id" : followRequest.ObjectId}).Error("Unfollow user, database failure")
		tracer.LogError(span, err)
		return err
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
												 "object_id" : followRequest.ObjectId}).Info("User unfollowed")

	return err
}

func (f *followService) IsUserFollowed(ctx context.Context, followRequest *model.FollowRequest) (interface{}, error) {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceIsUserFollowed")
	defer span.Finish()

	if err := validator.New().Struct(followRequest); err != nil {
		return false, err
	}

	exists, err := f.FollowRepository.IsUserFollowed(followRequest)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": followRequest.SubjectId,
													 "object_id" : followRequest.ObjectId}).Error("Check if follows, database failure")
		return false, err
	}

	return exists, nil
}

func (f *followService) IsMuted(ctx context.Context, mute *model.Mute) (interface{}, error) {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceIsUserIsMuted")
	defer span.Finish()

	if err := validator.New().Struct(mute); err != nil {
		return false, err
	}

	exists, err := f.FollowRepository.IsMuted(mute)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (f *followService) CreateUser(ctx context.Context, user *model.User) error {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceCreateUser")
	defer span.Finish()

	if err := validator.New().Struct(user); err != nil {
		return err
	}
	if err := f.FollowRepository.CreateUser(user); err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": user.Id}).Error("Create user node, database failure")
		return err
	}

	return nil
}

func (f *followService) AcceptFollowRequest(ctx context.Context, bearer string, userId string) error {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceAcceptFollowRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	loggedId, err := f.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return  err
	}

	err = f.FollowRepository.AcceptFollowRequest(&model.FollowRequest{
		SubjectId: userId,
		ObjectId:  loggedId,
	})
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": userId,
													 "object_id" : loggedId}).Error("Accept follow request, database failure")
		return  err
	}
	logger.LoggingEntry.WithFields(logrus.Fields{"subject_id": userId,
													"object_id" : loggedId}).Info("Follow request accepted")
	return err
}

func (f *followService) ReturnFollowedUsers(ctx context.Context, user *model.User) (interface{}, error) {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceReturnFollowedUsers")
	defer span.Finish()

	retVal, err := f.FollowRepository.ReturnFollowedUsers(user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": user.Id}).Error("Get followed users, database fetch failure")
		return retVal, err
	}
	return retVal, err
}

func (f *followService) ReturnUnmutedFollowedUsers(ctx context.Context, user *model.User) (interface{}, error) {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceReturnUnmutedFollowedUsers")
	defer span.Finish()

	retVal, err := f.FollowRepository.ReturnUnmutedFollowedUsers(user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": user.Id}).Error("Get followed users, database fetch failure")
		return retVal, err
	}
	return retVal, err
}

func (f *followService) ReturnFollowingUsers(ctx context.Context, user *model.User) (interface{}, error) {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceReturnFollowingUsers")
	defer span.Finish()

	retVal, err := f.FollowRepository.ReturnFollowingUsers(user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": user.Id}).Error("Get following users, database fetch failure")
		return retVal, err
	}
	return retVal, err
}

func (f *followService) ReturnFollowRequests(ctx context.Context, bearer string) (interface{}, error) {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceReturnFollowRequests")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	loggedId, err := f.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return false, err
	}

	retVal, err := f.FollowRepository.ReturnFollowRequests(&model.User{Id: loggedId})
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"logged_user_id": loggedId}).Error("Get follow requests for user, database fetch failure")
		return retVal, err
	}
	return retVal, err
}

func (f *followService) ReturnFollowRequestsForUser(ctx context.Context, bearer string, objectId string) (interface{}, error) {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceReturnFollowRequestsForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	loggedId, err := f.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return false, err
	}

	retVal, err := f.FollowRepository.ReturnFollowRequestsForUser(&model.User{Id: objectId}, loggedId)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"logged_user_id": loggedId, "object_id" : objectId}).Error("Get follow requests for user, database fetch failure")
		return retVal, err
	}
	return retVal, err
}

func (f *followService) ReturnRecommendedUsers(ctx context.Context, user *model.User) (interface{}, error) {
	span := tracer.StartSpanFromContext(ctx, "FollowServiceReturnRecommendedUsers")
	defer span.Finish()

	retVal, err := f.FollowRepository.ReturnRecommendedUsers(user)
	if err != nil {
		return retVal, err
	}

	len := len(retVal.Users)
	if len<20{
		retPopularUsers,err := f.FollowRepository.GetPopularUsers(user, 20-len, retVal.Users)
		if err != nil {
			return retVal, err
		}

		retVal.Users = append(retVal.Users, retPopularUsers.Users...)
	}

	return retVal, err
}

func (f *followService) RedisConnection() {
	// create client and ping redis
	var err error
	client := redis.NewClient(&redis.Options{Addr: conf.Current.RedisDatabase.Host+":"+ conf.Current.RedisDatabase.Port, Password: "", DB: 0})
	if _, err = client.Ping().Result(); err != nil {
		log.Fatalf("error creating redis client %s", err)
	}

	// subscribe to the required channels
	pubsub := client.Subscribe(saga.RelationshipChannel, saga.ReplyChannel)
	if _, err = pubsub.Receive(); err != nil {
		log.Fatalf("error subscribing %s", err)
	}
	defer func() { _ = pubsub.Close() }()
	ch := pubsub.Channel()

	log.Println("starting the stock service")
	for {
		select {
		case msg := <-ch:
			m := saga.RegisterUserMessage{}
			err := json.Unmarshal([]byte(msg.Payload), &m)
			if err != nil {
				log.Println(err)
				continue
			}

			switch msg.Channel {
			case saga.RelationshipChannel:

				// Happy Flow
				if m.Action == saga.ActionStart {
					// Check quantity of product

					userDTO := model.User{Id: m.User.Id}
					err := f.CreateUser(context.Background(), &userDTO)

					if err != nil{
						sendToReplyChannel(client, &m, saga.ActionError, saga.ServiceAuth, saga.ServiceRelationship)
					}else {
						sendToReplyChannel(client, &m, saga.ActionDone, saga.ServiceUser, saga.ServiceRelationship)
						//skinuti komentar za testiranje
						//sendToReplyChannel(client, &m, saga.ActionError, saga.ServiceAuth, saga.ServiceRelationship)
					}
				}

				// Rollback flow
				if m.Action == saga.ActionRollback {
					userDTO := model.User{Id: m.User.Id}
					f.DeleteUser(&userDTO)
					sendToReplyChannel(client, &m, saga.ActionError, saga.ServiceAuth, saga.ServiceRelationship)
				}
			}
		}
	}
}

func sendToReplyChannel(client *redis.Client, m *saga.RegisterUserMessage, action string, service string, senderService string) {
	var err error
	m.Action = action
	m.Service = service
	m.SenderService = senderService
	if err = client.Publish(saga.ReplyChannel, m).Err(); err != nil {
		log.Printf("error publishing done-message to %s channel", saga.ReplyChannel)
	}
	log.Printf("done message published to channel :%s", saga.ReplyChannel)
}
