package model

import (
	"errors"
	"html"

	"github.com/beevik/guid"
)

type User struct {
	Id                   string                             `bson:"_id,omitempty"`
	Username             string                             `bson:"username" validate:"required,min=1"`
	Name                 string                             `bson:"name" validate:"required,min=2"`
	Email                string                             `bson:"email" validate:"required,email"`
	Surname              string                             `bson:"surname" validate:"required,min=2"`
	Website              string                             `bson:"website" `
	Bio                  string                             `bson:"bio" `
	Number               string                             `bson:"number" `
	Gender               Gender                             `bson:"gender"`
	ImageUrl             string                             `bson:"imageUrl"`
	FavouritePosts       map[string][]IdWithMedia           `bson:"favouritePosts"`
	HighlightsStory      map[string]HighlightImageWithMedia `bson:"highlightsStory"`
	IsPrivate            bool                               `bson:"private_profile"`
	LikedPosts           []string                           `bson:"liked_posts"`
	DislikedPosts        []string                           `bson:"disliked_posts"`
	BlockedUsers         []string                           `bson:"blocked_users"`
	NotificationSettings NotificationSettings               `bson:"notification_settings" json:"notificationSettings"`
	PrivacySettings      PrivacySettings                    `bson:"privacy_settings" json:"privacySettings"`
	Category             Category                           `bson:"category"`
	IsVerified           bool                               `bson:"verified_profile"`
	IsDeleted            bool                               `bson:"deleted"`
}

type NotificationSettings struct {
	NotifyLike                string `json:"notifyLike"`
	NotifyDislike             string `json:"notifyDislike"`
	NotifyFollow              bool   `json:"notifyFollow"`
	NotifyFollowRequest       bool   `json:"notifyFollowRequest"`
	NotifyAcceptFollowRequest bool   `json:"notifyAcceptFollowRequest"`
	NotifyComment             string `json:"notifyComments"`
}

type NotificationSettingsUpdateReq struct {
	NotifyLike                int  `json:"notifyLike"`
	NotifyDislike             int  `json:"notifyDislike"`
	NotifyFollow              bool `json:"notifyFollow"`
	NotifyFollowRequest       bool `json:"notifyFollowRequest"`
	NotifyAcceptFollowRequest bool `json:"notifyAcceptFollowRequest"`
	NotifyComment             int  `json:"notifyComments"`
}

type PrivacySettings struct {
	IsPrivate       bool `json:"isPrivate"`
	ReceiveMessages bool `json:"receiveMessages"`
	IsTaggable      bool `json:"isTaggable"`
}

type Category string

const (
	INFLUENCER = iota
	SPORTS
	NEWS //MEDIA
	BUSINESS
	BRAND
	ORGANIZATION
	MUSIC
	ACTOR
)

type UserProfileResponse struct {
	Username             string                        `json:"username"`
	Name                 string                        `json:"name"`
	Surname              string                        `json:"surname"`
	Website              string                        `json:"website" `
	Bio                  string                        `json:"bio"`
	Email                string                        `json:"email"`
	Number               string                        `json:"number"`
	Gender               Gender                        `json:"gender"`
	ImageUrl             string                        `json:"imageUrl"`
	Following            bool                          `json:"following"`
	Blocked              bool                          `json:"blocked"`
	Muted                bool                          `json:"muted"`
	PostNumber           int                           `json:"postNumber"`
	FollowersNumber      int                           `json:"followersNumber"`
	FollowingNumber      int                           `json:"followingNumber"`
	SentFollowRequest    bool                          `json:"sentFollowRequest"`
	PrivacySettings      PrivacySettings               `json:"privacySettings"`
	NotificationSettings NotificationSettingsUpdateReq `json:"notificationSettings"`
}

type PostProfileResponse struct {
	Id    string `json:"id"`
	Media Media  `json:"media"`
}

type Gender string

const (
	MALE = iota
	FEMALE
)

type HighlightImageWithMedia struct {
	Url   string        `json:"url"`
	Media []IdWithMedia `json:"media"`
}

type HighlightRequest struct {
	Name     string   `json:"name"`
	StoryIds []string `json:"storyIds"`
}

type HighlightProfileResponse struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type StoryHighlight struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	ImageURL string `json:"imageUrl"`
}

type UserFollowingResponse struct {
	Following bool      `json:"following"`
	UserInfo  *UserInfo `json:"userInfo"`
}

type UserInfo struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	ImageURL string `json:"imageUrl"`
}

type FollowRequest struct {
	SubjectId string `json:"subjectId"`
	ObjectId  string `json:"objectId"`
}

type Mute struct {
	SubjectId string `json:"subjectId"`
	ObjectId  string `json:"objectId"`
}

