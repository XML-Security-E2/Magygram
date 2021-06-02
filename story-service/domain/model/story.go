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
	Id string
	Username string
	ImageURL string
}

type StoryResponse struct {
	Id string
	ContentType ContentType
	Media Media
	UserInfo UserInfo
}



func NewStoryResponse(story *Story) (*StoryResponse, error) {
	return &StoryResponse{Id: story.Id,
		Media: story.Media,
		UserInfo: story.UserInfo,
		ContentType: story.ContentType,
	}, nil
}