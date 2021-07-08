package model

import (
	"errors"
	"github.com/beevik/guid"
	"time"
)

type Tag struct {
	Id string `bson:"_id,omitempty"`
	Username string `bson:"username"`
}

type Story struct {
	Id string `bson:"_id,omitempty"`
	ContentType ContentType `bson:"content_type"`
	Media Media `bson:"media"`
	UserInfo UserInfo `bson:"user_info"`
	VisitedBy []UserInfo `bson:"visited_by"`
	CreatedTime time.Time `bson:"created_time"`
	Tags []Tag `bson:"tags"`
	Website string `bson:"website"`
	IsDeleted bool `bson:"deleted"`
}

type InfluencerRequest struct {
	PostIdInfluencer string `json:"postIdInfl"`
	UserId string `json:"userId"`
	Username string `json:"username"`
}

type ContentType string

const(
	REGULAR = iota
	CAMPAIGN
)

type CampaignRequest struct {
	ContentId string `json:"contentId"`
	ExposeOnceDate time.Time `json:"exposeOnceDate"`
	MinDisplaysForRepeatedly int `json:"minDisplaysForRepeatedly"`
	Frequency CampaignFrequency `json:"frequency"`
	TargetGroup TargetGroup `json:"targetGroup"`
	DisplayTime int `json:"displayTime"`
	DateFrom time.Time `json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
	Type string `json:"campaignType"`
}

type TargetGroup struct {
	MinAge int `json:"minAge"`
	MaxAge int `json:"maxAge"`
	Gender GenderType `json:"gender"`
}

type GenderType string

const (
	MALE = iota
	FEMALE
	ANY
)

type CampaignFrequency string

const(
	ONCE = iota
	REPEATEDLY
)

func NewStory(postOwner UserInfo, storyType ContentType, media Media, tags []Tag, website string) (*Story, error) {
	err := validateStoryTypeEnums(storyType)
	if err != nil {
		return nil, err
	}

	err = validateMediaTypeEnums(media)
	if err != nil {
		return nil, err
	}

	return &Story{Id: guid.New().String(),
				ContentType: storyType,
				Media: media,
				UserInfo: postOwner,
				VisitedBy: []UserInfo{},
				CreatedTime: time.Now(),
				Tags: tags,
				IsDeleted: false,
				Website: website,
	}, nil
}

func NewStoryInfluencer( story *Story, postOwner UserInfo) (*Story, error) {
	err := validateStoryTypeEnums(story.ContentType)
	if err != nil {
		return nil, err
	}

	err = validateMediaTypeEnums(story.Media)
	if err != nil {
		return nil, err
	}

	return &Story{Id: guid.New().String(),
		ContentType: story.ContentType,
		Media: story.Media,
		UserInfo: postOwner,
		VisitedBy: []UserInfo{},
		CreatedTime: time.Now(),
		Tags: story.Tags,
		IsDeleted: false,
		Website: story.Website,
	}, nil
}

func validateStoryTypeEnums(pt ContentType) error {
	switch pt {
	case "REGULAR", "CAMPAIGN":
		return nil
	}
	return errors.New("Invalid post type")
}

func validateMediaTypeEnums(md Media) error {

	if md.MediaType != "IMAGE" && md.MediaType !="VIDEO" {
			return errors.New("Invalid media type")
	}
	return nil
}

type Media struct {
	Url string
	MediaType string
}

type MediaType string

const(
	IMAGE = iota
	VIDEO
)

type UserInfo struct {
	Id string `bson:"id"`
	Username string
	ImageURL string
}

type AgentInfo struct {
	Id string
	Username string
	ImageURL string
	Website string
}

type UsersStoryResponseWithUserInfo struct {
	Id string `json:"id"`
	UserInfo UserInfo `json:"userInfo"`
	ContentType ContentType `json:"contentType"`
	Media Media `json:"media"`
	DateTime string `json:"dateTime"`
	Website string
}

type UsersStoryResponse struct {
	Id string `json:"id"`
	ContentType ContentType `json:"contentType"`
	Media Media `json:"media"`
	DateTime string `json:"dateTime"`
}

type StoryResponse struct {
	ContentType ContentType
	Media []MediaContent
	UserInfo UserInfo
	FirstUnvisitedStory int
}

type IdMediaWebsiteResponse struct {
	Id string
	Media Media
	Website string
	Likes int
	Dislikes int
	Comments int
	StoryViews int
}


type MediaContent struct{
	Url string
	MediaType string
	StoryId string
	Tags []Tag
	ContentType ContentType
	Website string
}

type HighlightImageWithMedia struct {
	Url  string `json:"url"`
	Media  []IdWithMedia `json:"media"`
}
type IdWithMedia struct {
	Id string `json:"id"`
	Media Media `json:"media"`
}

type HighlightRequest struct {
	Name  string `json:"name"`
	StoryIds  []string `json:"storyIds"`
}
type StoryResponseForAdmin struct {
	ContentType ContentType
	Media []MediaContent
	FirstUnvisitedStory int
}

func NewStoryResponse(story *Story, media []MediaContent,firstUnvisitedStory int) (*StoryResponse, error) {
	return &StoryResponse{
		Media: media,
		UserInfo: story.UserInfo,
		ContentType: story.ContentType,
		FirstUnvisitedStory: firstUnvisitedStory,
	}, nil
}

func NewStoryResponseForAdmin(story *Story, media []MediaContent,firstUnvisitedStory int) (*StoryResponseForAdmin, error) {
	return &StoryResponseForAdmin{
		Media: media,
		ContentType: story.ContentType,
		FirstUnvisitedStory: firstUnvisitedStory,
	}, nil
}

type StoryInfoResponse struct {
	Id string
	UserInfo UserInfo
	Visited bool
}

func NewStoryInfoResponse(story *Story, visited bool) (*StoryInfoResponse, error) {
	return &StoryInfoResponse{Id: story.Id,
		UserInfo: story.UserInfo,
		Visited: visited,
	}, nil
}

type FollowedUsersResponse struct {
	Users []string
}