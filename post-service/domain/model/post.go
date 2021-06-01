package model

import (
	"errors"
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
	ContentType ContentType `bson:"post_type"`
	Tags []string `bson:"tags"`
	HashTags []string `bson:"hashTags"`
	Media []Media `bson:"media"`
	UserInfo UserInfo `bson:"user_info"`
	LikedBy []UserInfo `bson:"liked_by"`
	DislikedBy []UserInfo `bson:"disliked_by"`
	Comments []Comment `bson:"comments"`
}

type ContentType string

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

func NewPost(postRequest *PostRequest, postOwner UserInfo, postType ContentType, media []Media) (*Post, error) {
	err := validatePostTypeEnums(postType)
	if err != nil {
		return nil, err
	}

	err = validateMediaTypeEnums(media)
	if err != nil {
		return nil, err
	}

	return &Post{Id: guid.New().String(),
		Description:   postRequest.Description,
		Location:    postRequest.Location,
		HashTags: getHashTagsFromDescription(postRequest.Description),
		UserInfo: postOwner,
		Media: media,
		Tags: postRequest.Tags,
		ContentType : postType,
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

func validatePostTypeEnums(pt ContentType) error {
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

type PostResponse struct {
	Id string
	Description string
	Location string
	ContentType ContentType
	Tags []string
	HashTags []string
	Media []Media
	UserInfo UserInfo
	LikedBy []UserInfo
	DislikedBy []UserInfo
	Comments []Comment
	Liked bool
	Disliked bool
}



func NewPostResponse(post *Post, liked bool, disliked bool) (*PostResponse, error) {
	return &PostResponse{Id: post.Id,
		Description:   post.Description,
		Location:    post.Location,
		HashTags: post.HashTags,
		Media: post.Media,
		UserInfo: post.UserInfo,
		LikedBy: post.LikedBy,
		DislikedBy: post.DislikedBy,
		Comments: post.Comments,
		Liked: liked,
		Disliked: disliked,
	}, nil
}

type PostId struct {
	Id string
}