type FollowedUsersResponse struct {
	Users []string
}

type UserRequest struct {
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
}

type EditUserRequest struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Website  string `json:"website"`
	Bio      string `json:"bio"`
	Number   string `json:"number"`
	Gender   Gender `json:"gender"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type ActivateLinkRequest struct {
	Email string `json:"email"`
}

type ChangeNewPasswordRequest struct {
	ResetPasswordId string `json:"resetPasswordId"`
	Password        string `json:"password"`
	PasswordRepeat  string `json:"passwordRepeat"`
}

func validateMediaTypeEnums(md Media) error {

	if md.MediaType != "IMAGE" && md.MediaType != "VIDEO" {
		return errors.New("Invalid media type")
	}
	return nil
}

type IdWithMedia struct {
	Id    string `json:"id"`
	Media Media  `json:"media"`
}

type FavouritePostRequest struct {
	PostId         string `json:"postId"`
	CollectionName string `json:"collectionName"`
}

type PostIdFavouritesFlag struct {
	Id         string `json:"id"`
	Favourites bool   `json:"favourites"`
}

type Media struct {
	Url       string `json:"url"`
	MediaType string `json:"mediaType"`
}

type MediaType string

const (
	IMAGE = iota
	VIDEO
)

var (
	MutE              = "Mute"
	FromPeopleIFollow = "FromPeopleIFollow"
	FromEveryOne      = "FromEveryOne"
)

func NewUser(userRequest *UserRequest) (*User, error) {
	return &User{Id: guid.New().String(),
		Name:            html.EscapeString(userRequest.Name),
		Surname:         html.EscapeString(userRequest.Surname),
		Username:        html.EscapeString(userRequest.Username),
		Email:           html.EscapeString(userRequest.Email),
		FavouritePosts:  map[string][]IdWithMedia{},
		HighlightsStory: map[string]HighlightImageWithMedia{},
		IsPrivate:       true,
		LikedPosts:      []string{},
		DislikedPosts:   []string{},
		BlockedUsers:    []string{},
		IsDeleted:       false,
		NotificationSettings: NotificationSettings{
			NotifyLike:                FromEveryOne,
			NotifyDislike:             FromEveryOne,
			NotifyFollow:              true,
			NotifyFollowRequest:       true,
			NotifyComment:             FromEveryOne,
			NotifyAcceptFollowRequest: true,
		},
		PrivacySettings: PrivacySettings{
			IsPrivate:       false,
			ReceiveMessages: true,
			IsTaggable:      true,
		},
		IsVerified: false,
	}, nil
}

func NewAgent(agentRegistrationDTO *AgentRegistrationDTO) (*User, error) {
	return &User{Id: guid.New().String(),
		Name:            html.EscapeString(agentRegistrationDTO.Name),
		Surname:         html.EscapeString(agentRegistrationDTO.Surname),
		Username:        html.EscapeString(agentRegistrationDTO.Username),
		Email:           html.EscapeString(agentRegistrationDTO.Email),
		Website: 		 html.EscapeString(agentRegistrationDTO.Website),
		FavouritePosts:  map[string][]IdWithMedia{},
		HighlightsStory: map[string]HighlightImageWithMedia{},
		IsPrivate:       true,
		LikedPosts:      []string{},
		DislikedPosts:   []string{},
		BlockedUsers:    []string{},
		IsDeleted:       false,
		NotificationSettings: NotificationSettings{
			NotifyLike:                FromEveryOne,
			NotifyDislike:             FromEveryOne,
			NotifyFollow:              true,
			NotifyFollowRequest:       true,
			NotifyComment:             FromEveryOne,
			NotifyAcceptFollowRequest: true,
		},
		PrivacySettings: PrivacySettings{
			IsPrivate:       false,
			ReceiveMessages: true,
			IsTaggable:      true,
		},
		IsVerified: false,
	}, nil
}


func validateGenderEnums(pt Gender) error {
	switch pt {
	case "MALE", "FEMALE":
		return nil
	}
	return errors.New("Invalid gender type")
}

var (
	DefaultCollection = "all posts"
)

type RegisteredUserResponseDTO struct {
	userId string
	qrCode string
}

type VerifyAccountDTO struct {
	UserId   string
	Category string
}

type RecommendedUsersResponse struct {
	Users []string
}

type FollowRecommendationResponse struct {
	Name string
	Surname string
	Username string
	ImageURL string
	RecommendedUsers []*RecommendUserInfo
}

type RecommendUserInfo struct {
	Id       string
	Username string
	ImageURL string
	SendedRequest bool
	Followed bool
}

type AgentRegistrationDTO struct {
	Username 	string
	Name	string
	Email	 string
	Surname	 string
	Website	string
	Password string
}