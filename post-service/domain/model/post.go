package model

import (
	"errors"
	"fmt"
	"github.com/beevik/guid"
	"mime/multipart"
	"strings"
)

/* Za postmana
{
    "Description" : "Description",
    "Location" : "Kraljevo",
    "PostType" : "REGULAR",
    "Media": [
        { "Url":"images/image1", "MediaType": "IMAGE" },
        { "Url":"images/image2", "MediaType": "IMAGE" },
        { "Url":"videos/video1", "MediaType": "VIDEO" }
    ],
    "UserInfo" : {
        "Id" : "2a5a1dc2-f5ad-4eb8-ae69-f4ea160422ff",
        "Username" : "testniusername",
        "ImageURL" : "images/image253"
    }
}
 */

type Post struct {
	Id string `bson:"_id,omitempty"`
	Description string `bson:"description"`
	Location string `bson:"location"`
	PostType PostType `bson:"post_type"`
	Tags []string `bson:"tags"`
	HashTags []string `bson:"hashTags"`
	Media []Media `bson:"media"`
	UserInfo UserInfo `bson:"user_info"`
	LikedBy []UserInfo `bson:"liked_by"`
	DislikedBy []UserInfo `bson:"disliked_by"`
	Comments []Comment `bson:"comments"`
}

type PostType string

const(
	REGULAR = iota
	CAMPAIGN
)

type PostRequest struct {
	Description string `json:"description"`
	Location string `json:"location"`
	Media []*multipart.FileHeader `json:"media"`
	Tags []string `json:"tags"`
}

func NewPost(postRequest *PostRequest, postOwner UserInfo, postType PostType, media []Media) (*Post, error) {
	err := validatePostTypeEnums(postType)
	if err != nil {
		return nil, err
	}

	err = validateMediaTypeEnums(media)
	if err != nil {
		return nil, err
	}

	fmt.Println(len(postRequest.Tags))

	return &Post{Id: guid.New().String(),
		Description:   postRequest.Description,
		Location:    postRequest.Location,
		HashTags: getHashTagsFromDescription(postRequest.Description),
		UserInfo: postOwner,
		LikedBy: []UserInfo{},
		DislikedBy: []UserInfo{},
		Comments: []Comment{},
	}, nil
}

func getHashTagsFromDescription(description string) []string {
	var hashTags []string
	words := strings.Fields(description)
	for _, w := range words {
		if strings.HasPrefix(w, "#") {
			hashTags = append(hashTags, strings.TrimPrefix(w, "#"))
		}
	}
	return hashTags
}

func validatePostTypeEnums(pt PostType) error {
	switch pt {
	case "REGULAR", "CAMPAIGN":
		return nil
	}
	return errors.New("Invalid post type")
}

func validateMediaTypeEnums(md []Media) error {
	if len(md)==0 {
		return nil
	}

	for _, media := range md {
		if media.MediaType != "IMAGE" && media.MediaType !="VIDEO" {
			return errors.New("Invalid media type")
		}
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

type Comment struct {
	Id string
	CreatedBy UserInfo
	Content string
}