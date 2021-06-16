package model

import (
	"fmt"
	"github.com/beevik/guid"
	"time"
)

type Notification struct {
	Id string `json:"id"`
	Username string `json:"username"`
	UserId string `json:"userId"`
	NotifyUrl string `json:"notifyUrl"`
	ImageUrl string `json:"imageUrl"`
	Note  string `json:"note"`
	Viewed  bool `json:"viewed"`
	Timestamp time.Time `json:"time"`
}

type NotificationRequest struct {
	Username string `json:"username"`
	UserId string `json:"userId"`
	NotifyUrl string `json:"notifyUrl"`
	ImageUrl string `json:"imageUrl"`
	Note  string `json:"note"`
}

var (
	Prefix = "notifications"
)

func NewNotification(notificationReq *NotificationRequest) *Notification {
	return &Notification{
		Id:        fmt.Sprintf("%s/%s/%s/false", Prefix, notificationReq.UserId, guid.New().String()),
		Username:  notificationReq.Username,
		UserId:    notificationReq.UserId,
		NotifyUrl: notificationReq.NotifyUrl,
		ImageUrl:  notificationReq.ImageUrl,
		Note:      notificationReq.Note,
		Viewed:    false,
		Timestamp: time.Now(),
	}
}
