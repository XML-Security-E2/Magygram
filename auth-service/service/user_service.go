package service

import (
	"auth-service/conf"
	"auth-service/domain/model"
	"auth-service/domain/repository"
	"auth-service/domain/service-contracts"
	"auth-service/logger"
	"auth-service/saga"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis"
	"github.com/pquerna/otp/totp"
	"github.com/sirupsen/logrus"
	"image/jpeg"
	"log"
)

type userService struct {
	repository.UserRepository
	repository.LoginEventRepository
}

func NewUserService(r repository.UserRepository,a repository.LoginEventRepository) service_contracts.UserService {
	return &userService{r,a}
}

func (u userService) RegisterUser(ctx context.Context, userRequest *model.UserRequest) (string,[]byte, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Magygram",
		AccountName: userRequest.Email,
	})

	img, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}

	buffer := new(bytes.Buffer)
	err = jpeg.Encode(buffer, img, nil)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email" : userRequest.Email}).Error("TOTP QR code not created")
		return "",[]byte{}, err
	}

	imageInBytes := buffer.Bytes()

//	logger.LoggingEntry.WithFields(logrus.Fields{"email" : userRequest.Email}).Info("TOTP QR code created")

	user, err := model.NewUser(userRequest,key.Secret())
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email" : userRequest.Email}).Warn("User registration validation failure")
		return "",[]byte{}, err
	}

	if err := validator.New().Struct(user); err!= nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email" : userRequest.Email}).Warn("User registration validation failure")
		return "",[]byte{}, err
	}

	result, err := u.UserRepository.Create(ctx, user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email" : userRequest.Email}).Error("User database create failure")
		return "",[]byte{}, err
	}

	if userId, ok := result.InsertedID.(string); ok {
		//logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : userId}).Info("User registered")
		return userId,imageInBytes, nil
	}
	return "",imageInBytes, err
}

func (u userService) ActivateUser(ctx context.Context, userId string) (bool, error) {

	user, err := u.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return false, err
	}
	user.Active = true
	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userId}).Error("User database update failure")
		return false, err
	}

	u.HandleLoginEventAndAccountActivation(ctx, user.Email, true, model.ActivatedAccount)
	logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : userId}).Info("User activated")

	return true, err
}

func (u userService) HandleLoginEventAndAccountActivation(ctx context.Context, userEmail string, successful bool, eventType string) {
	if successful {
		_, err := u.LoginEventRepository.Create(ctx, model.NewLoginEvent(userEmail, eventType, 0))
		if err != nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"user_email" : userEmail}).Warn("Create success login event, database failure")
		}
		return
	}
}

func (u userService) DeactivateUser(ctx context.Context, userEmail string) (bool, error) {
	user, err := u.UserRepository.GetByEmail(ctx, userEmail)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email": userEmail}).Warn("Invalid email address")
		return false, err
	}
	user.Active = false
	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": user.Id}).Error("User database update failure")
		return false, err
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"user_id": user.Id}).Info("User deactivated")

	return true, err
}

func (u userService) ResetPassword(ctx context.Context, changePasswordRequest *model.PasswordChangeRequest) (bool, error) {
	hashAndSalt, err := model.HashAndSaltPasswordIfStrongAndMatching(changePasswordRequest.Password, changePasswordRequest.PasswordRepeat)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : changePasswordRequest.UserId}).Warn("Passwords not valid")
		return false, err
	}

	user, err := u.UserRepository.GetByID(ctx, changePasswordRequest.UserId)
	if err != nil {
		return false, err
	}
	user.Password = hashAndSalt
	user.Active = true
	_, err = u.UserRepository.Update(ctx, user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": user.Id}).Error("User database update failure")
		return false, err
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : user.Id}).Info("Users password changed")
	return true, err
}

func (u userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := u.UserRepository.GetByEmail(ctx, email)

	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email": email}).Warn("Invalid email address")
		return nil, errors.New("invalid user id")
	}

	return user, err
}

func (u userService) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	user, err := u.UserRepository.GetByID(ctx, userId)

	if err != nil {
		return nil, errors.New("invalid user id")
	}

	return user, err
}

func (u userService) GetAllRolesByUserId(ctx context.Context, userId string) ([]model.Role, error) {
	return u.UserRepository.GetAllRolesByUserId(ctx, userId)
}

func (u userService) RegisterAgent(ctx context.Context, userRequest *model.UserRequest) (string,[]byte, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Magygram",
		AccountName: userRequest.Email,
	})

	img, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}

	buffer := new(bytes.Buffer)
	err = jpeg.Encode(buffer, img, nil)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email" : userRequest.Email}).Error("TOTP QR code not created")
		return "",[]byte{}, err
	}

	imageInBytes := buffer.Bytes()

	logger.LoggingEntry.WithFields(logrus.Fields{"email" : userRequest.Email}).Info("TOTP QR code created")

	user, err := model.NewAgent(userRequest,key.Secret())
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email" : userRequest.Email}).Warn("User registration validation failure")
		return "",[]byte{}, err
	}

	if err := validator.New().Struct(user); err!= nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email" : userRequest.Email}).Warn("User registration validation failure")
		return "",[]byte{}, err
	}

	result, err := u.UserRepository.Create(ctx, user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"email" : userRequest.Email}).Error("User database create failure")
		return "",[]byte{}, err
	}

	if userId, ok := result.InsertedID.(string); ok {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : userId}).Info("User registered")
		return userId,imageInBytes, nil
	}
	return "",imageInBytes, err
}

func (u userService) RedisConnection() {
	// create client and ping redis
	var err error
	client := redis.NewClient(&redis.Options{Addr: conf.Current.RedisDatabase.Host+":"+ conf.Current.RedisDatabase.Port, Password: "", DB: 0})
	if _, err = client.Ping().Result(); err != nil {
		log.Fatalf("error creating redis client %s", err)
	}

	// subscribe to the required channels
	pubsub := client.Subscribe(saga.AuthChannel, saga.ReplyChannel)
	if _, err = pubsub.Receive(); err != nil {
		log.Fatalf("error subscribing %s", err)
	}
	defer func() { _ = pubsub.Close() }()
	ch := pubsub.Channel()

	log.Println("starting the auth service")
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
			case saga.AuthChannel:

				// Happy Flow
				if m.Action == saga.ActionStart {
					// Check quantity of product

					userRequest := model.UserRequest{Id: m.User.Id, Email: m.User.Email, Password: m.User.Password, RepeatedPassword: m.User.RepeatedPassword}
					_,bufer, err := u.RegisterUser(context.TODO(),&userRequest)

					if err != nil{
						m.ImageByte = nil
						sendToReplyChannel(client, &m, saga.ActionError, saga.ServiceUser, saga.ServiceAuth)
					}else {
						m.ImageByte = bufer
						sendToReplyChannel(client, &m, saga.ActionDone, saga.ServiceRelationship, saga.ServiceAuth)
						//skinuti komentar za testiranje
						//sendToReplyChannel(client, &m, saga.ActionError, saga.ServiceUser, saga.ServiceAuth)
					}
				}

				// Rollback flow
				if m.Action == saga.ActionRollback {
					u.UserRepository.PhysicalDelete(context.TODO(), m.User.Id)
					sendToReplyChannel(client, &m, saga.ActionError, saga.ServiceUser, saga.ServiceAuth)
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
