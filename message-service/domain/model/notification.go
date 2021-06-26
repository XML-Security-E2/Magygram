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
	UserFromId string `json:"userFromId"`
	NotifyUrl string `json:"notifyUrl"`
	ImageUrl string `json:"imageUrl"`
	Type  string `json:"type"`
}

type UserInfo struct {
	Id string
	Username string
	ImageURL string
}

var (
	Prefix = "notifications"
	Liked = "Liked"
	Disliked = "Disliked"
	Commented = "Commented"
	Followed = "Followed"
	FollowRequest = "FollowRequest"
	AcceptedFollowRequest = "AcceptedFollowRequest"
	PublishedStory = "PublishedStory"
	PublishedPost = "PublishedPost"
)

func NewNotification(notificationReq *NotificationRequest) *Notification {
	return &Notification{
		Id:        guid.New().String(),
		Username:  notificationReq.Username,
		UserId:    notificationReq.UserId,
		NotifyUrl: notificationReq.NotifyUrl,
		ImageUrl:  notificationReq.ImageUrl,
		Note:      createNote(notificationReq.Type, notificationReq.Username),
		Viewed:    false,
		Timestamp: time.Now(),
	}
}

func createNote(notificationType string, username string) string {
	if notificationType == Liked {
		return fmt.Sprintf("%s liked your post", username)
	} else if notificationType == Disliked {
		return fmt.Sprintf("%s disliked your post", username)
	} else if notificationType == Commented {
		return fmt.Sprintf("%s commented on your post", username)
	} else if notificationType == Followed {
		return fmt.Sprintf("%s started following you", username)
	} else if notificationType == PublishedStory {
		return fmt.Sprintf("%s published story", username)
	} else if notificationType == PublishedPost {
		return fmt.Sprintf("%s published post", username)
	} else if notificationType == FollowRequest {
		return fmt.Sprintf("%s wants to follow you", username)
	} else if notificationType == AcceptedFollowRequest {
		return fmt.Sprintf("%s accepted follow request", username)
	}
	return ""
}
