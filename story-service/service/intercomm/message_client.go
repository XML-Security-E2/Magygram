package intercomm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"story-service/conf"
	"story-service/tracer"
)

type MessageClient interface {
	CreateNotifications(ctx context.Context, request *NotificationRequest) error
}

type NotificationRequest struct {
	Username string `json:"username"`
	UserId string `json:"userId"`
	UserFromId string `json:"userFromId"`
	NotifyUrl string `json:"notifyUrl"`
	ImageUrl string `json:"imageUrl"`
	Type  string `json:"type"`
}

type messageClient struct {}

func NewMessageClient() MessageClient {
	baseMessageUrl = fmt.Sprintf("%s%s:%s/api/notifications", conf.Current.Messageservice.Protocol, conf.Current.Messageservice.Domain, conf.Current.Messageservice.Port)
	return &messageClient{}
}

var (
	baseMessageUrl = ""
	PublishedStory = "PublishedStory"
)


func (m messageClient) CreateNotifications(ctx context.Context, request *NotificationRequest) error {
	span := tracer.StartSpanFromContext(ctx, "MessageClientCreateNotifications")
	defer span.Finish()

	jsonStr, err:= json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx,"POST", fmt.Sprintf("%s/multiple", baseMessageUrl), bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			return err
		}
		fmt.Println(resp.StatusCode)

		return err
	}

	fmt.Println(resp.StatusCode)
	return nil
}
