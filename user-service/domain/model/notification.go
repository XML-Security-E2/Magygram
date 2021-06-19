package model

type PostStoryNotifications struct {
	Id string `bson:"_id,omitempty"`
	UserId  string `bson:"userId" `
	NotificationsFromId  string `bson:"notificationsFromId" `
	PostNotifications bool `bson:"postNotifications"`
	StoryNotifications bool `bson:"storyNotifications"`
}

type SettingsRequest struct {
	PostNotifications bool `json:"notifyPost"`
	StoryNotifications bool `json:"notifyStory"`
}
