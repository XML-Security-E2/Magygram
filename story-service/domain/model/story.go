package model

import (
	"errors"
	"github.com/beevik/guid"
)

type Story struct {
	Id string `bson:"_id,omitempty"`
	ContentType ContentType `bson:"content_type"`
	Media Media `bson:"media"`
	UserInfo UserInfo `bson:"user_info"`
	VisitedBy []UserInfo `bson:"visited_by"`
}

type ContentType string

const(
	REGULAR = iota
	CAMPAIGN
)

func NewStory(postOwner UserInfo, storyType ContentType, media Media) (*Story, error) {
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

type MediaContent struct{
	Url string
	MediaType string
	StoryId string
}


func NewStoryResponse(story *Story, media []MediaContent,firstUnvisitedStory int) (*StoryResponse, error) {
	return &StoryResponse{
		Media: media,
		UserInfo: story.UserInfo,
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